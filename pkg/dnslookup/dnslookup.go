package dnslookup

import (
	"fmt"
	"log"
	"net"
	Url "net/url"
)

// DNSLookup returns the result of a dns lookup on the given url
func DNSLookup(url string) bool {
	u, err := Url.Parse(url)
	if err != nil {
		log.Fatal(fmt.Errorf("Invalid URL: %s. %v", url, err))
	}
	host := u.Hostname()
	ip, err := net.LookupIP(host)
	if err == nil && len(ip) != 0 {
		return true
	}
	return false
}
