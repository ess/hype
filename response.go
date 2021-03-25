package hype

type Response struct {
	Data  []byte
	Error error
}

func (response Response) Okay() bool {
	if response.Error == nil {
		return true
	}
	return false
}
