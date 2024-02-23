package jumio

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

func (a *API) CreateClient(ctx context.Context, clientId string, clientSecret string) (*http.Client, error) {
	config := clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     "https://auth.emea-1.jumio.ai/oauth2/token",
		AuthStyle:    oauth2.AuthStyleAutoDetect,
	}

	a.Client = config.Client(ctx)
	return a.Client, nil
}
