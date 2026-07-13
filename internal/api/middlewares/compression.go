package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

// gunzip response writer will wrap http response writer to write gunzip responses
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func hasGzip(header string) bool {
	return strings.Contains(header, "gzip")
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts gzip encoding
		encodingHeader := r.Header.Get("Accept-Encoding")

		if !hasGzip(encodingHeader) {
			next.ServeHTTP(w, r)
			return
		}

		// First set the response header
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap response writer here
		w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}

		next.ServeHTTP(w, r)
		fmt.Println("Sent response from compression middleware")
	})
}
