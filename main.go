package main                                        

import (    

	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"                                        
)                                                   

func main() {                                       
	const (
		seleniumPath    = "vendor/selenium-server-standalone-3.4.jar"
		geckoDriverPath = "some_path/geckodriver-v0.18.0-linux64"
		port		= 8080
	)

	opts := []selenium.ServiceOption {
//		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.

	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	// connect webdriver instance running local
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	//Navigate to simple blacktop interfacer
	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
		panic (err)
	}

	//Get a reference to textbox containing code
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
	if err != nil {
		panic(err)
	}

	//Remove the boilerplate crap not so trusty in testbox
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	err = elem.SendKeys(`
		package main
		import "fmt"

		func main() {
			fmt.Println("Hello Webdriver !!!\n")
		}
	`)

	if err != nil {
		panic(err)
	}

	//Run button pressed
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}

	if err := btn.Click(); err != nil {
		panic(err)
	}

	//Wait introduced to the program finish running and get output
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}

		if output != "Waiting for local destination (not server)..." {
			break
		}

		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	//Example Output:
	//Hello Webdriver !!!
	//
	//Program COmpleted
}
