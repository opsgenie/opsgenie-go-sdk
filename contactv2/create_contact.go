package contactv2

type CreateContactRequest struct {
	ApiKey
	*Identifier
	Method ContactMethod `json:"method"`
	To     string        `json:"to"`
}

type CreateContactResponse struct {
	ResponseMeta
	Status string `json:"result"`
	Result struct {
		ID string `json:"id"`
	} `json:"data"`
}
