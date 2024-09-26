package client

type Client struct {
	Host    string
	Port    uint16
	Timeout int
}

func NewClient(host string, port uint16, timeout int) *Client {
	return &Client{Host: host, Port: port, Timeout: timeout}
}
