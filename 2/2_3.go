package main
import "fmt"
import "math/cmplx"
import "math"

//import (
//    "fmt"
//    "math"
//    "math/cmplx"
//)

var aa = 3
var ss = "kkk"
// bb := true不可以

var (
    d = 4
    e = "eee"
    f = true
)

func variableZeroValue() {
    //
    var a int
    var s string
    fmt.Println(a, s)
    fmt.Printf("%d %q\n", a, s)
}

func valirableInitiablValue() {
    var a,b int = 3, 4
    var s string = "abc"
    fmt.Println(a, b, s)
    fmt.Printf("%d %q\n", a, s)
}

func variableTypeDeduction() {
    var a, b, c, s = 3, 4, true, "def"
    fmt.Println(a, b, c, s)
}

func variableShorter() {
    a, b, c, s := 3, 4, true, "def"
    b = 5
    fmt.Println(a, b, c, s)
}

func euler() {
    c := 3+4i
    fmt.Println(
        cmplx.Exp(1i * math.Pi) + 1, cmplx.Pow(math.E, 1i * math.Pi) + 1, cmplx.Abs(c))
    fmt.Printf(
        "%.3f, %.3f, %d\n",
        cmplx.Exp(1i * math.Pi) + 1, cmplx.Pow(math.E, 1i * math.Pi) + 1, cmplx.Abs(c))
}

func triangle() {
    var a, b int = 3, 4
    var c int
    c = int(math.Sqrt(float64(a * a + b * b)))
    fmt.Println(c)
}

func consts() {
    const file_name = "abc.txt"
    const c, d int = 3, 4
    const (
        filename = "abc.txt"
        a, b = 3,4
    )
    var e int = int(math.Sqrt(a * a + b * b))
    e  = int(math.Sqrt(float64(c * c + d * d)))
    fmt.Println(filename, file_name, e)
}

func enums() {
    const(
        cpp = iota
        java
        python
        golang
        javascript
    )
    fmt.Println(cpp, java, python, golang, javascript);
    const (
        b = 1 << (10*iota)
        kb
        gb
        tb
        pb
    )
    fmt.Println(b, kb, gb, tb, pb)
}

func main() {
    fmt.Println("Hello world")
    variableZeroValue()
    valirableInitiablValue()
    variableTypeDeduction()
    variableShorter()
    fmt.Println(ss, d, e, f)
    euler()
    triangle()

    consts()
    enums()
}
