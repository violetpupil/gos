package job

import (
	"fmt"
)

// 交替打印数字和字母
// https://github.com/lifei6671/interview-go/blob/master/question/q001.md
func PrintNumberAlphabet() {
	letter, number := make(chan bool), make(chan bool)
	exit := make(chan bool)

	go func() {
		i := 1
		for range number {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- true
		}
	}()

	go func() {
		i := 'A'
		for range letter {
			if i >= 'Z' {
				exit <- true
				return
			}

			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			number <- true
		}
	}()
	number <- true
	<-exit
}
