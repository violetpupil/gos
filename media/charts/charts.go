package charts

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/sirupsen/logrus"
)

// BarData 转成条形图数据
func BarData(values []any) []opts.BarData {
	data := make([]opts.BarData, 0)
	for _, v := range values {
		d := opts.BarData{Value: v}
		data = append(data, d)
	}
	return data
}

// NewBar 生成条形图 x为x轴
// series 为源数据，键为数据名
func NewBar(title string, x any, series map[string][]any) error {
	// 设置标题等
	bar := charts.NewBar()
	titleOpt := opts.Title{Title: title}
	globalOpts := charts.WithTitleOpts(titleOpt)
	bar.SetGlobalOptions(globalOpts)

	// 设置数据
	bar.SetXAxis(x)
	for name, values := range series {
		bar.AddSeries(name, BarData(values))
	}

	// 生成html文件
	f, err := os.Create("tmp.html")
	if err != nil {
		logrus.Errorln("create file error", err)
		return err
	}
	err = bar.Render(f)
	return err
}
