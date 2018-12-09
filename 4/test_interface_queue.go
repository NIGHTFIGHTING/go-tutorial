//package queue
package main

import (
    "fmt"
)

// 表示任意类型 interface{}
type Queue []interface{}

func (q *Queue) Push(v int) {
    *q = append(*q, v.int)
}

func (q *Queue) Pop() int {
    head := (*q)[0]
    *q = (*q)[1:]
    // 取出int值
    return head.(int)
}

func (q *Queue) IsEmpty() bool {
    return len(*q) == 0
}

func main() {
    q := Queue{1}
    q.Push(2)
    q.Push(3)

    fmt.Println(q.Pop())
    fmt.Println(q.Pop())
    fmt.Println(q.IsEmpty())
    fmt.Println(q.Pop())
    fmt.Println(q.IsEmpty())

    // q.Push("abc")
    // fmt.Println(q.Pop())
}
