package common

// IdReq 单个id请求
type IdReq struct {
	ID string `form:"id"`
}

// IdsReq 多个id请求
type IdsReq struct {
	IDs []string `json:"ids"`
}
