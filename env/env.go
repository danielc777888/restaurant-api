// The env package is a wrapper around the .env file and returns config values with the correct types not just strings.
package env

import (
	"os"
	"strings"
)

func DbDsn() string {
	return os.Getenv("DB_DSN")
}

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func RedisAddress() string {
	return os.Getenv("REDIS_ADDRESS")
}

func GeminiAPIKey() string {
	return os.Getenv("GEMINI_API_KEY")
}

func CacheEnabled() bool {
	cacheEnabled := os.Getenv("CACHE_ENABLED")
	return strings.ToLower(cacheEnabled) == "true"
}

func LLMEnabled() bool {
	LLMEnabled := os.Getenv("LLM_ENABLED")
	return strings.ToLower(LLMEnabled) == "true"
}
