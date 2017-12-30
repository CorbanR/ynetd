package config

// Service holds string representations of Service attributes.
type Service struct {
	Proxy          map[string]string
	Command        []string
	StopAfter      string
	StopSignal     string
	Timeout        string
	WaitAfterStart string
}
