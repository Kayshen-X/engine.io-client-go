package engine_io_client_go

import "sync"

type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel is a channel that can receive DataEvent
type DataChannel chan DataEvent

// DataChannelSlice is a slice containing DataChannels data
type DataChannelSlice []DataChannel

// EventBus Store information about specific topics of interest to subscribers
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if channelArry, found := eb.subscribers[topic]; found {

		channels := append(DataChannelSlice{}, channelArry...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}
