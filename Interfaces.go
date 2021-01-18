package searcher

// Searcher interface
type Searcher interface {
	Search(Filter) ([]Item, error)
}
