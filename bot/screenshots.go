package bot

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func captureScreenshot(urlStr string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	// create a timeout as a safety net to prevent any infinite wait loops
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var screenshot []byte
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// Disable headless mode
		chromedp.Flag("headless", false),
		// Make sure the browser window is maximized
		chromedp.Flag("start-maximized", true),
	)

	if err := chromedp.Run(
		ctx, elementScreenshot(urlStr, &screenshot),
	); err != nil {
		return nil, err
	}

	return screenshot, nil
}

// ======= ChromeDP Actions ========
func elementScreenshot(urlStr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		chromedp.WaitVisible(`#fight-details--2-0 .last-pull-label`, chromedp.ByQuery),
		chromedp.Click(`#fight-details--2-0 .last-pull-label`, chromedp.ByQuery),
		chromedp.WaitVisible(`#filter-damage-done-tab`, chromedp.ByID),
		chromedp.Click(`#filter-damage-done-tab`, chromedp.ByQuery),
		chromedp.WaitVisible(`#main-table-0`, chromedp.ByID),
		chromedp.Screenshot(`#main-table-0`, res, chromedp.ByID),
	}
}
