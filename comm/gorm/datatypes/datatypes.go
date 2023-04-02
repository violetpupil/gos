package datatypes

import "gorm.io/datatypes"

// ToJson 将json字符串转为datatypes.JSON
func ToJson(s string) datatypes.JSON {
	return datatypes.JSON([]byte(s))
}
