package middlewares

import (
	"bytes"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// publicHotCacheTTL 热点公开只读接口的进程内缓存时长（发布后最多延迟该时间对外可见）。
const publicHotCacheTTL = 45 * time.Second

type responseCacheEntry struct {
	status int
	body   []byte
	expire time.Time
}

// 按「方法 + 路径 + 查询串」缓存；仅用于无用户态、无副作用的 GET。
var publicHotCache sync.Map

type captureWriter struct {
	gin.ResponseWriter
	buf    *bytes.Buffer
	status int
}

func (w *captureWriter) WriteHeader(code int) {
	if w.status == 0 {
		w.status = code
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w *captureWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	w.buf.Write(b)
	return w.ResponseWriter.Write(b)
}

// PublicHotResponseCache 为高频公开 GET 提供短 TTL 响应体缓存，降低数据库压力。
// 不得用于带按请求副作用的接口（例如公开文章详情会记录阅读量）。
func PublicHotResponseCache() gin.HandlerFunc {
	return publicResponseCacheWithTTL(publicHotCacheTTL)
}

func publicResponseCacheWithTTL(ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}
		key := c.Request.Method + " " + c.Request.URL.Path
		if rq := c.Request.URL.RawQuery; rq != "" {
			key += "?" + rq
		}
		if v, ok := publicHotCache.Load(key); ok {
			ent := v.(*responseCacheEntry)
			if time.Now().Before(ent.expire) {
				c.Data(ent.status, "application/json; charset=utf-8", ent.body)
				c.Abort()
				return
			}
			publicHotCache.Delete(key)
		}

		buf := &bytes.Buffer{}
		cw := &captureWriter{ResponseWriter: c.Writer, buf: buf}
		c.Writer = cw
		c.Next()

		if cw.status != http.StatusOK || buf.Len() == 0 || c.IsAborted() {
			return
		}
		publicHotCache.Store(key, &responseCacheEntry{
			status: cw.status,
			body:   append([]byte(nil), buf.Bytes()...),
			expire: time.Now().Add(ttl),
		})
	}
}
