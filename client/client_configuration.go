package client

import (
	"fmt"
	"time"
)

type ClientProxyConfiguration struct {
	Host     string
	Port     int
	Username string
	Password string
	ProxyUri string
	Protocol string
}

type HttpTransportSettings struct {
	ConnectionTimeout time.Duration
	RequestTimeout    time.Duration
	MaxRetryAttempts  int
}

func (proxy *ClientProxyConfiguration) ToString() string {
	if proxy.ProxyUri != "" {
		return proxy.ProxyUri
	}
	if proxy.Protocol == "" {
		proxy.Protocol = "http"
	}
	if proxy.Username != "" && proxy.Password != "" {
		return fmt.Sprintf("%s://%s:%s@%s:%d", proxy.Protocol, proxy.Username, proxy.Password, proxy.Host, proxy.Port)
	} else {
		return fmt.Sprintf("%s://%s:%d", proxy.Protocol, proxy.Host, proxy.Port)
	}
	return ""
}
