package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

//var wg sync.WaitGroup

func filter(url string, filter string, c chan string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Get(url)

	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic occurred:", err)
				c <- "app_health{instance=\"" + url + "\"}" + "  0"
			}
		}()
		response.Body.Close()

	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		newStr := strings.ToLower(buf.String())
		re := regexp.MustCompile(filter)
		list := re.FindAllString(newStr, -1)
		if len(list) > 0 {
			c <- "app_health{instance=\"" + url + "\"}" + "  1"
		} else {
			c <- "app_health{instance=\"" + url + "\"}" + "  0"
		}
		response.Body.Close()

	}
}

func handleRequests() {
	http.HandleFunc("/metrics", homePage)
	log.Fatal(http.ListenAndServe(":9101", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	c := make(chan string)
	file, err := os.Open("./url.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		go filter(scanner.Text(), `head><meta content`, c)
		fmt.Fprintf(w, <-c+"\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	handleRequests()
}
