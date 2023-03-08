// 科大讯飞api调用
// 必须先 Init 初始化
// 不同服务的密钥不同
// https://www.xfyun.cn/doc/
package xfyun

type xfyun struct {
	appid string
}

var Xfyun *xfyun

// Init 初始化讯飞api调用
func Init(appid string) {
	Xfyun = &xfyun{appid}
}
