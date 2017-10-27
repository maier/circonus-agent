// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

type Win32_PerfRawData_Tcpip_IPv4 struct {
	DatagramsForwardedPersec         uint32
	DatagramsOutboundDiscarded       uint32
	DatagramsOutboundNoRoute         uint32
	DatagramsPersec                  uint32
	DatagramsReceivedAddressErrors   uint32
	DatagramsReceivedDeliveredPersec uint32
	DatagramsReceivedDiscarded       uint32
	DatagramsReceivedHeaderErrors    uint32
	DatagramsReceivedPersec          uint32
	DatagramsReceivedUnknownProtocol uint32
	DatagramsSentPersec              uint32
	FragmentationFailures            uint32
	FragmentedDatagramsPersec        uint32
	FragmentReassemblyFailures       uint32
	FragmentsCreatedPersec           uint32
	FragmentsReassembledPersec       uint32
	FragmentsReceivedPersec          uint32
	Name                             string
}

type Win32_PerfRawData_Tcpip_IPv6 struct {
	DatagramsForwardedPersec         uint32
	DatagramsOutboundDiscarded       uint32
	DatagramsOutboundNoRoute         uint32
	DatagramsPersec                  uint32
	DatagramsReceivedAddressErrors   uint32
	DatagramsReceivedDeliveredPersec uint32
	DatagramsReceivedDiscarded       uint32
	DatagramsReceivedHeaderErrors    uint32
	DatagramsReceivedPersec          uint32
	DatagramsReceivedUnknownProtocol uint32
	DatagramsSentPersec              uint32
	FragmentationFailures            uint32
	FragmentedDatagramsPersec        uint32
	FragmentReassemblyFailures       uint32
	FragmentsCreatedPersec           uint32
	FragmentsReassembledPersec       uint32
	FragmentsReceivedPersec          uint32
	Name                             string
}
