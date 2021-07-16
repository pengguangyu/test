package main

// radix树实现-go


import "fmt"

//radix Node
type Node struct {
	Nodes []*Node
	Index string //节点数据
	Path  string //节点path
}

var root = NewRootNode("")

func NewRootNode(s string) *Node {
	return &Node{
		Nodes: make([]*Node, 0),
		Index: "",
		Path:  "",
	}
}

func findInertPreNode(p string, pIndex int, node *Node) (n *Node, PreFix string, index string, isExists bool) {
	_childs := node.Nodes
	if len(_childs) == 0 {
		return node, findComonPrefix(p, node.Path), string(p[pIndex]), false
	}

	for _, v := range _childs {
		if v.Index == string(p[pIndex]) {
			if pIndex == len(p) {
				return &Node{}, "", "", true
			}
			return findInertPreNode(p, pIndex+1, v)
		}
	}

	return node, findComonPrefix(p, node.Path), string(p[pIndex]), false
}

func (s *Node) Insert(p string) {
	_node, prefix, index, isExists := findInertPreNode(p, 0, root)
	if isExists {
		//整个字符串都找到了 还往下走啥
		return
	}
	fmt.Println(prefix)
	_node.Nodes = append(_node.Nodes, &Node{
		Nodes: make([]*Node, 0),
		Index: index,
		Path:  p,
	})
}

//查找最大前缀
func findComonPrefix(s string, s1 string) string {
	_temp := make([]byte, 0)
	s1_len := len(s1)
	s_len := len(s)
	_len1 := s1_len
	if s1_len > s_len {
		_len1 = s_len
	}
	for i := 0; i < _len1; i++ {
		if s[i] == s1[i] {
			_temp = append(_temp, s[i])
		} else {
			break
		}
	}
	return string(_temp)
}

func main() {
	fmt.Println("[radix树实现-go]")
	root.Insert("romane")
	root.Insert("romans")
	root.Insert("abc")
	root.Insert("abcd")
	fmt.Println(root.Nodes[0].Nodes[0])
}

// 原文链接：https://blog.csdn.net/lzx2766478/article/details/115185643 radix树实现-go
// https://blog.csdn.net/phpduang/article/details/108516475 逐步解释
