package zero

import (
	"fmt"
	"os"
	"testing"
)

func TestMustLoad(t *testing.T) {
	os.Setenv("RUN_MODE", "dev")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("REDIS_DB", "0")
	os.Setenv("BOOL", "true")
	var c Config
	MustLoad("./config.yaml", &c)
	fmt.Printf("%+v\n", c)
}
