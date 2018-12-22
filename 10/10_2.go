package main

import (
    "fmt"
)

type worker struct {
    in chan int
    done chan bool
}

func createWorker(id int) worker {
    w := worker {
        in : make(chan int),
        done : make(chan bool),
    }
    // 这里需要增加go func,否则会死循环这里
    go doWorker(id, w.in, w.done)
    return w
}

func chanDemo() {
    var workers [10]worker
    for i := 0; i < 10; i++ {
        workers[i] = createWorker(i)
    }
//    for i := 0; i < 10; i++ {
//        workers[i].in <- 'a' + i
//        <-workers[i].done
//    }
//    for i := 0; i < 10; i++ {
//        workers[i].in <- 'A' + i
//        <-workers[i].done
//    }
    for i, worker := range workers {
        worker.in <- 'a' + i
    }
    for i, worker := range workers {
        worker.in <- 'A' + i
    }
    for _, worker := range workers {
        <-worker.done
        <-worker.done
    }
}

func doWorker(id int, c chan int,
        done chan bool) {
    for n := range c {
        fmt.Printf("Worker %d receivec %c \n", id, n)
        go func() {
            done <- true
        }()
    }
}

func main() {
    chanDemo()
}
