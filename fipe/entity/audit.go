package entity

type Audit struct {
	Payload    []byte
	Response   []byte
	StatusCode int
}

func (a Audit) PayloadString() string {
	return string(a.Payload)
}

func (a Audit) ResponseString() string {
	return string(a.Response)
}
