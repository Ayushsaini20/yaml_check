package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v53/github"
)

const (
	webhookSecret  = "12345678"
	owner          = "Ayushsaini20"
	repo           = "yaml_check"
	port           = ":8080"
	appID          = 1238632
	installationID = 65705633
	privateKeyPath = "/mnt/c/Users/AyushSaini/Downloads/pr-merge-automation.2025-05-02.private-key.pem"
)

var (
	processedEvents = make(map[string]time.Time)
	client          *github.Client
)

func main() {
	// Initialize GitHub App client
	var err error
	client, err = initGitHubAppClient()
	if err != nil {
		log.Fatalf("Failed to initialize GitHub client: %v", err)
	}

	http.HandleFunc("/webhook", webhookHandler)
	log.Printf("Starting webhook server on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func initGitHubAppClient() (*github.Client, error) {
	transport, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appID, installationID, privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create GitHub App transport: %v", err)
	}
	return github.NewClient(&http.Client{Transport: transport}), nil
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Verify webhook signature
	if !verifyWebhookSignature(r, webhookSecret) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// Log all headers for debugging
	log.Printf("Request headers: %+v\n", r.Header)

	// Get event type
	eventType := r.Header.Get("X-GitHub-Event")
	log.Printf("Received event: type=%s\n", eventType)

	// Decode payload to get PR number and action
	var payload github.PullRequestEvent
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body)) // Restore body

	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Printf("Error decoding payload: %v\n", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	pr := payload.GetPullRequest()
	action := payload.GetAction()
	prNumber := pr.GetNumber()

	// Create a deduplication key based on PR number and action
	dedupKey := fmt.Sprintf("PR%d-%s", prNumber, action)
	log.Printf("Deduplication key: %s\n", dedupKey)

	// Check for duplicate events
	if _, exists := processedEvents[dedupKey]; exists {
		log.Printf("Ignoring duplicate event: key=%s\n", dedupKey)
		return
	}
	processedEvents[dedupKey] = time.Now()

	// Clean up processed events older than 5 minutes
	cleanupProcessedEvents()

	// Log the PR number and action
	log.Printf("Processing PR #%d, action=%s\n", prNumber, action)

	if action != "opened" && action != "synchronize" {
		log.Printf("Ignoring action: %s for PR #%d\n", action, prNumber)
		return
	}

	prInfo, _, err := client.PullRequests.Get(context.Background(), owner, repo, prNumber)
	if err != nil {
		log.Printf("Error getting PR info for PR #%d: %v\n", prNumber, err)
		return
	}

	log.Printf("PR #%d: Mergeable=%v, MergeState=%s\n", prNumber, prInfo.GetMergeable(), prInfo.GetMergeableState())

	if prInfo.GetMergeable() && prInfo.GetMergeableState() == "clean" {
		ok, err := runTests(client, owner, repo, prNumber)
		if err != nil {
			log.Printf("Error running tests for PR #%d: %v\n", err)
			return
		}
		if ok {
			mergePR(client, prNumber)
		} else {
			log.Printf("Tests failed for PR #%d\n", prNumber)
		}
	} else {
		log.Printf("PR #%d not mergeable or not clean\n", prNumber)
	}
}

func verifyWebhookSignature(r *http.Request, secret string) bool {
	signature := r.Header.Get("X-Hub-Signature-256")
	if signature == "" {
		return false
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		return false
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return "sha256="+expectedMAC == signature
}

func cleanupProcessedEvents() {
	for eventKey, timestamp := range processedEvents {
		if time.Since(timestamp) > 5*time.Minute {
			log.Printf("Cleaning up old event: key=%s\n", eventKey)
			delete(processedEvents, eventKey)
		}
	}
}

func mergePR(client *github.Client, prNumber int) {
	commitMsg := "Auto-merged by PR bot"
	options := &github.PullRequestOptions{
		MergeMethod: "squash",
	}
	_, _, err := client.PullRequests.Merge(context.Background(), owner, repo, prNumber, commitMsg, options)
	if err != nil {
		log.Printf("Failed to merge PR #%d: %v\n", prNumber, err)
		return
	}
	log.Printf("Successfully merged PR #%d\n", prNumber)
}

func runTests(client *github.Client, owner, repo string, prNumber int) (bool, error) {
	tempDir, err := os.MkdirTemp("", "pr-files-test")
	if err != nil {
		return false, err
	}
	defer os.RemoveAll(tempDir)

	prInfo, _, err := client.PullRequests.Get(context.Background(), owner, repo, prNumber)
	if err != nil {
		return false, err
	}
	headSHA := prInfo.GetHead().GetSHA()

	files, _, err := client.PullRequests.ListFiles(context.Background(), owner, repo, prNumber, nil)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		filename := file.GetFilename()
		log.Printf("Processing file in PR #%d: %s\n", prNumber, filename)
		content, _, _, err := client.Repositories.GetContents(
			context.Background(), owner, repo, filename,
			&github.RepositoryContentGetOptions{Ref: headSHA},
		)
		if err != nil {
			log.Printf("Error fetching file %s: %v\n", filename, err)
			continue
		}

		text, err := content.GetContent()
		if err != nil {
			log.Printf("Error decoding file %s: %v\n", filename, err)
			continue
		}

		fullPath := filepath.Join(tempDir, filename)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return false, err
		}
		if err := os.WriteFile(fullPath, []byte(text), 0644); err != nil {
			return false, err
		}
	}

	hasPython := false
	hasYaml := false
	hasGo := false
	for _, file := range files {
		name := file.GetFilename()
		if strings.HasSuffix(name, ".py") {
			hasPython = true
		}
		if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") {
			hasYaml = true
		}
		if strings.HasSuffix(name, ".go") {
			hasGo = true
		}
	}

	if hasPython {
		cmd := exec.Command("pytest", tempDir)
		cmd.Dir = tempDir
		output, err := cmd.CombinedOutput()
		log.Printf("Pytest output:\n%s", output)
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 5 {
				log.Printf("No Python tests found; treating as pass")
			} else {
				log.Printf("Python tests failed: %v\n", err)
				return false, nil
			}
		}
	}

	if hasYaml {
		cmd := exec.Command("yamllint", tempDir)
		cmd.Dir = tempDir
		output, err := cmd.CombinedOutput()
		log.Printf("YAML lint output:\n%s", output)
		if err != nil {
			log.Printf("YAML linting failed: %v\n", err)
			return false, nil
		}
	}

	if hasGo {
		cmd := exec.Command("/snap/bin/golangci-lint", "run", "--path-prefix", tempDir)
		cmd.Dir = tempDir
		output, err := cmd.CombinedOutput()
		log.Printf("golangci-lint output:\n%s", output)
		if err != nil {
			log.Printf("Go linting failed: %v\n", err)
			return false, nil
		}
	}

	return true, nil
}