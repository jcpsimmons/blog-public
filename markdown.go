package main

import (
	"embed"
	"log"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
)

//go:embed **/*.md
var files embed.FS

func FileNameToTitle(fileName string) string {
	humanReadableName := strings.Split(fileName, ".")[0]
	humanReadableName = strings.Split(humanReadableName, "_")[1]
	humanReadableName = strings.ReplaceAll(humanReadableName, "-", " ")
	return humanReadableName
}

func GetAllEssaysAsListItems() ([]string, error) {
	files, err := files.ReadDir("essay")
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	var fileNames []string
	for _, file := range files {
		fileName := file.Name()
		humanReadableName := FileNameToTitle(fileName)
		if strings.HasSuffix(fileName, ".md") {
			// wrap in <a> and <li> tags
			htmlFileName := "<li><p class='offbit-medium'><a href=\"/essay/" + fileName[:len(fileName)-3] + "\">" + humanReadableName + "</p></a></li>"
			fileNames = append(fileNames, htmlFileName)
		}
	}
	return fileNames, nil
}

func GetMarkdown(fileName string) ([]byte, error) {
	log.Println("fileName:", fileName)

	post, err := files.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return markdown.ToHTML(post, nil, nil), nil
}
