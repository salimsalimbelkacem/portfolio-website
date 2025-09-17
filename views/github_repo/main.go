package githubrepo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repo struct {

	// name, url, description, language
	Name string `json:"name"`
	Htmlurl string `json:"html_url"`
	Descritption string `json:"description"`
	Language string `json:"language"`
}

type CacheEntry struct {
	ETag string
	Data []Repo
}

var RepoCache = map[string]*CacheEntry{}

func GetReposList(username string) []Repo {
	url := "https://api.github.com/users/" + username + "/repos"

	req, _ := http.NewRequest("GET", url, nil)

	if entry, ok := RepoCache[username]; ok && entry.ETag != "" {
		req.Header.Set("If-None-Match", entry.ETag)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		fmt.Println("Cache hit (ETag in memory)")
		return RepoCache[username].Data
	}

	body, _ := io.ReadAll(resp.Body)

	var repos []Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	etag := resp.Header.Get("ETag")
	RepoCache[username] = &CacheEntry{
		ETag: etag,
		Data: repos,
	}

	return repos
}
