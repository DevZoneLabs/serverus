package logScreenShots

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Launch the browser in non-headless mode for debugging
	url := launcher.New().Headless(false).MustLaunch()

	// Connect to the browser
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.Close()

	// Create a new page and navigate to the desired URL
	page := browser.MustPage("https://www.warcraftlogs.com/reports/wQFW1d3tfnmDzKgv")

	// Wait for the page to fully load
	page.MustWaitLoad()

	// Add a delay to ensure elements are rendered (debug purposes)
	time.Sleep(2 * time.Second)

	// Click on the button inside the div with ID 'fight-details--2-0'
	button := page.MustElement("#fight-details--2-0 .last-pull-label")
	if button != nil {
		log.Println("Button found inside the specified div, attempting to click it...")
		button.MustClick()
	} else {
		log.Fatal("Button not found inside the specified div, check the selector!")
	}

	// Add a delay to ensure that the page finishes loading after the click
	time.Sleep(2 * time.Second)

	// Take a screenshot of a specific element (e.g., '#summary-content')
	element := page.MustElement("#summary-content")
	if element != nil {
		log.Println("Element found, taking screenshot...")
		screenshotBytes := element.MustScreenshot()

		// Save the screenshot to a file
		if err := os.WriteFile("screenshot.png", screenshotBytes, 0644); err != nil {
			log.Fatalf("Error saving screenshot: %v", err)
		}

		// Decode the image and print the image data
		img, err := png.Decode(bytes.NewReader(screenshotBytes))
		if err != nil {
			log.Fatalf("Error decoding image: %v", err)
		}

		// Print image dimensions
		fmt.Printf("Image Width: %d\n", img.Bounds().Dx())
		fmt.Printf("Image Height: %d\n", img.Bounds().Dy())

		log.Println("Screenshot saved as 'screenshot.png'")
	} else {
		log.Fatal("Element not found, check the selector!")
	}

	// Close the browser
	browser.MustClose()
}
