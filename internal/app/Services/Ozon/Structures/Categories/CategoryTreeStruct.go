package Categories

type AttrOneTreeBody struct {
	CategoryID int    `json:"category_id"`
	Language   string `json:"language"`
}

type AttrOneTreeResult struct {
	Result struct {
		CategoryID int           `json:"category_id"`
		Title      string        `json:"title"`
		Children   []interface{} `json:"children"`
	} `json:"result"`
}
