package stat

type GetRequestPayload struct {
	From string `validate:"required,yyyymmdd"`
	To   string `validate:"required,yyyymmdd"`
}
