package bot

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func captureScreenshot(urlStr string) ([]byte, string, error) {

	log.Println("capturing screenshot for ", urlStr)

	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	// create a timeout as a safety net to prevent any infinite wait loops
	ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	var screenshot []byte

	title := ""

	if err := chromedp.Run(
		ctx, elementScreenshot(urlStr, &screenshot, &title),
	); err != nil {
		return nil, "", err
	}

	// Remove newline characters and the plus sign
	trimmed := strings.ReplaceAll(title, "\n", "")
	trimmed = strings.ReplaceAll(trimmed, "+", "")

	return screenshot, trimmed, nil
}

// ======= ChromeDP Actions ========
func elementScreenshot(urlStr string, res *[]byte, title *string) chromedp.Tasks {
	*title = "This is a test"
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		chromedp.Click(`#fight-details--2-0`),
		chromedp.Click(`#filter-damage-done-tab`, chromedp.ByQuery),
		chromedp.WaitVisible(`#main-table-0`, chromedp.ByID),
		chromedp.Evaluate(`document.querySelector("#ap-ea8a4fe5-container").remove();`, nil),
		chromedp.Text(`#filter-fight-boss-text`, title, chromedp.ByID),
		chromedp.Screenshot(`#main-table-0`, res, chromedp.ByID),
	}
}
