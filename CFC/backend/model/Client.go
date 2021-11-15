package model

type Client struct {
	ClientID int
	UserID   int
}

func NewClient(userID int) *Client {
	return &Client{UserID: userID}
}
func NewClientBoth(clientID int, userID int) *Client {
	return &Client{ClientID: clientID, UserID: userID}
}

func (c *Client) GetClientID() int {
	return c.ClientID
}

func (c *Client) SetClientID(clientID int) {
	c.ClientID = clientID
}

func (c *Client) GetUserID() int {
	return c.UserID
}

func (c *Client) SetUserID(userID int) {
	c.UserID = userID
}
