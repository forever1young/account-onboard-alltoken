package service

// OAuthFlow drives the Claude.ai OAuth authorization: generating the
// authorization link (bound to a fingerprint template) and exchanging the
// state#code callback for tokens.
type OAuthFlow struct {
	clientID    string
	redirectURI string
}

func NewOAuthFlow(clientID, redirectURI string) *OAuthFlow {
	return &OAuthFlow{clientID: clientID, redirectURI: redirectURI}
}

func (f *OAuthFlow) GenerateAuthLink(fingerprintTemplateID uint) (link, state string, err error) {
	panic("not implemented")
}

func (f *OAuthFlow) ExchangeCallback(state, code string) (refreshToken, accessToken string, err error) {
	panic("not implemented")
}
