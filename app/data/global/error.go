package global

var Errors chan error

func init() {
	Errors = make(chan error, 65535)
}
