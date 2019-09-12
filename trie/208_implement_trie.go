// Package ImplemetTrie provides ...
package ImplemetTr

type Trie struct {
	Children map[rune]*Trie
	IsLeaf   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Children: make(map[rune]*Trie), IsLeaf: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, c := range word {
		if child, ok := parent.Children[c]; ok {
			parent = child
		} else {
			newChild := &Trie{Children: make(map[rune]*Trie), IsLeaf: false}
			parent.Children[c] = newChild
			parent = newChild
		}
	}
	parent.IsLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, c := range word {
		if child, ok := parent.Children[c]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.IsLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for _, c := range prefix {
		if child, ok := parent.Children[c]; ok {
			parent = child
			continue
		}
		return false
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
