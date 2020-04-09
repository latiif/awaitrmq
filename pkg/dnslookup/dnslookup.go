package dnslookup

import (
	"net"
)

// DNSLookup returns the result of a dns lookup on the given url
func DNSLookup(url string) bool {
	ip, err := net.LookupIP(url)
	if err == nil && len(ip) != 0 {
		return true
	}
	return false
}
