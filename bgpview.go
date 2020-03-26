package bgpview

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func query(url string) ([]byte, error) {
	conn := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := conn.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetASN return data about ASN request
func GetASN(asNumber int) (*ASN, error) {
	urlStr := URL + "asn/" + strconv.Itoa(asNumber)
	respBody, err := query(urlStr)
	var asn ASN
	err = json.Unmarshal(respBody, &asn)
	if err != nil {
		return nil, err
	}
	return &asn, nil
}

// GetASNPrefixes return data about ASNPrefixes request
func GetASNPrefixes(asNumber int) (*ASNPrefixes, error) {
	urlStr := URL + "asn/" + strconv.Itoa(asNumber) + "/prefixes"
	respBody, err := query(urlStr)
	var asnPrefixes ASNPrefixes
	err = json.Unmarshal(respBody, &asnPrefixes)
	if err != nil {
		return nil, err
	}
	return &asnPrefixes, nil
}

// GetASNPeers return data about ASNPeers request
func GetASNPeers(asNumber int) (*ASNPeers, error) {
	urlStr := URL + "asn/" + strconv.Itoa(asNumber) + "/peers"
	respBody, err := query(urlStr)
	var asnPeers ASNPeers
	err = json.Unmarshal(respBody, &asnPeers)
	if err != nil {
		return nil, err
	}
	return &asnPeers, nil
}

// GetASNUpstreams return data about ASNUpstreams request
func GetASNUpstreams(asNumber int) (*ASNUpstreams, error) {
	urlStr := URL + "asn/" + strconv.Itoa(asNumber) + "/upstreams"
	respBody, err := query(urlStr)
	var asnUpstreams ASNUpstreams
	err = json.Unmarshal(respBody, &asnUpstreams)
	if err != nil {
		return nil, err
	}
	return &asnUpstreams, nil
}

// GetASNDownstreams return data about ASNDownstreams request
func GetASNDownstreams(asNumber int) (*ASNDownstreams, error) {
	urlStr := URL + "asn/" + strconv.Itoa(asNumber) + "/downstreams"
	respBody, err := query(urlStr)
	var asnDownstreams ASNDownstreams
	err = json.Unmarshal(respBody, &asnDownstreams)
	if err != nil {
		return nil, err
	}
	return &asnDownstreams, nil
}

// GetASNIXs return data about ASNIXs request
func GetASNIXs(asNumber int) (*ASNIXs, error) {
	urlStr := URL + "asn/" + strconv.Itoa(asNumber) + "/ixs"
	respBody, err := query(urlStr)
	var asnIXs ASNIXs
	err = json.Unmarshal(respBody, &asnIXs)
	if err != nil {
		return nil, err
	}
	return &asnIXs, nil
}

// GetPrefix return data about Prefix request
func GetPrefix(ipAddress string, cidr int) (*Prefix, error) {
	urlStr := URL + "prefix/" + ipAddress + "/" + strconv.Itoa(cidr)
	respBody, err := query(urlStr)
	var prefix Prefix
	err = json.Unmarshal(respBody, &prefix)
	if err != nil {
		return nil, err
	}
	return &prefix, nil
}

// GetIP return data about IP request
func GetIP(ipAddress string) (*IP, error) {
	urlStr := URL + "ip/" + ipAddress
	respBody, err := query(urlStr)
	var ip IP
	err = json.Unmarshal(respBody, &ip)
	if err != nil {
		return nil, err
	}
	return &ip, nil
}

// GetIX return data about IX request
func GetIX(ixId int) (*IX, error) {
	urlStr := URL + "ix/" + strconv.Itoa(ixId)
	respBody, err := query(urlStr)
	var ix IX
	err = json.Unmarshal(respBody, &ix)
	if err != nil {
		return nil, err
	}
	return &ix, nil
}

// GetSearch return data about Search request
func GetSearch(queryTerm string) (*Search, error) {
	urlStr := URL + "search?query_term=" + queryTerm
	respBody, err := query(urlStr)
	var search Search
	err = json.Unmarshal(respBody, &search)
	if err != nil {
		return nil, err
	}
	return &search, nil
}
