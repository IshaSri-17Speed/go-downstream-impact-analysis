package dependents

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Repo represents a downstream repository.
type Repo struct {
	FullName string
	URL      string
}

const githubSearchAPI = "https://api.github.com/search/code"

// FetchFromGitHub discovers downstream Go modules via GitHub search.
func FetchFromGitHub(ctx context.Context, modulePath string, maxResults int) ([]Repo, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN is not set")
	}

	query := fmt.Sprintf(`"%s" filename:go.mod`, modulePath)
	encodedQuery := url.QueryEscape(query)

	endpoint := fmt.Sprintf(
		"%s?q=%s&per_page=%d",
		githubSearchAPI,
		encodedQuery,
		maxResults,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("github request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned %s", resp.Status)
	}

	var raw struct {
		Items []struct {
			Repository struct {
				FullName string `json:"full_name"`
				HTMLURL  string `json:"html_url"`
			} `json:"repository"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	results := make([]Repo, 0, len(raw.Items))
	seen := make(map[string]struct{})

	for _, item := range raw.Items {
		name := item.Repository.FullName
		if name == "" {
			continue
		}
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}

		results = append(results, Repo{
			FullName: name,
			URL:      item.Repository.HTMLURL,
		})
	}

	return results, nil
}
