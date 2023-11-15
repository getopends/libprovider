package libprovider

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
}
