package engine_io_client_go

type Polling struct {
	Transport
	polling bool
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

func (p *Polling) doPoll() {

}
