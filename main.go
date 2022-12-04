package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {

	f, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	print(countWords(f))
}

type bt struct {
	val         int
	left, right *bt
}

func nbt(v int) *bt          { return &bt{v, nil, nil} }
func (root *bt) l(n *bt) *bt { root.left = n; return root }
func (root *bt) r(n *bt) *bt { root.right = n; return root }

func hasDublicates(root *bt) bool {
	contains := func(s []int, val int) bool {
		for _, el := range s {
			if el == val {
				return true
			}
		}
		return false
	}
	var values []int
	var hd func(*bt) bool
	hd = func(r *bt) bool {
		if r == nil {
			return false
		}
		if contains(values, r.val) {
			return true
		}
		values = append(values, r.val)
		return hd(r.left) || hd(r.right)
	}
	return hd(root)
}

func levelMax(root *bt) []int {
	var values [][]int
	var traverse func(*bt, int)
	traverse = func(root *bt, lvl int) {
		if root == nil {
			return
		}
		if len(values) <= lvl {
			tmp := make([]int, 1)
			values = append(values, tmp)
		}
		values[lvl] = append(values[lvl], root.val)
		traverse(root.left, lvl+1)
		traverse(root.right, lvl+1)
	}
	traverse(root, 0)

	lm := make([]int, len(values))
	for i, lvl := range values {
		for _, val := range lvl {
			if lm[i] < val {
				lm[i] = val
			}
		}
	}

	return lm
}

type bst struct {
	val         string
	count       int
	right, left *bst
}

func countWords(file []byte) *bst {
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(string(file), f)

	var root *bst
	for _, word := range words {
		insert(&root, word)
	}
	return root
}

func insert(root **bst, in string) {
	if (*root) == nil {
		*root = &bst{in, 1, nil, nil}
	} else if in == (*root).val {
		(*root).count++
	} else if in < (*root).val {
		insert(&(*root).left, in)
	} else /*in > root.val*/ {
		insert(&(*root).right, in)
	}
}

func print(root *bst) {
	if root == nil {
		return
	}
	print(root.left)
	fmt.Println(root.val, ":", root.count)
	//fmt.Printf("%v:%v", root.val, root.count)
	print(root.right)
}
