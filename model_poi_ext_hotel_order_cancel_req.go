/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// PoiExtHotelOrderCancelReq
type PoiExtHotelOrderCancelReq struct {
	// 抖音订单号
	OrderId string `json:"order_id,omitempty"`
	// 订单状态。0 - 未支付, 1 - 已支付
	OrderStatus int64 `json:"order_status,omitempty"`
	// 接入方商铺ID
	SupplierExtId string `json:"supplier_ext_id,omitempty"`
	// 接入方订单号
	OrderExtId string                               `json:"order_ext_id,omitempty"`
	DatePrice  []PoiExtHotelOrderCommitReqDatePrice `json:"date_price,omitempty"`
}
