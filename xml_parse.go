package main

import (
  "fmt"
  "xml"
  "http"
  "os"
  "io"
)

type Route struct {
  Tag   string "attr"
  Title string "attr"
}

type RouteList struct {
  XMLName xml.Name "body"
  Routes []Route "route>"
}

func ParseRouteList(x io.Reader) *RouteList {
  r := new(RouteList)
  xml.Unmarshal(x, r)
  return r
}

func main() {
  fmt.Printf("hello, world\n")
  url := "http://webservices.nextbus.com/service/publicXMLFeed?command=routeList&a=sf-muni"
  r, _, err := http.Get(url)

  if err != nil {
    println(err.String())
    os.Exit(1)
  }
  defer r.Body.Close()

  routes := new(RouteList)
  xml.Unmarshal(r.Body, routes)
  fmt.Println(routes)

  // Now we have an instance of RouteList that we can use

}
