package main

import (
    "fmt"
    "os"
    "bufio"
    "bytes"
    "regexp"
    "log"
    "net/http"
    "strings"
)

func filter( url string , filter string, c chan string ) {
    response, _ := http.Get(url)
    buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    newStr := strings.ToLower(buf.String())
    re := regexp.MustCompile(filter)
    list := re.FindAllString(newStr,-1)
    if len(list) > 0 {
        c <- "unifonic_health{instance=\"" + url + "\"}" + "  1\n"
    } else {
        c <- "unifonic_health{instance=\"" + url + "\"}" + "  0\n"
    }
  }

  func handleRequests() {
    http.HandleFunc("/metrics", homePage)
    log.Fatal(http.ListenAndServe(":9101", nil))
}


func homePage(w http.ResponseWriter, r *http.Request) {
    c := make(chan string,2)
    file, err := os.Open("./url.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        go filter(scanner.Text(),`^{"status":(\s+)?"up"`,c)
        fmt.Fprintf(w,<-c)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

  }

func main() {
    handleRequests()
}
