package resolver

import (
	"fmt"
	"net"
)

type Resolver struct {
	storage ResolverStorage
}

// IPLookup takes a ip string and returns domains assocatiated.
func(r *Resolver) IPLookup(ip string) (domains []string, err error) {
	addr, err := net.LookupAddr(ip)
	if err != nil {
		return addr, fmt.Errorf("lookup ip error: %s", err.Error())
	}

	//do a merge of the domains for newer ones from the net lookup
	r.storage.GetByIP(ip)
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
