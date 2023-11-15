package providersdk

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
