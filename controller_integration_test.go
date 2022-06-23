package one

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNetworkAndAuthorize(t *testing.T) {
	c, err := NewClientFromDefaultKey()
	require.NoError(t, err)

	//create network

	newNet, err := c.ControllerCreateNetwork(&Network{
		Name: "TestCreateNetworkAndAuthorize",
		Routes: []Route{
			Route{Target: "10.88.239.0/24"},
		},
		V4AssignMode: V4AssignMode{
			Zt: true,
		},
		IPAssignmentPools: []IPAssignmentPool{
			IPAssignmentPool{
				IPRangeStart: "10.88.239.1",
				IPRangeEnd:   "10.88.239.254",
			},
		},
	})
	require.NoError(t, err)

	require.Equal(t, "TestCreateNetworkAndAuthorize", newNet.Name)
	require.Equal(t, "10.88.239.0/24", newNet.Routes[0].Target)
	require.Equal(t, true, newNet.V4AssignMode.Zt)
	require.Equal(t, "10.88.239.1", newNet.IPAssignmentPools[0].IPRangeStart)
	require.Equal(t, "10.88.239.254", newNet.IPAssignmentPools[0].IPRangeEnd)

	//check that the network is in the list of networks
	networkIds, err := c.ControllerListNetworkIds()
	require.NoError(t, err)
	require.Subset(t, *networkIds, []string{newNet.ID})

	//authorize
	member, err := c.ControllerAuthorizeMember(newNet.ID, "000000000a", true)
	require.NoError(t, err)

	require.Equal(t, "000000000a", member.ID)
	require.Equal(t, true, member.Authorized)
}
