package searcher

// Engine doc ...
type Engine struct {
	searchers map[string]Searcher
}

// GetSearchersRegistryNames ...
func (e Engine) GetSearchersRegistryNames() []string {
	var searchersNames []string
	for searcherName := range e.searchers {
		searchersNames = append(searchersNames, searcherName)
	}
	return searchersNames
}
