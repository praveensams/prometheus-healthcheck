package main

import (
    "fmt"
    "os"
    "bufio"
    "bytes"
    "regexp"
    "log"
    "net/http"
)

func filter( url string , filter string) string {
    response, _ := http.Get(url)
    buf := new(bytes.Buffer)
    buf.ReadFrom(response.Body)
    newStr := buf.String()
    re := regexp.MustCompile(filter)
    list := re.FindAllString(newStr,-1)
    if len(list) > 0 {
        return "unifonic_health{instance=\"" + url + "\"}" + "  1"
    } else {
        return "unifonic_health{instance=\"" + url + "\"}" + "  0"
    }
  }



func main() {
    prints:=[]string{}
    file, err := os.Open("./url.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        prints=append(prints,filter(scanner.Text(),`^{"status":"UP"`))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    i := 0
    for i < len(prints) {
        fmt.Println(prints[i])
        i=i+1
        }

  }

