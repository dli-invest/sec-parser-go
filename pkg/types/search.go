type AutoGenerated struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
	Query    Query  `json:"query"`
}
type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}
type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}
type Source struct {
	Entity      string `json:"entity"`
	EntityWords string `json:"entity_words"`
	Tickers     string `json:"tickers"`
	Rank        int    `json:"rank"`
}

type Hits struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source,omitempty"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}
type Entity struct {
	Query    string `json:"query"`
	Operator string `json:"operator"`
}
type Match struct {
	Entity Entity `json:"entity"`
}
type Tickers struct {
	Query    string `json:"query"`
	Operator string `json:"operator"`
	Boost    int    `json:"boost"`
}
type Match0 struct {
	Tickers Tickers `json:"tickers"`
}
type Exists struct {
	Field string `json:"field"`
	Boost int    `json:"boost"`
}
type Should struct {
	Match  Match  `json:"match,omitempty"`
	Match0 Match0 `json:"match,omitempty"`
	Exists Exists `json:"exists,omitempty"`
}
type ID struct {
	Query string `json:"query"`
}
type Match0 struct {
	ID ID `json:"_id"`
}
type Should struct {
	Match  Match  `json:"match,omitempty"`
	Match0 Match0 `json:"match,omitempty"`
}
type Bool struct {
	Should []Should `json:"should"`
}
type Must struct {
	Bool Bool `json:"bool"`
}
type Bool struct {
	Should []Should `json:"should"`
	Must   []Must   `json:"must"`
}
type Query struct {
	Bool Bool `json:"bool"`
}
type Query struct {
	Query Query `json:"query"`
	From  int   `json:"from"`
	Size  int   `json:"size"`
}