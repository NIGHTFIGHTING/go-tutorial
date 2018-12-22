package main

import (
    "fmt"
    "time"
)

func createWorker(id int) chan int {
    c := make(chan int)
    // 这里需要增加go func,否则会死循环这里
    go func() {
        for {
            fmt.Printf("Worker %d receivec %c \n", id, <-c)
        }
    }()
    return c
}

func chanDemo() {
    var channels [10]chan int
    for i := 0; i < 10; i++ {
        channels[i] = createWorker(i)
    }
    for i := 0; i < 10; i++ {
        channels[i] <- 'a' + i
    }
    for i := 0; i < 10; i++ {
        channels[i] <- 'A' + i
    }
    time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
    // <-c 表示可以接收数据
//    for {
//        n, ok := <- c
//        if !ok {
//            break
//        }
//        fmt.Printf("Worker %d receivec %c \n", id, n)
//    }
    for n := range c {
        fmt.Printf("Worker %d receivec %c \n", id, n)
    }
}

func bufferedChannel() {
    // 设置一个缓冲区,可以缓冲3个元素,可以不接受
    c := make(chan int, 3)
    go worker(0, c)
    c <- 'a'
    c <- 'b'
    c <- 'c'
    c <- 'd'
    time.Sleep(time.Millisecond)
}

func channelClose() {
    c := make(chan int, 3)
    go worker(0, c)
    c <- 'a'
    c <- 'b'
    c <- 'c'
    c <- 'd'
    close(c)
    time.Sleep(time.Millisecond)
}

func main() {
    //chanDemo()
    //bufferedChannel()
    channelClose()
}
