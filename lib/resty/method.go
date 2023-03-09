// 非200响应码会正常返回响应，不会返回错误
package resty

import "os"

// Execute 使用指定方法请求 http.MethodGet
func Execute(method, url string, hook ReqHook) (*Response, error) {
	if client == nil {
		Init()
	}
	req := client.R()
	if hook != nil {
		req = hook(req)
	}

	res, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}
	return ToResponse(res), nil
}

func Get(url string, hook ReqHook) (*Response, error) {
	if client == nil {
		Init()
	}
	req := client.R()
	if hook != nil {
		req = hook(req)
	}

	res, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	return ToResponse(res), nil
}

func Post(url string, hook ReqHook) (*Response, error) {
	if client == nil {
		Init()
	}
	req := client.R()
	if hook != nil {
		req = hook(req)
	}

	res, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return ToResponse(res), nil
}

// PostFile 上传文件，直接读取文件字节
// Content-Type 根据内容检测
func PostFile(url string, hook ReqHook, name string) (*Response, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	if client == nil {
		Init()
	}
	req := client.R().SetBody(bytes)
	if hook != nil {
		req = hook(req)
	}

	res, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	return ToResponse(res), nil
}
