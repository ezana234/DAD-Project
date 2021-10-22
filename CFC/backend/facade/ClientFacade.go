package facade

type Client struct {
	clientID int
}

func newClient(clientID int) *Client {
	return &Client{
		clientID,
	}
}

func (c *Client) getClientByID(clientID int) error {
	return nil
}
