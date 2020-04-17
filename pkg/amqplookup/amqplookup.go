package amqplookup

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

// AMQPLookup returns the result of an amqp lookup on the given url
func AMQPLookup(url string, lookupTimeout time.Duration) bool {
	conn, err := amqp.DialConfig(fmt.Sprintf("amqp://%s", url), amqp.Config{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, lookupTimeout)
		},
	})
	if err == nil && conn != nil {
		defer conn.Close()
		return true
	}
	if err != nil && conn != nil {
		errMessage := fmt.Sprint(err)
		if strings.Contains(errMessage, `Exception (403) Reason: "username or password not allowed`) {
			return true
		}
	}
	return false
}
