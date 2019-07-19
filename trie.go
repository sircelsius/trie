package main

import (
	"fmt"
	"sync"
)

type Node struct {
	IsWord   bool
	Value    rune
	Children map[rune]*Node
}

type Trie struct {
	root *Node
	lock sync.RWMutex
}

func New() *Trie {
	return &Trie{
		root: &Node{
			IsWord:   false,
			Value:    rune(0),
			Children: make(map[rune]*Node),
		},
	}
}

func (t *Trie) InsertWord(word string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.root.insertWord(word)
}

func (n *Node) insertWord(word string) {
	if len(word) == 0 || n == nil {
		return
	}

	firstRune := []rune(word)[0]

	if n.Children[firstRune] == nil {
		n.Children[firstRune] = &Node{
			Value:    firstRune,
			IsWord:   len(word) == 1,
			Children: make(map[rune]*Node),
		}
	}

	n.Children[firstRune].insertWord(word[1:])
}

func (t *Trie) ContainsWord(word string) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.containsWord(word)
}

func (n *Node) containsWord(word string) bool {
	if len(word) == 0 {
		return false
	}

	firstRune := []rune(word)[0]

	if n.Children[firstRune] == nil {
		return false
	}

	if len(word) == 1 && n.Children[firstRune].IsWord {
		return true
	}

	return n.Children[firstRune].containsWord(word[1:])
}

func (n *Node) String() string {
	value := fmt.Sprintf("{  Value: %v, IsWord: %v, Children: {  ", string(n.Value), n.IsWord)
	for _, v := range n.Children {
		value += fmt.Sprintf("%v", v)
	}
	value += "  }  }"
	return value
}

func (t *Trie) String() string {
	return t.root.String()
}
