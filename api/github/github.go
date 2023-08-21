// https://docs.github.com/en
package github

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// SearchRepositories Find repositories via various criteria.
func SearchRepositories(lang string) error {
	res, err := resty.New().R().
		SetQueryParam("q", "language:"+lang).
		SetQueryParam("sort", "stars").
		Get("https://api.github.com/search/repositories")
	if err != nil {
		logrus.Errorln("search repositories error", err)
		return err
	}

	fmt.Println(res.String())
	return nil
}
