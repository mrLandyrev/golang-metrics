package rest

import (
	"compress/gzip"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		end := time.Since(start)

		logger.Info(
			"Request info",
			zap.String("URI", context.Request.RequestURI),
			zap.String("Method", context.Request.Method),
			zap.Duration("Time", end),
		)

		logger.Info(
			"Response info",
			zap.Int("Status", context.Writer.Status()),
			zap.Int("Size", context.Writer.Size()),
		)
	}
}

type gzipWriter struct {
	gin.ResponseWriter
	writer io.Writer
}

func (w *gzipWriter) Write(data []byte) (int, error) {
	return w.writer.Write(data)
}

type gzipReader struct {
	io.ReadCloser
	reader io.ReadCloser
}

func (r *gzipReader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Accept-Encoding") == "gzip" {
			c.Header("Content-Encoding", "gzip")
			gz := gzip.NewWriter(c.Writer)
			c.Writer = &gzipWriter{ResponseWriter: c.Writer, writer: gz}
			defer func() {
				gz.Close()
				c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
			}()
		}

		if c.Request.Header.Get("Content-Encoding") == "gzip" {
			gz, _ := gzip.NewReader(c.Request.Body)
			c.Request.Header.Del("Content-Encoding")
			c.Request.Header.Del("Content-Length")
			c.Request.Body = &gzipReader{ReadCloser: c.Request.Body, reader: gz}
			defer func() {
				gz.Close()
			}()
		}

		c.Next()
	}
}
