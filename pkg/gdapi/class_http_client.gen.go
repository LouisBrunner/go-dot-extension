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
  HTTPClientMethodMethodGet HTTPClientMethod = 0
  HTTPClientMethodMethodHead HTTPClientMethod = 1
  HTTPClientMethodMethodPost HTTPClientMethod = 2
  HTTPClientMethodMethodPut HTTPClientMethod = 3
  HTTPClientMethodMethodDelete HTTPClientMethod = 4
  HTTPClientMethodMethodOptions HTTPClientMethod = 5
  HTTPClientMethodMethodTrace HTTPClientMethod = 6
  HTTPClientMethodMethodConnect HTTPClientMethod = 7
  HTTPClientMethodMethodPatch HTTPClientMethod = 8
  HTTPClientMethodMethodMax HTTPClientMethod = 9
)

type HTTPClientStatus int
const (
  HTTPClientStatusStatusDisconnected HTTPClientStatus = 0
  HTTPClientStatusStatusResolving HTTPClientStatus = 1
  HTTPClientStatusStatusCantResolve HTTPClientStatus = 2
  HTTPClientStatusStatusConnecting HTTPClientStatus = 3
  HTTPClientStatusStatusCantConnect HTTPClientStatus = 4
  HTTPClientStatusStatusConnected HTTPClientStatus = 5
  HTTPClientStatusStatusRequesting HTTPClientStatus = 6
  HTTPClientStatusStatusBody HTTPClientStatus = 7
  HTTPClientStatusStatusConnectionError HTTPClientStatus = 8
  HTTPClientStatusStatusTlsHandshakeError HTTPClientStatus = 9
)

type HTTPClientResponseCode int
const (
  HTTPClientResponseCodeResponseContinue HTTPClientResponseCode = 100
  HTTPClientResponseCodeResponseSwitchingProtocols HTTPClientResponseCode = 101
  HTTPClientResponseCodeResponseProcessing HTTPClientResponseCode = 102
  HTTPClientResponseCodeResponseOk HTTPClientResponseCode = 200
  HTTPClientResponseCodeResponseCreated HTTPClientResponseCode = 201
  HTTPClientResponseCodeResponseAccepted HTTPClientResponseCode = 202
  HTTPClientResponseCodeResponseNonAuthoritativeInformation HTTPClientResponseCode = 203
  HTTPClientResponseCodeResponseNoContent HTTPClientResponseCode = 204
  HTTPClientResponseCodeResponseResetContent HTTPClientResponseCode = 205
  HTTPClientResponseCodeResponsePartialContent HTTPClientResponseCode = 206
  HTTPClientResponseCodeResponseMultiStatus HTTPClientResponseCode = 207
  HTTPClientResponseCodeResponseAlreadyReported HTTPClientResponseCode = 208
  HTTPClientResponseCodeResponseImUsed HTTPClientResponseCode = 226
  HTTPClientResponseCodeResponseMultipleChoices HTTPClientResponseCode = 300
  HTTPClientResponseCodeResponseMovedPermanently HTTPClientResponseCode = 301
  HTTPClientResponseCodeResponseFound HTTPClientResponseCode = 302
  HTTPClientResponseCodeResponseSeeOther HTTPClientResponseCode = 303
  HTTPClientResponseCodeResponseNotModified HTTPClientResponseCode = 304
  HTTPClientResponseCodeResponseUseProxy HTTPClientResponseCode = 305
  HTTPClientResponseCodeResponseSwitchProxy HTTPClientResponseCode = 306
  HTTPClientResponseCodeResponseTemporaryRedirect HTTPClientResponseCode = 307
  HTTPClientResponseCodeResponsePermanentRedirect HTTPClientResponseCode = 308
  HTTPClientResponseCodeResponseBadRequest HTTPClientResponseCode = 400
  HTTPClientResponseCodeResponseUnauthorized HTTPClientResponseCode = 401
  HTTPClientResponseCodeResponsePaymentRequired HTTPClientResponseCode = 402
  HTTPClientResponseCodeResponseForbidden HTTPClientResponseCode = 403
  HTTPClientResponseCodeResponseNotFound HTTPClientResponseCode = 404
  HTTPClientResponseCodeResponseMethodNotAllowed HTTPClientResponseCode = 405
  HTTPClientResponseCodeResponseNotAcceptable HTTPClientResponseCode = 406
  HTTPClientResponseCodeResponseProxyAuthenticationRequired HTTPClientResponseCode = 407
  HTTPClientResponseCodeResponseRequestTimeout HTTPClientResponseCode = 408
  HTTPClientResponseCodeResponseConflict HTTPClientResponseCode = 409
  HTTPClientResponseCodeResponseGone HTTPClientResponseCode = 410
  HTTPClientResponseCodeResponseLengthRequired HTTPClientResponseCode = 411
  HTTPClientResponseCodeResponsePreconditionFailed HTTPClientResponseCode = 412
  HTTPClientResponseCodeResponseRequestEntityTooLarge HTTPClientResponseCode = 413
  HTTPClientResponseCodeResponseRequestUriTooLong HTTPClientResponseCode = 414
  HTTPClientResponseCodeResponseUnsupportedMediaType HTTPClientResponseCode = 415
  HTTPClientResponseCodeResponseRequestedRangeNotSatisfiable HTTPClientResponseCode = 416
  HTTPClientResponseCodeResponseExpectationFailed HTTPClientResponseCode = 417
  HTTPClientResponseCodeResponseImATeapot HTTPClientResponseCode = 418
  HTTPClientResponseCodeResponseMisdirectedRequest HTTPClientResponseCode = 421
  HTTPClientResponseCodeResponseUnprocessableEntity HTTPClientResponseCode = 422
  HTTPClientResponseCodeResponseLocked HTTPClientResponseCode = 423
  HTTPClientResponseCodeResponseFailedDependency HTTPClientResponseCode = 424
  HTTPClientResponseCodeResponseUpgradeRequired HTTPClientResponseCode = 426
  HTTPClientResponseCodeResponsePreconditionRequired HTTPClientResponseCode = 428
  HTTPClientResponseCodeResponseTooManyRequests HTTPClientResponseCode = 429
  HTTPClientResponseCodeResponseRequestHeaderFieldsTooLarge HTTPClientResponseCode = 431
  HTTPClientResponseCodeResponseUnavailableForLegalReasons HTTPClientResponseCode = 451
  HTTPClientResponseCodeResponseInternalServerError HTTPClientResponseCode = 500
  HTTPClientResponseCodeResponseNotImplemented HTTPClientResponseCode = 501
  HTTPClientResponseCodeResponseBadGateway HTTPClientResponseCode = 502
  HTTPClientResponseCodeResponseServiceUnavailable HTTPClientResponseCode = 503
  HTTPClientResponseCodeResponseGatewayTimeout HTTPClientResponseCode = 504
  HTTPClientResponseCodeResponseHttpVersionNotSupported HTTPClientResponseCode = 505
  HTTPClientResponseCodeResponseVariantAlsoNegotiates HTTPClientResponseCode = 506
  HTTPClientResponseCodeResponseInsufficientStorage HTTPClientResponseCode = 507
  HTTPClientResponseCodeResponseLoopDetected HTTPClientResponseCode = 508
  HTTPClientResponseCodeResponseNotExtended HTTPClientResponseCode = 510
  HTTPClientResponseCodeResponseNetworkAuthRequired HTTPClientResponseCode = 511
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

func  (me *HTTPClient) ConnectToHost(host String, port int64, tls_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 504540374) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , tls_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPClient) SetConnection(connection StreamPeer, )  {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_connection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3281897016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{connection.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPClient) GetConnection() StreamPeer {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2741655269) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStreamPeer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *HTTPClient) RequestRaw(method HTTPClientMethod, url String, headers PackedStringArray, body PackedByteArray, ) Error {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_raw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 540161961) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method) , url.AsCTypePtr(), headers.AsCTypePtr(), body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&method)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPClient) Request(method HTTPClientMethod, url String, headers PackedStringArray, body String, ) Error {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3778990155) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method) , url.AsCTypePtr(), headers.AsCTypePtr(), body.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&method)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPClient) Close()  {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPClient) HasResponse() bool {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_response")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPClient) IsResponseChunked() bool {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_response_chunked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPClient) GetResponseCode() int64 {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_response_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPClient) GetResponseHeaders() PackedStringArray {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_response_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *HTTPClient) GetResponseHeadersAsDictionary() Dictionary {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_response_headers_as_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *HTTPClient) GetResponseBodyLength() int64 {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_response_body_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPClient) ReadResponseBodyChunk() PackedByteArray {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("read_response_body_chunk")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2115431945) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *HTTPClient) SetReadChunkSize(bytes int64, )  {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_read_chunk_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPClient) GetReadChunkSize() int64 {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_read_chunk_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPClient) SetBlockingMode(enabled bool, )  {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blocking_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPClient) IsBlockingModeEnabled() bool {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_blocking_mode_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPClient) GetStatus() HTTPClientStatus {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1426656811) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret HTTPClientStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPClient) Poll() Error {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPClient) SetHttpProxy(host String, port int64, )  {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_http_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPClient) SetHttpsProxy(host String, port int64, )  {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_https_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPClient) QueryStringFromDict(fields Dictionary, ) String {
  classNameV := StringNameFromStr("HTTPClient")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("query_string_from_dict")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2538086567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{fields.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
