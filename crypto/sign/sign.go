// api签名 - 科大讯飞
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
package sign

import (
	"github.com/violetpupil/gos/std/hmac"
	"github.com/violetpupil/gos/std/md5"
)

// Sign 计算签名 HmacSha1(secret, md5(appid + ts))
func Sign(appid, ts, secret string) string {
	baseS := appid + ts
	baseS = md5.Sum(baseS)
	baseS = hmac.HmacSha1Base64(secret, baseS)
	return baseS
}
