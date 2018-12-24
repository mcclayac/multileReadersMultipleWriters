package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*

Anthonys-MacBook-Pro:go mcclayac$ godoc net/http Get
func Get(url string) (resp *Response, err error)
    Get issues a GET to the specified URL. If the response is one of the
    following redirect codes, Get follows the redirect, up to a maximum of
    10 redirects:

	301 (Moved Permanently)
	302 (Found)
	303 (See Other)
	307 (Temporary Redirect)
	308 (Permanent Redirect)

    An error is returned if there were too many redirects or if there was an
    HTTP protocol error. A non-2xx response doesn't cause an error. Any
    returned error will be of type *url.Error. The url.Error value's Timeout
    method will report true if request timed out or was canceled.

    When err is nil, resp always contains a non-nil resp.Body. Caller should
    close resp.Body when done reading from it.

    Get is a wrapper around DefaultClient.Get.

    To make a request with custom headers, use NewRequest and
    DefaultClient.Do.


Anthonys-MacBook-Pro:go mcclayac$ godoc io/ioutil ReadAll
func ReadAll(r io.Reader) ([]byte, error)
    ReadAll reads from r until an error or EOF and returns the data it read.
    A successful call returns err == nil, not err == EOF. Because ReadAll is
    defined to read from src until EOF, it does not treat an EOF from Read
    as an error to be reported.




 */

func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	//bodyString := string(body)
	//fmt.Printf(bodyString , "\n\n")
	return len(body), nil
}

/*func getter(url string, size chan int) {
	length, err := getPage(url)
	if err == nil {
		size <- length
	}

}*/

func getter(url string, size chan string) {
	length, err := getPage(url)
	if err == nil {
		size <- fmt.Sprintf("%s has length %d", url, length)
	}
}


// getter refactored to worker
func worker(urlch chan string, sizeCh chan string, id int) {
	for {
		url := <-urlch
		length, err := getPage(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("id:%d .. %s has length %d", id, url, length)
		} else {
			sizeCh <- fmt.Sprintf("id:%d Error getting %s: %s", id, url, err)
		}
	}
}


func generator(url string, urlch chan string) {

	urlch <- url

}


func main() {

	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
		"http://www.bbc.co.uk",
		"http://www.cnn.com",
		"https://learning.oreilly.com/videos/introduction-to-go/9781491913871/9781491913871-video191857",
		"https://learning.oreilly.com/videos/introduction-to-go/9781491913871/9781491913871-video191856",
		"https://learning.oreilly.com/videos/introduction-to-go/9781491913871/9781491913871-video191855",
		"https://learning.oreilly.com/videos/introduction-to-go/9781491913871/9781491913871-video191854?autoplay=false",
		"https://github.com/mcclayac/goConcurrency/blob/master/main.go",
		"https://learning.oreilly.com/videos/introduction-to-go/9781491913871/9781491913871-video191854",
		"https://us-west-1.signin.aws.amazon.com/oauth?response_type=code&client_id=arn%3Aaws%3Aiam%3A%3A015428540659%3Auser%2Fec2&redirect_uri=https%3A%2F%2Fus-west-1.console.aws.amazon.com%2Fec2%2Fv2%2Fhome%3Fregion%3Dus-west-1%26state%3DhashArgs%2523Instances%253A%26isauthcode%3Dtrue&forceMobileLayout=0&forceMobileApp=0",
		"https://grafeas.io/",
		"https://www.open-scap.org/",
		"https://anchore.io/"}
	//url := "http://www.google.com"

	//urls = urls

	//size := make(chan int)
	//size := make(chan string)
	urlCh := make(chan string)
	sizeCh := make(chan string)

	for i := 0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	for _, url := range urls {
		go generator(url, urlCh)
		//urlCh <- url
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}

	//urlCh <- "http://www.oreilly.com"

	//fmt.Printf("%s\n", <-sizeCh)



	/*for _, url := range urls {
		//pageLen, err := getPage(url)
		//if err != nil {
		//	os.Exit(1)
		//}
		go getter(url, size)
		//fmt.Printf("Len of %s = %d\n", url, pageLen)
	}

	for i := 0 ; i < len(urls); i++ {
		fmt.Printf("%s\n", <-size)

	}*/
}

/*
Anthonys-MacBook-Pro:go mcclayac$ godoc fmt Sprintf
func Sprintf(format string, a ...interface{}) string
    Sprintf formats according to a format specifier and returns the
    resulting string.


 */