package cidr

import "net"

// Parse is a convenience function that returns only the IPNet
// This function ignores errors since it is primarily a test helper, the result could be nil
func Parse(s string) *net.IPNet {
	_, c, _ := net.ParseCIDR(s)
	return c
}

func IsIPV4(ip net.IP) (net.IP, bool) {
	if len(ip) == net.IPv4len {
		return ip, true
	}

	if len(ip) == net.IPv6len && IsZeros(ip[0:10]) && ip[10] == 0xff && ip[11] == 0xff {
		return ip[12:16], true
	}

	return ip, false
}

func IsZeros(p net.IP) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != 0 {
			return false
		}
	}
	return true
}
