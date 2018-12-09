package main

import (
    "fmt"
    "go-tutorial/5/mock"
    "go-tutorial/5/real"
    "time"
)

type Retriever interface {
    Get(url string) string
}

type Poster interface {
    Post(url string,
        form map[string]string) string
}

func download(r Retriever) string {
    return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
    poster.Post("http://www.imooc.com",
        map[string]string {
            "name": "ccmouse",
            "course": "golang",
        })
}

type RetrieverPoster interface {
    Retriever
    Poster
}

const url = "http://www.imooc.com"
func session(s RetrieverPoster) string {
    s.Post(url, map[string]string {
        "contents": "another faked imooc.com",
    })
    return s.Get(url)
}

func main() {
    var r Retriever
    r = &mock.Retriever{"this is a fake imooc.com"}
    fakedRetriever := mock.Retriever{"this is a fake imooc.com"}
    inspect(r)
    // Type assertion
    if mockRetriever, ok := r.(*mock.Retriever); ok {
        fmt.Println(mockRetriever.Contents)
    } else {
        fmt.Println("not a mock retriever")
    }
    // 
    fmt.Println("Try a session with mockRetriever")
    fmt.Println(session(&fakedRetriever))


    r = &real.Retriever{}
    inspect(r)
    r = &real.Retriever{
        UserAgent: "Mozilla/5.0",
        TimeOut: time.Minute,
    }
    inspect(r)
    // Type assertion
    if realRetriever, ok := r.(*real.Retriever); ok {
        fmt.Println("Timeout: ", realRetriever.TimeOut)
    }
    if mockRetriever, ok := r.(*mock.Retriever); ok {
        fmt.Println(mockRetriever.Contents)
    } else {
        fmt.Println("not a mock retriever")
    }

    // fmt.Println(download(r))
    // fmt.Println(download(mock.Retriever{"i am a good coder"}))
    // fmt.Println(download(real.Retriever{}))
}

func inspect(r Retriever) {
    fmt.Println("Inspecting", r)
    fmt.Printf("> %T %v\n", r, r)
    fmt.Println("> Type Switch:")
    switch v := r.(type) {
    case *mock.Retriever:
        fmt.Println("Contents:", v.Contents)
    case *real.Retriever:
        fmt.Println("UserAgent:", v.UserAgent)
    }
    fmt.Println()
}

