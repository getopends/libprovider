package providerd

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
