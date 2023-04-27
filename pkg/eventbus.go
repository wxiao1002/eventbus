package pkg

import "sync"

var (
	sysEventBus = NewEventBus()
	PubLish     = sysEventBus.Pub
	SubScribe   = sysEventBus.Sub
	DeSubScribe = sysEventBus.DeSub
)

type EventBus interface {
	Pub(topic string, data any)
	Sub(topic string, f func(data any)) string
	DeSub(topic string, subId string)
}

type SubCallBack struct {
	SubID    string
	Callback func(data any)
}

type DefaultEventBus struct {
	topicSubMap map[string][]*SubCallBack
	lock        sync.RWMutex
}

func NewEventBus() *DefaultEventBus {
	e := &DefaultEventBus{}
	e.topicSubMap = make(map[string][]*SubCallBack, 10)
	return e
}

func (e *DefaultEventBus) Pub(topic string, data any) {
	go func() {
		e.lock.RLock()
		defer e.lock.RUnlock()
		subs, ok := e.topicSubMap[topic]
		if !ok {
			return
		}
		for _, c := range subs {
			go c.Callback(data)
		}
	}()
}

func (e *DefaultEventBus) Sub(topic string, f func(data any)) string {
	e.lock.RLock()
	defer e.lock.RUnlock()
	subId := randStr(10)
	sub := &SubCallBack{
		SubID:    subId,
		Callback: f}
	e.topicSubMap[topic] = append(e.topicSubMap[topic], sub)
	return subId
}

func (e *DefaultEventBus) DeSub(topic string, subId string) {
	e.lock.RLock()
	defer e.lock.RUnlock()

	subs, ok := e.topicSubMap[topic]
	if !ok {
		return
	}
	for i, c := range subs {
		if c.SubID == subId {
			subs = append(subs[:i], subs[i+1:]...)
		}
	}
	e.topicSubMap[topic] = subs
}
