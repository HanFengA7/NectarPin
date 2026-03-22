package bootstrap

import "strings"

var defaultCorsAllowedOrigins = []string{
	"http://localhost:5173",
	"http://127.0.0.1:5173",
	"http://localhost:4173",
	"http://127.0.0.1:4173",
}

// EffectiveCorsAllowedOrigins 返回实际用于 CORS 校验的 Origin 列表。
// 若配置中非空则仅用配置项（去空白、去尾部斜杠）；否则使用本地开发默认列表。
func EffectiveCorsAllowedOrigins(s *ServerConfig) []string {
	if s == nil || len(s.CorsAllowedOrigins) == 0 {
		out := make([]string, len(defaultCorsAllowedOrigins))
		copy(out, defaultCorsAllowedOrigins)
		return out
	}
	out := make([]string, 0, len(s.CorsAllowedOrigins))
	for _, o := range s.CorsAllowedOrigins {
		t := strings.TrimSpace(o)
		if t == "" {
			continue
		}
		out = append(out, strings.TrimSuffix(t, "/"))
	}
	if len(out) == 0 {
		out = append(out, defaultCorsAllowedOrigins...)
	}
	return out
}

// DefaultCorsOriginsHint 内置本地开发 CORS 白名单的展示用字符串（向导提示等）
func DefaultCorsOriginsHint() string {
	return strings.Join(defaultCorsAllowedOrigins, ", ")
}
