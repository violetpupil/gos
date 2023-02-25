package resty

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Response http响应
type Response struct {
	RestyRequest  *resty.Request
	RestyResponse *resty.Response
	HttpResponse  *http.Response

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

	// 统计信息
	// TotalTime ≈ ConnTime + ServerTime + ResponseTime
	// ConnTime ≈ DNSLookup + TCPConnTime + TLSHandshake
	TotalTime  time.Duration // 请求响应总时间
	SentAt     time.Time     // 请求发送时间
	ReceivedAt time.Time     // 响应接收时间

	resty.TraceInfo               // 请求追踪信息
	RemoteAddr      string        // 目标网络地址 ip:host
	DNSLookup       time.Duration // dns查找时间
	ConnTime        time.Duration // 总连接时间
	TCPConnTime     time.Duration // tcp连接时间
	TLSHandshake    time.Duration // tls握手时间
	ServerTime      time.Duration // 服务处理耗时
	ResponseTime    time.Duration // 响应持续耗时
	IsConnReused    bool          // tcp连接是否之前被使用过
	IsConnWasIdle   bool          // tcp连接是否从空闲池获取
	ConnIdleTime    time.Duration // 如果tcp连接从空闲池获取，空闲时间
	RequestAttempt  int           // 请求尝试次数
}

// ToResponse 将 resty.Response 转为 Response
func ToResponse(src *resty.Response) *Response {
	dst := new(Response)
	dst.RestyRequest = src.Request
	dst.RestyResponse = src
	dst.HttpResponse = src.RawResponse

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
