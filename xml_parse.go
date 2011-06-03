package main

import (
  "fmt"
  "strings"
  "xml"
)

type Route struct {
  Tag   string "attr"
  Title string "attr"
}

type RouteList struct {
  XMLName xml.Name "body"
  Routes []Route "route>"
}

func ParseRouteList(x string) *RouteList {
  r := new(RouteList)
  reader := strings.NewReader(x)
  xml.Unmarshal(reader, r)
  return r
}

func main() {
  fmt.Printf("hello, world\n")
}
