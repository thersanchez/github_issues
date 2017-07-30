package main

import "fmt"

func main() {
	user := "thersanchez"
	repo := "caesar"
	urlFmt := "https://api.github.com/repos/%s/%s/issues"
	url := fmt.Sprintf(urlFmt, user, repo)
	fmt.Println(url)
}
