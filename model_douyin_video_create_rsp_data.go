/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DouyinVideoCreateRspData struct for DouyinVideoCreateRspData
type DouyinVideoCreateRspData struct {
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	ItemId      string `json:"item_id,omitempty"`
}
