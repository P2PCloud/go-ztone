package one

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type ControllerNetworkMember struct {
	ActiveBridge                 bool          `json:"activeBridge,omitempty"`
	Address                      string        `json:"address,omitempty"`
	AuthenticationExpiryTime     int           `json:"authenticationExpiryTime,omitempty"`
	Authorized                   bool          `json:"authorized,omitempty"`
	Capabilities                 []interface{} `json:"capabilities,omitempty"`
	CreationTime                 int64         `json:"creationTime,omitempty"`
	ID                           string        `json:"id,omitempty"`
	IPAssignments                []interface{} `json:"ipAssignments,omitempty"`
	LastAuthorizedCredential     interface{}   `json:"lastAuthorizedCredential,omitempty"`
	LastAuthorizedCredentialType string        `json:"lastAuthorizedCredentialType,omitempty"`
	LastAuthorizedTime           int64         `json:"lastAuthorizedTime,omitempty"`
	LastDeauthorizedTime         int           `json:"lastDeauthorizedTime,omitempty"`
	NoAutoAssignIps              bool          `json:"noAutoAssignIps,omitempty"`
	Nwid                         string        `json:"nwid,omitempty"`
	Objtype                      string        `json:"objtype,omitempty"`
	RemoteTraceLevel             int           `json:"remoteTraceLevel,omitempty"`
	RemoteTraceTarget            interface{}   `json:"remoteTraceTarget,omitempty"`
	Revision                     int           `json:"revision,omitempty"`
	SsoExempt                    bool          `json:"ssoExempt,omitempty"`
	Tags                         []interface{} `json:"tags,omitempty"`
	VMajor                       int           `json:"vMajor,omitempty"`
	VMinor                       int           `json:"vMinor,omitempty"`
	VProto                       int           `json:"vProto,omitempty"`
	VRev                         int           `json:"vRev,omitempty"`
}

//Returns public id from status
func (c *Client) GetPublicId() (string, error) {
	if c.publicId == "" {
		status, err := c.Status()
		if err != nil {
			return "", err
		}
		c.publicId = strings.Split(status.PublicIdentity, ":")[0]
	}
	return c.publicId, nil
}

func (c *Client) ControllerGetNetwork(networkId string) (*Network, error) {
	url := fmt.Sprintf("/controller/network/%s", networkId)
	var value Network
	return &value, c.wrapJSON(url, &value)
}

func (c *Client) ControllerCreateNetwork(network *Network) (*Network, error) {
	publicId, err := c.GetPublicId()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(network)
	if err != nil {
		return nil, err
	}

	req, err := c.makeBaseReq("POST", "/controller/network/"+publicId+"______", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status was not 200, was %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Network{}
	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//shorthand for a quick member authorization
func (c *Client) ControllerAuthorizeMember(networkId string, memberId string, authorized bool) (*ControllerNetworkMember, error) {
	return c.ControllerUpdateNetworkMember(networkId, &ControllerNetworkMember{
		ID:         memberId,
		Authorized: true,
		Address:    memberId,
	})
}

func (c *Client) ControllerUpdateNetworkMember(networkId string, netMember *ControllerNetworkMember) (*ControllerNetworkMember, error) {
	data, err := json.Marshal(netMember)
	if err != nil {
		return nil, err
	}

	if netMember.ID == "" {
		return nil, fmt.Errorf("netMember.ID is empty")
	}

	req, err := c.makeBaseReq(
		"POST",
		fmt.Sprintf("/controller/network/%s/member/%s", networkId, netMember.ID),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status was not 200, was %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &ControllerNetworkMember{}
	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Lists network ids on this controller
func (c *Client) ControllerListNetworkIds() ([]string, error) {
	var values []string
	return values, c.wrapJSON("/controller/network", &values)
}

func (c *Client) ControllerGetNetworkMember(networkId string, memberId string) (*ControllerNetworkMember, error) {
	url := fmt.Sprintf("/controller/network/%s/member/%s", networkId, memberId)
	var value ControllerNetworkMember
	return &value, c.wrapJSON(url, &value)
}

// Lists network ids on this controller
func (c *Client) ControllerListNetworkMemberIds(networkId string) ([]string, error) {
	var values map[string]int
	url := fmt.Sprintf("/controller/network/%s/member", networkId)
	err := c.wrapJSON(url, &values)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(values))

	i := 0
	for k := range values {
		ids[i] = k
		i++
	}

	return ids, nil
}

var SaneNetworkDefaults = &Network{
	Name: "default",
	Routes: []Route{
		Route{Target: "10.244.0.0/16"},
	},
	V4AssignMode: V4AssignMode{
		Zt: true,
	},
	IPAssignmentPools: []IPAssignmentPool{
		IPAssignmentPool{
			IPRangeStart: "10.244.0.1",
			IPRangeEnd:   "10.244.255.254",
		},
	},
}
