// 数据处理
// https://help.aliyun.com/document_detail/410762.html
package oss

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// VideoSnapshot 视频截帧参数值
// https://help.aliyun.com/document_detail/64555.html
// 使用视频截帧时，按视频截帧截取的图片数量计费。
// 截图宽高都指定为0，使用视频宽高，一个指定为0，则自动计算
// 如果视频可以公开读，截图地址是在视频地址后加 ?x-oss-process=参数值
func VideoSnapshot(width, height int) string {
	// 截图时间为第1秒，jpg格式
	s := "video/snapshot,t_1000,f_jpg,m_fast,w_%d,h_%d"
	s = fmt.Sprintf(s, width, height)
	return s
}

// VideoSnapshotSignURL 生成私有文件视频截帧URL
func VideoSnapshotSignURL(objectKey string, expiredInSec int64, width, height int) (string, error) {
	p := VideoSnapshot(width, height)
	s, err := SignURL(objectKey, expiredInSec, oss.Process(p))
	return s, err
}
