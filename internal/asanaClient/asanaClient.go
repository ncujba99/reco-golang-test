package asanaClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Workspace struct {
	GID          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type WorkspaceResponse struct {
	Data []Workspace `json:"data"`
}

var token = os.Getenv("YOUR_TOKEN")

func GetWorkspaces() ([]Workspace, error) {

	u := url.URL{
		Scheme: "https",
		Host:   "app.asana.com",
		Path:   "api/1.0/workspaces",
	}

	req, _ := http.NewRequest("GET", u.String(), nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Print("error")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response WorkspaceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, nil
	}
	return response.Data, err

}

type Project struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

// Define a struct to represent the entire response
type ProjectsResponse struct {
	Data     []Project `json:"data"`
	NextPage *string   `json:"next_page"` // Use a pointer for optional field
}

func GetProjects(limit int, offset int, workspace string, team string, archived bool, optFields []string) ([]Project, *string) {

	u := url.URL{
		Scheme: "https",
		Host:   "app.asana.com",
		Path:   "/api/1.0/projects",
	}

	params := url.Values{}
	params.Add("limit", fmt.Sprint(limit))
	// params.Add("offset", fmt.Sprint(offset))
	params.Add("workspace", workspace)
	if team != "" {
		params.Add("team", team)
	}
	params.Add("archived", fmt.Sprint(archived))
	for _, field := range optFields {
		params.Add("opt_fields", field)
	}

	u.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Print("aaa")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response ProjectsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, nil
	}
	return response.Data, response.NextPage

}
