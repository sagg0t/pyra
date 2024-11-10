package log

const Key = "event"

const (
	ServerStartEvent    = "ServerStart"
	ServerShutdownEvent = "ServerShutdown"
	ServerStopEvent     = "ServerStop"
	PanicEvent          = "Panic"
	RequestEvent        = "RequestLog"
	DBTraceEvent        = "DBTrace"
)
