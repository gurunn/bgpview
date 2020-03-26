[![GoDoc](https://godoc.org/github.com/gurunn/bgpview?status.svg)](http://godoc.org/github.com/gurunn/bgpview)

BGPView client library 
---------------------
BGPView is a simple API allowing consumers to view all sort of analytics data about the current state and structure of the internet.
Link to BGPView: 

* [official site](https://bgpview.docs.apiary.io/)

Usage example
------------------------------------------------

```sh
go get -u github.com/gurunn/bgpview
```

Simple example using library in cmd/bgpview/main.go

```go
package main

import (
	"fmt"
	"github.com/gurunn/bgpview"
)

func main() {
	asn, err := bgpview.GetASN(61138)
	asnPrefixes, err := bgpview.GetASNPrefixes(61138)
	asnPeers, err := bgpview.GetASNPeers(61138)
	asnUpstreams, err := bgpview.GetASNUpstreams(62240)
	asnDownstreams, err := bgpview.GetASNDownstreams(55470)
	asnIXs, err := bgpview.GetASNIXs(18119)
	prefix, err := bgpview.GetPrefix("192.209.63.0", 24)
	ip, err := bgpview.GetIP("2a05:dfc7:60::")
	ix, err := bgpview.GetIX(450)
	search, err := bgpview.GetSearch("digitalocean")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(asn)
	fmt.Println(asnPrefixes)
	fmt.Println(asnPeers)
	fmt.Println(asnUpstreams)
	fmt.Println(asnDownstreams)
	fmt.Println(asnIXs)
	fmt.Println(prefix)
	fmt.Println(ip)
	fmt.Println(ix)
	fmt.Println(search)
}
```