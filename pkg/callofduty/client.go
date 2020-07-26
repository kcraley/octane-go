package callofduty

// Client performs all interactions against the Call of Duty site and APIs
type Client struct {
	Authentication *Authentication
}

// NewClient creates and returns a new client
// to interact with the Call of Duty APIs
func NewClient() *Client {
	return &Client{
		Authentication: &Authentication{},
	}
}

// Login authenticates to the Call of Duty APIS
// with the existing client.
func (c *Client) Login(username, password string) error {
	// Set username and password for authentication
	c.Authentication.Username = username
	c.Authentication.Password = password

	// Perform a login
	err := c.Authentication.login()
	if err != nil {
		return err
	}
	return nil
}
