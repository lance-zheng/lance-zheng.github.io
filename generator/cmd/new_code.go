package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/spf13/cobra"
)

var newCode = &cobra.Command{
	Use: "newcode",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("请指定题目")
		}

		titleSlug := args[0]
		output := fmt.Sprintf("./sources/leetcode/%s.md", titleSlug)
		if _, err := os.Stat(output); err == nil {
			log.Fatalln("文件已存在")
		}
		handleNewCode(titleSlug, output)
	},
}

func init() {
	rootCmd.AddCommand(newCode)
}

type Response struct {
	Data struct {
		Question struct {
			QuestionId         string `json:"questionId"`
			QuestionFrontendId string `json:"questionFrontendId"`
			TitleSlug          string `json:"titleSlug"`
			TranslatedTitle    string `json:"translatedTitle"`
			TranslatedContent  string `json:"translatedContent"`
			TopicTags          []struct {
				TranslatedName string `json:"translatedName"`
			} `json:"topicTags"`
		} `json:"question"`
	} `json:"data"`
}

func handleNewCode(titleSlug, output string) {
	// load question
	resp := loadQuestionByTitleSlugMustExists(titleSlug)

	// join tags
	var topicNames []string
	for _, tag := range resp.Data.Question.TopicTags {
		topicNames = append(topicNames, tag.TranslatedName)
	}
	tags := strings.Join(topicNames, ",")

	markdown, err := converterHtmlToMarkdown(resp.Data.Question.TranslatedContent)
	if err != nil {
		log.Fatalln("html 转换 markdown 失败")
	}

	data := map[string]interface{}{
		"Seq":       resp.Data.Question.QuestionFrontendId,
		"Title":     resp.Data.Question.TranslatedTitle,
		"TitleSlug": titleSlug,
		"Question":  markdown,
		"Tags":      tags,
	}

	bytes := renderData(data)

	MustOverrideFile(output, bytes)
}

const questionTmpl = `<!-- markdownlint-disable -->
<!-- customize-tags:{{.Tags}} -->

# {{.Seq}}. {{.Title}}

> [题目链接](https://leetcode.cn/problems/{{.TitleSlug}}/)

{{.Question}}

<!-- markdownlint-restore -->
<!--------------------------------->
<!-- generate by new_leetcode.go -->
`

func renderData(data map[string]interface{}) []byte {
	t := template.Must(template.New("question").Parse(questionTmpl))
	buf := &bytes.Buffer{}
	t.Execute(buf, data)
	return buf.Bytes()
}

func converterHtmlToMarkdown(html string) (string, error) {
	converter := md.NewConverter("", true, nil)
	converter.Use(plugin.GitHubFlavored())

	return converter.ConvertString(html)
}

const graphqlTmpl = `
{
    "operationName": "questionData",
    "variables": {
        "titleSlug": "%s"
    },
    "query": "query questionData($titleSlug: String!){question(titleSlug: $titleSlug) {questionId questionFrontendId titleSlug translatedTitle translatedContent topicTags { translatedName  }}}"
}
`

func loadQuestionByTitleSlugMustExists(titleSlug string) *Response {
	query := fmt.Sprintf(graphqlTmpl, titleSlug)
	req, err := http.NewRequest("POST", "https://leetcode.cn/graphql/", bytes.NewBuffer([]byte(query)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		panic(err)
	}

	if res.Data.Question.QuestionId == "" {
		log.Fatalln("题目不存在")
	}

	return &res
}
