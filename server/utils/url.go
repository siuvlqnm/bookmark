package utils

import (
	"errors"
	"net/url"
	"strings"
)

func ParseUrl(link string) (err error, P map[string]interface{}) {
	P = make(map[string]interface{}, 5)
	u, err := url.Parse(link)
	if err != nil {
		return errors.New("域名解析失败"), P
	}
	h := strings.Split(u.Host, ":")
	P["protocol"] = u.Scheme
	P["domain"] = h[0]
	P["port"] = h[1]
	P["path"] = u.Path
	P["query"] = u.RawQuery
	return err, P
}
