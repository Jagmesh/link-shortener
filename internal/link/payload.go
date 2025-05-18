package link

type CreateRequestPayload struct {
	Url string `json:"url" validate:"required,url"`
}

type DeleteRequestPayload struct {
	Hash string `json:"hash"`
	Url  string `json:"url" validate:"omitempty,url"`
	Id   uint   `json:"Id" validate:"number"`
}
