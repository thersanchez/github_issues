package main

import "fmt"

func main() {
	user := "thersanchez"
	repo := "caesar"
	url := issueURL(user, repo)
	fmt.Println(url)
}

// issueURL returns the unquoted URL to access the issues from the Github
// repository of the given user with the given repo name.
func issueURL(user, repo string) string {
	const urlFmt = "https://api.github.com/repos/%s/%s/issues"
	return fmt.Sprintf(urlFmt, user, repo)
}
