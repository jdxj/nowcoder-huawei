package question

import (
	"net"
)

func LegalIP(param string) bool {
	ip := net.ParseIP(param)
	if ip.String() == "<nil>" {
		return false
	}
	return true
}
