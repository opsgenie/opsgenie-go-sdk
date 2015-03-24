package client

type ClientProxyConfiguration struct {
	Host		string
	Port 		int
	Username	string
	Password	string
}

type HttpTransportSettings struct {
	SocketTimeOut		int
	ConnectionTimeOut	int
	MaxRetryAttempts	int
	MaxOpenConnections 	int
}
