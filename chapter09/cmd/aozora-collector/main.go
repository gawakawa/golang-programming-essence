package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	SiteURL  string
	ZipURL   string
}

func findAuthorAndZIP(siteURL string) (string, string) {
	log.Println("query", siteURL)
	res, err := http.Get(siteURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf(
			"status code error: %d %s",
			res.StatusCode,
			res.Status,
		)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", ""
	}

	author := doc.Find("table[summary=作家データ] tr:nth-child(1) td:nth-child(2)").
		Text()

	zipURL := ""
	doc.Find("table.download a").
		Each(func(n int, s *goquery.Selection) {
			href := s.AttrOr("href", "")
			if strings.HasSuffix(href, ".zip") {
				zipURL = href
			}
		})

	if zipURL == "" {
		return author, ""
	}
	if strings.HasPrefix(zipURL, "http://") ||
		strings.HasPrefix(zipURL, "https://") {
		return author, ""
	}

	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}
	u.Path = path.Join(path.Dir(u.Path), zipURL)
	return author, u.String()
}

func findEntries(siteURL string) ([]Entry, error) {
	res, err := http.Get(siteURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf(
			"status code error: %d %s",
			res.StatusCode,
			res.Status,
		)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	entries := []Entry{}
	doc.Find("ol li a").Each(func(i int, s *goquery.Selection) {
		token := pat.FindStringSubmatch(s.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		title := s.Text()
		pageURL := fmt.Sprintf(
			"https://www.aozora.gr.jp/cards/%s/card%s.html",
			token[1],
			token[2],
		)
		author, zipURL := findAuthorAndZIP(pageURL)
		if zipURL != "" {
			entries = append(entries, Entry{
				AuthorID: token[1],
				Author:   author,
				TitleID:  token[2],
				Title:    title,
				SiteURL:  siteURL,
				ZipURL:   zipURL,
			})
		}
	})

	return entries, nil
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Title, entry.ZipURL)
	}
}
