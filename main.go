package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	numContainers := 5 // Change this to the desired number of containers

	composeData := map[string]interface{}{
		"version":  "3",
		"services": make(map[string]interface{}),
		"volumes":  make(map[string]interface{}),
	}

	db := make(map[string]interface{})
	var dbEnvVars []string
	dbEnvVars = append(dbEnvVars, fmt.Sprintf("POSTGRES_USER=%v", os.Getenv("POSTGRES_USER")))
	dbEnvVars = append(dbEnvVars, fmt.Sprintf("POSTGRES_PASSWORD=%v", os.Getenv("POSTGRES_PASS")))

	dbName := "konseouldb"
	db["image"] = "postgres:14.1-alpine"
	db["restart"] = "always"
	db["environment"] = dbEnvVars
	db["ports"] = []string{fmt.Sprintf("5432:5432")}
	db["volumes"] = []string{fmt.Sprintf("konseouldb:/var/lib/postgresql/data")}
	composeData["services"].(map[string]interface{})[dbName] = db

	volume := make(map[string]interface{})
	volumeName := dbName
	volume["driver"] = "local"
	composeData["volumes"].(map[string]interface{})[volumeName] = volume

	for i := 1; i <= numContainers; i++ {
		service := make(map[string]interface{})
		serviceName := fmt.Sprintf("cornflake-%d", i)
		port := 8080 + i

		service["build"] = map[string]interface{}{
			"context": "./cornflake", // Change the context path as needed
		}
		service["ports"] = []string{fmt.Sprintf("%d:8080", port)}
		service["environment"] = []string{fmt.Sprintf("PORT=%d", 8080), fmt.Sprintf("INDEX=%d", i)}
		service["env_file"] = []string{fmt.Sprintf(".env")}
		composeData["services"].(map[string]interface{})[serviceName] = service
	}

	yamlData, err := yaml.Marshal(composeData)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("../docker-compose.yml", yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Docker Compose file generated successfully.")
}
