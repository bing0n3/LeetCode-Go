// Package trie provides ...
package trie

type MapSum struct {
	Next map[rune]*MapSum
	Val  int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{Next: make(map[rune]*MapSum), Val: 0}
}

func (this *MapSum) Insert(key string, val int) {
	parent := this
	for _, r := range key {
		if _, ok := parent.Next[r]; !ok {
			parent.Next[r] = &MapSum{Next: make(map[rune]*MapSum), Val: 0}
		}
		parent = parent.Next[r]
	}
	parent.Val = val
}

func (this *MapSum) Sum(prefix string) int {
	parent := this
	for _, r := range prefix {
		if _, ok := parent.Next[r]; !ok {
			return 0
		}
		parent = parent.Next[r]
	}
	// sum from this
	sum := sumChild(parent)
	return sum
}

func sumChild(this *MapSum) int {
	sum := 0
	for _, v := range this.Next {
		sum += sumChild(v)
	}
	return sum + this.Val
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
