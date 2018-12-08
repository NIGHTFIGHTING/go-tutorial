package tree

import "fmt"

type Node struct {
    Value int
    Left, Right *Node
}

func CreateNode(Value int) *Node {
    return &Node{Value: Value}
}

func (node Node) Print() {
    fmt.Println()
    fmt.Print(node.Value)
    fmt.Println()
}

func (node Node) SetByValue(Value int) {
    node.Value = Value
}

func (node *Node) SetValue(Value int) {
    if nil == node {
        fmt.Println("Setting Value to nil " +
                "node. Ignored.")
        return
    }
    node.Value = Value 
}
