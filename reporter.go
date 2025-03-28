package reporter

import (
	"log/slog"
	"net/http"
	"runtime"
)

type RuntimeError struct{}

func (r RuntimeError) Reporter(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown file"
		line = 0
	}

	funcName := "unknown function"
	if function := runtime.FuncForPC(pc); function != nil {
		funcName = function.Name()
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError),
		(http.StatusInternalServerError))

	slog.Error("Internal server error",
		"error", err,
		"funcName", funcName,
		"file", file,
		"line", line,
	)
}
