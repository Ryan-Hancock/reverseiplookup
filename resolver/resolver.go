package resolver

import (
	"fmt"
	"net"
)

// IPLookup takes a ip string and returns domains assocatiated.
func IPLookup(ip string) (domains []string, err error) {
	addr, err := net.LookupAddr(ip)
	if err != nil {
		return addr, fmt.Errorf("lookup ip error: %s", err.Error())
	}
	return addr, err
}

// HostLookup takes a domain string and returns ips assocatiated.
func HostLookup(domain string) (ips []string, err error) {
	addr, err := net.LookupHost(domain)
	if err != nil {
		return addr, fmt.Errorf("lookup domain error: %s", err.Error())
	}
	return addr, err
}
