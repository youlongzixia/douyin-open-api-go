/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiSkuSyncReqSkus struct for PoiSkuSyncReqSkus
type PoiSkuSyncReqSkus struct {
	// 价格(人民币分)
	Price int64 `json:"price,omitempty"`
	// 在线状态 1 - 在线; 2 - 下线
	Status int64 `json:"status,omitempty"`
	// 库存数量
	Stock      int64                   `json:"stock,omitempty"`
	Attributes PoiSkuSyncReqAttributes `json:"attributes,omitempty"`
}
