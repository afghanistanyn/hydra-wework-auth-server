package globals

import "errors"

var Config = struct {
	CookieSecret string
	WeworkConfig weworkConfig
	HydraConfig  hydraConfig
}{}

type weworkConfig struct {
	WeworkCropID  string
	WeworkAgentID string
	WeworkSecret  string
}

func (c *weworkConfig) Validate() error {

	if c.WeworkCropID == "" {
		return errors.New("wework corp id is missing")
	}

	if c.WeworkAgentID == "" {
		return errors.New("wework agent id is missing")
	}

	if c.WeworkSecret == "" {
		return errors.New("wework secret is missing")
	}

	return nil
}

type hydraConfig struct {
	HydraHost         string
	HydraMethod       string
	HydraAdminPort    string
	HydraClientID     string
	HydraClientSecret string
}

func (c *hydraConfig) Validate() error {

	if c.HydraHost == "" {
		return errors.New("hydra host is missing")
	}
	if c.HydraAdminPort == "" {
		return errors.New("hydra admin port is missing")
	}
	if c.HydraClientID == "" {
		return errors.New("hydra client id is missing")
	}
	if c.HydraClientSecret == "" {
		return errors.New("hydra client secret is missing")
	}

	return nil
}
