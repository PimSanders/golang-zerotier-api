package golangzerotierapi

// Client represents a client for interacting with the API.
type Client struct {
	BaseURL string
	Token   string
}

// NnwClient creates a new instance of Client with the provided base URL and API token.
func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: baseURL,
		Token:   token,
	}
}
