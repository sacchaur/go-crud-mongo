package configs

// package env reads environment variables from file or OS and creates env.Map

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

var Map map[string]string

var config ApiConfig

func Load(source string) {
	var env map[string]string

	if source == "file" {
		wd, err := os.Getwd() // Get current working directory
		if err != nil {
			log.Fatalf("Error getting working directory: %v", err)
		}

		envFile := filepath.Join(wd, ".env")
		env, err := godotenv.Read(envFile)
		if err != nil {
			log.Fatalf("Error loading %s: %v", envFile, err)
		}
		Map = env
		log.Printf("%s loaded", envFile)
		return
	}

	// load OS variables
	variables := os.Environ()
	env = make(map[string]string)
	for _, variable := range variables {
		// split by equal sign
		keyVal := strings.Split(variable, "=")
		env[keyVal[0]] = keyVal[1]
	}
	Map = env
	log.Println("OS variables loaded")
}

// ParseToStruct parses the environment variables from the Map into
// the fields of the provided struct.
func ParseToStruct[T any](obj *T) {
	// set environment variables before parsing
	for key, value := range Map {
		os.Setenv(key, value)
	}

	objValue := reflect.ValueOf(obj).Elem()
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		envVar, exists := Map[field.Tag.Get("env")]
		if exists {
			fieldValue := objValue.Field(i)
			if fieldValue.CanSet() {
				fieldValue.SetString(envVar)
			}
		}
	}
}

// ApiConfig stores env configuration settings for the application.
type ApiConfig struct {
	AppCode             string `env:"APP_CODE"`
	AppEnv              string `env:"APP_ENV"`
	AppVersion          string `env:"APP_VERSION"`
	AppPort             string `env:"PORT"`
	MongoDBURI          string `env:"MONGODB_URI"`
	MongoDBName         string `env:"MONGODB_DB_NAME"`
	MongoUserCollection string `env:"MONGODB_USER_COLLECTION"`
	MongoUsername       string `env:"MONGODB_USERNAME"`
	MongoPassword       string `env:"MONGODB_PASSWORD"`
	MongoTimeout        string `env:"MONGODB_TIMEOUT"`
	// AppKey             string `env:"APP_KEY"`
	// LocalRedisPort     string `env:"LOCAL_REDIS_PORT"`
	// LocalRedisHost     string `env:"LOCAL_REDIS_HOST"`
	// LocalRedisPass     string `env:"LOCAL_REDIS_PASS"`
	// CentralRedisPort   string `env:"CENTRAL_REDIS_PORT"`
	// CentralRedisHost   string `env:"CENTRAL_REDIS_HOST"`
	// CentralRedisPass   string `env:"CENTRAL_REDIS_PASS"`
	// TimeZone           string `env:"TIMEZONE"`
	// AWSBucket          string `env:"AWS_BUCKET"`
	// AppBaseUrl         string `env:"BASE_URL"`
	// LogPath            string `env:"LOG_PATH"`
	// NewRelicLicenseKey string `env:"NEW_RELIC_LICENSE_KEY"`
	// TokenBankHealth    string `env:"TOKEN_BANK_HEALTH"`
	// AppServicesHealth  string `env:"APPSERVICES_HEALTH"`
	// UserApiHealth      string `env:"USER_API_HEALTH"`
	// LogFormat          string `env:"LOG_FORMAT"`
}

func NewApiConfig() (ApiConfig, error) {
	Load("file")
	ParseToStruct(&config)

	return config, nil
}

func GetConfig() ApiConfig {
	return config
}
