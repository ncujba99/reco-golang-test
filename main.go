package main

import (
	"encoding/json"
	"log"
	"os"
	"reco-golang-test/internal/asanaClient"
)

func SaveStructAsJSON(data any, filename string) error {
	// Marshal the struct into a byte slice
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Open the file for writing with appropriate permissions (0644)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed even on errors

	_, err = file.Write(jsonData)
	return err
}

func main() {

	chunk_size := 10

	workspaces, err := asanaClient.GetWorkspaces()
	if err != nil {
		log.Printf("Error occurred while processing data: %v\n", err)
	}
	for _, workspace := range workspaces {

		projects, offset, err := asanaClient.GetProjects(
			chunk_size,    // Limit: Number of projects to retrieve
			nil,           // Offset: Starting index for pagination
			workspace.GID, // Asana workspace ID
			nil,           // Team: Optional team ID (leave empty if not needed)
			false,         // Archived: Include archived projects (false for active)
			nil,           // optFields: Optional list of additional fields to retrieve (leave nil for defaults)
		)
		if err != nil {
			log.Printf("Error occurred while processing data: %v\n", err)
		}

		for _, project := range projects {
			SaveStructAsJSON(project, project.Gid+".json")
		}

		for *offset != "" {
			projects, offset, err = asanaClient.GetProjects(
				chunk_size,    // Limit: Number of projects to retrieve
				offset,        // Offset: Starting index for pagination
				workspace.GID, // Asana workspace ID
				nil,           // Team: Optional team ID (leave empty if not needed)
				false,         // Archived: Include archived projects (false for active)
				nil,           // optFields: Optional list of additional fields to retrieve (leave nil for defaults)
			)

			if err != nil {
				log.Printf("Error occurred while processing data: %v\n", err)
			}
			for _, project := range projects {
				SaveStructAsJSON(project, project.Gid+".json")
			}

		}

	}

}
