package asanaClient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Project struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	StartDate   time.Time  `json:"startDate,omitempty"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
}

func GetProjects(limit int, offset int, workspace string, team string, archived bool, optFields []string) ([]Project, error) {

	u := url.URL{
		Scheme: "https",
		Host:   "app.asana.com",
		Path:   "/api/1.0/projects",
	}

	params := url.Values{}
	params.Add("limit", fmt.Sprint(limit))
	params.Add("offset", fmt.Sprint(offset))
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
	req.Header.Add("authorization", "Bearer 7112637efe2ee31f7e89a254fcd37c97")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Print("aaa")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	return nil, nil

}
