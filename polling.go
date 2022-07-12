package engine_io_client_go

type Polling struct {
	Transport
	polling        bool
	supportsBinary bool
}

func (p *Polling) name() string {
	return "polling"
}

func (p *Polling) doOpen() {
	p.poll()
}

func pause(p *Polling, onPause func()) {
	p.readyState = "paused"
	onPause()
}

func (p *Polling) pause(onPause func()) {
	p.readyState = "pausing"
	if p.polling == true || p.writable == false {
		total := 0
		if p.polling == true {
			total++
			//	once pollComplete
			//	--total || pause();
		}
		if p.writable == false {
			total++
			//  once drain
			//	--total || pause();
		}
	} else {
		pause(p, onPause)
	}

}

func (p *Polling) poll() {
	p.polling = true
	p.doPoll()
	//	emit poll
}

func (p *Polling) onData(data any) {
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

	if p.readyState == "closed" {
		p.polling = false
		//	emit pollComplete
		if p.readyState == "open" {
			p.poll()
		}
	}
}
func close(p *Polling) {
	packet := Send{typeName: "close"}
	packets := []Send{packet}
	p.write(packets)
}

func (p *Polling) doClose() {
	if p.readyState == "open" {
		close(p)
	} else {
		//	once open, close
	}
}

func (p *Polling) doPoll() {

}

func (p *Polling) write() {
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
}

func (p *Polling) uri() string {
	query := p.query
	var schema string
	if p.secure != nil {
		schema = "https"
	} else {
		schema = "http"
	}
	port := ""
	// cache busting is forced
	//if (false !== this.timestampRequests) {
	//	query[this.timestampParam] = yeast();
	//}

	if _, ok := query["sid"]; ok && p.supportsBinary == true {
		query["b64"] = "1"
	}

	//query = parseqs.encode(query);
	if p.port != "0" && ((schema == "https" && p.port != "443") || (schema == "http" && p.port != "80")) {
		port = ":" + p.port
	}
	queryStr := ""
	if len(queryStr) != 0 {
		queryStr = "?" + queryStr
	}
	return ""
}
