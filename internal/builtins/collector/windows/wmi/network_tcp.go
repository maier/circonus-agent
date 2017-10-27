// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

type Win32_PerfRawData_Tcpip_TCPv4 struct {
	ConnectionFailures          uint32
	ConnectionsActive           uint32
	ConnectionsEstablished      uint32
	ConnectionsPassive          uint32
	ConnectionsReset            uint32
	Name                        string
	SegmentsPersec              uint32
	SegmentsReceivedPersec      uint32
	SegmentsRetransmittedPersec uint32
	SegmentsSentPersec          uint32
}

type Win32_PerfRawData_Tcpip_TCPv6 struct {
	ConnectionFailures          uint32
	ConnectionsActive           uint32
	ConnectionsEstablished      uint32
	ConnectionsPassive          uint32
	ConnectionsReset            uint32
	Name                        string
	SegmentsPersec              uint32
	SegmentsReceivedPersec      uint32
	SegmentsRetransmittedPersec uint32
	SegmentsSentPersec          uint32
}
