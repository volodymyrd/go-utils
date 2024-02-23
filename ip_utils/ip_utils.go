package ip_utils

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func GetIPAddress(r *http.Request) string {
	// Check for X-Forwarded-For header which contains the IP address from proxy
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For can contain a comma-separated list of IPs; the first one is client's IP
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	return ip
}

func GetLocation(ip string) string {
	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", "https://ipapi.co/"+ip+"/json/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := ipapiClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Location got successfully")
	return string(body)
}
