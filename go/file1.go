package main

import (
	"context"
	"fmt"
	"log"

)

func main() {
	
	token := "ghp_1234567890abcdefEXAMPLETOKEN"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Example API call
	repo, _, err := client.Repositories.Get(ctx, "octocat", "Hello-World")
	if err != nil {
		log.Fatalf("Error fetching repo: %v", err)
	}

	fmt.Printf("Repository name: %s\n", *repo.Name)
}
