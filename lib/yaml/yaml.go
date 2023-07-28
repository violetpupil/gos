package yaml

import (
	"os"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func UnmarshalFile(f string, v any) error {
	bs, err := os.ReadFile(f)
	if err != nil {
		logrus.Error("read file error ", err)
		return err
	}
	err = yaml.Unmarshal(bs, v)
	return err
}
