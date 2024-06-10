package asanaClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ProjectResponse struct {
	Data     []AsanaProject `json:"data"`
	NextPage NextPageInfo   `json:"next_page"`
}

type AsanaProject struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type NextPageInfo struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}

func GetProjects(limit int, offset *string, workspace string, team *string, archived bool, optFields []string) ([]AsanaProject, *string, error) {

	u := url.URL{
		Scheme: "https",
		Host:   "app.asana.com",
		Path:   "/api/1.0/projects",
	}

	params := url.Values{}
	params.Add("limit", fmt.Sprint(limit))
	if offset != nil {
		params.Add("offset", *offset)
	}

	params.Add("workspace", workspace)
	params.Add("archived", fmt.Sprint(archived))
	if team != nil {
		params.Add("team", *team)
	}
	for _, field := range optFields {
		params.Add("opt_fields", field)
	}

	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error making request: %w", err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Printf("error closing response body: %v\n", err)
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var response ProjectResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return response.Data, &response.NextPage.Offset, nil
}
