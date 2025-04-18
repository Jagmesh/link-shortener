package link

type CreateRequestPayload struct {
	Url string `json:"url" validate:"required,url"`
}

type DeleteRequestPayload struct {
	Hash string `json:"hash"`
	Url  string `json:"url" validate:"url"`
	Id   uint   `json:"id" validate:"number"`
}
