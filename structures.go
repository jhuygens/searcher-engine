package searcher

// Item doc ...
type Item struct {
	Type    string `json:"type"`
	Library string `json:"library"`
	Name    string `json:"name"`
	Artwork string `json:"artwork"`
	Info    Info   `json:"info"`
}

// Info item struct
type Info struct {
	PreviewURL  string   `json:"preview_url"`
	Title       string   `json:"title"`
	Collection  string   `json:"collection"`
	Artist      string   `json:"artist"`
	Languages   []string `json:"languages"`
	RatingAvg   float64  `json:"rating_avg"`
	Genres      []string `json:"genres"`
	Description string   `json:"description"`
	MoreInfo    string   `json:"more_info"`
	ReleaseDate string   `json:"release_date"`
	Country     string   `json:"country"`
	Price       float64  `json:"price"`
	RentalPrice float64  `json:"rental_price"`
	Currency    string   `json:"currency"`
	URL         string   `json:"url"`
}

// Filter search
type Filter struct {
	Name    []FieldValue `json:"name,omitempty"`
	Artist  []FieldValue `json:"artist,omitempty"`
	Album   []FieldValue `json:"album,omitempty"`
	Genre   []FieldValue `json:"genre,omitempty"`
	Year    []FieldValue `json:"year,omitempty"`
	Country []FieldValue `json:"country,omitempty"`
	Type    string       `json:"type,omitempty"`
	Library string       `json:"library,omitempty"`
}

// FieldValue doc ...
type FieldValue struct {
	Value   string `json:"value,omitempty"`
	Exclude bool   `json:"exclude,omitempty"`
}
