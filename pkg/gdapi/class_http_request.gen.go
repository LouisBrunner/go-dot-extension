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
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3215244323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), custom_headers.AsCTypePtr(), gdc.ConstTypePtr(&method) , request_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&method)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPRequest) RequestRaw(url String, custom_headers PackedStringArray, method HTTPClientMethod, request_data_raw PackedByteArray, ) Error {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_raw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2714829993) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), custom_headers.AsCTypePtr(), gdc.ConstTypePtr(&method) , request_data_raw.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&method)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPRequest) CancelRequest()  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cancel_request")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) SetTlsOptions(client_options TLSOptions, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tls_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2210231844) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetHttpClientStatus() HTTPClientStatus {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_http_client_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1426656811) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret HTTPClientStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *HTTPRequest) SetUseThreads(enable bool, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_threads")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) IsUsingThreads() bool {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_threads")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetAcceptGzip(enable bool, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_accept_gzip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) IsAcceptingGzip() bool {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_accepting_gzip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetBodySizeLimit(bytes int64, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_body_size_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetBodySizeLimit() int64 {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_body_size_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetMaxRedirects(amount int64, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_redirects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetMaxRedirects() int64 {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_redirects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetDownloadFile(path String, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_download_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetDownloadFile() String {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_download_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *HTTPRequest) GetDownloadedBytes() int64 {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_downloaded_bytes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) GetBodySize() int64 {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_body_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetTimeout(timeout float64, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetTimeout() float64 {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetDownloadChunkSize(chunk_size int64, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_download_chunk_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&chunk_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) GetDownloadChunkSize() int64 {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_download_chunk_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *HTTPRequest) SetHttpProxy(host String, port int64, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_http_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *HTTPRequest) SetHttpsProxy(host String, port int64, )  {
  classNameV := StringNameFromStr("HTTPRequest")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_https_proxy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

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
