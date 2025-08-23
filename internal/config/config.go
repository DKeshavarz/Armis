package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init(){
	Load(".env")
}

func Load(file string){
	err := godotenv.Load(file)
	if err != nil {
		log.Printf("Error in loading .env : %s", err)
	}
}


func GetEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
    valueStr := GetEnv(key, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultValue
}

func GetEnvAsBool(key string, defaultValue bool) bool {
    valueStr := GetEnv(key, "")
    if value, err := strconv.ParseBool(valueStr); err == nil {
        return value
    }
    return defaultValue
}

func GetEnvAsSlice(key string, defaultValue []string, separator string) []string {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	// Split by separator and trim whitespace from each element
	values := strings.Split(valueStr, separator)
	result := make([]string, 0, len(values))
	
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	
	return result
}