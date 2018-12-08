package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
    fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map
    fmt.Println(m2)

	var m3 map[string]int // m3 == nil
    fmt.Println(m3)

	fmt.Println("Traversing map m")
    for k, v := range m {
        fmt.Println("key:", k , " value:", v)
    }
    for k := range m {
        fmt.Println("key:", k)
    }
    for _, v := range m {
        fmt.Println("value:", v)
    }

	fmt.Println("Getting values")
    courseName, ok := m["course"]
    fmt.Println(courseName, ok)
    if causeName, ok := m["cause"]; ok {
        fmt.Println(causeName, ok)
    } else {
        fmt.Println("key does not exist")
    }

	fmt.Println("Deleting values")
    delete(m, "site")
    fmt.Println(m)

	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}
