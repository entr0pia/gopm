package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"runtime"
	"strings"
)

type urlBase struct {
	url  string
	name string
	page int
}

func (t *urlBase) toString() string {
	return fmt.Sprintf("%s?q=%s&page=%d", t.url, t.name, t.page)
}

func (t *urlBase) query() error {
	resp, err := http.Get(t.toString())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 400 {
		return errors.New("Search page number too large.")
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	resultList := doc.Find("div.LegacySearchSnippet")
	resultList.Each(func(_ int, selection *goquery.Selection) {
		pkgName := strings.TrimSpace(selection.Find("[data-test-id='snippet-title']").Text())
		pkgVer := strings.TrimSpace(selection.Find("[data-test-id='snippet-version']").Text())
		pkgDate := strings.TrimSpace(selection.Find("[data-test-id='snippet-published']").Text())
		pkgPop := strings.TrimSpace(selection.Find("[data-test-id='snippet-importedby']").Text())
		pkgLis := strings.TrimSpace(selection.Find("[data-test-id='snippet-license']").Text())
		pkgSynop := selection.Find("[data-test-id='snippet-synopsis']").Text()
		if runtime.GOOS != "windows" {
			pkgName = "\033[32m" + pkgName + "\033[0m"
		}
		fmt.Printf("%s # %s; %s; %s; %s\n  %s\n", pkgName, pkgVer, pkgDate, pkgPop, pkgLis,
			pkgSynop)
		if pkgSynop != "" {
			fmt.Println()
		}
	})
	return nil
}
