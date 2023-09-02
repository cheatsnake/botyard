package helpers

func JsonMessage(body string) interface{} {
	return struct {
		Message string `json:"message"`
	}{
		Message: body,
	}
}
