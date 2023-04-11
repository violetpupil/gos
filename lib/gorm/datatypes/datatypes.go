package datatypes

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
)

// Marshal json编码
func Marshal(v any) (datatypes.JSON, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		logrus.Errorln("json marshal error", err)
		return nil, err
	}
	return datatypes.JSON(bs), nil
}
