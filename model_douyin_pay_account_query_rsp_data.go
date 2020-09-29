/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// DouyinPayAccountQueryRspData struct for DouyinPayAccountQueryRspData
type DouyinPayAccountQueryRspData struct {
	// 金额
	Balance int64 `json:"balance,omitempty"`
	// 附加信息
	Ext string `json:"ext,omitempty"`
	// 返回码
	RetCode string `json:"ret_code,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty"`
	// 返回码
	SubCode string `json:"sub_code,omitempty"`
	// 返回信息
	SubMsg string `json:"sub_msg,omitempty"`
	// 错误描述
	Description string `json:"description,omitempty"`
	// 返回错误码
	ErrorCode int64 `json:"error_code,omitempty"`
}