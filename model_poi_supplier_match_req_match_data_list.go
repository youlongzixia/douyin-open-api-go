/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiSupplierMatchReqMatchDataList struct for PoiSupplierMatchReqMatchDataList
type PoiSupplierMatchReqMatchDataList struct {
	// POI地址
	Address string `json:"address,omitempty"`
	// 高德POI ID
	AmapId string `json:"amap_id,omitempty"`
	// POI所在城市
	City string `json:"city,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty"`
	// POI名称
	PoiName string `json:"poi_name,omitempty"`
	// POI所在省份
	Province string `json:"province,omitempty"`
	// 第三方商户ID
	SupplierExtId string `json:"supplier_ext_id,omitempty"`
}
