package engine_io_client_go

import "errors"

type Transport struct {
	path              string
	hostname          string
	port              string
	secure            interface{}
	query             string
	timestampParam    string
	timestampRequests interface{}
	readyState        string
	agent             interface{}
	socket            interface{} // Socket
	enableXDR         bool
	withCredentials   bool
	writable          bool

	//	SSL options For Client
	pfx                string
	key                string
	passphrase         string
	cert               string
	ca                 string
	ciphers            string
	rejectUnauthorized bool
	forceNode          bool

	// results of ReactNative environment detection
	isReactNative bool

	// other options for client
	extraHeaders map[string]string
	localAddress string
}

type Send struct {
	typeName string
	data     string
}

func (t *Transport) open() {
	if t.readyState == "closed" || t.readyState == "" {
		t.readyState = "opening"
		t.doOpen()
	}
}

func (t *Transport) close() {
	if t.readyState == "opening" || t.readyState == "open" {
		t.doClose()
		t.onClose()
	}
}

func (t *Transport) send(packets []Send) (bool, error) {
	if t.readyState == "open" {
		t.write(packets)
		return true, nil
	} else {
		return false, errors.New("transport not open")
	}
}

func (t *Transport) write(packets []Send) {

}

func (t *Transport) doClose() {

}

func (t *Transport) doOpen() {
	t.readyState = "open"
	t.writable = true
	//	emit open
}

func (t *Transport) onPacket(packet interface{}) {

}

func (t *Transport) onClose() {
	t.readyState = "closed"
	//	emit close
}

func (t *Transport) onData(data string) {
	//packet := Parser{}
	//this.onPacket(packet)
}

func (t *Transport) onPakcet() {
	//	emit packet
}
