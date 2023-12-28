package setting

import (
)

type EngineSetting struct {
	HostPort    string
	ServiceName string
}


func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}