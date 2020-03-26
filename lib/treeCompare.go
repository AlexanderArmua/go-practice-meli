package lib

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

// Same dice si dos arboles binarios son iguales
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	close(ch1)
	close(ch2)

	for i := 0; i < len(ch1); i++ {
		node1 := <-ch1
		node2 := <-ch2
		if node1 != node2 {
			return false
		}
	}

	return true
}

// Walk recorre un arbol binario y retorna sus elementos en un chanel
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value // Envio el valor medio en el chanell

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// TestTreeCompare hace el test del archivo y cosas varias
func TestTreeCompare() {
	tree1 := tree.New(1)
	tree2 := tree.New(1)

	if Same(tree1, tree2) {
		fmt.Println("Los Trees son iguales")
	} else {
		fmt.Println("Los Trees son distintos")
	}
}
