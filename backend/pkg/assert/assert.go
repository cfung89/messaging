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
	if input.A != input.B {
		slog.Error(fmt.Sprintf("Mismatch: %s != %s", input.A, input.B), "error", input.Err)
	}
	os.Exit(1)
}
