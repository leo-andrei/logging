package log

type Transaction struct {
	TraceID    string
	Logger     *Logger
	Attributes map[string]interface{}
}

func NewTransaction(traceID string, logger *Logger, attributes map[string]interface{}) *Transaction {
	return &Transaction{
		TraceID:    traceID,
		Logger:     logger,
		Attributes: attributes,
	}
}

func (t *Transaction) Log(level LogLevel, message string) {
	attrs := make(map[string]interface{})
	for k, v := range t.Attributes {
		attrs[k] = v
	}
	attrs["TraceID"] = t.TraceID
	t.Logger.Log(level, message, attrs)
}
