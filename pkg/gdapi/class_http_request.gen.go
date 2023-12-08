// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HTTPRequest struct {
  obj gdc.ObjectPtr
}

func (me *HTTPRequest) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HTTPRequest) BaseClass() string {
  return "HTTPRequest"
}

type HTTPRequestResult int
const (
  HTTPRequestResultSuccess HTTPRequestResult = 0
  HTTPRequestResultChunkedBodySizeMismatch HTTPRequestResult = 1
  HTTPRequestResultCantConnect HTTPRequestResult = 2
  HTTPRequestResultCantResolve HTTPRequestResult = 3
  HTTPRequestResultConnectionError HTTPRequestResult = 4
  HTTPRequestResultTlsHandshakeError HTTPRequestResult = 5
  HTTPRequestResultNoResponse HTTPRequestResult = 6
  HTTPRequestResultBodySizeLimitExceeded HTTPRequestResult = 7
  HTTPRequestResultBodyDecompressFailed HTTPRequestResult = 8
  HTTPRequestResultRequestFailed HTTPRequestResult = 9
  HTTPRequestResultDownloadFileCantOpen HTTPRequestResult = 10
  HTTPRequestResultDownloadFileWriteError HTTPRequestResult = 11
  HTTPRequestResultRedirectLimitReached HTTPRequestResult = 12
  HTTPRequestResultTimeout HTTPRequestResult = 13
)
