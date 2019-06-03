// Copyright 2019 Catena cyber All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package layers

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func FuzzPacket(data []byte) int {
	p := gopacket.NewPacket(data, layers.LinkTypeEthernet, gopacket.DecodeStreamsAsDatagrams)
	for _, l := range p.Layers() {
		gopacket.LayerString(l)
	}
	if p.ErrorLayer() != nil {
		return 0
	}
	return 1
}
