package bearychat

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// WebhookResponse represents a response.
type WebhookResponse struct {
	StatusCode int              `json:"-"`
	Code       int              `json:"code"`
	Error      string           `json:"error,omitempty"`
	Result     *json.RawMessage `json:"result"`
}

func (w WebhookResponse) IsOk() bool {
	return w.Code == 0
}

// WebhookClient represents any webhook client can send message to BearyChat.
type WebhookClient interface {
	// Set webhook webhook.
	SetWebhook(webhook string) WebhookClient

	// Set http client.
	SetHTTPClient(client *http.Client) WebhookClient

	// Send webhook payload.
	Send(payload io.Reader) (*WebhookResponse, error)
}

type webhookClient struct {
	httpClient *http.Client

	Webhook string
}

// Creates a new incoming webhook client.
//
// For full documentation, visit https://bearychat.com/integrations/incoming .
func NewIncomingWebhookClient(webhook string) *webhookClient {
	return &webhookClient{
		httpClient: http.DefaultClient,

		Webhook: webhook,
	}
}

func (w *webhookClient) SetWebhook(webhook string) WebhookClient {
	w.Webhook = webhook
	return w
}

func (w *webhookClient) SetHTTPClient(c *http.Client) WebhookClient {
	w.httpClient = c
	return w
}

func (w *webhookClient) Send(payload io.Reader) (*WebhookResponse, error) {
	if w.Webhook == "" {
		return nil, errors.New("webhook url is required")
	}
	if w.httpClient == nil {
		return nil, errors.New("http client is required")
	}

	resp, err := w.httpClient.Post(w.Webhook, "application/json", payload)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	webhookResponse := new(WebhookResponse)
	webhookResponse.StatusCode = resp.StatusCode
	if err := json.NewDecoder(resp.Body).Decode(webhookResponse); err != nil {
		return nil, err
	}

	return webhookResponse, nil
}
