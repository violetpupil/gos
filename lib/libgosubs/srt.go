// 读写srt字幕文件
// https://docs.fileformat.com/zh/video/srt/
package libGoSubs

import "github.com/wargarblgarbl/libgosubs/srt"

// Subtitle 单条字幕
// Id 字幕id，从1开始
// Start 开始时间 00:00:00,000
// End 结束时间 00:00:00,000
// Line 多行字幕
type Subtitle = srt.Subtitle

// WriteSrt 写srt字幕文件
func WriteSrt(subs []srt.Subtitle, name string) error {
	var sub srt.SubRip
	sub.Subtitle.Content = subs
	err := srt.WriteSrt(&sub, name)
	return err
}
