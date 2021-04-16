package hydra_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type HydraClient struct {
	adminHost               string
	publicHost              string
	acceptLoginRequestURL   string
	acceptConsentRequestURL string
	introspectURL           string
}

func NewHydraClient(adminHost string, publicHost string) *HydraClient {
	return &HydraClient{
		adminHost:               adminHost,
		publicHost:              publicHost,
		acceptLoginRequestURL:   adminHost + "/oauth2/auth/requests/login/accept",
		acceptConsentRequestURL: adminHost + "/oauth2/auth/requests/consent/accept",
		introspectURL:           adminHost + "/oauth2/introspect",
	}
}

func (h *HydraClient) AcceptLoginRequest(loginChallenge string, request AcceptLoginRequest) (*AcceptLoginResponse, error) {
	url := fmt.Sprintf("%s?login_challenge=%s", h.acceptLoginRequestURL, loginChallenge)
	b, _ := json.Marshal(request)

	httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[DB] do http request error")
		return nil, err
	}

	defer resp.Body.Close()

	var response AcceptLoginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("[DB] par body error")
		return nil, err
	}

	return &response, nil
}

func (h *HydraClient) AcceptConsentRequest(loginChallenge string, request AcceptConsentRequest) (*AcceptConsentResponse, error) {
	url := fmt.Sprintf("%s?consent_challenge=%s", h.acceptConsentRequestURL, loginChallenge)
	b, _ := json.Marshal(request)

	httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[DB] do http request error")
		return nil, err
	}

	defer resp.Body.Close()

	var response AcceptConsentResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("[DB] par body error")
		return nil, err
	}

	if len(response.Error) > 0 {
		log.Printf("[DB] accept consent error: %+v", response)
		return nil, errors.New(response.ErrorDescription)
	}

	return &response, nil
}

func (h *HydraClient) Introspect(token string) (*IntrospectResponse, error) {
	data := url.Values{}
	data.Set("token", token)
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, h.introspectURL , strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal("[DB] make http request error", err)
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		log.Fatal("[DB] do http request error", err)
		return nil, err
	}

	defer resp.Body.Close()

	var response IntrospectResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("[DB] par body error")
		return nil, err
	}

	if len(response.Error) > 0 {
		log.Printf("[DB] accept consent error: %+v", response)
		return nil, errors.New(response.ErrorDescription)
	}

	if !response.Active{
		log.Println("[DB] token not active")
		return nil, errors.New("token not active")
	}

	return &response, nil
}
