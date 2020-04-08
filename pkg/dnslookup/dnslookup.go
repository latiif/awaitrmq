package dnslookup

import (
	"net"
)

// DNSLookup returns the result of a dns lookup on the given url
func DNSLookup(url string) bool {
	_, err := net.LookupIP(url)
	return err != nil
}
