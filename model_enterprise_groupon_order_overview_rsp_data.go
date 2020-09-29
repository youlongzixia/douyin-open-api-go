/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// EnterpriseGrouponOrderOverviewRspData struct for EnterpriseGrouponOrderOverviewRspData
type EnterpriseGrouponOrderOverviewRspData struct {
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// 错误码描述
	Description    string                                        `json:"description,omitempty"`
	RefundingCount int64                                         `json:"refunding_count,omitempty"`
	Refunded       EnterpriseGrouponOrderOverviewRspDataRefunded `json:"refunded,omitempty"`
	Sold           EnterpriseGrouponOrderOverviewRspDataRefunded `json:"sold,omitempty"`
}