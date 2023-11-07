package common

import "github.com/sirupsen/logrus"

// IdReq 单个id请求
type IdReq struct {
	ID string `form:"id"`
}

func (r IdReq) Logger() *logrus.Entry {
	return logrus.WithField("id", r.ID)
}

// IdsReq 多个id请求
type IdsReq struct {
	IDs []string `json:"ids"`
}

func (r IdsReq) Logger() *logrus.Entry {
	return logrus.WithField("ids", r.IDs)
}
