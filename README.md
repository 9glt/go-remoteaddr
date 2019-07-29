# go-remoteaddr
get remote addr from http.Request


```go
package main

import (
    "github.com/9glt/go-remoteaddr"
    "net/http"
    "fmt"
)

func main() {
    request, _ := http.NewRequest("GET", "/path", nil)
    request.Header.Set("X-Forwarded-For", "192.168.100.100")
    request.Header.Set("X-Real-IP", "192.168.100.100")
    request.RemoteAddr = "127.0.0.1:123123"
    
    addr := remoteaddr.IP(request, []string{"proxy-ip-address1","proxy-ip-addres2"})
    fmt.Printf("%+v", addr)
}
```