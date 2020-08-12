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
  generated by IDE. for [MobileGatewayAPI]
************************************************/

import (
	"github.com/sacloud/libsacloud/sacloud"
)

/************************************************
   To support fluent interface for Find()
************************************************/

// Reset 検索条件のリセット
func (api *MobileGatewayAPI) Reset() *MobileGatewayAPI {
	api.reset()
	return api
}

// Offset オフセット
func (api *MobileGatewayAPI) Offset(offset int) *MobileGatewayAPI {
	api.offset(offset)
	return api
}

// Limit リミット
func (api *MobileGatewayAPI) Limit(limit int) *MobileGatewayAPI {
	api.limit(limit)
	return api
}

// Include 取得する項目
func (api *MobileGatewayAPI) Include(key string) *MobileGatewayAPI {
	api.include(key)
	return api
}

// Exclude 除外する項目
func (api *MobileGatewayAPI) Exclude(key string) *MobileGatewayAPI {
	api.exclude(key)
	return api
}

// FilterBy 指定キーでのフィルター
func (api *MobileGatewayAPI) FilterBy(key string, value interface{}) *MobileGatewayAPI {
	api.filterBy(key, value, false)
	return api
}

// FilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *MobileGatewayAPI) FilterMultiBy(key string, value interface{}) *MobileGatewayAPI {
	api.filterBy(key, value, true)
	return api
}

// WithNameLike 名称条件
func (api *MobileGatewayAPI) WithNameLike(name string) *MobileGatewayAPI {
	return api.FilterBy("Name", name)
}

// WithTag タグ条件
func (api *MobileGatewayAPI) WithTag(tag string) *MobileGatewayAPI {
	return api.FilterBy("Tags.Name", tag)
}

// WithTags タグ(複数)条件
func (api *MobileGatewayAPI) WithTags(tags []string) *MobileGatewayAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *MobileGatewayAPI) WithSizeGib(size int) *MobileGatewayAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *MobileGatewayAPI) WithSharedScope() *MobileGatewayAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *MobileGatewayAPI) WithUserScope() *MobileGatewayAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

// SortBy 指定キーでのソート
func (api *MobileGatewayAPI) SortBy(key string, reverse bool) *MobileGatewayAPI {
	api.sortBy(key, reverse)
	return api
}

// SortByName 名称でのソート
func (api *MobileGatewayAPI) SortByName(reverse bool) *MobileGatewayAPI {
	api.sortByName(reverse)
	return api
}

// func (api *MobileGatewayAPI) SortBySize(reverse bool) *MobileGatewayAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
   To support Setxxx interface for Find()
************************************************/

// SetEmpty 検索条件のリセット
func (api *MobileGatewayAPI) SetEmpty() {
	api.reset()
}

// SetOffset オフセット
func (api *MobileGatewayAPI) SetOffset(offset int) {
	api.offset(offset)
}

// SetLimit リミット
func (api *MobileGatewayAPI) SetLimit(limit int) {
	api.limit(limit)
}

// SetInclude 取得する項目
func (api *MobileGatewayAPI) SetInclude(key string) {
	api.include(key)
}

// SetExclude 除外する項目
func (api *MobileGatewayAPI) SetExclude(key string) {
	api.exclude(key)
}

// SetFilterBy 指定キーでのフィルター
func (api *MobileGatewayAPI) SetFilterBy(key string, value interface{}) {
	api.filterBy(key, value, false)
}

// SetFilterMultiBy 任意項目でのフィルタ(完全一致 OR条件)
func (api *MobileGatewayAPI) SetFilterMultiBy(key string, value interface{}) {
	api.filterBy(key, value, true)
}

// SetNameLike 名称条件
func (api *MobileGatewayAPI) SetNameLike(name string) {
	api.FilterBy("Name", name)
}

// SetTag タグ条件
func (api *MobileGatewayAPI) SetTag(tag string) {
	api.FilterBy("Tags.Name", tag)
}

// SetTags タグ(複数)条件
func (api *MobileGatewayAPI) SetTags(tags []string) {
	api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *MobileGatewayAPI) SetSizeGib(size int) {
// 	api.FilterBy("SizeMB", size*1024)
// }

// func (api *MobileGatewayAPI) SetSharedScope() {
// 	api.FilterBy("Scope", "shared")
// }

// func (api *MobileGatewayAPI) SetUserScope() {
// 	api.FilterBy("Scope", "user")
// }

// SetSortBy 指定キーでのソート
func (api *MobileGatewayAPI) SetSortBy(key string, reverse bool) {
	api.sortBy(key, reverse)
}

// SetSortByName 名称でのソート
func (api *MobileGatewayAPI) SetSortByName(reverse bool) {
	api.sortByName(reverse)
}

// func (api *MobileGatewayAPI) SortBySize(reverse bool) {
// 	api.sortBy("SizeMB", reverse)
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *MobileGatewayAPI) New() *sacloud.MobileGateway {
// 	return &sacloud.MobileGateway{}
// }

// func (api *MobileGatewayAPI) Create(value *sacloud.MobileGateway) (*sacloud.MobileGateway, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *MobileGatewayAPI) Read(id string) (*sacloud.MobileGateway, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *MobileGatewayAPI) Update(id string, value *sacloud.MobileGateway) (*sacloud.MobileGateway, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *MobileGatewayAPI) Delete(id string) (*sacloud.MobileGateway, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *MobileGatewayAPI) setStateValue(setFunc func(*sacloud.Request)) *MobileGatewayAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *MobileGatewayAPI) request(f func(*sacloud.Response) error) (*sacloud.MobileGateway, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.MobileGateway, nil
//}
//
//func (api *MobileGatewayAPI) createRequest(value *sacloud.MobileGateway) *sacloud.Request {
//	req := &sacloud.Request{}
//	req.MobileGateway = value
//	return req
//}