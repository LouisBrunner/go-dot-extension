// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  HTTPRequestResultResultSuccess HTTPRequestResult = 0
  HTTPRequestResultResultChunkedBodySizeMismatch HTTPRequestResult = 1
  HTTPRequestResultResultCantConnect HTTPRequestResult = 2
  HTTPRequestResultResultCantResolve HTTPRequestResult = 3
  HTTPRequestResultResultConnectionError HTTPRequestResult = 4
  HTTPRequestResultResultTlsHandshakeError HTTPRequestResult = 5
  HTTPRequestResultResultNoResponse HTTPRequestResult = 6
  HTTPRequestResultResultBodySizeLimitExceeded HTTPRequestResult = 7
  HTTPRequestResultResultBodyDecompressFailed HTTPRequestResult = 8
  HTTPRequestResultResultRequestFailed HTTPRequestResult = 9
  HTTPRequestResultResultDownloadFileCantOpen HTTPRequestResult = 10
  HTTPRequestResultResultDownloadFileWriteError HTTPRequestResult = 11
  HTTPRequestResultResultRedirectLimitReached HTTPRequestResult = 12
  HTTPRequestResultResultTimeout HTTPRequestResult = 13
)

func  (me *HTTPRequest) Request(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data String, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) RequestRaw(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data_raw PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) CancelRequest() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetTlsOptions(client_options TLSOptions, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetHttpClientStatus() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetUseThreads(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) IsUsingThreads() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetAcceptGzip(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) IsAcceptingGzip() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetBodySizeLimit(bytes int, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetBodySizeLimit() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetMaxRedirects(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetMaxRedirects() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetDownloadFile(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetDownloadFile() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetDownloadedBytes() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetBodySize() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetTimeout(timeout float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetTimeout() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetDownloadChunkSize(chunk_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) GetDownloadChunkSize() { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetHttpProxy(host String, port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *HTTPRequest) SetHttpsProxy(host String, port int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
