package middlewares

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/pkg/logging"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *configs.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}
func structuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		if(strings.Contains(c.Request.URL.Path, "swagger")){
			c.Next()
			return
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw


		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = time.Since(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		keys := make(map[logging.ExtraKey]interface{})
		keys[logging.ClientIP] = param.ClientIP
		keys[logging.StatusCode] = param.StatusCode
		keys[logging.Latency] = param.Latency
		keys[logging.Path] = param.Path
		keys[logging.BodySize] = param.BodySize
		keys[logging.ErrorMessage] = param.ErrorMessage
		keys[logging.Method] = param.Method
		keys[logging.RequestBody] = string(bodyBytes)
		keys[logging.ResponseBody] = blw.body.String()
		
		logger.Info(logging.RequestResponse, logging.API, "", keys)
	}
}
