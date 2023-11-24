package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var DotEnvFile = os.Getenv("DOTENV")
var DefaultDotEnvFile = os.Getenv("DOTENV_DEFAULT")

func init() {
	if DotEnvFile == "" {
		DotEnvFile = ".env"
	}
	if DefaultDotEnvFile == "" {
		DefaultDotEnvFile = ".env.default"
	}
	changeDir()
	if _, err := os.Stat(DotEnvFile); err == nil {
		err = godotenv.Load(DotEnvFile)
		if err != nil {
			log.Fatal("Error loading "+DotEnvFile+" file\n", err)
		}
	}
	if _, err := os.Stat(DefaultDotEnvFile); err == nil {
		err := godotenv.Load(DefaultDotEnvFile)
		if err != nil {
			log.Fatal("Error loading "+DefaultDotEnvFile+" file\n", err)
		}
	}
}

func changeDir() {
	wd, _ := os.Getwd()
	for filepath.Dir(wd) != wd {
		if FileExists(filepath.Join(wd, DefaultDotEnvFile)) || FileExists(filepath.Join(wd, DotEnvFile)) {
			break
		}
		wd = filepath.Dir(wd)
	}
	if !FileExists(filepath.Join(wd, DefaultDotEnvFile)) && !FileExists(filepath.Join(wd, DotEnvFile)) {
		log.Panicf("Can not find %s or %s file", DefaultDotEnvFile, DotEnvFile)
	}
	fmt.Println("Find .env file in " + wd)
	os.Chdir(wd)

}

func Get(key string, fallback ...string) string {
	value := os.Getenv(key)
	if len(value) == 0 && len(fallback) > 0 {
		return fallback[0]
	}
	return value
}

func GetInt(key string, fallback ...int) (int, error) {
	value := os.Getenv(key)
	if value == "" && len(fallback) > 0 {
		return fallback[0], nil
	}
	return strconv.Atoi(value)
}

func GetIntMust(key string, fallback ...int) int {
	value := os.Getenv(key)
	if value == "" && len(fallback) > 0 {
		return fallback[0]
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return v
}

func GetMust(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(key + " should be setted")
	}
	return value
}
func GetFloat(key string, fallback ...float64) (float64, error) {
	value := os.Getenv(key)
	if value == "" && len(fallback) > 0 {
		return fallback[0], nil
	}
	return strconv.ParseFloat(value, 64)
}

func GetFloatMust(key string, fallback ...float64) float64 {
	value := os.Getenv(key)
	if len(value) == 0 && len(fallback) > 0 {
		return fallback[0]
	}
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return v
}
func GetBool(key string, fallback ...bool) (bool, error) {
	value := os.Getenv(key)
	if value == "" && len(fallback) > 0 {
		return fallback[0], nil
	}
	return strconv.ParseBool(value)
}

func GetBoolMust(key string, fallback bool) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return v
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func Getwd() string {
	wd, _ := os.Getwd()
	return wd
}

// GetPath return absolute path if p is relative path
func GetPath(p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(Getwd(), p)
}
