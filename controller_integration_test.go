package one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNetworkAndAuthorize(t *testing.T) {
	c, err := NewClientFromDefaultKey()
	require.NoError(t, err)

	//create network

	newNet, err := c.ControllerCreateNetwork(SaneNetworkDefaults)
	require.NoError(t, err)

	require.Equal(t, "default", newNet.Name)
	require.Equal(t, "10.244.0.0/16", newNet.Routes[0].Target)
	require.Equal(t, true, newNet.V4AssignMode.Zt)
	require.Equal(t, "10.244.0.1", newNet.IPAssignmentPools[0].IPRangeStart)
	require.Equal(t, "10.244.255.254", newNet.IPAssignmentPools[0].IPRangeEnd)

	//get network
	newNet, err = c.ControllerGetNetwork(newNet.ID)
	require.NoError(t, err)

	require.Equal(t, "10.244.0.1", newNet.IPAssignmentPools[0].IPRangeStart)
	require.Equal(t, "10.244.255.254", newNet.IPAssignmentPools[0].IPRangeEnd)

	//check that the network is in the list of networks
	networkIds, err := c.ControllerListNetworkIds()
	require.NoError(t, err)
	require.Subset(t, networkIds, []string{newNet.ID})

	//authorize
	member, err := c.ControllerAuthorizeMember(newNet.ID, "000000000a", true)
	require.NoError(t, err)

	require.Equal(t, "000000000a", member.ID)
	require.Equal(t, true, member.Authorized)
}
