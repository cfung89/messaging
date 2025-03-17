package assert

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

func Assert(truth bool, msg string) {
	if !truth {
		pc, _, _, _ := runtime.Caller(1)
		slog.Error(fmt.Sprintf("(%s): %s", runtime.FuncForPC(pc).Name(), msg))
		os.Exit(1)
	}
}

func Error(err error) {
	if err != nil {
		pc, _, _, _ := runtime.Caller(1)
		slog.Error(fmt.Sprintf("(%s): %s", runtime.FuncForPC(pc).Name(), err.Error()))
		os.Exit(1)
	}
}
