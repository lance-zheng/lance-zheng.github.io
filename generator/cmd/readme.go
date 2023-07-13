package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

type Readme struct {
	Categories       map[string]Category
	SortedCategories []Category
}
type Category struct {
	Name  string
	Posts []Post
}
type Post struct {
	Title string
	Path  string
}

func init() {
	var readme = &cobra.Command{
		Use: "readme",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("generate README.md file")
			handlerReadmePage()
		},
	}

	rootCmd.AddCommand(readme)
}

func handlerReadmePage() {
	isMdFile := regexp.MustCompile(".md$")
	ignore := regexp.MustCompile("leetcode/*")
	files := Ls("./sources", func(path string) bool {
		return !ignore.Match([]byte(path)) && isMdFile.Match([]byte(path))
	})
	readme := generateReadmeDef(files)
	generateReadmePage(readme)
}

func generateReadmePage(readme *Readme) {
	var readmeStr strings.Builder
	var allPage strings.Builder
	readmeStr.WriteString("<!-- markdownlint-disable -->\n\n")
	readmeStr.WriteString("- [All](sources/generated-sources/all.md)\n")
	allPage.WriteString("<!-- markdownlint-disable -->\n\n# Notes\n")
	for _, c := range readme.SortedCategories {
		path := fmt.Sprintf("sources/generated-sources/%s.md", base64.URLEncoding.EncodeToString([]byte(c.Name)))

		var sb strings.Builder
		sb.WriteString("<!-- markdownlint-disable -->\n\n")
		sb.WriteString(fmt.Sprintf("## %s\n\n", c.Name))
		allPage.WriteString(fmt.Sprintf("\n### %s\n\n", c.Name))
		for _, p := range c.Posts {
			link := fmt.Sprintf("- [%s](/%s)\n", p.Title, strings.ReplaceAll(p.Path, " ", "%20"))
			sb.WriteString(link)
			allPage.WriteString(link)
		}
		MustOverrideFile(path, []byte(sb.String()))

		readmeStr.WriteString(fmt.Sprintf("- [%s](/%s)<sup>%d</sup>\n", c.Name, strings.ReplaceAll(path, " ", "%20"), len(c.Posts)))
	}

	MustOverrideFile("sources/generated-sources/all.md", []byte(allPage.String()))
	MustOverrideFile("README.md", []byte(readmeStr.String()))
}

func generateReadmeDef(files []string) *Readme {
	var readme = &Readme{
		Categories: map[string]Category{},
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		s := bufio.NewScanner(f)
		var (
			title    string
			category string
		)
		for s.Scan() {
			if category != "" && title != "" {
				break
			}
			line := strings.TrimSpace(s.Text())
			if category == "" && strings.HasPrefix(line, "<!-- customize-category:") {
				category = line[len("<!-- customize-category:") : len(line)-3]
			}
			if title == "" && strings.HasPrefix(line, "# ") {
				title = line[len("# "):]
			}
		}
		if category == "" {
			category = "Other"
		}
		category = strings.TrimSpace(category)
		if title == "" {
			idx := strings.Index(f.Name(), "/")
			title = f.Name()[idx+1 : len(f.Name())-3]
		}

		c := readme.Categories[category]
		c.Name = category
		c.Posts = append(c.Posts, Post{
			Title: title,
			Path:  f.Name(),
		})
		readme.Categories[category] = c
	}

	for k, c := range readme.Categories {
		if k == "Other" {
			defer func(c Category) {
				readme.SortedCategories = append(readme.SortedCategories, c)
			}(c)
		} else {
			readme.SortedCategories = append(readme.SortedCategories, c)
		}

		sort.Slice(readme.SortedCategories, func(i, j int) bool {
			return strings.ToLower(readme.SortedCategories[i].Name) < strings.ToLower(readme.SortedCategories[j].Name)
		})
	}
	return readme
}
