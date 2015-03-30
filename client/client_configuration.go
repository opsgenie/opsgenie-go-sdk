package client

import "fmt"

type ClientProxyConfiguration struct {
	Host		string
	Port 		int
	Username	string
	Password	string
	ProxyUri 	string
	Secured		bool
}

type HttpTransportSettings struct {
	SocketTimeOut		int
	ConnectionTimeOut	int
	MaxRetryAttempts	int
	MaxOpenConnections 	int
}

func (proxy *ClientProxyConfiguration) ToString() string {
	protocol := "http"
	if proxy.Secured {
		protocol = "https"
	}

	if proxy.ProxyUri != "" {
		return proxy.ProxyUri
	}	
	if proxy.Username != "" && proxy.Password != "" {		
		return fmt.Sprintf("%s://%s:%s@%s:%d", protocol, proxy.Username, proxy.Password, proxy.Host, proxy.Port)
	} else {
		return fmt.Sprintf("%s://%s:%d", protocol, proxy.Host, proxy.Port )
	}
	return ""
}