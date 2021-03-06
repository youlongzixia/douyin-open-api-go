/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// ClientTokenRspData struct for ClientTokenRspData
type ClientTokenRspData struct {
	// 接口调用凭证
	AccessToken string `json:"access_token,omitempty"`
	// 错误码描述
	Description string `json:"description,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty"`
	// access_token接口调用凭证超时时间，单位（秒
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
