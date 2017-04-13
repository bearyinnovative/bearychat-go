package bearychat

import "time"

type RTMContext struct {
	Loop RTMLoop

	uid string
}

func (c *RTMContext) UID() string {
	return c.uid
}

func NewRTMContext(token string) (*RTMContext, error) {
	rtmClient, err := NewRTMClient(token)
	if err != nil {
		return nil, err
	}

	user, wsHost, err := rtmClient.Start()
	if err != nil {
		return nil, err
	}

	rtmLoop, err := NewRTMLoop(wsHost)
	if err != nil {
		return nil, err
	}

	return &RTMContext{
		Loop: rtmLoop,
		uid:  user.Id,
	}, nil
}

func (c *RTMContext) Run() (error, chan RTMMessage, chan error) {
	err := c.Loop.Start()
	if err != nil {
		return err, nil, nil
	}
	defer c.Loop.Stop()

	go c.Loop.Keepalive(time.NewTicker(10 * time.Second))

	errC := c.Loop.ErrC()
	messageC, err := c.Loop.ReadC()
	if err != nil {
		return err, nil, nil
	}

	return nil, messageC, errC
}
