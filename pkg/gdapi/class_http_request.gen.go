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



// Enums

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

func (me *HTTPRequest) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HTTPRequest) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HTTPRequest) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *HTTPRequest) Request(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data String, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) RequestRaw(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data_raw PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) CancelRequest()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetTlsOptions(client_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetHttpClientStatus()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetUseThreads(enable bool, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) IsUsingThreads()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetAcceptGzip(enable bool, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) IsAcceptingGzip()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetBodySizeLimit(bytes int, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetBodySizeLimit()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetMaxRedirects(amount int, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetMaxRedirects()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetDownloadFile(path String, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetDownloadFile()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetDownloadedBytes()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetBodySize()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetTimeout(timeout float32, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetTimeout()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetDownloadChunkSize(chunk_size int, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) GetDownloadChunkSize()  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetHttpProxy(host String, port int, )  {
  panic("TODO: implement")
}

func  (me *HTTPRequest) SetHttpsProxy(host String, port int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
