package main

func main() {
  // create int channel
  c1 := make(chan int)
  // assign c1 to c2 variable
  var chan int c2 = c1

  // 자료를 뺄 수만있는 채널의 자료형 (receive)
  var <-cahn int c1 = c1
  // 자료를 넣을 수만있는 채널의 자료형 (send)
  var chan<- int c4 = c1
}
