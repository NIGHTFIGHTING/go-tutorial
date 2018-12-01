package main
import "fmt"

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

func consts() {
    const (
        filename = "abc.txt"
        a, b = 3,4
    )
}


func main() {
    fmt.Println("Hello world")
    variableZeroValue()
    valirableInitiablValue()
    variableTypeDeduction()
    variableShorter()
    fmt.Println(ss, d, e, f)
    consts()
}
