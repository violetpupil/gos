package resty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type Response = resty.Response

// Res http响应
type Res struct {
	RestyRequest  *resty.Request
	RestyResponse *resty.Response
	HttpResponse  *http.Response

	Url        string         // 请求链接
	StatusCode int            // 状态码
	Status     string         // 状态字符串
	Proto      string         // 协议版本
	IsSuccess  bool           // 状态码是200~299
	IsError    bool           // 状态码大于等于400
	Header     http.Header    // 响应头
	Cookies    []*http.Cookie // 响应cookie
	Size       int64          // 响应体字节数
	Body       []byte         // 响应体字节
	String     string         // 响应体字符串，去掉首尾空白
	// 响应体自动解码，类型断言用指针
	Result any // 响应码200~299，与Request.SetResult配合使用
	Error  any // 响应码>399，与Request.SetError配合使用

	IsJson bool // 响应体是否为json编码

	// 统计信息
	// TotalTime ≈ ConnTime + ServerTime + ResponseTime
	// ConnTime ≈ DNSLookup + TCPConnTime + TLSHandshake
	TotalTime  time.Duration // 请求响应总时间
	SentAt     time.Time     // 请求发送时间
	ReceivedAt time.Time     // 响应接收时间

	// 请求追踪信息，EnableTrace后才会有值
	resty.TraceInfo
	RemoteAddr     string        // 目标网络地址 ip:host
	DNSLookup      time.Duration // dns查找时间
	ConnTime       time.Duration // 总连接时间
	TCPConnTime    time.Duration // tcp连接时间
	TLSHandshake   time.Duration // tls握手时间
	ServerTime     time.Duration // 服务处理耗时
	ResponseTime   time.Duration // 响应持续耗时
	IsConnReused   bool          // tcp连接是否之前被使用过
	IsConnWasIdle  bool          // tcp连接是否从空闲池获取
	ConnIdleTime   time.Duration // 如果tcp连接从空闲池获取，空闲时间
	RequestAttempt int           // 请求尝试次数
}

// ToResponse 将 resty.Response 转为 Response
// 这个函数是为了说明两个结构体的关系
func ToResponse(src *resty.Response) *Res {
	dst := new(Res)
	dst.RestyRequest = src.Request
	dst.RestyResponse = src
	dst.HttpResponse = src.RawResponse

	dst.Url = src.Request.URL
	dst.StatusCode = src.StatusCode()
	dst.Status = src.Status()
	dst.Proto = src.Proto()
	dst.IsSuccess = src.IsSuccess()
	dst.IsError = src.IsError()
	dst.Header = src.Header()
	dst.Cookies = src.Cookies()
	dst.Size = src.Size()
	dst.Body = src.Body()
	dst.String = src.String()
	dst.Result = src.Result()
	dst.Error = src.Error()

	dst.IsJson = json.Valid(dst.Body)

	dst.TotalTime = src.Time()
	dst.SentAt = src.Request.Time
	dst.ReceivedAt = src.ReceivedAt()

	ti := src.Request.TraceInfo()
	dst.TraceInfo = ti
	if ti.RemoteAddr != nil {
		dst.RemoteAddr = ti.RemoteAddr.String()
	}
	dst.DNSLookup = ti.DNSLookup
	dst.ConnTime = ti.ConnTime
	dst.TCPConnTime = ti.TCPConnTime
	dst.TLSHandshake = ti.TLSHandshake
	dst.ServerTime = ti.ServerTime
	dst.ResponseTime = ti.ResponseTime
	dst.IsConnReused = ti.IsConnReused
	dst.IsConnWasIdle = ti.IsConnWasIdle
	dst.ConnIdleTime = ti.ConnIdleTime
	dst.RequestAttempt = ti.RequestAttempt
	return dst
}

// ToError 请求失败时，返回错误
func (r *Res) ToError() error {
	return fmt.Errorf("response fail: %s %s", r.Status, r.String)
}

// ToError 请求失败时，返回错误
func ToError(r *resty.Response) error {
	return fmt.Errorf("response fail: %s %s", r.Status(), r.String())
}
