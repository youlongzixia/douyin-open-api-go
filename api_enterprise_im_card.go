/*
 * 抖音开放API
 *
 * douyin open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package douyin

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// EnterpriseImCardApiService EnterpriseImCardApi service
type EnterpriseImCardApiService service

// EnterpriseImCardDeleteOpts Optional parameters for the method 'EnterpriseImCardDelete'
type EnterpriseImCardDeleteOpts struct {
	Body optional.Interface
}

/*
EnterpriseImCardDelete 删除消息卡片
删除消息卡片
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accessToken 调用/oauth/access_token/生成的token，此token需要用户授权。
 * @param openId 通过/oauth/access_token/获取，用户唯一标志
 * @param optional nil or *EnterpriseImCardDeleteOpts - Optional Parameters:
 * @param "Body" (optional.Interface of EnterpriseImCardDeleteReq) -
@return EnterpriseImCardDeleteRsp
*/
func (a *EnterpriseImCardApiService) EnterpriseImCardDelete(ctx _context.Context, accessToken string, openId string, localVarOptionals *EnterpriseImCardDeleteOpts) (EnterpriseImCardDeleteRsp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnterpriseImCardDeleteRsp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/enterprise/im/card/delete"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("access_token", parameterToString(accessToken, ""))
	localVarQueryParams.Add("open_id", parameterToString(openId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarOptionalBody, localVarOptionalBodyok := localVarOptionals.Body.Value().(EnterpriseImCardDeleteReq)
		if !localVarOptionalBodyok {
			return localVarReturnValue, nil, reportError("body should be EnterpriseImCardDeleteReq")
		}
		localVarPostBody = &localVarOptionalBody
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
EnterpriseImCardList 获取消息卡片列表
获取消息卡片列表
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accessToken 调用/oauth/access_token/生成的token，此token需要用户授权。
 * @param openId 通过/oauth/access_token/获取，用户唯一标志
 * @param cursor 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
 * @param count 每页数量
@return EnterpriseImCardListRsp
*/
func (a *EnterpriseImCardApiService) EnterpriseImCardList(ctx _context.Context, accessToken string, openId string, cursor int64, count int64) (EnterpriseImCardListRsp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnterpriseImCardListRsp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/enterprise/im/card/list"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("access_token", parameterToString(accessToken, ""))
	localVarQueryParams.Add("open_id", parameterToString(openId, ""))
	localVarQueryParams.Add("cursor", parameterToString(cursor, ""))
	localVarQueryParams.Add("count", parameterToString(count, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// EnterpriseImCardSaveOpts Optional parameters for the method 'EnterpriseImCardSave'
type EnterpriseImCardSaveOpts struct {
	Body optional.Interface
}

/*
EnterpriseImCardSave 创建/更新消息卡片
创建/更新消息卡片
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accessToken 调用/oauth/access_token/生成的token，此token需要用户授权。
 * @param openId 通过/oauth/access_token/获取，用户唯一标志
 * @param optional nil or *EnterpriseImCardSaveOpts - Optional Parameters:
 * @param "Body" (optional.Interface of EnterpriseImCardSaveReq) -
@return EnterpriseImCardSaveRsp
*/
func (a *EnterpriseImCardApiService) EnterpriseImCardSave(ctx _context.Context, accessToken string, openId string, localVarOptionals *EnterpriseImCardSaveOpts) (EnterpriseImCardSaveRsp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnterpriseImCardSaveRsp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/enterprise/im/card/save"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("access_token", parameterToString(accessToken, ""))
	localVarQueryParams.Add("open_id", parameterToString(openId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarOptionalBody, localVarOptionalBodyok := localVarOptionals.Body.Value().(EnterpriseImCardSaveReq)
		if !localVarOptionalBodyok {
			return localVarReturnValue, nil, reportError("body should be EnterpriseImCardSaveReq")
		}
		localVarPostBody = &localVarOptionalBody
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
