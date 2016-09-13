package bearychat

import "fmt"

type RTMUserService struct {
	rtm *RTMClient
}

func newRTMUserService(rtm *RTMClient) error {
	rtm.User = &RTMUserService{rtm}
	return nil
}

func (s *RTMUserService) Info(userId string) (*User, error) {
	user := new(User)
	resource := fmt.Sprintf("v1/user.info?user_id=%s", userId)
	_, err := s.rtm.Get(resource, user)
	return user, err
}
