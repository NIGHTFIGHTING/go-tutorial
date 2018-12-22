package main

import (
    "fmt"
    "time"
)

func chanDemo() {
    c := make(chan int)
    go func() {
        for {
            n := <-c
            fmt.Println(n)
        }
    }()
    c <- 1
    c <- 2
    time.Sleep(time.Millisecond)
}

func worker1(id int, c chan int) {
    for {
        n := <-c
        fmt.Printf("Worker %d receivec %d \n", id, n)
    }
}

func chanDemo1() {
    c := make(chan int)
    go worker1(0, c)
    c <- 1
    c <- 2
    time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
    // <-c 表示可以接收数据
    for {
        fmt.Printf("Worker %d receivec %c \n", id, <-c)
    }
}

func chanDemo2() {
    var channels [10]chan int
    for i := 0; i < 10; i++ {
        channels[i] = make(chan int)
        go worker(i, channels[i])
    }
    for i := 0; i < 10; i++ {
        channels[i] <- 'a' + i
    }
    for i := 0; i < 10; i++ {
        channels[i] <- 'A' + i
    }
    time.Sleep(time.Millisecond)
}

// chan<-表示可以发数据
func createWorker(id int) chan<- int {
    c := make(chan int)
    go worker(id, c)
    return c
}

func chanDemo3() {
    // chan<-表示可以发数据
    var channels [10]chan<- int
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

func bufferedChannel() {
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
    chanDemo1()
//    chanDemo()
    chanDemo2()
//    chanDemo3()
//    bufferedChannel()

}
