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

	// Create a headless Chrome context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),              // Ensure headless mode is enabled
		chromedp.Flag("disable-gpu", false),          // Disable GPU use
		chromedp.Flag("no-sandbox", true),            // Bypass OS security model
		chromedp.Flag("disable-dev-shm-usage", true), // Prevent Chrome from crashing on some systems
		chromedp.WindowSize(1920, 1080),              // init with a desktop view
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
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
		chromedp.WaitReady(`#fight-details-header--2-0`, chromedp.ByQuery),
		chromedp.Click(`#fight-details-header--2-0 .all-fights-entry:last-child`, chromedp.ByQuery),
		chromedp.Click(`#filter-damage-done-tab`, chromedp.ByID),
		chromedp.Text(`#filter-fight-boss-text`, title, chromedp.ByID),
		chromedp.WaitReady(`#main-table-0`, chromedp.ByID),
		chromedp.Evaluate(`document.querySelector("#ap-ea8a4fe5-container")?.remove();`, nil),
		chromedp.Evaluate(`document.querySelector("#corner_ad_video")?.remove()`, nil),
		chromedp.Screenshot(`#main-table-0`, res, chromedp.ByID),
	}
}
