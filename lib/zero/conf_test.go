package zero

import "testing"

func TestMustLoad(t *testing.T) {
	var c Config
	MustLoad("./config.yaml", &c)
}
