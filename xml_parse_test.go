package main
import (
  "testing"
)

const routeXml = `<?xml version="1.0" encoding="utf-8" ?> 
<body copyright="All data copyright San Francisco Muni 2011."> 
<route tag="F" title="F-Market &amp; Wharves"/> 
<route tag="J" title="J-Church"/> 
</body>`

func TestParseRouteListXml(t *testing.T) {
  r := ParseRouteList(routeXml)

  if len(r.Routes) != 2 {
    t.Errorf("RouteList should have 2 elements. Seen: %v", len(r.Routes))
  }

  if r.Routes[0].Tag != "F" && r.Routes[0].Tag != "F-Market & Wharves" {
    t.Errorf("First element of RouteList should equal F-Market & Wharves. Received: %v", r.Routes[0])
  }

  if r.Routes[1].Tag != "J" && r.Routes[0].Tag != "J-Church" {
    t.Errorf("First element of RouteList should equal J-Church. Received: %v", r.Routes[1])
  }
}
