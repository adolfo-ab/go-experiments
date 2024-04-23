package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GitHubUser struct {
	Login   string `json:"login"`
	Name    string `json:"name"`
	NumRepo int    `json:"public_repos"`
}

func userInfo(login string) (*GitHubUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30000*time.Millisecond)
	defer cancel()

	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	fmt.Println(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil

}

func main() {
	user, err := userInfo("adolfo-ab")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
