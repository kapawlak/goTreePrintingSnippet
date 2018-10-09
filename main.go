package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Tree struct {
	root *Node
	len  int
}

type Node struct {
	key         int
	left, right *Node
	height      int8
	children    int
}

func (k *Node) RandInit() *Node {
	k.key = rand.Intn(150)
	k.height = 0
	return k
}

func PrintTree(kt *Tree, data string) {
	//height of tree *2 to account for "|"
	height := int(2 * (kt.root.height + 1))

	//make a [][]interface{} to hold strings and tree data
	treearray := make([][]interface{}, height)

	//calculate the number of maximum elements in the last row
	width := int(math.Pow(2, float64(kt.root.height-1)))

	//make an array with width double the max elements + the height to allow for padding
	for i := range treearray {
		treearray[i] = make([]interface{}, height+2*width+1)
		for j := range treearray[i] {
			//the array is initialized with only space
			treearray[i][j] = " "
		}
	}
	//CoordPrint take a pointer to the array and overwrites the spaces where appropriate
	CoordPrint(kt.root, data, 0, (height+2*width)/2+1, 0, &treearray, (height+2*width)/2)
	//Print the array element by element
	for _, ta := range treearray {
		for _, r := range ta {
			//I set to %3v because I have 3 digit numbers. If this is changed to e.g. %2v
			//you also have to change --- to -- etc in CoordPrint. I could automate but
			//I'm lazy
			fmt.Printf("%3v", r)
		}
		fmt.Println()
	}
	fmt.Println()

}

func CoordPrint(k *Node, data string, y, x, px int, treearray *[][]interface{}, width int) {
	//Coord print takes the Node, data, the y level, x position, parent's x position,
	//array pointer, and the current spacing

	//check that you aren't accessing a nil node
	if k != nil {
		//print the desired data
		switch data {
		case "key":
			(*treearray)[y][x] = k.key
		case "child":
			(*treearray)[y][x] = k.children
		}
		//for all children nodes, put characters above.
		if y > 0 {
			//if this is a left child, put left chars
			if px > x {
				(*treearray)[y-2][x] = "┌"
				(*treearray)[y-1][x] = "│"
				for i := x + 1; i < px; i++ {
					(*treearray)[y-2][i] = "───"
				}
			} else { //else put right chars
				(*treearray)[y-2][x] = "──┐"
				(*treearray)[y-1][x] = "│"
				for i := px + 1; i < x; i++ {
					(*treearray)[y-2][i] = "───"
				}

			}
		}
		//lazy way to avoid index error
		if y+2 < len(*treearray) {
			//reduce width by a factor of 2
			w := width / 2
			//pass to left and right nodes with new coordinates
			CoordPrint(k.left, data, y+2, x-w, x, treearray, w)
			CoordPrint(k.right, data, y+2, x+w, x, treearray, w)
		}
	}

}

func main() {

	//Barebones example of use
	tree := new(Tree)
	tree.root = (&Node{}).RandInit()
	PrintTree(tree, "key")

	tree.root.left = (&Node{}).RandInit()
	tree.root.height = 1
	PrintTree(tree, "key")

	tree.root.right = (&Node{}).RandInit()
	PrintTree(tree, "key")

	tree.root.left.left = (&Node{}).RandInit()
	tree.root.height = 2
	PrintTree(tree, "key")

	tree.root.right.right = (&Node{}).RandInit()
	PrintTree(tree, "key")

	tree.root.right.left = (&Node{}).RandInit()
	PrintTree(tree, "key")

	tree.root.right.right.left = (&Node{}).RandInit()
	tree.root.height = 3
	PrintTree(tree, "key")

	tree.root.right.right.right = (&Node{}).RandInit()
	PrintTree(tree, "key")

}
