package remoteaddr

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	request, _ := http.NewRequest("GET", "/path", nil)
	request.Header.Set("X-Forwarded-For", "192.168.100.100")
	request.Header.Set("X-Real-IP", "192.168.100.100")
	request.RemoteAddr = "127.0.0.1:123123"

	addr := IP(request, []string{})
	if addr.IP != "127.0.0.1" {
		t.Fatal()
	}
	if addr.XForIP != "192.168.100.100" {
		t.Fatal()
	}

	addr1 := IP(request, []string{"127.0.0.1"})
	if addr1.IP != "192.168.100.100" {
		t.Fatal()
	}
	if addr1.XForIP != "192.168.100.100" {
		t.Fatal()
	}
	if addr1.BehindProxy != true {
		t.Fatal()
	}

	request, _ = http.NewRequest("GET", "/path", nil)
	request.RemoteAddr = "127.0.0.5:12312"

	addr2 := IP(request, nil)
	if addr2.IP != "127.0.0.5" {
		t.Fatal()
	}
	if addr2.BehindProxy == true {
		t.Fatal()
	}

	request, _ = http.NewRequest("GET", "/path", nil)
	request.RemoteAddr = "[::1]:12312"

	addr3 := IP(request, nil)
	if addr3.IP != "::1" {
		t.Fatal()
	}

}
