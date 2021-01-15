package utils

import (
	"net/url"
	"strconv"
	"strings"
)

type UrlInfo struct {
	Protocol string
	Domain   string
	Port     int64
	Path     string
	Query    string
}

func ParseUrl(link string) (err error, I *UrlInfo) {
	u, err := url.Parse(link)
	if err != nil {
		return err, I
	}
	h := strings.Split(u.Host, ":")
	if len(h) == 2 {
		port, _ := strconv.ParseInt(h[1], 10, 0)
		I = &UrlInfo{Protocol: u.Scheme, Domain: h[0], Port: port, Path: u.Path, Query: u.RawQuery}
		return err, I
	}
	I = &UrlInfo{Protocol: u.Scheme, Domain: h[0], Path: u.Path, Query: u.RawQuery}
	return err, I
}
