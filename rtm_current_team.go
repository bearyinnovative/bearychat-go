package bearychat

type RTMCurrentTeamService struct {
	rtm *RTMClient
}

func newRTMCurrentTeamService(rtm *RTMClient) error {
	rtm.CurrentTeam = &RTMCurrentTeamService{rtm}

	return nil
}

// Retrieves current team's information.
func (s *RTMCurrentTeamService) Info() (*Team, error) {
	team := new(Team)
	_, err := s.rtm.Get("v1/current_team.info", team)
	return team, err
}

// Retrieves current team's members.
func (s *RTMCurrentTeamService) Members() ([]*User, error) {
	members := []*User{}
	_, err := s.rtm.Get("v1/current_team.members?all=true", &members)
	return members, err
}

// Retrieves current team's channels.
func (s *RTMCurrentTeamService) Channels() ([]*Channel, error) {
	channels := []*Channel{}
	_, err := s.rtm.Get("v1/current_team.channels", &channels)
	return channels, err
}
