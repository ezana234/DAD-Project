package model

type Client struct {
	ClientID int
	userID   int
}

func NewClient(userID int) *Client {
	return &Client{userID: userID}
}

func (c *Client) GetClientID() int {
	return c.ClientID
}

func (c *Client) SetClientID(clientID int) {
	c.ClientID = clientID
}

func (c *Client) UserID() int {
	return c.userID
}

func (c *Client) SetUserID(userID int) {
	c.userID = userID
}
