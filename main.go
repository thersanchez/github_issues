package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	user := "thersanchez"
	repo := "caesar"
	url := issueURL(user, repo)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("expected status 200, got %d", resp.StatusCode)
	}
	printReader(resp.Body)
}

// issueURL returns the unquoted URL to access the issues from the Github
// repository of the given user with the given repo name.
func issueURL(user, repo string) string {
	const urlFmt = "https://api.github.com/repos/%s/%s/issues"
	return fmt.Sprintf(urlFmt, user, repo)
}

// printReader prints the bytes read from r to stdout in chunks of 1024 bytes
func printReader(r io.Reader) {
	chunk := make([]byte, 1024)
	var err error
	var n int
	for err != io.EOF {
		n, err = r.Read(chunk)
		if n > 0 {
			fmt.Print(string(chunk))
		}
		if err != nil && err != io.EOF {
			log.Fatalf("reading chunk: %s", err)
		}
	}
}
