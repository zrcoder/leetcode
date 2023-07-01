package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Meta struct {
	FrontendID string `json:"frontendQuestionId"`
	Title      string `json:"title"`
	Referer    string
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
	PaidOnly   bool   `json:"paidOnly"`
}

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) < 3 {
		log.Fatal("must pass the go source and markdown destination directories")
	}

	src := os.Args[1]
	dst := os.Args[2]
	data, err := os.ReadFile(filepath.Join(src, "question.json"))
	fatalIfError(err)
	question := &Meta{}
	err = json.Unmarshal(data, question)
	fatalIfError(err)
	mdData, err := os.ReadFile(filepath.Join(src, "question.md"))
	fatalIfError(err)
	codeData, err := os.ReadFile(filepath.Join(src, "solution.go"))
	fatalIfError(err)

	dst = filepath.Join(dst, question.TitleSlug+".md")
	_, err = os.Stat(dst)
	if err == nil {
		log.Fatal("already exist: ", dst)
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString("---\n")
	buf.WriteString(fmt.Sprintf("title: %s. %s\n", question.FrontendID, question.Title))
	buf.WriteString(fmt.Sprintf("date: %s\n", time.Now().Format(time.RFC3339)))
	buf.WriteString("---\n\n")
	buf.Write(mdData)
	buf.WriteString("## 分析\n\n")
	code, note := parseCodeAndNotes(codeData)
	if note != nil {
		buf.Write(note)
		buf.WriteString("\n")
	}
	buf.WriteString("```go\n")
	buf.Write(code)
	buf.WriteString("```\n")
	testData, err := os.ReadFile(filepath.Join(src, "solution_test.go"))
	fatalIfError(err)
	if !bytes.Contains(testData, []byte("TODO")) {
		buf.WriteString("\n测试用例:\n\n```go\n\n")
		i := bytes.Index(testData, []byte("func "))
		if i != -1 {
			buf.Write(testData[i:])
			buf.WriteString("\n```\n\n")
		}
	}

	err = os.WriteFile(dst, buf.Bytes(), 0640)
	fatalIfError(err)
}

func parseCodeAndNotes(data []byte) (code, note []byte) {
	const (
		noteStart = "/* @note start\n"
		noteEnd   = "@note end */\n"
		codeStart = "// @submit start\n"
		codeEnd   = "// @submit end\n"
	)
	i := bytes.Index(data, []byte(noteStart))
	if i != -1 {
		data = data[i+len(noteStart):]
		i = bytes.Index(data, []byte(noteEnd))
		if i != -1 {
			note = data[:i]
			data = data[i+len(noteEnd):]
		}
	}
	i = bytes.Index(data, []byte(codeStart))
	if i == -1 {
		code = data
		return
	}
	data = data[i+len(codeStart):]
	i = bytes.LastIndex(data, []byte(codeEnd))
	if i == -1 {
		code = data
		return
	}
	code = data[:i]
	return
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
