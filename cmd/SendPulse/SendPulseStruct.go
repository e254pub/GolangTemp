package SendPulse

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ErrorCode   int    `json:"error_code"`
}

type emailArray struct {
	Content email `json:"email"`
}

type Recipient struct {
	Email string `json:"email"`
}

type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Template struct {
	Id        int              `json:"id"`
	Variables TemplateVariable `json:"variables"`
}

type TemplateVariable struct {
	Token string `json:"token"`
}

type email struct {
	Subject  string      `json:"subject"`
	Template Template    `json:"template"`
	From     Sender      `json:"from"`
	To       []Recipient `json:"to"`
}

type sendEmailResponse struct {
	Result bool `json:"result"`
}
