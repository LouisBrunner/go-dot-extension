// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type HTTPClientMethod int
const (
  HTTPClientMethodGet HTTPClientMethod = 0
  HTTPClientMethodHead HTTPClientMethod = 1
  HTTPClientMethodPost HTTPClientMethod = 2
  HTTPClientMethodPut HTTPClientMethod = 3
  HTTPClientMethodDelete HTTPClientMethod = 4
  HTTPClientMethodOptions HTTPClientMethod = 5
  HTTPClientMethodTrace HTTPClientMethod = 6
  HTTPClientMethodConnect HTTPClientMethod = 7
  HTTPClientMethodPatch HTTPClientMethod = 8
  HTTPClientMethodMax HTTPClientMethod = 9
)

type HTTPClientStatus int
const (
  HTTPClientStatusDisconnected HTTPClientStatus = 0
  HTTPClientStatusResolving HTTPClientStatus = 1
  HTTPClientStatusCantResolve HTTPClientStatus = 2
  HTTPClientStatusConnecting HTTPClientStatus = 3
  HTTPClientStatusCantConnect HTTPClientStatus = 4
  HTTPClientStatusConnected HTTPClientStatus = 5
  HTTPClientStatusRequesting HTTPClientStatus = 6
  HTTPClientStatusBody HTTPClientStatus = 7
  HTTPClientStatusConnectionError HTTPClientStatus = 8
  HTTPClientStatusTlsHandshakeError HTTPClientStatus = 9
)

type HTTPClientResponseCode int
const (
  HTTPClientResponseContinue HTTPClientResponseCode = 100
  HTTPClientResponseSwitchingProtocols HTTPClientResponseCode = 101
  HTTPClientResponseProcessing HTTPClientResponseCode = 102
  HTTPClientResponseOk HTTPClientResponseCode = 200
  HTTPClientResponseCreated HTTPClientResponseCode = 201
  HTTPClientResponseAccepted HTTPClientResponseCode = 202
  HTTPClientResponseNonAuthoritativeInformation HTTPClientResponseCode = 203
  HTTPClientResponseNoContent HTTPClientResponseCode = 204
  HTTPClientResponseResetContent HTTPClientResponseCode = 205
  HTTPClientResponsePartialContent HTTPClientResponseCode = 206
  HTTPClientResponseMultiStatus HTTPClientResponseCode = 207
  HTTPClientResponseAlreadyReported HTTPClientResponseCode = 208
  HTTPClientResponseImUsed HTTPClientResponseCode = 226
  HTTPClientResponseMultipleChoices HTTPClientResponseCode = 300
  HTTPClientResponseMovedPermanently HTTPClientResponseCode = 301
  HTTPClientResponseFound HTTPClientResponseCode = 302
  HTTPClientResponseSeeOther HTTPClientResponseCode = 303
  HTTPClientResponseNotModified HTTPClientResponseCode = 304
  HTTPClientResponseUseProxy HTTPClientResponseCode = 305
  HTTPClientResponseSwitchProxy HTTPClientResponseCode = 306
  HTTPClientResponseTemporaryRedirect HTTPClientResponseCode = 307
  HTTPClientResponsePermanentRedirect HTTPClientResponseCode = 308
  HTTPClientResponseBadRequest HTTPClientResponseCode = 400
  HTTPClientResponseUnauthorized HTTPClientResponseCode = 401
  HTTPClientResponsePaymentRequired HTTPClientResponseCode = 402
  HTTPClientResponseForbidden HTTPClientResponseCode = 403
  HTTPClientResponseNotFound HTTPClientResponseCode = 404
  HTTPClientResponseMethodNotAllowed HTTPClientResponseCode = 405
  HTTPClientResponseNotAcceptable HTTPClientResponseCode = 406
  HTTPClientResponseProxyAuthenticationRequired HTTPClientResponseCode = 407
  HTTPClientResponseRequestTimeout HTTPClientResponseCode = 408
  HTTPClientResponseConflict HTTPClientResponseCode = 409
  HTTPClientResponseGone HTTPClientResponseCode = 410
  HTTPClientResponseLengthRequired HTTPClientResponseCode = 411
  HTTPClientResponsePreconditionFailed HTTPClientResponseCode = 412
  HTTPClientResponseRequestEntityTooLarge HTTPClientResponseCode = 413
  HTTPClientResponseRequestUriTooLong HTTPClientResponseCode = 414
  HTTPClientResponseUnsupportedMediaType HTTPClientResponseCode = 415
  HTTPClientResponseRequestedRangeNotSatisfiable HTTPClientResponseCode = 416
  HTTPClientResponseExpectationFailed HTTPClientResponseCode = 417
  HTTPClientResponseImATeapot HTTPClientResponseCode = 418
  HTTPClientResponseMisdirectedRequest HTTPClientResponseCode = 421
  HTTPClientResponseUnprocessableEntity HTTPClientResponseCode = 422
  HTTPClientResponseLocked HTTPClientResponseCode = 423
  HTTPClientResponseFailedDependency HTTPClientResponseCode = 424
  HTTPClientResponseUpgradeRequired HTTPClientResponseCode = 426
  HTTPClientResponsePreconditionRequired HTTPClientResponseCode = 428
  HTTPClientResponseTooManyRequests HTTPClientResponseCode = 429
  HTTPClientResponseRequestHeaderFieldsTooLarge HTTPClientResponseCode = 431
  HTTPClientResponseUnavailableForLegalReasons HTTPClientResponseCode = 451
  HTTPClientResponseInternalServerError HTTPClientResponseCode = 500
  HTTPClientResponseNotImplemented HTTPClientResponseCode = 501
  HTTPClientResponseBadGateway HTTPClientResponseCode = 502
  HTTPClientResponseServiceUnavailable HTTPClientResponseCode = 503
  HTTPClientResponseGatewayTimeout HTTPClientResponseCode = 504
  HTTPClientResponseHttpVersionNotSupported HTTPClientResponseCode = 505
  HTTPClientResponseVariantAlsoNegotiates HTTPClientResponseCode = 506
  HTTPClientResponseInsufficientStorage HTTPClientResponseCode = 507
  HTTPClientResponseLoopDetected HTTPClientResponseCode = 508
  HTTPClientResponseNotExtended HTTPClientResponseCode = 510
  HTTPClientResponseNetworkAuthRequired HTTPClientResponseCode = 511
)
