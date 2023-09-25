package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type File struct {
	Name        string
	ContentType string
	Reader      io.Reader
}

func ServeFile(w http.ResponseWriter, file File) error {

	gw := NewGzipResponseWriter(w)
	defer gw.Flush()

	gw.Header().Set("Content-Type", file.ContentType)
	gw.Header().Set("Content-Encoding", "gzip")

	// Set the content disposition header
	gw.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", file.Name))
	gw.Header().Set("Cache-Control", "no-store")

	gw.WriteHeader(http.StatusOK)

	_, err := io.Copy(gw, file.Reader)
	if err != nil {
		return errors.New("failed to serve file")
	}

	return nil
}
