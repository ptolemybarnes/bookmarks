package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
)

type Bookmark struct {
	Description string
	Url         string
}

func loadBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark
	filePath := "./bookmarks.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return bookmarks, err
	}

	err = json.Unmarshal(data, &bookmarks)
	if err != nil {
		return bookmarks, err
	}
	return bookmarks, nil
}

func main() {
	bookmarks, err := loadBookmarks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	idx, err := fuzzyfinder.FindMulti(
		bookmarks,
		func(i int) string {
			return bookmarks[i].Description
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("Bookmark: %s \nUrl: %s",
				bookmarks[i].Description,
				bookmarks[i].Url)
		}))
	if err != nil {
		log.Fatal(err)
	}
	selectedUrl := bookmarks[idx[0]].Url
	fmt.Printf("selected: %v\n", selectedUrl)
	cmd := exec.Command("open", selectedUrl)
	cmd.Run()
}
