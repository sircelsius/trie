package main

import (
	"math/rand"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func TestTrie_New(t *testing.T) {
	trie := New()
	if trie.root == nil {
		t.Errorf("Expected non null trie, got %v", trie)
	}
}

func TestTrie_ContainsWord(t *testing.T) {
	trie := New()
	if trie.ContainsWord("foo") {
		t.Errorf("Trie should not contain word!")
	}

	trie.InsertWord("foo")
	if !trie.ContainsWord("foo") {
		t.Errorf("Trie should contain the word foo, got %v", trie)
	}
}

func TestTrie_InsertWord(t *testing.T) {
	trie := New()
	trie.InsertWord("foo")
	trie.InsertWord("bar")
	if len(trie.root.Children) != 2 {
		t.Errorf("Root should have children, got %v", trie)
	}
}

func BenchmarkTrie_InsertWord10(b *testing.B) {
	benchMarkInsert(10, b)
}

func BenchmarkTrie_InsertWord100(b *testing.B) {
	benchMarkInsert(100, b)
}

func BenchmarkTrie_InsertWord1000(b *testing.B) {
	benchMarkInsert(1000, b)
}

func BenchmarkTrie_InsertWord10000(b *testing.B) {
	benchMarkInsert(10000, b)
}

func BenchmarkTrie_ContainsWord100X100(b *testing.B) {
	randomTrie(100, 100, b)
}

func randomTrie(size int, count int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie := New()
		for j := 0; j < count; j++ {
			word := String(size)
			trie.InsertWord(word)
		}
		word := String(size)
		trie.ContainsWord(word)
	}
}

func benchMarkInsert(size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie := New()
		word := String(size)
		trie.InsertWord(word)
	}
}