package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var leetcode = &cobra.Command{
	Use: "leetcode",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("generate LeetCode.md file")
		handleLeetCodePage()
	},
}

func init() {
	rootCmd.AddCommand(leetcode)
}

const OUTPUT = "sources/LeetCode.md"

func handleLeetCodePage() {
	data, titleMap := parse("./sources/leetcode/")
	tags := []string{}
	for k := range data {
		tags = append(tags, k)
	}

	// sort by name
	sort.Slice(tags, func(i, j int) bool {
		return tags[i] < tags[j]
	})

	content := `<!-- markdownlint-disable -->
<!-- customize-category:LeetCode-->
`

	for _, tag := range tags {
		content += fmt.Sprintf("\n- [%s](#%s)", tag, tag)
	}
	content += "\n"

	for _, tag := range tags {
		content += fmt.Sprintf("\n## %s\n\n", tag)

		for _, name := range data[tag] {
			content += fmt.Sprintf("- [%s](/sources/leetcode/%s)\n", titleMap[name], name)
		}
	}

	MustOverrideFile(OUTPUT, []byte(content))
}

const TAG = "<!-- customize-tags:"

func parse(dir string) (map[string][]string, map[string]string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	result := map[string][]string{}
	titleMap := map[string]string{}

	for _, file := range files {
		if file.IsDir() || path.Ext(file.Name()) != ".md" {
			continue
		}
		f, err := os.Open(filepath.Join(dir, file.Name()))

		if err != nil {
			log.Fatalln(err)
		}

		s := bufio.NewScanner(f)
		title := ""
		tags := []string{}
		for s.Scan() {
			line := strings.TrimSpace(s.Text())

			if strings.Contains(line, TAG) {
				line = strings.ReplaceAll(line, TAG, "")
				line = strings.ReplaceAll(line, "-->", "")
				tags = strings.Split(line, ",")
			}
			if strings.HasPrefix(line, "# ") && title == "" {
				title = strings.ReplaceAll(line, "# ", "")
			}

			if title != "" && len(tags) > 0 {
				break
			}
		}

		if title == "" {
			title = file.Name()
		}

		for _, tag := range tags {
			tag = strings.TrimSpace(tag)
			fileList := append(result[tag], file.Name())
			result[tag] = fileList
		}

		titleMap[file.Name()] = title
		f.Close()
	}
	return result, titleMap
}
