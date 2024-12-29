package assert

import (
	"log/slog"
	"os"
)

type EqualIn struct {
	A   any
	B   any
	Err string
}

func Equal(input *EqualIn) {
	if input.A != input.B {
		slog.Error(input.Err)
	}
	os.Exit(1)
}
