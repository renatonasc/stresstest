package entity

type Request struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    string
}
