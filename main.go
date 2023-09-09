package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	numContainers := 5 // Change this to the desired number of containers

	composeData := map[string]interface{}{
		"version":  "3",
		"services": make(map[string]interface{}),
	}

	for i := 1; i <= numContainers; i++ {
		service := make(map[string]interface{})
		serviceName := fmt.Sprintf("web%d", i)
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
