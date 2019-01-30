package gointercom

// Config ..
type Config struct {
	AccessToken string
}

var config *Config

// GetConfig ...
func GetConfig() *Config {
	return config
}

// SetToken ...
func (c *Config) SetToken(token string) {
	c.AccessToken = token
}

// GetToken ...
func (c *Config) GetToken() string {
	return c.AccessToken
}

func init() {
	config = New()
}
