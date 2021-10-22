package model

type Client struct {
	clientID int
	userID   int
}

func NewClient(userID int) *Client {
	return &Client{userID: userID}
}

func (c *Client) ClientID() int {
	return c.clientID
}

func (c *Client) SetClientID(clientID int) {
	c.clientID = clientID
}

func (c *Client) UserID() int {
	return c.userID
}

func (c *Client) SetUserID(userID int) {
	c.userID = userID
}
