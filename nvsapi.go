package one

import (
	"bytes"
	"fmt"
)

// NetworkStatus is the API response to a /status call.
type NetworkStatus struct {
	Address           string `json:"address"`
	PublicIdentity    string `json:"publicIdentity"`
	WorldID           int64  `json:"worldId"`
	WorldTimestamp    int64  `json:"worldTimestamp"`
	Online            bool   `json:"online"`
	TCPFallbackActive bool   `json:"tcpFallbackActive"`
	// Enum: ALWAYS / TRUSTED / NEVER
	RelayPolicy  string `json:"relayPolicy"`
	VersionMajor int64  `json:"versionMajor"`
	VersionMinor int64  `json:"versionMinor"`
	VersionRev   int64  `json:"versionRev"`
	Version      string `json:"version"`
	Clock        int64  `json:"clock"`
}

// Network is the data structure that encapsulates a network.
type Network struct {
	ID                string             `json:"id,omitempty"`
	LegacyNetworkID   string             `json:"nwid,omitempty"`
	MAC               string             `json:"mac,omitempty"`
	Name              string             `json:"name,omitempty"`
	Status            string             `json:"status,omitempty"`
	Type              string             `json:"type,omitempty"`
	MTU               int                `json:"mtu,omitempty"`
	DHCP              bool               `json:"dhcp,omitempty"`
	Bridge            bool               `json:"bridge,omitempty"`
	BroadcastEnabled  bool               `json:"broadcastEnabled,omitempty"`
	PortError         int                `json:"portError,omitempty"`
	NetconfRevision   int64              `json:"netconfRevision,omitempty"`
	AssignedAddresses []string           `json:"assignedAddresses,omitempty"`
	Routes            []Route            `json:"routes,omitempty"`
	PortDeviceName    string             `json:"portDeviceName,omitempty"`
	AllowManaged      bool               `json:"allowManaged,omitempty"`
	AllowGlobal       bool               `json:"allowGlobal,omitempty"`
	AllowDefault      bool               `json:"allowDefault,omitempty"`
	AllowDNS          bool               `json:"allowDNS,omitempty"`
	V4AssignMode      V4AssignMode       `json:"v4AssignMode,omitempty"`
	IPAssignmentPools []IPAssignmentPool `json:"ipAssignmentPools,omitempty"`
}

type V4AssignMode struct {
	Zt bool `json:"zt"`
}

type IPAssignmentPool struct {
	IPRangeEnd   string `json:"ipRangeEnd"`
	IPRangeStart string `json:"ipRangeStart"`
}

// Route encapsulates network routes. See Network.
type Route struct {
	Target string `json:"target,omitempty"`
	Via    string `json:"via"`
	Flags  int64  `json:"flags"`
	Metric int64  `json:"metric"`
}

// Peer encapsulates ZeroTier One peers.
type Peer struct {
	Address      string `json:"address"`
	VersionMajor int64  `json:"versionMajor"`
	VersionMinor int64  `json:"versionMinor"`
	VersionRev   int64  `json:"versionRev"`
	Version      string `json:"version"`
	Latency      int64  `json:"latency"`
	// Enum: LEAF / UPSTREAM / ROOT / PLANET
	Role  string `json:"role"`
	Paths []Path `json:"paths"`
}

// Path is the path on the network.
type Path struct {
	Address       string `json:"address"`
	LastSend      int64  `json:"lastSend"`
	LastReceive   int64  `json:"lastReceive"`
	Active        bool   `json:"active"`
	Expired       bool   `json:"expired"`
	Preferred     bool   `json:"preferred"`
	TrustedPathID int64  `json:"trustedPathId"`
}

// Status returns the status of the ZeroTier One instance
func (c *Client) Status() (*NetworkStatus, error) {
	ns := &NetworkStatus{}
	return ns, c.wrapJSON("/status", ns)
}

// ListNetworks returns all networks that ZeroTier One knows about.
func (c *Client) ListNetworks() ([]*Network, error) {
	nws := []*Network{}
	return nws, c.wrapJSON("/network", &nws)
}

// GetNetwork queries a specific network.
func (c *Client) GetNetwork(id string) (*Network, error) {
	nw := &Network{}
	return nw, c.wrapJSON("/network/"+id, nw)
}

// JoinNetwork attempts to join a network. It will return an error if there is
// a problem contacting the service.
func (c *Client) JoinNetwork(id string) error {
	req, err := c.makeBaseReq("POST", "/network/"+id, bytes.NewBufferString("{}"))
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Response status was not 200, was %d", resp.StatusCode)
	}

	return nil
}

// LeaveNetwork attempts to leave a network. It will return an error if there is
// a problem contacting the service.
func (c *Client) LeaveNetwork(id string) error {
	req, err := c.makeBaseReq("DELETE", "/network/"+id, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("response status was not 200, was %d", resp.StatusCode)
	}

	return nil
}

// ListPeers queries the peers that ZeroTier One knows about.
func (c *Client) ListPeers() ([]*Peer, error) {
	peers := []*Peer{}
	return peers, c.wrapJSON("/peer", &peers)
}

// GetPeer queries a specific peer by address.
func (c *Client) GetPeer(address string) (*Peer, error) {
	p := &Peer{}
	return p, c.wrapJSON("/peer/"+address, p)
}
