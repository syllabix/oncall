package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type SlackSettings struct {
	AuthToken         string
	AppToken          string
	WebHook           string
	SigningSecret     string
	VerificationToken string
	DebugMode         bool
}

type ServerSettings struct {
	Host         string
	Port         string
	ProfilerPort string
	DocsURL      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Load() (ServerSettings, SlackSettings) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("unable to load .env file: reason %v", err))
	}

	return ServerSettings{
			Host:         os.Getenv("HOST"),
			Port:         os.Getenv("PORT"),
			DocsURL:      os.Getenv("DOCS_PATH"),
			ProfilerPort: os.Getenv("PPROF"),
			ReadTimeout:  getEnvAsDur("READ_TIMEOUT", time.Minute*2),
			WriteTimeout: getEnvAsDur("WRITE_TIMEOUT", time.Minute*2),
		}, SlackSettings{
			AuthToken:         os.Getenv("SLACK_AUTH_TOKEN"),
			AppToken:          os.Getenv("SLACK_APP_TOKEN"),
			WebHook:           os.Getenv("SLACK_WEB_HOOK"),
			SigningSecret:     os.Getenv("SLACK_SIGNING_SECRET"),
			VerificationToken: os.Getenv("SLACK_VERIFICATION_TOKEN"),
			DebugMode:         asBool("SLACK_ENABLE_DEBUG_MODE", true),
		}
}

// getEnvAsDur returns a duration value for and environment variable key
// if the key is empty or value is empty (or not in valid duration syntax)
// this func returns the provided default
func getEnvAsDur(key string, def time.Duration) time.Duration {
	val := os.Getenv(key)
	dur, err := time.ParseDuration(val)
	if err != nil {
		return def
	}
	return dur
}

func asBool(key string, def bool) bool {
	val := os.Getenv(key)
	switch strings.ToLower(val) {
	case "true":
		return true

	case "false":
		return false

	default:
		return def
	}
}
