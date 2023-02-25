package resty

import "github.com/go-resty/resty/v2"

// client http客户端
// 必须先调用 Init 初始化
var client *resty.Client

// Init 初始化客户端
// trace 表示是否追踪请求信息
func Init(trace bool) {
	client = resty.New()
	if trace {
		client = client.EnableTrace()
	}
}

// Get get请求
// trace 表示是否追踪请求信息
func Get(url string, trace bool) (*Response, error) {
	req := client.R()
	if trace {
		req = req.EnableTrace()
	}

	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	return ToResponse(res), nil
}
