package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Release struct {
	Name string `json:"name"`
}

func main() {
	org := flag.String("org", "", "github org")
	repo := flag.String("repo", "", "github repo")

	flag.Parse()

	if len(*org) < 1 {
		fmt.Println("missing required option: org")
		return
	}

	if len(*repo) < 1 {
		fmt.Println("missing required option: repo")
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", *org, *repo)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var releases []Release

	err = json.Unmarshal(body, &releases)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	count := len(releases)

	for index := count - 1; index >= 0; index-- {
		release := releases[index]
		fmt.Println(release.Name[1:])
	}
}
