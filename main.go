package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const github_url_base string = "https://api.github.com/users/"

type user struct {
	name   string
	events []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("github-activity <Username>")
		return
	}

	user_1 := user{os.Args[1], nil}

	fmt.Printf("Fetching statistics for user: %s\n", user_1.name)

	target_url := github_url_base + user_1.name + "/events"
	fmt.Printf("Target URL: %s\n", target_url)

	resp, err := http.Get(target_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("User events:")

	fmt.Print(string(body))
}
