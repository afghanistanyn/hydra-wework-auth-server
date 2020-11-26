package wework

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type UserStatus int

const (
	UserActive   UserStatus = 1
	UserDisabled UserStatus = 2
	UserInactive UserStatus = 3
)

type GetUserResponse struct {
	Code           int        `json:"errcode,omitempty"`
	Message        string     `json:"errmsg,omitempty"`
	UserID         string     `json:"userid,omitempty"`
	Name           string     `json:"name,omitempty"`
	Mobile         string     `json:"mobile,omitempty"`
	ThumbAvatar    string     `json:"avatar,omitempty"`
	Email          string     `json:"email,omitempty"`
	Position       string     `json:"position,omitempty"`
	Gender         string     `json:"gender,omitempty"`
	Status         UserStatus `json:"status,omitempty"`
	MainDepartment uint8      `json:"main_department"`
}

func (c *Client) GetUser(uid string) (*GetUserResponse, error) {
	q := url.Values{}
	q.Set("userid", uid)

	//自建应用可以读取该应用设置的可见范围内的成员信息
	//应用须拥有指定成员的查看权限
	u := "https://qyapi.weixin.qq.com/cgi-bin/user/get?" + q.Encode()

	var resp GetUserResponse
	if err := c.getJSON(u, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("Get user error: %v %v", resp.Code, resp.Message)
	}

	return &resp, nil
}

func (c *Client) collectUserInfo(uid string, userInfo map[string]interface{}) error {
	userResp, err := c.GetUser(uid)
	if err != nil {
		return fmt.Errorf("Get wework user failed. %v", err)
	}

	if userResp.Status != UserActive {
		return fmt.Errorf("User is not active: %s", uid)
	}

	userInfo["username"] = userResp.UserID
	userInfo["name"] = userResp.Mobile
	userInfo["email"] = userResp.Email
	userInfo["position"] = userResp.Position
	userInfo["email_verified"] = true
	return nil
}

func (s *Client) getTokenVars(uid string) (map[string]interface{}, error) {
	vars := make(map[string]interface{})
	if err := s.collectUserInfo(uid, vars); err != nil {
		return nil, err
	}
	log.Infof("User authenticated. %v", vars)
	return vars, nil
}
