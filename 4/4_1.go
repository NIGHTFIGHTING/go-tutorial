package main

import (
    "fmt"
)

type treeNode struct {
    value int
    left, right *treeNode
}

func createNode(value int) *treeNode {
    return &treeNode{value: value}
}

func (node treeNode) print() {
    fmt.Println()
    fmt.Print(node.value)
    fmt.Println()
}

func (node treeNode) setByValue(value int) {
    node.value = value
}

func (node *treeNode) setValue(value int) {
    if nil == node {
        fmt.Println("Setting value to nil " +
                "node. Ignored.")
        return
    }
    node.value = value 
}

// 中序遍历
func (node *treeNode) traverse() {
    if node == nil {
        return
    }
    node.left.traverse()
    node.print()
    node.right.traverse()
}

func main() {
    var root treeNode

    root = treeNode{value: 3}
    root.left = &treeNode{}
    root.right = &treeNode{5, nil, nil}
    root.right.left = new(treeNode)
    root.left.right = createNode(2)

    nodes := []treeNode {
        {value: 3},
        {},
        {6, nil, &root},
    }
    fmt.Println("nodes: ", nodes)

    root.print()
    fmt.Println()

    // 值传递
    root.right.setByValue(6)
    root.right.print()
    fmt.Println()
    // 指针传递
    root.right.setValue(6)
    root.right.print()
    fmt.Println()

    pRoot := &root
    pRoot.print()
    pRoot.setValue(200)
    pRoot.print()

    pRoot = nil
    pRoot.setValue(200)
    // pRoot.print()
    root.traverse()

}
