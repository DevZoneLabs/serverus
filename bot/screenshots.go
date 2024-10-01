package bot

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func captureScreenshot(urlStr string) ([]byte, string, error) {
	// Create a headless Chrome context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),              // Ensure headless mode is enabled
		chromedp.Flag("disable-gpu", true),           // Disable GPU use
		chromedp.Flag("no-sandbox", true),            // Bypass OS security model
		chromedp.Flag("disable-dev-shm-usage", true), // Prevent Chrome from crashing on some systems
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx,
		chromedp.WithLogf(log.Printf))
	defer cancel()

	// create a timeout as a safety net to prevent any infinite wait loops
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
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
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		chromedp.WaitReady(`#fight-details--2-0 .last-pull-label`, chromedp.ByQuery),
		chromedp.Click(`#fight-details--2-0 .last-pull-label`, chromedp.ByQuery),
		chromedp.WaitReady(`#filter-damage-done-tab`, chromedp.ByQuery),
		chromedp.Click(`#filter-damage-done-tab`, chromedp.ByQuery),
		chromedp.WaitVisible(`#main-table-0`, chromedp.ByID),
		chromedp.Evaluate(`document.querySelector("#ap-ea8a4fe5-container").remove();`, nil),
		chromedp.Text(`#filter-fight-boss-text`, title, chromedp.ByID),
		chromedp.Screenshot(`#main-table-0`, res, chromedp.ByID),
	}
}
