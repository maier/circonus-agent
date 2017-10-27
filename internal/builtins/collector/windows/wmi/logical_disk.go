// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

// +build windows

package wmi

/*

__GENUS                 : 2
__CLASS                 : Win32_PerfFormattedData_PerfDisk_LogicalDisk
__SUPERCLASS            : Win32_PerfFormattedData
__DYNASTY               : CIM_StatisticalInformation
__RELPATH               : Win32_PerfFormattedData_PerfDisk_LogicalDisk.Name="_Total"
__PROPERTY_COUNT        : 32
__DERIVATION            : {Win32_PerfFormattedData, Win32_Perf, CIM_StatisticalInformation}
__SERVER                : DESKTOP-BILB01C
__NAMESPACE             : root\cimv2
__PATH                  : \\DESKTOP-BILB01C\root\cimv2:Win32_PerfFormattedData_PerfDisk_LogicalDisk.Name="_Total"
AvgDiskBytesPerRead     : 0
AvgDiskBytesPerTransfer : 4096
AvgDiskBytesPerWrite    : 4096
AvgDiskQueueLength      : 0
AvgDiskReadQueueLength  : 0
AvgDisksecPerRead       : 0
AvgDisksecPerTransfer   : 0
AvgDisksecPerWrite      : 0
AvgDiskWriteQueueLength : 0
Caption                 :
CurrentDiskQueueLength  : 0
Description             :
DiskBytesPersec         : 16172
DiskReadBytesPersec     : 0
DiskReadsPersec         : 0
DiskTransfersPersec     : 3
DiskWriteBytesPersec    : 16172
DiskWritesPersec        : 3
FreeMegabytes           : 201468
Frequency_Object        :
Frequency_PerfTime      :
Frequency_Sys100NS      :
Name                    : _Total
PercentDiskReadTime     : 0
PercentDiskTime         : 4
PercentDiskWriteTime    : 4
PercentFreeSpace        : 82
PercentIdleTime         : 95
SplitIOPerSec           : 0
Timestamp_Object        :
Timestamp_PerfTime      :
Timestamp_Sys100NS      :
PSComputerName          : DESKTOP-BILB01C

__GENUS                 : 2
__CLASS                 : Win32_PerfFormattedData_PerfDisk_LogicalDisk
__SUPERCLASS            : Win32_PerfFormattedData
__DYNASTY               : CIM_StatisticalInformation
__RELPATH               : Win32_PerfFormattedData_PerfDisk_LogicalDisk.Name="C:"
__PROPERTY_COUNT        : 32
__DERIVATION            : {Win32_PerfFormattedData, Win32_Perf, CIM_StatisticalInformation}
__SERVER                : DESKTOP-BILB01C
__NAMESPACE             : root\cimv2
__PATH                  : \\DESKTOP-BILB01C\root\cimv2:Win32_PerfFormattedData_PerfDisk_LogicalDisk.Name="C:"
AvgDiskBytesPerRead     : 0
AvgDiskBytesPerTransfer : 4096
AvgDiskBytesPerWrite    : 4096
AvgDiskQueueLength      : 0
AvgDiskReadQueueLength  : 0
AvgDisksecPerRead       : 0
AvgDisksecPerTransfer   : 0
AvgDisksecPerWrite      : 0
AvgDiskWriteQueueLength : 0
Caption                 :
CurrentDiskQueueLength  : 0
Description             :
DiskBytesPersec         : 16172
DiskReadBytesPersec     : 0
DiskReadsPersec         : 0
DiskTransfersPersec     : 3
DiskWriteBytesPersec    : 16172
DiskWritesPersec        : 3
FreeMegabytes           : 201014
Frequency_Object        :
Frequency_PerfTime      :
Frequency_Sys100NS      :
Name                    : C:
PercentDiskReadTime     : 0
PercentDiskTime         : 9
PercentDiskWriteTime    : 9
PercentFreeSpace        : 82
PercentIdleTime         : 95
SplitIOPerSec           : 0
Timestamp_Object        :
Timestamp_PerfTime      :
Timestamp_Sys100NS      :
PSComputerName          : DESKTOP-BILB01C

__GENUS                 : 2
__CLASS                 : Win32_PerfFormattedData_PerfDisk_LogicalDisk
__SUPERCLASS            : Win32_PerfFormattedData
__DYNASTY               : CIM_StatisticalInformation
__RELPATH               : Win32_PerfFormattedData_PerfDisk_LogicalDisk.Name="HarddiskVolume4"
__PROPERTY_COUNT        : 32
__DERIVATION            : {Win32_PerfFormattedData, Win32_Perf, CIM_StatisticalInformation}
__SERVER                : DESKTOP-BILB01C
__NAMESPACE             : root\cimv2
__PATH                  : \\DESKTOP-BILB01C\root\cimv2:Win32_PerfFormattedData_PerfDisk_LogicalDisk.Name="HarddiskVolume4"
AvgDiskBytesPerRead     : 0
AvgDiskBytesPerTransfer : 0
AvgDiskBytesPerWrite    : 0
AvgDiskQueueLength      : 0
AvgDiskReadQueueLength  : 0
AvgDisksecPerRead       : 0
AvgDisksecPerTransfer   : 0
AvgDisksecPerWrite      : 0
AvgDiskWriteQueueLength : 0
Caption                 :
CurrentDiskQueueLength  : 0
Description             :
DiskBytesPersec         : 0
DiskReadBytesPersec     : 0
DiskReadsPersec         : 0
DiskTransfersPersec     : 0
DiskWriteBytesPersec    : 0
DiskWritesPersec        : 0
FreeMegabytes           : 454
Frequency_Object        :
Frequency_PerfTime      :
Frequency_Sys100NS      :
Name                    : HarddiskVolume4
PercentDiskReadTime     : 0
PercentDiskTime         : 0
PercentDiskWriteTime    : 0
PercentFreeSpace        : 54
PercentIdleTime         : 96
SplitIOPerSec           : 0
Timestamp_Object        :
Timestamp_PerfTime      :
Timestamp_Sys100NS      :
PSComputerName          : DESKTOP-BILB01C

*/
