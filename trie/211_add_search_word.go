package tireWithDot

type WordDictionary struct {
	Children [26]*WordDictionary
	IsLeaf   bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{Children: [26]*WordDictionary{}, IsLeaf: false}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	parent := this
	for i, _ := range word {
		if parent.Children[word[i]-'a'] != nil {
			parent = parent.Children[word[i]-'a']
		} else {
			newChild := &WordDictionary{Children: [26]*WordDictionary{}, IsLeaf: false}
			parent.Children[word[i]-'a'] = newChild
			parent = parent.Children[word[i]-'a']
		}
	}
	parent.IsLeaf = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

	return match(word, 0, this)
}

func match(word string, k int, root *WordDictionary) bool {
	if k == len(word) {
		return root.IsLeaf
	}
	if word[k] != '.' {
		return root.Children[word[k]-'a'] != nil && match(word, k+1, root.Children[word[k]-'a'])
	} else {
		for i := 0; i < len(root.Children); i++ {
			if root.Children[i] != nil {
				if match(word, k+1, root.Children[i]) {
					return true
				}
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
