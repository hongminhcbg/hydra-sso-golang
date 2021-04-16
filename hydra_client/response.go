package hydra_client

type BaseError struct {
	Debug            string `json:"debug,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
	StatusCode       int    `json:"status_code,omitempty"`
}

type AcceptLoginResponse struct {
	RedirectTo string `json:"redirect_to"`
}

type AcceptLoginRequest struct {
	Arc                    string `json:"arc,omitempty"`
	ForceSubjectIdentifier string `json:"force_subject_identifier,omitempty"`
	Remember               bool   `json:"remember"`
	Subject                string `json:"subject"`
}

type AccessToken struct {
	Subject string `json:"subject"`
}

type IdToken struct {
	Subject string `json:"subject"`
}

type Session struct {
	AccessToken AccessToken `json:"access_token"`
	IdToken     IdToken     `json:"id_token"`
}

type AcceptConsentRequest struct {
	GrantScope []string `json:"grant_scope"`
	Remember   bool     `json:"remember"`
	Session    Session  `json:"session"`
}

type AcceptConsentResponse struct {
	BaseError
	RedirectTo string `json:"redirect_to"`
}

type IntrospectResponse struct {
	BaseError
	Scope string `json:"scope"`
	Active bool `json:"active"`
}