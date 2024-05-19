// The env package is a wrapper around the .env file and returns config values with the correct types not just strings.
package env

import (
	"os"
	"strings"
)

func CacheEnabled() bool {
	cacheEnabled := os.Getenv("CACHE_ENABLED")
	return strings.ToLower(cacheEnabled) == "true"
}

func LLMEnabled() bool {
	LLMEnabled := os.Getenv("LLM_ENABLED")
	return strings.ToLower(LLMEnabled) == "true"
}
