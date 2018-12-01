package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
    const filename = "abc.txt"
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Printf("error: %s\n", err)
    } else {
        fmt.Printf("%s\n", contents)
    }
}
