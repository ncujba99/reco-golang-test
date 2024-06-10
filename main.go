package main

import (
	"fmt"
	"reco-golang-test/internal/asanaClient"
)

func main() {

	workspaces, err := asanaClient.GetWorkspaces()
	if err != nil {
		panic(err)
	}
	for _, workspace := range workspaces {
		fmt.Println(workspace.GID)

		projects, nextPage := asanaClient.GetProjects(
			1,             // Limit: Number of projects to retrieve
			1,             // Offset: Starting index for pagination
			workspace.GID, // Replace with your Asana workspace ID
			"",            // Team: Optional team ID (leave empty if not needed)
			false,         // Archived: Include archived projects (false for active)
			nil,           // optFields: Optional list of additional fields to retrieve (leave nil for defaults)
		)

		if nextPage != nil {

		}

		for _, project := range projects {
			fmt.Println(project)

		}
	}

}
