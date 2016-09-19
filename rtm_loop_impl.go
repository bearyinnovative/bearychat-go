package bearychat

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type rtmLoop struct {
	wsHost string
	conn   *websocket.Conn
	state  RTMLoopState
	callId uint64
	llock  *sync.RWMutex // lock for properties below

	rtmCBacklog int
	rtmC        chan RTMMessage
	errC        chan error
}

type rtmLoopSetter func(*rtmLoop) error

// Set RTM message chan backlog.
func WithRTMLoopBacklog(backlog int) rtmLoopSetter {
	return func(r *rtmLoop) error {
		r.rtmCBacklog = backlog
		return nil
	}
}

func NewRTMLoop(wsHost string, setters ...rtmLoopSetter) (*rtmLoop, error) {
	l := &rtmLoop{
		wsHost: wsHost,
		state:  RTMLoopStateClosed,
		callId: 0,
		llock:  &sync.RWMutex{},

		errC: make(chan error, 1024),
	}
	for _, setter := range setters {
		if err := setter(l); err != nil {
			return nil, err
		}
	}

	if l.rtmCBacklog <= 0 {
		l.rtmC = make(chan RTMMessage)
	} else {
		l.rtmC = make(chan RTMMessage, l.rtmCBacklog)
	}

	return l, nil
}

func (l *rtmLoop) Start() error {
	l.llock.Lock()
	defer l.llock.Unlock()

	conn, _, err := websocket.DefaultDialer.Dial(l.wsHost, nil)
	if err != nil {
		return err
	}

	l.conn = conn
	l.state = RTMLoopStateOpen

	go l.readMessage()

	return nil
}

func (l *rtmLoop) Stop() error {
	return nil
}

func (l *rtmLoop) State() RTMLoopState {
	l.llock.RLock()
	defer l.llock.RUnlock()

	return l.state
}

func (l *rtmLoop) Ping() error {
	return l.Send(RTMMessage{"type": RTMMessageTypePing})
}

func (l *rtmLoop) Keepalive(interval *time.Ticker) error {
	defer interval.Stop()
	for {
		select {
		case <-interval.C:
			if err := l.Ping(); err != nil {
				return errors.Wrap(err, "keepalive closed")
			}
		}
	}
}

func (l *rtmLoop) Send(m RTMMessage) error {
	if l.State() != RTMLoopStateOpen {
		return ErrRTMLoopClosed
	}

	if _, hasCallId := m["call_id"]; !hasCallId {
		m["call_id"] = l.advanceCallId()
	}

	rawMessage, err := json.Marshal(m)
	if err != nil {
		return errors.Wrap(err, "encode message failed")
	}

	if err := l.conn.WriteMessage(websocket.TextMessage, rawMessage); err != nil {
		return errors.Wrap(err, "write socket failed")
	}

	return nil
}

func (l *rtmLoop) ReadC() (chan RTMMessage, error) {
	if l.State() != RTMLoopStateOpen {
		return nil, ErrRTMLoopClosed
	}

	return l.rtmC, nil
}

func (l *rtmLoop) ErrC() chan error {
	return l.errC
}

// Listen & read message from BearyChat
func (l *rtmLoop) readMessage() {
	for {
		if l.State() == RTMLoopStateClosed {
			return
		}

		_, rawMessage, err := l.conn.ReadMessage()
		if err != nil {
			l.errC <- errors.Wrap(err, "read socket failed")
			continue
		}

		message := RTMMessage{}
		if err = json.Unmarshal(rawMessage, &message); err != nil {
			l.errC <- errors.Wrap(err, "decode message failed")
			continue
		}

		l.rtmC <- message
	}
}

func (l *rtmLoop) advanceCallId() uint64 {
	return atomic.AddUint64(&l.callId, 1)
}
