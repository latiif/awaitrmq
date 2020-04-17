package dnslookup

import (
	"net"
	"strings"
)

// DNSLookup returns the result of a dns lookup on the given url
func DNSLookup(url string) bool {
	//TODO find a better way to extract host
	host := strings.Split(url, ":")[0] //remove port number
	ip, err := net.LookupIP(host)
	if err == nil && len(ip) != 0 {
		return true
	}
	return false
}
