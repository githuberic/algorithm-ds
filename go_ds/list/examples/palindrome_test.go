package examples

import (
	"algorithm-ds/go_ds/list"
	"testing"
)

func TestPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := list.NewLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome1(l))
	}
}

func TestPalindrome2(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := list.NewLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome2(l))
		l.Print()
	}
}
