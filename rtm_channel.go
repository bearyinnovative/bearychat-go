package bearychat

import "fmt"

type RTMChannelService struct {
	rtm *RTMClient
}

func newRTMChannelService(rtm *RTMClient) error {
	rtm.Channel = &RTMChannelService{rtm}
	return nil
}

func (s *RTMChannelService) Info(channelId string) (*Channel, error) {
	channel := new(Channel)
	resource := fmt.Sprintf("v1/channel.info?channel_id=%s", channelId)
	_, err := s.rtm.Get(resource, channel)
	return channel, err
}
