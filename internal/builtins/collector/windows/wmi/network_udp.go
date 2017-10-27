// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

type Win32_PerfRawData_Tcpip_UDPv4 struct {
	DatagramsNoPortPersec   uint32
	DatagramsPersec         uint32
	DatagramsReceivedErrors uint32
	DatagramsReceivedPersec uint32
	DatagramsSentPersec     uint32
	Name                    string
}

type Win32_PerfRawData_Tcpip_UDPv6 struct {
	DatagramsNoPortPersec   uint32
	DatagramsPersec         uint32
	DatagramsReceivedErrors uint32
	DatagramsReceivedPersec uint32
	DatagramsSentPersec     uint32
	Name                    string
}
