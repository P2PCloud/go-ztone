# go-ztone: a client to the ZeroTier One local configuration socket

`go-ztone` is for using against your local [ZeroTier One](https://www.zerotier.com) instance. You could use the "secret authtoken" described in [this document](https://github.com/zerotier/zerotierone#running) or `NewClientFromDefaultKey` method which would load `/var/lib/zerotier-one/authtoken.secret`.

```go
package main

import (
	"fmt"
	"os"

	one "github.com/p2pcloud/go-ztone"
)

func main() {
	c, err := NewClientFromDefaultKey()
    if err != nil {
		panic(err)
	}

	networks, err := c.Networks()
	if err != nil {
		panic(err)
	}

	peers, err := c.Peers()
	if err != nil {
		panic(err)
	}

	fmt.Println("Networks w/ MAC:")
	for _, network := range networks {
		fmt.Println(network.ID, network.MAC)
	}

	fmt.Println("Peers w/ Latency:")
	for _, peer := range peers {
		fmt.Println(peer.Address, peer.Latency)
	}
}
```

## Functionality

`ztone` has a few basic functions, most of them [listed here](https://github.com/zerotier/ZeroTierOne/blob/master/service/README.md#network-virtualization-service-api), as well as on the [GoDoc](https://pkg.go.dev/github.com/p2pcloud/go-ztone).

## Other methods
Check out [openapi docs by zerotier](https://docs.zerotier.com/openapi/servicev1.json)

## Example Code

You can see examples in the `examples/` directory. There are two:

- `list-things`: lists a few different properties and takes no arguments.
- `query-things`: takes a network ID and returns a few properties about it.

# Author

Originally created by Erik Hollensbe <github@hollensbe.org>

Maintained by [p2pcloud team](https://p2pcloud.io)

# License

BSD 3-Clause
