package main

import (
	"log"

	zt "github.com/p2pcloud/go-ztone"
)

func main() {
	client, err := zt.NewClientFromDefaultKey()
	if err != nil {
		panic(err)
	}
	ids, err := client.ControllerListNetworkIds()
	if err != nil {
		panic(err)
	}
	for _, netId := range ids {
		log.Println("network id:", netId)
		memberIds, err := client.ControllerListNetworkMemberIds(netId)
		if err != nil {
			panic(err)
		}
		for _, memberId := range memberIds {
			member, err := client.ControllerGetNetworkMember(netId, memberId)
			if err != nil {
				panic(err)
			}

			peer, err := client.GetPeer(member.Address)
			if err != nil {
				continue // not online now
			}

			log.Println("---")
			log.Println("Member info:")
			log.Println("Address:", member.Address)
			log.Println("ipAssignments:", member.IPAssignments)
			log.Println("Authorized:", member.Authorized)

			log.Println("Peer info:")
			log.Println("Latency:", peer.Latency)
			//search for ip in peer.Paths
		}
	}
}
