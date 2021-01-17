package searcher

// Searcher interface
type Searcher interface {
	Search(Filter) ([]Item, error)
}

// Engine doc ...
type Engine struct {
	searchers map[string]Searcher
}
