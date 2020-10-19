/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// InlineObject34 struct for InlineObject34
type InlineObject34 struct {
	// 订单类型- 201 预约点餐订单, 202 餐厅预定订单, 203 餐厅排队订单, 9001 景区门票订单, 9101 团购券订单, 9201 在线预约订单, 9301 外卖订单, 140 住宿订单
	OrderType int64 `json:"order_type,omitempty"`
	// 订单状态。0 - 未支付, 1 - 已支付
	DateTime int64 `json:"date_time,omitempty"`
	// 订单的细节，不同的订单业务有不同的结构体，请具体询问业务方字段结构
	OrderDetail string                  `json:"order_detail,omitempty"`
	ExtShopInfo PoiOrderSyncExtShopInfo `json:"ext_shop_info,omitempty"`
	MiniApp     PoiOrderSyncMiniApp     `json:"mini_app,omitempty"`
}
