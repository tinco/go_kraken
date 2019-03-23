package websocket

import (
	"time"
)

// Parameters defines adapter behavior.
type Parameters struct {
	AutoReconnect     bool
	ReconnectInterval time.Duration
	ReconnectAttempts int
	reconnectTry      int
	ShutdownTimeout   time.Duration

	ResubscribeOnReconnect bool

	HeartbeatTimeout time.Duration
	LogTransport     bool

	URL string
}

func NewDefaultParameters() *Parameters {
	return &Parameters{
		AutoReconnect:          true,
		ReconnectInterval:      time.Second,
		reconnectTry:           0,
		ReconnectAttempts:      5,
		URL:                    prodBaseURL,
		ShutdownTimeout:        time.Second * 5,
		ResubscribeOnReconnect: true,
		HeartbeatTimeout:       time.Second * 15, // HB = 3s
		LogTransport:           false,            // log transport send/recv
	}
}

func NewDefaultSandboxParameters() *Parameters {
	return &Parameters{
		AutoReconnect:          true,
		ReconnectInterval:      time.Second,
		reconnectTry:           0,
		ReconnectAttempts:      5,
		URL:                    sandboxBaseURL,
		ShutdownTimeout:        time.Second * 5,
		ResubscribeOnReconnect: true,
		HeartbeatTimeout:       time.Second * 15, // HB = 3s
		LogTransport:           false,            // log transport send/recv
	}
}
