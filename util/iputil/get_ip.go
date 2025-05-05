package iputil

import (
	"net"
	"net/http"
	"strings"
)

// GetClientIp returns the client IP address from the http.Request
func GetClientIp(httpReq *http.Request) (string, error) {
	ip := httpReq.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0]), nil
	}

	ip = httpReq.Header.Get("X-Real-IP")
	if ip != "" {
		return ip, nil
	}

	ip, _, err := net.SplitHostPort(httpReq.RemoteAddr)
	if err != nil {
		return "", err
	}
	return ip, nil
}
