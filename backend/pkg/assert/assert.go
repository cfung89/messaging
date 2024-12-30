package assert

import (
	"fmt"
	"log/slog"
	"os"
)

type EqualIn struct {
	A   any
	B   any
	Err string
}

func Equal(input *EqualIn) {
	if input.A == nil || input.B == nil {
		slog.Error("Error: nil value in assert input")
		os.Exit(1)
	}
	if input.A != input.B {
		slog.Error(fmt.Sprintf("Mismatch: %s != %s", input.A, input.B), "error", input.Err)
		os.Exit(1)
	}
}

type LtIn struct {
	A   int64
	B   int64
	Err string
}

// date A < date B
func LessThan(input *LtIn) {
	if input.A > input.B {
		slog.Error(fmt.Sprintf("%d is not less than %d", input.A, input.B), "error", input.Err)
		os.Exit(1)
	}
}

func NotNil(val error) {
	if val != nil {
		slog.Error(val.Error())
		os.Exit(1)
	}
}
