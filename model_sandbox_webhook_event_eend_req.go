/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// SandboxWebhookEventEendReq
type SandboxWebhookEventEendReq struct {
	// 需要mock的事件类型, 开放平台会通过webhook发送一条mock数据给你
	EventType string `json:"event_type,omitempty"`
}
