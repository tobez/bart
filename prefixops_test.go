// Copyright (c) 2024 Karl Gaissmaier
// SPDX-License-Identifier: MIT

package bart

import (
	"fmt"
	"net/netip"
	"os"
	"testing"
)

var p = netip.MustParsePrefix

var poInput = []netip.Prefix{
	p("10.0.0.0/8"),
	p("10.0.0.0/16"),
	p("10.0.0.0/24"),
	p("10.0.1.0/24"),
	p("10.0.0.0/16"),
	p("10.0.0.0/24"),
	p("10.0.1.0/24"),
	p("10.0.0.0/24"),
	p("10.0.1.0/24"),
	p("10.0.1.0/24"),
}

func TestPrefixOps(t *testing.T) {
	it := new(Table[int])
	for _, item := range poInput {
		iv, found := it.GetOrInsert(item)
		if found {
			fmt.Printf("%v was there as %d\n", item, *iv)
		} else {
			fmt.Printf("%v was not there\n", item)
		}
		*iv++
	}
	it.Fprint(os.Stdout)

	fmt.Println()

	// Output:
	// ▼
	// ├─ 10.0.0.0/8 (9.9.9.9)
	// │  ├─ 10.0.0.0/24 (8.8.8.8)
	// │  └─ 10.0.1.0/24 (10.0.0.0)
	// ├─ 127.0.0.0/8 (127.0.0.1)
	// │  └─ 127.0.0.1/32 (127.0.0.1)
	// ├─ 169.254.0.0/16 (10.0.0.0)
	// ├─ 172.16.0.0/12 (8.8.8.8)
	// └─ 192.168.0.0/16 (9.9.9.9)
	//    └─ 192.168.1.0/24 (127.0.0.1)
}
