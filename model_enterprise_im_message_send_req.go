/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// EnterpriseImMessageSendReq
type EnterpriseImMessageSendReq struct {
	// 客服id，传则走客服会话，否则为普通会话
	PersonaId string `json:"persona_id,omitempty"`
	// 消息的接收方openid
	ToUserId string `json:"to_user_id,omitempty"`
	// 内部使用
	ClientMsgId string `json:"client_msg_id,omitempty"`
	// 消息体见下方schema
	Content string `json:"content,omitempty"`
	// 消息内容格式 `text` - 文本消息 `image` - 图片消息 `video` - 视频消息 `card` - 卡片消息
	MessageType string `json:"message_type,omitempty"`
}
