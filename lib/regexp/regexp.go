// More precisely, it is the syntax accepted by RE2 and described at https://golang.org/s/re2syntax, except for \C.
// go doc regexp/syntax
package regexp

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

// FindAllString 查找所有匹配的字符串
func FindAllString(expr, s string) ([]string, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		logrus.Errorln("regexp compile error", err)
		return nil, err
	}
	res := r.FindAllString(s, -1)
	return res, nil
}
