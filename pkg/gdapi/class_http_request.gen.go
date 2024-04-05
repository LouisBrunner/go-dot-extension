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

type ptrsForHTTPRequestList struct {
  fnRequest gdc.MethodBindPtr
  fnRequestRaw gdc.MethodBindPtr
  fnCancelRequest gdc.MethodBindPtr
  fnSetTlsOptions gdc.MethodBindPtr
  fnGetHttpClientStatus gdc.MethodBindPtr
  fnSetUseThreads gdc.MethodBindPtr
  fnIsUsingThreads gdc.MethodBindPtr
  fnSetAcceptGzip gdc.MethodBindPtr
  fnIsAcceptingGzip gdc.MethodBindPtr
  fnSetBodySizeLimit gdc.MethodBindPtr
  fnGetBodySizeLimit gdc.MethodBindPtr
  fnSetMaxRedirects gdc.MethodBindPtr
  fnGetMaxRedirects gdc.MethodBindPtr
  fnSetDownloadFile gdc.MethodBindPtr
  fnGetDownloadFile gdc.MethodBindPtr
  fnGetDownloadedBytes gdc.MethodBindPtr
  fnGetBodySize gdc.MethodBindPtr
  fnSetTimeout gdc.MethodBindPtr
  fnGetTimeout gdc.MethodBindPtr
  fnSetDownloadChunkSize gdc.MethodBindPtr
  fnGetDownloadChunkSize gdc.MethodBindPtr
  fnSetHttpProxy gdc.MethodBindPtr
  fnSetHttpsProxy gdc.MethodBindPtr
}

var ptrsForHTTPRequest ptrsForHTTPRequestList

func initHTTPRequestPtrs(iface gdc.Interface) {

  className := StringNameFromStr("HTTPRequest")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("request")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnRequest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3215244323))
  }
  {
    methodName := StringNameFromStr("request_raw")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnRequestRaw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2714829993))
  }
  {
    methodName := StringNameFromStr("cancel_request")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnCancelRequest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_tls_options")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetTlsOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2210231844))
  }
  {
    methodName := StringNameFromStr("get_http_client_status")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetHttpClientStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1426656811))
  }
  {
    methodName := StringNameFromStr("set_use_threads")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetUseThreads = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_threads")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnIsUsingThreads = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_accept_gzip")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetAcceptGzip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_accepting_gzip")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnIsAcceptingGzip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_body_size_limit")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetBodySizeLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_body_size_limit")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetBodySizeLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_max_redirects")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetMaxRedirects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_redirects")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetMaxRedirects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_download_file")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetDownloadFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_download_file")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetDownloadFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_downloaded_bytes")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetDownloadedBytes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_body_size")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetBodySize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_timeout")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_timeout")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_download_chunk_size")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetDownloadChunkSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_download_chunk_size")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnGetDownloadChunkSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_http_proxy")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetHttpProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
  }
  {
    methodName := StringNameFromStr("set_https_proxy")
    defer methodName.Destroy()
    ptrsForHTTPRequest.fnSetHttpsProxy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
  }
}

type HTTPRequest struct {
  Node
}

func (me *HTTPRequest) BaseClass() string {
  return "HTTPRequest"
}

func NewHTTPRequest() *HTTPRequest {
  str := StringNameFromStr("HTTPRequest") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &HTTPRequest{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *HTTPRequest) Request(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data String, ) Error {
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), custom_headers.AsCTypePtr(), gdc.ConstTypePtr(&method) , request_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&method)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnRequest), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPRequest) RequestRaw(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data_raw PackedByteArray, ) Error {
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), custom_headers.AsCTypePtr(), gdc.ConstTypePtr(&method) , request_data_raw.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&method)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnRequestRaw), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPRequest) CancelRequest()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnCancelRequest), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) SetTlsOptions(client_options TLSOptions, )  {
  cargs := []gdc.ConstTypePtr{client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetTlsOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetHttpClientStatus() HTTPClientStatus {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret HTTPClientStatus

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetHttpClientStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPRequest) SetUseThreads(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetUseThreads), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) IsUsingThreads() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnIsUsingThreads), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetAcceptGzip(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetAcceptGzip), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) IsAcceptingGzip() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnIsAcceptingGzip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetBodySizeLimit(bytes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetBodySizeLimit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetBodySizeLimit() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetBodySizeLimit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetMaxRedirects(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetMaxRedirects), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetMaxRedirects() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetMaxRedirects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetDownloadFile(path String, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetDownloadFile), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetDownloadFile() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetDownloadFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *HTTPRequest) GetDownloadedBytes() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetDownloadedBytes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) GetBodySize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetBodySize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetTimeout(timeout float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetTimeout), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetTimeout() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetTimeout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetDownloadChunkSize(chunk_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&chunk_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetDownloadChunkSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetDownloadChunkSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnGetDownloadChunkSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetHttpProxy(host String, port int64, )  {
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetHttpProxy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) SetHttpsProxy(host String, port int64, )  {
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForHTTPRequest.fnSetHttpsProxy), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type HTTPRequestRequestCompletedSignalFn func(result int, response_code int, headers PackedStringArray, body PackedByteArray, )

func (me *HTTPRequest) ConnectRequestCompleted(subs SignalSubscribers, fn HTTPRequestRequestCompletedSignalFn) {
  sig := StringNameFromStr("request_completed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *HTTPRequest) DisconnectRequestCompleted(subs SignalSubscribers, fn HTTPRequestRequestCompletedSignalFn) {
  sig := StringNameFromStr("request_completed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
