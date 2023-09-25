package http

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func GzipMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gw := NewGzipResponseWriter(w)
			gw.Header().Set("Content-Encoding", "gzip")
			defer gw.Flush()
			next.ServeHTTP(gw, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type GzipResponseWriter struct {
	w  http.ResponseWriter
	gw *gzip.Writer
}

func NewGzipResponseWriter(w http.ResponseWriter) *GzipResponseWriter {
	return &GzipResponseWriter{
		w:  w,
		gw: gzip.NewWriter(w),
	}
}

func (gw *GzipResponseWriter) Header() http.Header {
	return gw.w.Header()
}

func (gw *GzipResponseWriter) Write(p []byte) (int, error) {
	return gw.gw.Write(p)
}

func (gw *GzipResponseWriter) WriteHeader(statusCode int) {
	gw.w.WriteHeader(statusCode)
}

func (gw *GzipResponseWriter) Flush() error {
	if err := gw.gw.Flush(); err != nil {
		return err
	}
	if err := gw.gw.Close(); err != nil {
		return err
	}
	return nil
}
