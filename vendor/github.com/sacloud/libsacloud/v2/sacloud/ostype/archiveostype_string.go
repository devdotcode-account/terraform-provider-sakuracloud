// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "stringer -type=ArchiveOSType"; DO NOT EDIT.

package ostype

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Custom-0]
	_ = x[CentOS-1]
	_ = x[CentOS8-2]
	_ = x[CentOS7-3]
	_ = x[CentOS6-4]
	_ = x[Ubuntu-5]
	_ = x[Ubuntu1804-6]
	_ = x[Ubuntu1604-7]
	_ = x[Debian-8]
	_ = x[Debian10-9]
	_ = x[Debian9-10]
	_ = x[CoreOS-11]
	_ = x[RancherOS-12]
	_ = x[K3OS-13]
	_ = x[Kusanagi-14]
	_ = x[SophosUTM-15]
	_ = x[FreeBSD-16]
	_ = x[Netwiser-17]
	_ = x[OPNsense-18]
	_ = x[Windows2016-19]
	_ = x[Windows2016RDS-20]
	_ = x[Windows2016RDSOffice-21]
	_ = x[Windows2016SQLServerWeb-22]
	_ = x[Windows2016SQLServerStandard-23]
	_ = x[Windows2016SQLServer2017Standard-24]
	_ = x[Windows2016SQLServerStandardAll-25]
	_ = x[Windows2016SQLServer2017StandardAll-26]
	_ = x[Windows2019-27]
}

const _ArchiveOSType_name = "CustomCentOSCentOS8CentOS7CentOS6UbuntuUbuntu1804Ubuntu1604DebianDebian10Debian9CoreOSRancherOSK3OSKusanagiSophosUTMFreeBSDNetwiserOPNsenseWindows2016Windows2016RDSWindows2016RDSOfficeWindows2016SQLServerWebWindows2016SQLServerStandardWindows2016SQLServer2017StandardWindows2016SQLServerStandardAllWindows2016SQLServer2017StandardAllWindows2019"

var _ArchiveOSType_index = [...]uint16{0, 6, 12, 19, 26, 33, 39, 49, 59, 65, 73, 80, 86, 95, 99, 107, 116, 123, 131, 139, 150, 164, 184, 207, 235, 267, 298, 333, 344}

func (i ArchiveOSType) String() string {
	if i < 0 || i >= ArchiveOSType(len(_ArchiveOSType_index)-1) {
		return "ArchiveOSType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ArchiveOSType_name[_ArchiveOSType_index[i]:_ArchiveOSType_index[i+1]]
}
