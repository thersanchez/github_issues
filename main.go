package main

import "fmt"

func main() {
	user := "thersanchez"
	repo := "caesar"
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", user, repo)
	fmt.Println(url)
}
