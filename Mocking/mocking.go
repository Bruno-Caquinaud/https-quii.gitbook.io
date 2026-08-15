package mocking

import (
	"fmt"
	"io"
	"iter"
)

const (
	finalWord = "Go!"
)

type Sleeper interface {
	Sleep()
}

func CountDown(output io.Writer, sleeper Sleeper, start int) (caracterNumber int, err error) {
	for i := range countDownFrom(start) {
		localCaracterNumber, err := fmt.Fprintf(output, "%d\n", i)

		if err != nil {
			return caracterNumber, err
		}

		sleeper.Sleep()
		caracterNumber += localCaracterNumber
	}

	localCaracterNumber, err := fmt.Fprintf(output, "%s\n", finalWord)
	caracterNumber += localCaracterNumber

	return
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
