package hydra_client

type AcceptLoginResponse struct {
	RedirectTo string `json:"redirect_to"`
}

type AcceptLoginRequest struct {
	Arc string `json:"arc,omitempty"`
	ForceSubjectIdentifier string `json:"force_subject_identifier,omitempty"`
	Remember bool `json:"remember"`
	Subject string `json:"subject"`
}
