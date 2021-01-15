package browser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"../goiv"
)

type Reddit struct {
	Kind string `json:"kind"`
	Data Data   `json:"data"`
}
type Source struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
type Content struct {
	Selftext string `json:"selftext"`
	Title    string `json:"title"`
	Downs    int    `json:"downs"`
	Ups      int    `json:"ups"`
	Score    int    `json:"score"`
	URL      string `json:"url"`
}
type Children struct {
	Kind string  `json:"kind"`
	Data Content `json:"data"`
}
type Data struct {
	Children []Children `json:"children"`
}

//HTTPRequestCustomUserAgent bla bla
func HTTPRequestCustomUserAgent(url, userAgent string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(
			"resp.StatusCode: " +
				strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}

func Parse(url string) {

	var reddit Reddit
	b, err := HTTPRequestCustomUserAgent(url, "Mozilla/5.0 CK={} (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(b), &reddit)
	for i := 0; i < len(reddit.Data.Children); i++ {

		text := string(reddit.Data.Children[i].Data.Selftext)
		img := string(reddit.Data.Children[i].Data.URL)

		fmt.Println("Title :", reddit.Data.Children[i].Data.Title)
		fmt.Println("Upvots :", reddit.Data.Children[i].Data.Ups)
		fmt.Println("Downvotes :", reddit.Data.Children[i].Data.Downs)
		if text == "" {
			if strings.HasSuffix(img, "jpg") {
				fmt.Println("Content : ", img)
				goiv.Viewer(img, 480, 480)
			} else {
				fmt.Println("No images found in the subreddit")
			}
		} else {
			fmt.Println("Content : ", text)
		}
	}
}
