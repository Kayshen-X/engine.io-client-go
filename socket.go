package engine_io_client_go

import (
	"errors"
	"strings"
)

const Websocket = "websocket"

const Polling = "polling"

type SocketOptions struct {
	hostname           string
	port               int
	secure             bool
	query              map[string]string
	upgrade            bool
	path               string
	forceJSONP         bool
	jsonp              bool
	enablesXDR         bool
	withCredentials    bool
	timestampParam     string
	timestampRequests  interface{}
	transports         []string
	transportOptions   map[string]interface{}
	policyPort         int
	rememberUpgrade    bool
	onlyBinaryUpgrades bool
	pfx                string
	key                string
	passphrase         string
	cert               string
	ca                 string
	ciphers            string
	rejectUnauthorized bool
	forceNode          bool
	extraHeaders       map[string]interface{}
	localAddress       string
}

type Socket struct {
	hostname              string
	port                  int
	secure                bool
	query                 map[string]string
	upgrade               bool
	agent                 bool
	path                  string
	forceJSONP            bool
	jsonp                 bool
	enablesXDR            bool
	withCredentials       bool
	timestampParam        string
	timestampRequests     interface{}
	transports            []string
	transportOptions      map[string]interface{}
	readyState            string
	writeBuffer           []interface{}
	prevBufferLen         int
	policyPort            int
	rememberUpgrade       bool
	binaryType            any
	onlyBinaryUpgrades    bool
	pfx                   string
	key                   string
	passphrase            string
	cert                  string
	ca                    string
	ciphers               string
	rejectUnauthorized    bool
	forceNode             bool
	extraHeaders          map[string]interface{}
	localAddress          string
	id                    string
	upgrades              interface{}
	pingInterval          interface{}
	pingTimeout           interface{}
	pingIntervalTimer     interface{}
	pingTimeoutTimer      interface{}
	priorWebsocketSuccess bool
}

func (s Socket) open() error {
	var transportMethod string

	if len(s.transports) == 0 {
		return errors.New("no transports available")
	}

	if s.rememberUpgrade && s.priorWebsocketSuccess && IsContain(s.transports, Websocket) {
		transportMethod = "websocket"
	} else {
		transportMethod = s.transports[0]
	}

	s.readyState = "opening"

	transport, err := s.createTransport(transportMethod)

	if err != nil {
		s.transports = s.transports[1:]
		return s.open()
	}

	//transport.open()

	s.setTransport(transport)

	return nil
}

func (s Socket) createTransport(name string) (interface{}, error) {
	query := make(map[string]string)
	MapClone(query, s.query)
	//query["EIO"] = ""
	query["transport"] = name
	// per-transport options
	//options := s.transportOptions[name]

	if len(s.id) != 0 {
		query["sid"] = s.id
	}

	//transport := transports[name]({})

	return nil, nil
}

func (s Socket) setTransport(transport interface{}) {

}

func NewSocket(options SocketOptions) (*Socket, error) {
	port := options.port
	if len(options.hostname) != 0 && options.port == 0 {
		if options.secure {
			port = 443
		} else {
			port = 80
		}
	}

	path := "/engine.io/"

	if len(options.path) != 0 {
		path = options.path
		if !strings.HasSuffix(options.path, "/") {
			path = path + "/"
		}
	}

	transports := []string{Polling, Websocket}
	if len(options.transports) != 0 {
		transports = options.transports
	}
	writeBuffer := []interface{}{}

	policyPort := 843

	if options.policyPort != 0 {
		policyPort = options.policyPort
	}

	rememberUpgrade := false
	if options.rememberUpgrade {
		rememberUpgrade = options.rememberUpgrade
	}

	socket := &Socket{
		hostname:              options.hostname,
		port:                  port,
		secure:                options.secure,
		query:                 options.query,
		upgrade:               options.upgrade,
		path:                  path,
		forceJSONP:            options.forceJSONP,
		jsonp:                 options.jsonp,
		enablesXDR:            options.enablesXDR,
		withCredentials:       options.withCredentials,
		timestampParam:        options.timestampParam,
		timestampRequests:     options.timestampRequests,
		transports:            transports,
		transportOptions:      options.transportOptions,
		readyState:            "",
		writeBuffer:           writeBuffer,
		prevBufferLen:         0,
		policyPort:            policyPort,
		rememberUpgrade:       rememberUpgrade,
		binaryType:            nil,
		onlyBinaryUpgrades:    options.onlyBinaryUpgrades,
		pfx:                   options.pfx,
		key:                   options.key,
		passphrase:            options.passphrase,
		cert:                  options.cert,
		ca:                    options.ca,
		ciphers:               options.ciphers,
		rejectUnauthorized:    options.rejectUnauthorized,
		forceNode:             options.forceNode,
		extraHeaders:          options.extraHeaders,
		localAddress:          options.localAddress,
		id:                    "",
		upgrades:              nil,
		pingInterval:          nil,
		pingTimeout:           nil,
		pingIntervalTimer:     nil,
		pingTimeoutTimer:      nil,
		priorWebsocketSuccess: false,
	}
	err := socket.open()
	if nil != err {
		return nil, err
	}

	return socket, nil
}
