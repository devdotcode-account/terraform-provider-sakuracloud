// Copyright 2016-2020 The Libsacloud Authors
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

package api

/************************************************
  generated by IDE. for [SubnetAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *SubnetAPI) Reset() *SubnetAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *SubnetAPI) Offset(offset int) *SubnetAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *SubnetAPI) Limit(limit int) *SubnetAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *SubnetAPI) Include(key string) *SubnetAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *SubnetAPI) Exclude(key string) *SubnetAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *SubnetAPI) FilterBy(key string, value interface{}) *SubnetAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *SubnetAPI) FilterMultiBy(key string, value interface{}) *SubnetAPI {
	api.filterBy(key, value, true)
	return api
}

//func (api *SubnetAPI) WithNameLike(name string) *SubnetAPI {
//	return api.FilterBy("Name", name)
//}
//
//func (api *SubnetAPI) WithTag(tag string) *SubnetAPI {
//	return api.FilterBy("Tags.Name", tag)
//}
//func (api *SubnetAPI) WithTags(tags []string) *SubnetAPI {
//	return api.FilterBy("Tags.Name", []interface{}{tags})
//}

// func (api *SubnetAPI) WithSizeGib(size int) *SubnetAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *SubnetAPI) WithSharedScope() *SubnetAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *SubnetAPI) WithUserScope() *SubnetAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *SubnetAPI) SortBy(key string, reverse bool) *SubnetAPI {
	api.sortBy(key, reverse)
	return api
}

//func (api *SubnetAPI) SortByName(reverse bool) *SubnetAPI {
//	api.sortByName(reverse)
//	return api
//}

// func (api *SubnetAPI) SortBySize(reverse bool) *SubnetAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *SubnetAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *SubnetAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *SubnetAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *SubnetAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *SubnetAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルター
func (api *SubnetAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *SubnetAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

//func (api *SubnetAPI) SetNameLike(name string) {
//	api.FilterBy("Name", name)
//}
//
//func (api *SubnetAPI) SetTag(tag string) {
//	api.FilterBy("Tags.Name", tag)
//}
//func (api *SubnetAPI) SetTags(tags []string) {
//	api.FilterBy("Tags.Name", []interface{}{tags})
//}

// func (api *SubnetAPI) SetSizeGib(size int) {
// 	api.FilterBy("SizeMB", size*1024)
// }

// func (api *SubnetAPI) SetSharedScope() {
// 	api.FilterBy("Scope", "shared")
// }

// func (api *SubnetAPI) SetUserScope() {
// 	api.FilterBy("Scope", "user")
// }

// SetSortBy 指定キーでのソート
func (api *SubnetAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

//func (api *SubnetAPI) SetSortByName(reverse bool) {
//	api.sortByName(reverse)
//}

// func (api *SubnetAPI) SetSortBySize(reverse bool) {
// 	api.sortBy("SizeMB", reverse)
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

//func (api *SubnetAPI) New() *sacloud.Subnet {
//	return &sacloud.Subnet{}
//}

//func (api *SubnetAPI) Create(value *sacloud.Subnet) (*sacloud.Subnet, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.create(api.createRequest(value), res)
//	})
//}

// Read 読み取り
func (api *SubnetAPI) Read(id sacloud.ID) (*sacloud.Subnet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

//func (api *SubnetAPI) Update(id sacloud.ID, value *sacloud.Subnet) (*sacloud.Subnet, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.update(id, api.createRequest(value), res)
//	})
//}
//
//func (api *SubnetAPI) Delete(id sacloud.ID) (*sacloud.Subnet, error) {
//	return api.request(func(res *sacloud.Response) error {
//		return api.delete(id, nil, res)
//	})
//}

/************************************************
  Inner functions
************************************************/

func (api *SubnetAPI) setStateValue(setFunc func(*sacloud.Request)) *SubnetAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *SubnetAPI) request(f func(*sacloud.Response) error) (*sacloud.Subnet, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Subnet, nil
}

func (api *SubnetAPI) createRequest(value *sacloud.Subnet) *sacloud.Request {
	req := &sacloud.Request{}
	req.Subnet = value
	return req
}
