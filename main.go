package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func GetHttpHTML(url string, selector string) {
	userDataDir := "/Users/alichen/Library/Application Support/Google/Chrome"
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserDataDir(userDataDir),
		chromedp.Flag("headless", false), // headless mode
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string
	var currentURL string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.OuterHTML("html", &htmlContent),
		chromedp.Click(selector, chromedp.NodeVisible),
		chromedp.WaitNotPresent(selector, chromedp.ByQuery),
		chromedp.Location(&currentURL),
	)
	if err != nil {
		log.Fatal(err)
	}

	var ticketNumCss string = "div > span.ticket-quantity.ng-scope > input"
	var nextStep string = "div.form-actions.plain.align-center.register-new-next-button-area > button"
	// Wait for the new page to load
	err = chromedp.Run(ctx,
		chromedp.Navigate(currentURL),
		chromedp.WaitVisible(ticketNumCss, chromedp.ByQuery),
		chromedp.WaitReady(ticketNumCss, chromedp.ByQuery),
		chromedp.SendKeys(ticketNumCss, "1", chromedp.ByQuery),
		chromedp.WaitVisible("#person_agree_terms", chromedp.ByID), // 等待 #person_agree_terms 元素可见
		chromedp.Click("#person_agree_terms", chromedp.ByID),
		chromedp.WaitVisible(nextStep, chromedp.ByQuery),
		chromedp.OuterHTML("html", &nextStep),
		chromedp.Click(nextStep, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	// Example of extracting data
	doc.Find(nextStep).Each(func(i int, s *goquery.Selection) {
		log.Println(s.Text())
	})
}

func main() {
	GetHttpHTML("https://bigzero.kktix.cc/events/0f96c00c", "body > div.outer-wrapper > div.content-wrapper > div > div.tickets > a")
}
