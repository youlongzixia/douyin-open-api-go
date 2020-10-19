/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

// InlineObject23 struct for InlineObject23
type InlineObject23 struct {
	// 卡片id，创建时不传；更新时必传。同一个用户的卡片id不可重复
	CardId string `json:"card_id,omitempty"`
	// 卡片类型 `question_list` - 问题列表
	CardType string `json:"card_type,omitempty"`
	// 卡片内容字段 json序列化成string， { \"question_list\" { \"text\" \"有什么疑问呢\", \"questions\" [{ \"name\" \"问题1\", \"text\" \"关键词1\" }, { \"name\" \"问题2\", \"text\" \"关键词2\" } ] } }
	Content string `json:"content,omitempty"`
}
