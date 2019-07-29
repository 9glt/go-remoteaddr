package remoteaddr

import (
	"net/http"
	"strings"
)

func getClientIp(r *http.Request) string {
	remote_addr := r.RemoteAddr
	idx := strings.LastIndex(remote_addr, ":")
	if idx != -1 {
		remote_addr = remote_addr[0:idx]
		if remote_addr[0] == '[' && remote_addr[len(remote_addr)-1] == ']' {
			remote_addr = remote_addr[1 : len(remote_addr)-1]
		}
	}
	return remote_addr
}

// IP takes *http.Request, proxies ips and depth - how many proxies before my app
func IP(r *http.Request, proxies []string, depth int) Addr {
	clientIP := getClientIp(r)
	ack := Addr{IP: clientIP}
	ack.XForIP = r.Header.Get("X-Forwarded-For")
	ack.XRealIP = r.Header.Get("X-Real-IP")
	if inList(clientIP, proxies) {
		ack.BehindProxy = true
		xfips := strings.Split(ack.XForIP, ",")
		if (depth - 1) > -1 {
			depth = depth - 1
			ack.IP = strings.TrimSpace(xfips[len(xfips)-depth-1])
		}
		if ack.XRealIP != "" {
			ack.IP = ack.XRealIP
		}
	}
	return ack
}

func inList(f string, l []string) bool {
	for _, v := range l {
		if f == v {
			return true
		}
	}
	return false
}

type Addr struct {
	IP          string
	XForIP      string
	XRealIP     string
	BehindProxy bool
}
