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
		// Capture Selector
		// Click,
		// etc
		chromedp.Screenshot(`#content`, res, chromedp.NodeVisible),
	}
}
