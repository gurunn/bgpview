package bgpview

import (
	"testing"
)

func TestGetASN(t *testing.T) {
	asNumber := 61138
	asn, err := GetASN(asNumber)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if asn == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetASNPrefixes(t *testing.T) {
	asNumber := 61138
	asnPrefixes, err := GetASNPrefixes(asNumber)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if asnPrefixes == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetASNPeers(t *testing.T) {
	asNumber := 61138
	asnPeers, err := GetASNPeers(asNumber)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if asnPeers == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetASNUpstreams(t *testing.T) {
	asNumber := 62240
	asnUpstreams, err := GetASNUpstreams(asNumber)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if asnUpstreams == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetASNDownstreams(t *testing.T) {
	asNumber := 55470
	asnDownstreams, err := GetASNDownstreams(asNumber)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if asnDownstreams == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetASNIXs(t *testing.T) {
	asNumber := 18119
	asnIXs, err := GetASNIXs(asNumber)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if asnIXs == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetPrefix(t *testing.T) {
	ipAddress := "192.209.63.0"
	cidr := 24
	prefix, err := GetPrefix(ipAddress, cidr)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if prefix == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetIP(t *testing.T) {
	ipAddress := "2a05:dfc7:60::"
	ip, err := GetIP(ipAddress)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if ip == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetIX(t *testing.T) {
	ixId := 492
	ix, err := GetIX(ixId)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if ix == nil {
		t.Error("service return nil data")
		return
	}
}

func TestGetSearch(t *testing.T) {
	queryTerm := "digitalocean"
	search, err := GetSearch(queryTerm)
	if err != nil {
		t.Error("error requesting report: ", err)
		return
	}
	if search == nil {
		t.Error("service return nil data")
		return
	}
}
