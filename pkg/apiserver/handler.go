package apiserver

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
)

func recoverHandler(panicReason interface{}, w http.ResponseWriter) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("recover from panic situation: - %v\r\n", panicReason))
	for i := 2; ; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
	}
	fmt.Print(buffer.String())

	headers := http.Header{}
	if contentType := w.Header().Get("Content-Type"); len(contentType) > 0 {
		headers.Set("Accept", contentType)
	}
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
