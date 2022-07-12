package tranports

import "github.com/Kayshen-X/engine.io-client-go"

type Polling struct {
	engine_io_client_go.Transport
	Polling        bool
	SupportsBinary bool
	Name           string
}

type PollingOption struct {
	engine_io_client_go.Transport
	Polling            bool
	Path               string
	Hostname           string
	Port               string
	Secure             interface{}
	Query              map[string]string
	TimestampParam     string
	TimestampRequests  interface{}
	Agent              interface{}
	Socket             interface{} // Socket
	EnableXDR          bool
	WithCredentials    bool
	Writable           bool
	Pfx                string
	Key                string
	Passphrase         string
	Cert               string
	Ca                 string
	Ciphers            string
	RejectUnauthorized bool
	ForceNode          bool
	ExtraHeaders       map[string]string
	LocalAddress       string
}

func (p *Polling) DoOpen() {
	p.Poll()
}

func pause(p *Polling, onPause func()) {
	p.ReadyState = "paused"
	onPause()
}

func (p *Polling) Pause(onPause func()) {
	p.ReadyState = "pausing"
	if p.Polling == true || p.Writable == false {
		total := 0
		if p.Polling == true {
			total++
			//	once pollComplete
			//	--total || pause();
		}
		if p.Writable == false {
			total++
			//  once drain
			//	--total || pause();
		}
	} else {
		pause(p, onPause)
	}

}

func (p *Polling) Poll() {
	p.Polling = true
	p.DoPoll()
	//	emit poll
}

func (p *Polling) OnData(data any) {
	//cb := func(packet any, index int, total int) {
	//	if p.readyState == "opening" && packet.typeName == "open" {
	//		p.onOpen()
	//	}
	//	if packet.typeName == "close" {
	//		p.onClose()
	//	}
	//	p.onPacket(packet)
	//}

	// decode payload
	//parser.decodePayload(data, p.socket.binaryType, cb);

	if p.ReadyState == "closed" {
		p.Polling = false
		//	emit pollComplete
		if p.ReadyState == "open" {
			p.Poll()
		}
	}
}
func close(p *Polling) {
	packet := engine_io_client_go.Send{TypeName: "close"}
	packets := []engine_io_client_go.Send{packet}
	p.Write(packets)
}

func (p *Polling) DoClose() {
	if p.ReadyState == "open" {
		close(p)
	} else {
		//	once open, close
	}
}

func (p *Polling) DoPoll() {

}

//func (p *Polling) write() {
//var self = this;
//this.writable = false;
//var callbackfn = function () {
//	self.writable = true;
//	self.emit('drain');
//};
//
//parser.encodePayload(packets, this.supportsBinary, function (data) {
//	self.doWrite(data, callbackfn);
//});
//}

func (p *Polling) Uri() string {
	query := p.Query
	var schema string
	if p.Secure != nil {
		schema = "https"
	} else {
		schema = "http"
	}
	port := ""
	// cache busting is forced
	//if (false !== this.timestampRequests) {
	//	query[this.timestampParam] = yeast();
	//}

	if _, ok := query["sid"]; ok && p.SupportsBinary == true {
		query["b64"] = "1"
	}

	//query = parseqs.encode(query);
	if p.Port != "0" && ((schema == "https" && p.Port != "443") || (schema == "http" && p.port != "80")) {
		port = ":" + p.Port
	}
	queryStr := ""
	if len(queryStr) != 0 {
		queryStr = "?" + queryStr
	}
	return ""
}

func NewPolling(option PollingOption) Polling {
	supportsBinary := false
	return Polling{
		SupportsBinary: supportsBinary,
		Polling:        option.Polling,
		Name:           "polling",
		Transport: engine_io_client_go.Transport{
			Path:               option.Path,
			Hostname:           option.Hostname,
			Port:               option.Port,
			Secure:             option.Secure,
			Query:              option.Query,
			TimestampParam:     option.TimestampParam,
			TimestampRequests:  option.TimestampRequests,
			ReadyState:         "",
			Agent:              option.Agent,
			Socket:             option.Socket,
			EnableXDR:          option.EnableXDR,
			WithCredentials:    option.WithCredentials,
			Writable:           option.Writable,
			Pfx:                option.Pfx,
			Key:                option.Key,
			Passphrase:         option.Passphrase,
			Cert:               option.Cert,
			Ca:                 option.Ca,
			Ciphers:            option.Ciphers,
			RejectUnauthorized: option.RejectUnauthorized,
			ForceNode:          option.ForceNode,
			ExtraHeaders:       option.ExtraHeaders,
			LocalAddress:       option.LocalAddress,
		},
	}
}
