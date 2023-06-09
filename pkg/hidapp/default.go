package hidapp

var defaultProcessor = NewProcessor()

// AppendRegexp appends new regexps to the default instance,
// the arguments can be string or *regexp.Regexp.
func AppendRegexp(reg ...any) error {
	return defaultProcessor.AppendRegexp(reg...)
}
