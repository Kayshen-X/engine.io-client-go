package engine_io_client_go

import "errors"

type Transport struct {
	Path              string
	Hostname          string
	Port              string
	Secure            interface{}
	Query             map[string]string
	TimestampParam    string
	TimestampRequests interface{}
	ReadyState        string
	Agent             interface{}
	Socket            interface{} // Socket
	EnableXDR         bool
	WithCredentials   bool
	Writable          bool

	//	SSL options For Client
	Pfx                string
	Key                string
	Passphrase         string
	Cert               string
	Ca                 string
	Ciphers            string
	RejectUnauthorized bool
	ForceNode          bool

	// other options for client
	ExtraHeaders map[string]string
	LocalAddress string

	//	被继承者实现
	OnPacket interface{}
	Write    interface{}
	DoClose  interface{}
}

type Send struct {
	TypeName string
	Data     string
}

func (t *Transport) Open() {
	if t.ReadyState == "closed" || t.ReadyState == "" {
		t.ReadyState = "opening"
		t.DoOpen()
	}
}

func (t *Transport) Close() {
	if t.ReadyState == "opening" || t.ReadyState == "open" {
		t.DoClose()
		t.OnClose()
	}
}

func (t *Transport) Send(packets []Send) (bool, error) {
	if t.ReadyState == "open" {
		t.Write(packets)
		return true, nil
	} else {
		return false, errors.New("transport not open")
	}
}

func (t *Transport) OnOpen() {
	t.ReadyState = "open"
	t.Writable = true
	//	emit onOpen
}

//func (t *Transport) Write(packets []Send) {
//
//}
//
//func (t *Transport) DoClose() {
//
//}

func (t *Transport) DoOpen() {
	t.ReadyState = "open"
	t.Writable = true
	//	emit open
}

//func (t *Transport) OnPacket(packet interface{}) {
//
//}

func (t *Transport) OnClose() {
	t.ReadyState = "closed"
	//	emit close
}

func (t *Transport) OnData(data string) {
	//packet := Parser{}
	//this.onPacket(packet)
}

func (t *Transport) OnPakcet() {
	//	emit packet
}
