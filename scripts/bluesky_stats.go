// bluesky_stats.go is a script to fetch the number of likes, quotes, and reposts for a Bluesky post.
// It uses the atproto SDK (github.com/bluesky-social/indigo).
//
// Usage:
//
//	cd scripts
//	go run bluesky_stats.go <post_url_or_at_uri>
//
// Example:
//
//	go run bluesky_stats.go https://bsky.app/profile/bsky.app/post/3lbvlb537322q
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/xrpc"
)

type BlueskyStats struct {
	URI     string `json:"uri"`
	Likes   int64  `json:"likes"`
	Reposts int64  `json:"reposts"`
	Quotes  int64  `json:"quotes"`
	Replies int64  `json:"replies"`
}

type BlogPost struct {
	FilePath       string
	BlueskyPostURL string
}

func main() {
	postsDir := "../content/posts"

	ctx := context.Background()
	client := &xrpc.Client{
		Host: "https://public.api.bsky.app",
	}

	// Scan all posts mode
	posts, err := findPostsWithBlueskyRefs(postsDir)
	if err != nil {
		log.Fatalf("Error scanning posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts with 'bluesky' attribute found.")
		return
	}

	fmt.Printf("Found %d posts with Bluesky URLs\n", len(posts))
	allStats := make(map[string]BlueskyStats)

	for _, p := range posts {
		fmt.Printf("Fetching stats for %s...\n", p.BlueskyPostURL)
		stats, err := fetchStats(ctx, client, p.BlueskyPostURL)
		if err != nil {
			fmt.Printf("Warning: failed to fetch stats for %s: %v\n", p.BlueskyPostURL, err)
			continue
		}

		// Write individual JSON file in the post directory
		outPath := filepath.Join(filepath.Dir(p.FilePath), "bluesky-stats.json")
		if err := writeJSON(outPath, stats); err != nil {
			fmt.Printf("Warning: failed to write JSON to %s: %v\n", outPath, err)
		} else {
			fmt.Printf("Wrote stats to %s\n", outPath)
		}

		allStats[p.BlueskyPostURL] = *stats
	}
}

func writeJSON(path string, data interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func fetchStats(ctx context.Context, client *xrpc.Client, input string) (*BlueskyStats, error) {
	atUri := input
	if strings.HasPrefix(input, "https://bsky.app/profile/") {
		atUri = urlToAtUri(ctx, client, input)
	}

	// FeedGetPostThread returns the post and its replies/parents.
	// depth=0 and parentHeight=0 to just get the post itself.
	res, err := bsky.FeedGetPostThread(ctx, client, 0, 0, atUri)
	if err != nil {
		return nil, fmt.Errorf("failed to get post thread: %w", err)
	}

	if res.Thread.FeedDefs_ThreadViewPost == nil {
		return nil, fmt.Errorf("post not found or inaccessible (type: %T)", res.Thread)
	}

	post := res.Thread.FeedDefs_ThreadViewPost.Post

	stats := &BlueskyStats{
		URI: atUri,
	}
	if post.LikeCount != nil {
		stats.Likes = *post.LikeCount
	}
	if post.RepostCount != nil {
		stats.Reposts = *post.RepostCount
	}
	if post.QuoteCount != nil {
		stats.Quotes = *post.QuoteCount
	}
	if post.ReplyCount != nil {
		stats.Replies = *post.ReplyCount
	}
	return stats, nil
}

func findPostsWithBlueskyRefs(root string) ([]BlogPost, error) {
	var posts []BlogPost
	uniqueURLs := make(map[string]bool)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		inFrontMatter := false
		for scanner.Scan() {
			line := scanner.Text()
			if line == "---" {
				if !inFrontMatter {
					inFrontMatter = true
					continue
				} else {
					break // end of frontmatter
				}
			}
			if inFrontMatter {
				if strings.HasPrefix(line, "bluesky:") {
					url := strings.TrimSpace(strings.TrimPrefix(line, "bluesky:"))
					if url != "" && !uniqueURLs[url] {
						posts = append(posts, BlogPost{FilePath: path, BlueskyPostURL: url})
						uniqueURLs[url] = true
					}
				}
			}
		}
		return scanner.Err()
	})
	return posts, err
}

func urlToAtUri(ctx context.Context, client *xrpc.Client, url string) string {
	// Format: https://bsky.app/profile/<handle>/post/<postId>
	trimmed := strings.TrimPrefix(url, "https://bsky.app/profile/")
	parts := strings.Split(trimmed, "/")
	if len(parts) < 3 || parts[1] != "post" {
		log.Fatalf("Error: invalid Bluesky post BlueskyPostURL format: %s", url)
	}

	handle := parts[0]
	postId := parts[2]

	// If handle is already a DID, use it.
	if strings.HasPrefix(handle, "did:") {
		return fmt.Sprintf("at://%s/app.bsky.feed.post/%s", handle, postId)
	}

	// Resolve handle to DID
	res, err := atproto.IdentityResolveHandle(ctx, client, handle)
	if err != nil {
		log.Fatalf("Error: failed to resolve handle %s: %v", handle, err)
	}

	return fmt.Sprintf("at://%s/app.bsky.feed.post/%s", res.Did, postId)
}
