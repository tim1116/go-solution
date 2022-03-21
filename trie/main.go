package main

import (
	"fmt"
	"trie"
)

func main() {
	t := trie.NewTrie()
	// 插入
	value := trie.Value{"aaa", "bbb"}
	t.Insert(value)
	t.Insert(trie.Value{"aaa", "bbb"})
	t.Insert(trie.Value{"aaa", "ddd"})
	t.Insert(trie.Value{"aaa", "bbb", "ccc"})

	// 查找
	res := t.Search(trie.Value{"aaa", "bbb"})
	showSearch(res)
	// 必须完全匹配
	res = t.Search(trie.Value{"aaa"})
	showSearch(res)

	res = t.Search(trie.Value{"aaa", "bbb", "ccc", "ddd"})
	showSearch(res)
}

func showSearch(node *trie.Trie) {
	if node != nil {
		fmt.Println(node.Value())
	} else {
		fmt.Println("没有查找到")
	}
}
