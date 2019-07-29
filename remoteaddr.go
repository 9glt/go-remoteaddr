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

func IP(r *http.Request, proxies []string) Addr {
	clientIP := getClientIp(r)
	ack := Addr{IP: clientIP}
	ack.XForIP = r.Header.Get("X-Forwarded-For")
	ack.XRealIP = r.Header.Get("X-Real-IP")
	if inList(clientIP, proxies) {
		ack.BehindProxy = true
		ack.IP = strings.TrimSpace(strings.Split(ack.XForIP, ",")[0])
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
