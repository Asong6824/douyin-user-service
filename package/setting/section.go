package setting

import (
	"time"
)

type EngineSetting struct {
	HostPort    string
	ServiceName string
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type CacheSetting struct {
	Host          string
  	Port          string
  	Password      string
  	DB            int
  	MaxIdleConns  int
  	axActiveConns int
  	IdleTimeout   time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}