package jumio

import "net/http"

type API struct {
	Client          *http.Client
	AccountsBaseUrl string
}
