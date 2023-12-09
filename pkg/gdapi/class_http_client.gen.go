// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HTTPClient struct {
  obj gdc.ObjectPtr
}

func (me *HTTPClient) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HTTPClient) BaseClass() string {
  return "HTTPClient"
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

func (me *HTTPClient) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HTTPClient) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *HTTPClient) ConnectToHost(host String, port int, tls_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) SetConnection(connection StreamPeer, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetConnection()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) RequestRaw(method HTTPClientMethod, url String, headers PackedStringArray, body PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) Request(method HTTPClientMethod, url String, headers PackedStringArray, body String, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) Close()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) HasResponse()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) IsResponseChunked()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetResponseCode()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetResponseHeaders()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetResponseHeadersAsDictionary()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetResponseBodyLength()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) ReadResponseBodyChunk()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) SetReadChunkSize(bytes int, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetReadChunkSize()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) SetBlockingMode(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) IsBlockingModeEnabled()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) GetStatus()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) Poll()  {
  panic("TODO: implement")
}

func  (me *HTTPClient) SetHttpProxy(host String, port int, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) SetHttpsProxy(host String, port int, )  {
  panic("TODO: implement")
}

func  (me *HTTPClient) QueryStringFromDict(fields Dictionary, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
