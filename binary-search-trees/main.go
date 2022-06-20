package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) PreOrderTraversal() {
	if tree != nil {
		fmt.Printf("%d ", tree.Value)
		tree.Left.PreOrderTraversal()
		tree.Right.PreOrderTraversal()
	}
	return
}

func main() {
	tree := &BST{Value: 12}
	tree.Insert(45)
	tree.Insert(13)
	tree.Insert(12)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(1)
	tree.PreOrderTraversal()
}
