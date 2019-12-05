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

package types

// EProxyLBProxyMode エンハンスドロードバランサでのプロキシ方式
type EProxyLBProxyMode string

// ProxyLBProxyModes エンハンスドロードバランサでのプロキシ方式
var ProxyLBProxyModes = struct {
	// Unknown 不明
	Unknown EProxyLBProxyMode
	// HTTP .
	HTTP EProxyLBProxyMode
	// HTTPS .
	HTTPS EProxyLBProxyMode
	// TCP .
	TCP EProxyLBProxyMode
}{
	Unknown: EProxyLBProxyMode(""),
	HTTP:    EProxyLBProxyMode("http"),
	HTTPS:   EProxyLBProxyMode("https"),
	TCP:     EProxyLBProxyMode("tcp"),
}

// String EProxyLBProxyModeの文字列表現
func (m EProxyLBProxyMode) String() string {
	return string(m)
}
