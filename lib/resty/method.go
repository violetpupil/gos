package resty

// Execute 使用指定方法请求 http.MethodGet
func Execute(method, url string, hook ReqHook) (*Response, error) {
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
