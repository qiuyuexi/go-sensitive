package ahocorasick

import "fmt"

/**
ac自动机节点
 */
type trieNode struct {
	value     string
	endString string
	fail      *trieNode
	next      map[string]*trieNode
}

type Ahocorasick struct {
	Tree map[int]*trieNode
}

func (trieTree *trieNode) insert(value string) {
	str := []rune(value)
	node := trieTree
	for _, v := range str {
		index := string(v)
		if node.next[index] != nil {
			node = node.next[index]
		} else {
			node.next[index] = new(trieNode)
			node.next[index].next = make(map[string]*trieNode)
			node.next[index].value = index
			node.next[index].endString = ""
			node = node.next[index]
		}
	}
	node.endString = value
}

func (root *trieNode) Search(value string) (hitString []string) {
	str := []rune(value)
	node := root
	hitString = []string{}
	hit := 0
	defer func() {
		hitString = []string{}
	}()
	for _, v := range str {

		index := string(v)

		//如果不存在，则一直回溯fial，直到回到root
		for {
			_, ok := node.next[index];
			if !ok && node != root {
				node = node.fail
			}
			if ok || node == root {
				break
			}

		}

		if _, ok := node.next[index]; !ok {
			node = root
		} else {
			node = node.next[index]
		}
		curNode := node
		for curNode != nil && curNode != root {
			if curNode.endString != "" {
				hitString = append(hitString, curNode.endString)
				hit++
			}
			curNode = curNode.fail
		}
	}
	return hitString
}

/**
构建失败指针
 */
func (root *trieNode) buildFail() {
	queue := make(map[int]*trieNode)
	tail := 0
	head := 0
	queue[head] = root

	for head <= tail {
		node := queue[head]
		head++

		//遍历当前节点的子节点
		for _, v := range node.next {

			//root的 第一层子节点 fial 全部指向root
			if node == root {
				v.fail = root
			} else {

				//root的 第二层及以下子节点，判断当前节点的fail节点 子节点是否存在该子节点
				// 例如 s1:abc,s2:bc
				// 那么 s1:a.fail -> root  s2.b.fial -> root
				// s1.a.b.fail ->  s1.a.fial=> root , root.child 包含b -> s1.a.b.fail ->s2.b
				// 根据这个可以递推
				pp := node.fail
				for pp != nil {
					if _, ok := pp.next[v.value]; ok {
						v.fail = pp.next[v.value]
						break
					}
					pp = pp.fail
				}
				if pp == nil {
					v.fail = root
				}
			}
			tail++
			queue[tail] = v
		}
	}
}

/**

 */
func (root *trieNode) debug() {
	queue := make(map[int]*trieNode)
	tail := 0
	head := 0
	queue[head] = root
	fmt.Println(&root)
	for head <= tail {
		node := queue[head]
		head++

		//遍历当前节点的子节点
		for _, v := range node.next {

			//root的 第一层子节点 fial 全部指向root
			fmt.Println(v.value, v.fail)
			tail++
			queue[tail] = v
		}
	}
}

func GetAhocorasick() *Ahocorasick {
	ahocorasick := new(Ahocorasick)
	ahocorasick.Tree = make(map[int]*trieNode)
	return ahocorasick
}

/**
获取ac自动机，
 */
func GetTire(words []string) (root *trieNode) {

	root = new(trieNode)
	root.fail = nil
	root.endString = ""
	root.value = ""
	root.next = make(map[string]*trieNode)

	for _, v := range words {
		root.insert(v)
	}
	root.buildFail()
	return root
}
