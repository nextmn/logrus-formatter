// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package httplog

import (
	"log/slog"
	"net/http"
)

// responseWriter is a wrapper for [http.ResponseWriter] storing the statusCode for logging purpose.
type responseWriter struct {
	w          http.ResponseWriter
	statusCode int
}

// WriteHeader implements the [http.ResponseWriter] interface and store the statusCode.
func (r *responseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.w.WriteHeader(statusCode)
}

// Header implements the [http.ResponseWriter] interface.
func (w *responseWriter) Header() http.Header {
	return w.w.Header()
}

// Write implements the [http.ResponseWriter] interface.
func (w *responseWriter) Write(data []byte) (int, error) {
	return w.w.Write(data)
}

// responseWriterFlusher is a wrapper for [http.ResponseWriterFlusher] storing the statuscode for logging purpose.
type responseWriterFlusher struct {
	*responseWriter
	http.Flusher
}

// Flush implements the [http.ResponseWriterFlusher] interface.
func (w *responseWriterFlusher) Flush() {
	w.Flusher.Flush()
}

// newResponseWriter returns an [http.ResponseWriter] or a [http.ResponseWriterFlusher] depending on the input parameter.
func newResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	writer := &responseWriter{
		w: w,
	}

	if flusher, ok := w.(http.Flusher); ok {
		return &responseWriterFlusher{
			responseWriter: writer,
			Flusher:        flusher,
		}
	}
	return writer
}

// NewRequestLoggerMiddlware creates an [http.Handler] which logs requests.
func NewRequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := newResponseWriter(w)
		defer func() {
			var statusCode int
			if wrapper, ok := rw.(*responseWriter); ok {
				statusCode = wrapper.statusCode
			} else if wrapper, ok := rw.(*responseWriterFlusher); ok {
				statusCode = wrapper.statusCode
			}
			slog.Info("HTTP Request",
				"ip-address", r.RemoteAddr,
				"method", r.Method,
				"uri", r.URL.Path,
				"status-code", statusCode,
			)
		}()
		next.ServeHTTP(rw, r)
	})
}
