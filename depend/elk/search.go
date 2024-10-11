package elk

type SearchResult struct {
	Hits struct {
		Total struct {
			Value float64 `json:"value"` // 总数
		} `json:"total"`
		Hits []struct {
			Source any    `json:"_source"` // 文档对象
			ID     string `json:"_id"`     // 文档id
			Sort   any    `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
	Took float64 `json:"took"` // 请求耗时，毫秒
}

type SearchParams struct {
	Size  int `json:"size"`
	Query struct {
		Match map[string]any `json:"match"`
	} `json:"query"`
	Sort        []any `json:"sort"`
	SearchAfter any   `json:"search_after,omitempty"`
}
