package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)
	src, dst := regurlarFiles()
	log.Println("src:", src)
	log.Println("dest:", dst)
	_, err := os.Stat(dst)
	if err == nil {
		log.Fatal("file already existed:", dst)
	}
	f, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := covert(f)
	err = os.WriteFile(dst, data, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func regurlarFiles() (string, string) {
	if len(os.Args) < 3 {
		log.Fatal("must pass the go source and markdown destination directories")
	}
	src := os.Args[1]
	dst := os.Args[2]
	slug := strings.TrimRight(src, "/")
	i := strings.LastIndex(slug, "/")
	if i > 0 {
		slug = slug[i+1:]
	}
	i = strings.Index(slug, ".")
	if i > 0 {
		slug = slug[i+1:]
	}
	dst += "/" + slug + ".md"
	src += "/solution.go"
	return src, dst
}

func covert(r io.Reader) []byte {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	buf := bytes.NewBuffer(nil)
	start := false
	titleWrited := false
	url := ""
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "// http"):
			url = line[3:]
		case strings.TrimSpace(line) == "/*":
			start = true
		case line == "package main" || line == "*/" || strings.HasPrefix(line, "// Created by"):
		default:
			if !start {
				continue
			}
			if !titleWrited {
				titleWrited = true
				title := strings.TrimSpace(line)
				level := ""
				i := strings.LastIndex(title, "(")
				if i > 0 {
					level = title[i:]
					title = title[:i]
					title = strings.TrimSpace(title)
				}
				buf.WriteString("---\n")
				buf.WriteString("title: ")
				buf.WriteString(strconv.Quote(title))
				buf.WriteString("\n---\n\n")
				buf.WriteString(fmt.Sprintf("[%s %s](%s)\n\n", title, level, url))
			} else {
				if line == "// @lc code=begin" {
					buf.WriteString("## 分析\n\n\n```go\n")
				} else if line == "// @lc code=end" {
					buf.WriteString("```\n\n")
				} else {
					buf.WriteString(line + "\n")
				}
			}
		}
	}
	return buf.Bytes()
}
