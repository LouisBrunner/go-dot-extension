// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForHTTPClientList struct {
	fnConnectToHost                  gdc.MethodBindPtr
	fnSetConnection                  gdc.MethodBindPtr
	fnGetConnection                  gdc.MethodBindPtr
	fnRequestRaw                     gdc.MethodBindPtr
	fnRequest                        gdc.MethodBindPtr
	fnClose                          gdc.MethodBindPtr
	fnHasResponse                    gdc.MethodBindPtr
	fnIsResponseChunked              gdc.MethodBindPtr
	fnGetResponseCode                gdc.MethodBindPtr
	fnGetResponseHeaders             gdc.MethodBindPtr
	fnGetResponseHeadersAsDictionary gdc.MethodBindPtr
	fnGetResponseBodyLength          gdc.MethodBindPtr
	fnReadResponseBodyChunk          gdc.MethodBindPtr
	fnSetReadChunkSize               gdc.MethodBindPtr
	fnGetReadChunkSize               gdc.MethodBindPtr
	fnSetBlockingMode                gdc.MethodBindPtr
	fnIsBlockingModeEnabled          gdc.MethodBindPtr
	fnGetStatus                      gdc.MethodBindPtr
	fnPoll                           gdc.MethodBindPtr
	fnSetHttpProxy                   gdc.MethodBindPtr
	fnSetHttpsProxy                  gdc.MethodBindPtr
	fnQueryStringFromDict            gdc.MethodBindPtr
}

var ptrsForHTTPClient ptrsForHTTPClientList

func initHTTPClientPtrs(iface gdc.Interface) {

	className := StringNameFromStr("HTTPClient")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("connect_to_host")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnConnectToHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 504540374))
	}
	{
		methodName := StringNameFromStr("set_connection")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnSetConnection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3281897016))
	}
	{
		methodName := StringNameFromStr("get_connection")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetConnection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2741655269))
	}
	{
		methodName := StringNameFromStr("request_raw")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnRequestRaw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 540161961))
	}
	{
		methodName := StringNameFromStr("request")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnRequest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3778990155))
	}
	{
		methodName := StringNameFromStr("close")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has_response")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnHasResponse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_response_chunked")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnIsResponseChunked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_response_code")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetResponseCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_response_headers")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetResponseHeaders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("get_response_headers_as_dictionary")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetResponseHeadersAsDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
	}
	{
		methodName := StringNameFromStr("get_response_body_length")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetResponseBodyLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("read_response_body_chunk")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnReadResponseBodyChunk = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2115431945))
	}
	{
		methodName := StringNameFromStr("set_read_chunk_size")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnSetReadChunkSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_read_chunk_size")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetReadChunkSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_blocking_mode")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnSetBlockingMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_blocking_mode_enabled")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnIsBlockingModeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_status")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnGetStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1426656811))
	}
	{
		methodName := StringNameFromStr("poll")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("set_http_proxy")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnSetHttpProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
	}
	{
		methodName := StringNameFromStr("set_https_proxy")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnSetHttpsProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
	}
	{
		methodName := StringNameFromStr("query_string_from_dict")
		defer methodName.Destroy()
		ptrsForHTTPClient.fnQueryStringFromDict = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2538086567))
	}

}

type HTTPClient struct {
	RefCounted
}

func (me *HTTPClient) BaseClass() string {
	return "HTTPClient"
}

func NewHTTPClient() *HTTPClient {
	str := StringNameFromStr("HTTPClient") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &HTTPClient{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type HTTPClientMethod int

const (
	HTTPClientMethodMethodGet     HTTPClientMethod = 0
	HTTPClientMethodMethodHead    HTTPClientMethod = 1
	HTTPClientMethodMethodPost    HTTPClientMethod = 2
	HTTPClientMethodMethodPut     HTTPClientMethod = 3
	HTTPClientMethodMethodDelete  HTTPClientMethod = 4
	HTTPClientMethodMethodOptions HTTPClientMethod = 5
	HTTPClientMethodMethodTrace   HTTPClientMethod = 6
	HTTPClientMethodMethodConnect HTTPClientMethod = 7
	HTTPClientMethodMethodPatch   HTTPClientMethod = 8
	HTTPClientMethodMethodMax     HTTPClientMethod = 9
)

type HTTPClientStatus int

const (
	HTTPClientStatusStatusDisconnected      HTTPClientStatus = 0
	HTTPClientStatusStatusResolving         HTTPClientStatus = 1
	HTTPClientStatusStatusCantResolve       HTTPClientStatus = 2
	HTTPClientStatusStatusConnecting        HTTPClientStatus = 3
	HTTPClientStatusStatusCantConnect       HTTPClientStatus = 4
	HTTPClientStatusStatusConnected         HTTPClientStatus = 5
	HTTPClientStatusStatusRequesting        HTTPClientStatus = 6
	HTTPClientStatusStatusBody              HTTPClientStatus = 7
	HTTPClientStatusStatusConnectionError   HTTPClientStatus = 8
	HTTPClientStatusStatusTlsHandshakeError HTTPClientStatus = 9
)

type HTTPClientResponseCode int

const (
	HTTPClientResponseCodeResponseContinue                     HTTPClientResponseCode = 100
	HTTPClientResponseCodeResponseSwitchingProtocols           HTTPClientResponseCode = 101
	HTTPClientResponseCodeResponseProcessing                   HTTPClientResponseCode = 102
	HTTPClientResponseCodeResponseOk                           HTTPClientResponseCode = 200
	HTTPClientResponseCodeResponseCreated                      HTTPClientResponseCode = 201
	HTTPClientResponseCodeResponseAccepted                     HTTPClientResponseCode = 202
	HTTPClientResponseCodeResponseNonAuthoritativeInformation  HTTPClientResponseCode = 203
	HTTPClientResponseCodeResponseNoContent                    HTTPClientResponseCode = 204
	HTTPClientResponseCodeResponseResetContent                 HTTPClientResponseCode = 205
	HTTPClientResponseCodeResponsePartialContent               HTTPClientResponseCode = 206
	HTTPClientResponseCodeResponseMultiStatus                  HTTPClientResponseCode = 207
	HTTPClientResponseCodeResponseAlreadyReported              HTTPClientResponseCode = 208
	HTTPClientResponseCodeResponseImUsed                       HTTPClientResponseCode = 226
	HTTPClientResponseCodeResponseMultipleChoices              HTTPClientResponseCode = 300
	HTTPClientResponseCodeResponseMovedPermanently             HTTPClientResponseCode = 301
	HTTPClientResponseCodeResponseFound                        HTTPClientResponseCode = 302
	HTTPClientResponseCodeResponseSeeOther                     HTTPClientResponseCode = 303
	HTTPClientResponseCodeResponseNotModified                  HTTPClientResponseCode = 304
	HTTPClientResponseCodeResponseUseProxy                     HTTPClientResponseCode = 305
	HTTPClientResponseCodeResponseSwitchProxy                  HTTPClientResponseCode = 306
	HTTPClientResponseCodeResponseTemporaryRedirect            HTTPClientResponseCode = 307
	HTTPClientResponseCodeResponsePermanentRedirect            HTTPClientResponseCode = 308
	HTTPClientResponseCodeResponseBadRequest                   HTTPClientResponseCode = 400
	HTTPClientResponseCodeResponseUnauthorized                 HTTPClientResponseCode = 401
	HTTPClientResponseCodeResponsePaymentRequired              HTTPClientResponseCode = 402
	HTTPClientResponseCodeResponseForbidden                    HTTPClientResponseCode = 403
	HTTPClientResponseCodeResponseNotFound                     HTTPClientResponseCode = 404
	HTTPClientResponseCodeResponseMethodNotAllowed             HTTPClientResponseCode = 405
	HTTPClientResponseCodeResponseNotAcceptable                HTTPClientResponseCode = 406
	HTTPClientResponseCodeResponseProxyAuthenticationRequired  HTTPClientResponseCode = 407
	HTTPClientResponseCodeResponseRequestTimeout               HTTPClientResponseCode = 408
	HTTPClientResponseCodeResponseConflict                     HTTPClientResponseCode = 409
	HTTPClientResponseCodeResponseGone                         HTTPClientResponseCode = 410
	HTTPClientResponseCodeResponseLengthRequired               HTTPClientResponseCode = 411
	HTTPClientResponseCodeResponsePreconditionFailed           HTTPClientResponseCode = 412
	HTTPClientResponseCodeResponseRequestEntityTooLarge        HTTPClientResponseCode = 413
	HTTPClientResponseCodeResponseRequestUriTooLong            HTTPClientResponseCode = 414
	HTTPClientResponseCodeResponseUnsupportedMediaType         HTTPClientResponseCode = 415
	HTTPClientResponseCodeResponseRequestedRangeNotSatisfiable HTTPClientResponseCode = 416
	HTTPClientResponseCodeResponseExpectationFailed            HTTPClientResponseCode = 417
	HTTPClientResponseCodeResponseImATeapot                    HTTPClientResponseCode = 418
	HTTPClientResponseCodeResponseMisdirectedRequest           HTTPClientResponseCode = 421
	HTTPClientResponseCodeResponseUnprocessableEntity          HTTPClientResponseCode = 422
	HTTPClientResponseCodeResponseLocked                       HTTPClientResponseCode = 423
	HTTPClientResponseCodeResponseFailedDependency             HTTPClientResponseCode = 424
	HTTPClientResponseCodeResponseUpgradeRequired              HTTPClientResponseCode = 426
	HTTPClientResponseCodeResponsePreconditionRequired         HTTPClientResponseCode = 428
	HTTPClientResponseCodeResponseTooManyRequests              HTTPClientResponseCode = 429
	HTTPClientResponseCodeResponseRequestHeaderFieldsTooLarge  HTTPClientResponseCode = 431
	HTTPClientResponseCodeResponseUnavailableForLegalReasons   HTTPClientResponseCode = 451
	HTTPClientResponseCodeResponseInternalServerError          HTTPClientResponseCode = 500
	HTTPClientResponseCodeResponseNotImplemented               HTTPClientResponseCode = 501
	HTTPClientResponseCodeResponseBadGateway                   HTTPClientResponseCode = 502
	HTTPClientResponseCodeResponseServiceUnavailable           HTTPClientResponseCode = 503
	HTTPClientResponseCodeResponseGatewayTimeout               HTTPClientResponseCode = 504
	HTTPClientResponseCodeResponseHttpVersionNotSupported      HTTPClientResponseCode = 505
	HTTPClientResponseCodeResponseVariantAlsoNegotiates        HTTPClientResponseCode = 506
	HTTPClientResponseCodeResponseInsufficientStorage          HTTPClientResponseCode = 507
	HTTPClientResponseCodeResponseLoopDetected                 HTTPClientResponseCode = 508
	HTTPClientResponseCodeResponseNotExtended                  HTTPClientResponseCode = 510
	HTTPClientResponseCodeResponseNetworkAuthRequired          HTTPClientResponseCode = 511
)

func (me *HTTPClient) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *HTTPClient) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *HTTPClient) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *HTTPClient) ConnectToHost(host String, port int64, tls_options TLSOptions) Error {
	cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port), tls_options.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnConnectToHost), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HTTPClient) SetConnection(connection StreamPeer) {
	cargs := []gdc.ConstTypePtr{connection.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnSetConnection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HTTPClient) GetConnection() StreamPeer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStreamPeer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetConnection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *HTTPClient) RequestRaw(method HTTPClientMethod, url String, headers PackedStringArray, body PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method), url.AsCTypePtr(), headers.AsCTypePtr(), body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&method)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnRequestRaw), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HTTPClient) Request(method HTTPClientMethod, url String, headers PackedStringArray, body String) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method), url.AsCTypePtr(), headers.AsCTypePtr(), body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&method)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnRequest), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HTTPClient) Close() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HTTPClient) HasResponse() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnHasResponse), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HTTPClient) IsResponseChunked() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnIsResponseChunked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HTTPClient) GetResponseCode() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetResponseCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HTTPClient) GetResponseHeaders() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetResponseHeaders), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *HTTPClient) GetResponseHeadersAsDictionary() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetResponseHeadersAsDictionary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *HTTPClient) GetResponseBodyLength() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetResponseBodyLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HTTPClient) ReadResponseBodyChunk() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnReadResponseBodyChunk), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *HTTPClient) SetReadChunkSize(bytes int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnSetReadChunkSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HTTPClient) GetReadChunkSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetReadChunkSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HTTPClient) SetBlockingMode(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnSetBlockingMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HTTPClient) IsBlockingModeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnIsBlockingModeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *HTTPClient) GetStatus() HTTPClientStatus {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HTTPClientStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnGetStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HTTPClient) Poll() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnPoll), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *HTTPClient) SetHttpProxy(host String, port int64) {
	cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnSetHttpProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HTTPClient) SetHttpsProxy(host String, port int64) {
	cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnSetHttpsProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *HTTPClient) QueryStringFromDict(fields Dictionary) String {
	cargs := []gdc.ConstTypePtr{fields.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPClient.fnQueryStringFromDict), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
