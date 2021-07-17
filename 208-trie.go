
type Node struct {
	bEnd bool
	iPrefixCount int
	AlphaMap map[int] *Node
}

type Trie struct {
	pRoot *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	pTrie := new (Trie)

	pNode := new(Node)
	pNode.AlphaMap = make(map[int] *Node)

	pTrie.pRoot = pNode
	return *pTrie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	pCurrentNode := this.pRoot
	pCurrentNode.iPrefixCount++

	iLen := len(word)
	for i, c := range word {
		pValue,_ok := pCurrentNode.AlphaMap[int(c)]
		if _ok{
			pCurrentNode = pValue
		}else{
			pNewNode := new(Node)
			pNewNode.iPrefixCount = 1
			pNewNode.AlphaMap = make(map[int] *Node)
			pCurrentNode.AlphaMap[int(c)] = pNewNode
			pCurrentNode = pNewNode
		}

		pCurrentNode.iPrefixCount++
		if i == iLen - 1{
			pCurrentNode.bEnd = true
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	bIn := true
	pCurrentNode := this.pRoot
	for i, c := range word {
		pValue,_ok := pCurrentNode.AlphaMap[int(c)]
		if !_ok{
			bIn = false
			break
		}else{
			pCurrentNode = pValue
		}

		if len(word)-1 == i{
			if !pCurrentNode.bEnd{
				bIn = false
			}
		}
	}
	return bIn
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	bHave := true
	pCurrentNode := this.pRoot
	for i, c := range prefix {
		pValue,_ok := pCurrentNode.AlphaMap[int(c)]
		if !_ok{
			bHave = false
			break
		}else{
			pCurrentNode = pValue
		}

		if len(prefix)-1 == i{
			if pCurrentNode.iPrefixCount <= 0{
				bHave = false
			}
		}
	}
	return bHave
}
