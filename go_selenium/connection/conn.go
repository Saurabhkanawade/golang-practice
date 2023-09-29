package conn

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"time"
)

func Connection() {
	// Set up a new Chrome driver instance
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, err := selenium.NewRemote(caps, "http://localhost:9515")
	if err != nil {
		log.Fatalf("Failed to start a new Chrome session: %v", err)
	}
	defer wd.Quit()

	// Navigate to the Google homepage

	if err := wd.Get("https://www.followmyhealth.com/login#!/default"); err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	// Find the login button and click it.

	// Wait for the login form to load.
	time.Sleep(time.Second * 10)

	// Find the username and password input fields and fill them in.
	username, err := wd.FindElements()
	if err != nil {
		fmt.Printf("Failed to find username field: %s\n", err)
		return
	}
	if err := username.SendKeys("your-username"); err != nil {
		fmt.Printf("Failed to enter username: %s\n", err)
		return
	}
	password, err := wd.FindElement(selenium.ByID, "password")
	if err != nil {
		fmt.Printf("Failed to find password field: %s\n", err)
		return
	}
	if err := password.SendKeys("your-password"); err != nil {
		fmt.Printf("Failed to enter password: %s\n", err)
		return
	}

	// Find the submit button and click it.
	submitButton, err := wd.FindElement(selenium.ByID, "submit-button")
	if err != nil {
		fmt.Printf("Failed to find submit button: %s\n", err)
		return
	}
	if err := submitButton.Click(); err != nil {
		fmt.Printf("Failed to click submit button: %s\n", err)
		return
	}

	// Wait for the dashboard to load.
	time.Sleep(time.Second * 5)

	// Print the title of the search results page
	title, err := wd.Title()
	if err != nil {
		log.Fatalf("Failed to get search results title: %v", err)
	}
	fmt.Println("Search results title:", title)
}
