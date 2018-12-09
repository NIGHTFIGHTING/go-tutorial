package main

import (
    "fmt"
    "go-tutorial/4/tree"
)

type myTreeNode struct {
    node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
    if myNode == nil || myNode.node == nil {
        return
    }
    left := myTreeNode{myNode.node.Left}
    right := myTreeNode{myNode.node.Right}
    left.postOrder()
    right.postOrder()
    myNode.node.Print()
}

func main() {
    var root tree.Node

    root = tree.Node{Value: 3}
    root.Left = &tree.Node{}
    root.Right = &tree.Node{5, nil, nil}
    root.Right.Left = new(tree.Node)
    root.Left.Right = tree.CreateNode(2)

    nodes := []tree.Node {
        {Value: 3},
        {},
        {6, nil, &root},
    }
    fmt.Println("nodes: ", nodes)

    root.Print()
    fmt.Println()

    // 值传递
    root.Right.SetByValue(6)
    root.Right.Print()
    fmt.Println()
    // 指针传递
    root.Right.SetValue(6)
    root.Right.Print()
    fmt.Println()

    pRoot := &root
    pRoot.Print()
    pRoot.SetValue(200)
    pRoot.Print()

    pRoot = nil
    pRoot.SetValue(200)
    // pRoot.Print()
    root.Traverse()

    nodeCount := 0
    root.TraverseFunc(func(node *tree.Node) {
        nodeCount++
    })
    fmt.Println("Node count:", nodeCount)
}
