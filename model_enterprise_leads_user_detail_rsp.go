/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// EnterpriseLeadsUserDetailRsp
type EnterpriseLeadsUserDetailRsp struct {
	Data    EnterpriseLeadsUserDetailRspData `json:"data,omitempty"`
	Message string                           `json:"message,omitempty"`
	Extra   FollowingListRspDataExtra        `json:"extra,omitempty"`
}