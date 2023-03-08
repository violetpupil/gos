// api签名 - 科大讯飞
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
package sign

import (
	"github.com/violetpupil/components/std/base64"
	"github.com/violetpupil/components/std/hmac"
	"github.com/violetpupil/components/std/md5"
	"github.com/violetpupil/components/std/time"
)

// Sign 计算签名 HmacSha1(secret, md5(appid + ts))
func Sign(appid string, secret string) string {
	baseS := appid + time.Ts()
	baseS = md5.Sum(baseS)
	baseB := hmac.HmacSha1(secret, baseS)
	baseS = base64.Encode(baseB)
	return baseS
}
