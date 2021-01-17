package searcher

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/jhuygens/cache"
)

var engine = Engine{
	searchers: make(map[string]Searcher),
}

// New doc ...
func New(searchers map[string]Searcher) *Engine {
	return &Engine{
		searchers: searchers,
	}
}

// RegisterSearcher doc ...
func RegisterSearcher(library string, searcher Searcher) {
	engine.searchers[library] = searcher
}

// ValidateRegisterImplement doc ...
func ValidateRegisterImplement() error {
	if engine.searchers != nil {
		return nil
	}
	return fmt.Errorf("The implementation of the 'Searcher' interface has not been registered")
}

// Search returns a key cache search
func Search(filter Filter) (string, error) {
	var items []Item
	if filter.Library == "" {
		for _, searcher := range engine.searchers {
			result, err := searcher.Search(filter)
			if err != nil {
				return "", err
			}
			items = append(
				items,
				result...,
			)
		}
	} else {
		result, err := engine.searchers[filter.Library].Search(filter)
		if err != nil {
			return "", err
		}
		items = append(
			items,
			result...,
		)
	}
	sort.Sort(ByName(items))
	keySearch, err := GenerateKeySearch(filter)
	if err != nil {
		return "", err
	}
	err = cache.Set(keySearch, items)
	if err != nil {
		return "", err
	}
	return keySearch, nil
}

// GenerateKeySearch ...
func GenerateKeySearch(filter Filter) (string, error) {
	buf, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	h.Write(buf)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}