package adapter

type Client struct {
}

func (c *Client) Get(http Http, url string) (interface{}, error) {
	return http.Get(url)
}
