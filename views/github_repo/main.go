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

func GetReposList(username string) (repos []Repo) {
	url := "https://api.github.com/users/" + username + "/repos?sort=updated&type=all"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	return
}

// func main() {
// 	for _,repo := range getReposList("salimsalimbelkacem"){
// 		println(repo.Name)
// 		println(repo.Descritption)
// 		println(repo.Htmlurl)
// 		println(repo.Language)
// 		println("-----------------")
// 	}
// }
