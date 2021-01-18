package searcher

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/jgolang/config"
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
func RegisterSearcher(library string, searcher Searcher) error {
	if library == "" {
		return fmt.Errorf("No set a library name")
	}
	engine.searchers[library] = searcher
	return nil
}

// ValidateRegisterImplement doc ...
func ValidateRegisterImplement() error {
	if len(engine.searchers) > 0 {
		return nil
	}
	return fmt.Errorf("The implementation of the 'Searcher' interface has not been registered")
}

// Search returns a key cache search
func Search(filter Filter) (string, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return "", err
	}
	var items []Item
	if filter.Library == "all" {
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
		if engine.searchers[filter.Library] == nil {
			return "", fmt.Errorf("Error: the searcher %v is not regiter", filter.Library)
		}
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
	searchKey, err := GenerateSearchKey(filter)
	if err != nil {
		return "", err
	}
	err = setCacheItems(searchKey, items)
	if err != nil {
		return "", err
	}
	err = cache.Expire(searchKey, config.GetInt("cache.expire_time"))
	if err != nil {
		return "", err
	}
	return searchKey, nil
}

// ByName items order
type ByName []Item

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return strings.ToLower(a[i].Name) < strings.ToLower(a[j].Name) }

// GenerateSearchKey ...
func GenerateSearchKey(filter Filter) (string, error) {
	buf, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}
	h := sha256.New()
	h.Write(buf)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// GetSearchersRegistryNames ...
func GetSearchersRegistryNames() []string {
	return engine.GetSearchersRegistryNames()
}

func setCacheItems(searchKey string, items []Item) error {
	buff, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return cache.Set(searchKey, string(buff))
}
