package autocomplete

import (
	trie "github.com/Ben352/goTrie"
)

var t *trie.Trie

func init() {
	t = trie.CreateNewTrie()
}

func LoadTrie(fileName string) {
	t.LoadTrie(fileName)
}

func SaveTrie(fileName string) {
	t.SerializeTrie(fileName)
}

func AddWord(word string) {
	t.InsertWord(word)
}

func GetSuggestions(prefix string) []string {
	suggestions := t.GetWords(prefix, 99)
	var results = []string{}
	for _, element := range suggestions {
		results = append(results, element.Suggestion)
	}
	return results
}
