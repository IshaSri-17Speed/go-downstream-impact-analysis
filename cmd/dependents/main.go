package main

import (
	"context"
	"flag"
	"log"

	"github.com/IshaSri-17Speed/go-downstream-impact-analysis/internal/dependents"
)

func main() {
	module := flag.String("module", "", "Go module path (e.g. go.opentelemetry.io/otel)")
	limit := flag.Int("limit", 20, "maximum number of downstream repositories")
	flag.Parse()

	if *module == "" {
		log.Fatal("usage: dependents -module <module-path>")
	}

	ctx := context.Background()

	repos, err := dependents.FetchFromGitHub(ctx, *module, *limit)
	if err != nil {
		log.Fatalf("failed to fetch dependents: %v", err)
	}

	log.Printf("Found %d downstream repositories using %s:\n", len(repos), *module)
	for _, r := range repos {
		log.Printf(" - %s (%s)", r.FullName, r.URL)
	}
}
