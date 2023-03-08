package resty

import (
	"github.com/go-resty/resty/v2"
)

// client http客户端
// 必须先调用 Init 初始化
var client *resty.Client

// Init 初始化客户端
func Init() {
	client = resty.New()
}
