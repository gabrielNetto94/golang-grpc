package configs

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Get enviroment by key
func GetEnv(key string) string {

	filename := "../../.env"
	env := make(map[string]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error read file", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error scan", err.Error())
	}

	val, ok := env[key]
	if !ok {
		log.Fatal("env variable -> ", key, " not found")
	}
	return val
}
