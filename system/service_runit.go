package system

import (
	"fmt"
	"os"

	"github.com/aelsabbahy/goss/util"
)

type ServiceRunit struct {
	service string
}

func NewServiceRunit(service string, system *System, config util.Config) Service {
	return &ServiceRunit{service: service}
}

func (s *ServiceRunit) Service() string {
	return s.service
}

func (s *ServiceRunit) Exists() (bool, error) {
	if invalidService(s.service) {
		return false, nil
	}
	if _, err := os.Stat(fmt.Sprintf("/etc/service/%s/run", s.service)); err == nil {
		return true, err
	}
	return false, nil
}

func (s *ServiceRunit) Enabled() (bool, error) {
	// All Runit services that are present are enabled.  You could test if /etc/service/<name>/run is executable,
	// but it would be an error for it not to be executable, so then Exists() should fail.
	return true, nil
}

func (s *ServiceRunit) Running() (bool, error) {
	if invalidService(s.service) {
		return false, nil
	}
	cmd := util.NewCommand("sv", "status", s.service)
	cmd.Run()
	if cmd.Status == 0 {
		return true, cmd.Err
	}
	return false, nil
}
