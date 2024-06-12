package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

type Node struct {
	children   map[rune]*Node
	numSubstr  int
	lastUpdIdx int
}

type Trie struct {
	root *Node
}

func NewNode(idx int, num int) *Node {
	return &Node{make(map[rune]*Node), num, idx}
}

func NewTrie() *Trie {
	return &Trie{NewNode(math.MinInt32, math.MinInt32)}
}

func (t *Trie) addSuffix(suffix string, marker rune, idx int, isWholeWord bool) {
	switch marker {
	case 'b':
		t.addBlueWord(suffix, idx, isWholeWord)
	case 'r':
		t.addRedWord(suffix, idx, isWholeWord)
	case 'w':
		t.addWhiteWord(suffix)
	case 'x':
		t.addBlackWord(suffix, idx)
	}
}

func (t *Trie) addBlueWord(suffix string, idx int, isWholeWord bool) {
	curNode := t.root
	for _, letter := range suffix {
		if _, ok := curNode.children[letter]; !ok {
			curNode.children[letter] = NewNode(idx, 1)
		}
		curNode = curNode.children[letter]
		if curNode.lastUpdIdx != idx && curNode.numSubstr != math.MinInt32 {
			curNode.numSubstr++
			curNode.lastUpdIdx = idx
		}
	}
	if isWholeWord {
		curNode.numSubstr = math.MinInt32
	}
}

func (t *Trie) addRedWord(suffix string, idx int, isWholeWord bool) {
	curNode := t.root
	wentWholeWord := true
	for _, letter := range suffix {
		if _, ok := curNode.children[letter]; !ok {
			wentWholeWord = false
			break
		}
		curNode = curNode.children[letter]
		if curNode.lastUpdIdx != idx && curNode.numSubstr != math.MinInt32 {
			curNode.numSubstr--
			curNode.lastUpdIdx = idx
		}
	}
	if isWholeWord && wentWholeWord {
		curNode.numSubstr = math.MinInt32
	}
}

func (t *Trie) addWhiteWord(word string) {
	curNode := t.root
	wentWholeWord := true
	for _, letter := range word {
		if _, ok := curNode.children[letter]; !ok {
			wentWholeWord = false
			break
		}
		curNode = curNode.children[letter]
	}
	if wentWholeWord {
		curNode.numSubstr = math.MinInt32
	}
}

func (t *Trie) addBlackWord(suffix string, idx int) {
	curNode := t.root
	for _, letter := range suffix {
		if _, ok := curNode.children[letter]; !ok {
			break
		}
		curNode = curNode.children[letter]
		if curNode.lastUpdIdx != idx {
			curNode.numSubstr = math.MinInt32
			curNode.lastUpdIdx = idx
		}
	}
}

func (t *Trie) findMoveWord(node *Node, path string) (string, int) {
	if node == nil {
		return "", math.MinInt32
	}
	word := path
	maxNumSubstr := node.numSubstr
	for letter, child := range node.children {
		childWord, childNum := t.findMoveWord(child, path+string(letter))
		if childNum > maxNumSubstr {
			word = childWord
			maxNumSubstr = childNum
		}
	}
	return word, maxNumSubstr
}

func genRandStr(length int) string {
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = 'a' + byte(time.Now().UnixNano()%26)
		time.Sleep(1)
	}
	return string(randStr)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var numTests int
	fmt.Fscan(in, &numTests)

	for i := 0; i < numTests; i++ {
		trie := NewTrie()

		var numWords, numBlueWords, numRedWords, idxBlackWord int
		fmt.Fscan(in, &numWords, &numBlueWords, &numRedWords, &idxBlackWord)

		for idx := 1; idx <= numWords; idx++ {
			var word string
			fmt.Fscan(in, &word)

			var wordMarker rune
			if idx <= numBlueWords {
				wordMarker = 'b'
			} else if numBlueWords < idx && idx <= numBlueWords+numRedWords {
				wordMarker = 'r'
			} else if numBlueWords+numRedWords < idx && idx != idxBlackWord {
				wordMarker = 'w'
			} else if numBlueWords+numRedWords < idx && idx == idxBlackWord {
				wordMarker = 'x'
			}

			trie.addSuffix(word, wordMarker, idx, true)

			if wordMarker != 'w' {
				for i := 1; i < len(word); i++ {
					trie.addSuffix(word[i:], wordMarker, idx, false)
				}
			}
		}

		moveWord, moveNum := trie.findMoveWord(trie.root, "")

		if moveNum < 0 {
			fmt.Fprintln(out, genRandStr(10), "0")
		} else {
			fmt.Fprintln(out, moveWord, moveNum)
		}
	}
}
