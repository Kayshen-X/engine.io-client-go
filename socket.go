package engine_io_client_go

type PerMessageDeflate struct {
	threshold int
}

type SocketOptions struct {
	hostname            string
	port                int
	secure              bool
	query               map[string]interface{}
	upgrade             bool
	forceBase64         bool
	timestampParam      string
	timestampRequests   bool
	transports          []string
	policyPost          int
	rememberUpgrade     bool
	onlyBinaryUpgrades  bool
	transportOptions    map[interface{}]interface{}
	pfx                 string
	key                 string
	passphrase          string
	cert                string
	ca                  []string
	ciphers             string
	rejectUnauthorized  bool
	extraHeaders        map[string]string
	requestTimeout      int
	withCredentials     bool
	closeOnBeforeunload bool
	useNativeTimers     bool
	autoUnref           bool
	perMessageDeflate   PerMessageDeflate
	path                string
	protocols           []string
}

type Socket struct {
	hostname string
	port     int
	agent    bool
	secure   bool

	query map[string]interface{}
}

func NewSocket(options SocketOptions) *Socket {
	port := options.port
	if len(options.hostname) != 0 && options.port == 0 {
		if options.secure {
			port = 443
		} else {
			port = 80
		}
	}

	return &Socket{
		hostname: options.hostname,
		port:     port,
		secure:   options.secure,
		query:    options.query,
	}
}
