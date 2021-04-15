package hydra_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HydraClient struct {
	adminHost          string
	publicHost         string
	acceptLoginRequest string
}

func NewHydraClient(adminHost string, publicHost string) *HydraClient {
	return &HydraClient{
		adminHost:          adminHost,
		publicHost:         publicHost,
		acceptLoginRequest: adminHost + "/oauth2/auth/requests/login/accept",
	}
}

func (h *HydraClient) AcceptLoginRequest(loginChallenge string, request AcceptLoginRequest) (*AcceptLoginResponse, error) {
	url := fmt.Sprintf("%s?login_challenge=%s", h.acceptLoginRequest, loginChallenge)
	b, _ := json.Marshal(request)

	httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

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