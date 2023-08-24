package otto

import (
	"github.com/robertkrimen/otto"
)

type (
	// Otto is the representation of the JavaScript runtime.
	Otto = otto.Otto
)

// Run 运行js代码
func Run(code string) (otto.Value, error) {
	vm := otto.New()
	value, err := vm.Run(code)
	return value, err
}
