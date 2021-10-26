package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

var urls = []string{
	"https://lh3.googleusercontent.com/proxy/F7ymq4vImpXO6X-459_w3F8s4RtlEm4Rle1TFpE6JK_rIQ-aD8lldT0cIJ9AzLieFHWXTG7t74UAyEp3u0jl2WTM85-Kk9z1zs_T80qknwcwObxKSZo",
	"http://image.dongascience.com/Photo/2020/03/5bddba7b6574b95d37b6079c199d7101.jpg",
	"https://www.ui4u.go.kr/depart/img/content/sub03/img_con03030100_01.jpg",
}

func main() {
	var wg sync.WaitGroup
	// 주의할점 go 루틴 내부에서 wg.Wait(1) 포함할 경우 메인 고루틴이 wg.Wait() 를 통과해버릴 가능성이 있습니다. 이런 상태를 레이스 컨디션이라고 합니다.
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	// wati group 이 없다면 파일이 다 다운로드 되기전에 압축을 시도 합니다.
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("images.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}

// download downloads url and return the contents and error
func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

// urlToFilename return the filename part from the rawurl
func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	// filepath.Base 함수는 경로에서 가장 마지막 부분을 return
	return filepath.Base(url.Path), nil
}

// writeZip writes a zip archive file
func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}
