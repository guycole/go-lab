package main

import "fmt"

type nodeType struct {
	ivalue int
	next   *nodeType
}

func newNodeType(ivalue int) *nodeType {
	result := nodeType{ivalue: ivalue}
	return &result
}

func main() {
	fmt.Printf("begin\n")

	rootNode := newNodeType(123)
	rootNode.next = newNodeType(234)

	var currentNode *nodeType

	for currentNode = rootNode; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("%d\n", currentNode.ivalue)
	}

	fmt.Printf("end\n")
}
