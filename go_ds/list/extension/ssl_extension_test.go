package extension

import (
	"algorithm-ds/go_ds/list"
	"testing"
)

var ssl *SSLExtension

func init() {
	l := list.NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToHead(i + 1)
	}

	ssl = NewSSLExtension(*l)
}

func TestReverse(t *testing.T) {
	ssl.Print()
	ssl.Reverse()
	ssl.Print()
}

func TestHasCycle(t *testing.T) {
	t.Log(ssl.HasCycle())
	ssl.Head.Next.Next.Next.Next.Next.Next = ssl.Head.Next.Next.Next
	t.Log(ssl.HasCycle())
}
