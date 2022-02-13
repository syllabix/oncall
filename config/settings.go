package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/syllabix/oncall/datastore/db"
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

func Load() (ServerSettings, SlackSettings, db.Settings) {
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
		}, db.Settings{
			DBName:             os.Getenv("DB_NAME"),
			SSLMode:            os.Getenv("SSL_MODE"),
			User:               os.Getenv("DB_USER"),
			Password:           os.Getenv("DB_PASSWORD"),
			Host:               os.Getenv("DB_HOST"),
			Port:               os.Getenv("DB_PORT"),
			MaxConnections:     getEnvAsInt("MAX_CONNS", 5),
			MaxIdleConnections: getEnvAsInt("MAX_IDLE_CONNS", 5),
			MaxConnLifetime:    getEnvAsDur("MAX_CONN_LIFETIME", time.Hour),
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

func getEnvAsInt(key string, def int) int {
	val := os.Getenv(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return num
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
