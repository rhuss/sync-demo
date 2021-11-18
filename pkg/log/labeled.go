package log

// LabeledLogger Will add a label to all messages.
type LabeledLogger struct {
	Label string
	Logger
}

func (l LabeledLogger) Println(v ...interface{}) {
	vars := l.prependLabel(v)
	l.Logger.Println(vars...)
}

func (l LabeledLogger) Printf(format string, v ...interface{}) {
	format = l.Label + " " + format
	l.Logger.Printf(format, v...)
}

func (l LabeledLogger) prependLabel(v []interface{}) []interface{} {
	vars := make([]interface{}, len(v)+1)
	vars[0] = l.Label
	for i := 0; i < len(v); i++ {
		vars[i+1] = v[i]
	}
	return vars
}
