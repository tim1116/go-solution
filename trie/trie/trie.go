package trie

type Value []string

// 前缀数Node 实现一个类似路由注册于匹配的功能
type Trie struct {
	children []*Trie
	isEnd    bool
	value    string
}

func NewTrie() *Trie {
	return &Trie{}
}

func emptyValue(s Value) bool {
	if s == nil || len(s) == 0 {
		return true
	}
	return false
}

func (t *Trie) Value() string {
	if t == nil {
		panic("空指针异常")
	}
	return t.value
}

// 插入元素
func (t *Trie) Insert(s Value) {
	if emptyValue(s) {
		return
	}
	root := t
	for k, str := range s {
		node := t.insertStr(root, str, len(s) == k+1)
		root = node
	}
}

// 查找元素 返回最后一个节点
func (t *Trie) Search(s Value) *Trie {
	if emptyValue(s) {
		return nil
	}
	root := t
	var res *Trie
	for _, str := range s {
		if root.children == nil {
			return nil
		}
		for _, child := range root.children {
			if child.value == str {
				root = child
				res = child
				break
			}
			return nil
		}
	}
	if res.isEnd {
		return res
	}
	return nil
}

// 插入子节点
// node 为父节点
func (t *Trie) insertStr(node *Trie, s string, isEnd bool) *Trie {
	if s == "" {
		panic("节点元素不能为空")
	}

	for _, child := range node.children {
		if child.value == s {
			if child.isEnd == false && child.isEnd != isEnd {
				// 存在以当前节点为结束的元素
				child.isEnd = true
			}
			return child
		}
	}
	child := &Trie{
		value: s,
		isEnd: isEnd,
	}
	node.children = append(node.children, child)
	return child
}
