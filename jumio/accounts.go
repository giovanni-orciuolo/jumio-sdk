package jumio

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"time"
)

type AccountRequest struct {
	UserReference             string `json:"userReference"`
	CustomerInternalReference string `json:"customerInternalReference"`
	CallbackUrl               string `json:"callbackUrl"`
	TokenLifetime             string `json:"tokenLifetime"`
	Web                       struct {
		SuccessUrl string `json:"successUrl"`
		ErrorUrl   string `json:"errorUrl"`
		Locale     string `json:"locale"`
	} `json:"web"`
	WorkflowDefinition struct {
		Key         string `json:"key"`
		Credentials []struct {
			Category string `json:"category"`
			Country  struct {
				PredefinedType string   `json:"predefinedType"`
				Values         []string `json:"values"`
			} `json:"country"`
			Type struct {
				PredefinedType string   `json:"predefinedType"`
				Values         []string `json:"values"`
			} `json:"type"`
		} `json:"credentials"`
		Capabilities struct {
			WatchlistScreening struct {
			} `json:"watchlistScreening"`
			DocumentVerification struct {
				EnableExtraction bool `json:"enableExtraction"`
			} `json:"documentVerification"`
			Ruleset struct {
			} `json:"ruleset"`
		} `json:"capabilities"`
	} `json:"workflowDefinition"`
	UserConsent struct {
		UserLocation struct {
			Country string `json:"country"`
		} `json:"userLocation"`
		Consent struct {
			Obtained   string    `json:"obtained"`
			ObtainedAt time.Time `json:"obtainedAt"`
		} `json:"consent"`
	} `json:"userConsent"`
}

type AccountResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Account   struct {
		Id string `json:"id"`
	} `json:"account"`
	Web struct {
		Href       string `json:"href"`
		SuccessUrl string `json:"successUrl"`
		ErrorUrl   string `json:"errorUrl"`
	} `json:"web"`
	Sdk struct {
		Token string `json:"token"`
	} `json:"sdk"`
	WorkflowExecution struct {
		Id          string `json:"id"`
		Credentials []struct {
			Id              string   `json:"id"`
			Category        string   `json:"category"`
			AllowedChannels []string `json:"allowedChannels"`
			Api             struct {
				Token string `json:"token"`
				Parts struct {
					Front string `json:"front"`
					Back  string `json:"back"`
				} `json:"parts"`
				WorkflowExecution string `json:"workflowExecution"`
			} `json:"api"`
		} `json:"credentials"`
	} `json:"workflowExecution"`
}

func (a *API) CreateAccount(req AccountRequest) (*AccountResponse, error) {
	if a.Client == nil {
		return nil, errors.New("http client is not set, create it by calling CreateClient")
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshalling request")
	}

	res, err := a.Client.Post(a.AccountsBaseUrl+"/api/v1/accounts", "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, errors.Wrap(err, "making request to Jumio")
	}
	defer res.Body.Close()

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading response body")
	}

	var accountRes AccountResponse
	err = json.Unmarshal(rbody, &accountRes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling response")
	}

	return &accountRes, nil
}
