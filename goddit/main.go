package main

import (
	"fmt"
	"log"

	"./browser"
	"moul.io/banner"

	"github.com/dixonwille/wmenu"
)

func main() {
	fmt.Println(banner.Inline("goddit"))
	questions := question()
	menu := wmenu.NewMenu("Choose a category from above")
	menu.Action(func(opts []wmenu.Opt) error {
		url := "https://www.reddit.com/r/" + questions + "/" + opts[0].Text + ".json?limit=999999"
		fmt.Println("\nHit CTRL-C on the terminal to escape")
		browser.Parse(url)
		return nil
	})
	menu.Option("hot", nil, true, nil)
	menu.Option("rising", nil, false, nil)
	menu.Option("top", nil, false, nil)
	menu.Option("new", nil, false, nil)
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
func question() string {
	var sub string
	fmt.Println("Enter the Subreddit to browse :")
	fmt.Scanln(&sub)
	return sub
}
