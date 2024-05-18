package env

import (
	"os"
	"strings"
)

func CacheEnabled() bool {
	cacheEnabled := os.Getenv("CACHE_ENABLED")
	return strings.ToLower(cacheEnabled) == "true"
}
