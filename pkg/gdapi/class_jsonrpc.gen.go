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

type ptrsForJSONRPCList struct {
	fnSetScope          gdc.MethodBindPtr
	fnProcessAction     gdc.MethodBindPtr
	fnProcessString     gdc.MethodBindPtr
	fnMakeRequest       gdc.MethodBindPtr
	fnMakeResponse      gdc.MethodBindPtr
	fnMakeNotification  gdc.MethodBindPtr
	fnMakeResponseError gdc.MethodBindPtr
}

var ptrsForJSONRPC ptrsForJSONRPCList

func initJSONRPCPtrs(iface gdc.Interface) {

	className := StringNameFromStr("JSONRPC")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_scope")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnSetScope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2572618360))
	}
	{
		methodName := StringNameFromStr("process_action")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnProcessAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2963479484))
	}
	{
		methodName := StringNameFromStr("process_string")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnProcessString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1703090593))
	}
	{
		methodName := StringNameFromStr("make_request")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnMakeRequest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3423508980))
	}
	{
		methodName := StringNameFromStr("make_response")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnMakeResponse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5053918))
	}
	{
		methodName := StringNameFromStr("make_notification")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnMakeNotification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2949127017))
	}
	{
		methodName := StringNameFromStr("make_response_error")
		defer methodName.Destroy()
		ptrsForJSONRPC.fnMakeResponseError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 928596297))
	}
}

type JSONRPC struct {
	Object
}

func (me *JSONRPC) BaseClass() string {
	return "JSONRPC"
}

func NewJSONRPC() *JSONRPC {
	str := StringNameFromStr("JSONRPC") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &JSONRPC{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type JSONRPCErrorCode int

const (
	JSONRPCErrorCodeParseError     JSONRPCErrorCode = -32700
	JSONRPCErrorCodeInvalidRequest JSONRPCErrorCode = -32600
	JSONRPCErrorCodeMethodNotFound JSONRPCErrorCode = -32601
	JSONRPCErrorCodeInvalidParams  JSONRPCErrorCode = -32602
	JSONRPCErrorCodeInternalError  JSONRPCErrorCode = -32603
)

func (me *JSONRPC) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *JSONRPC) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *JSONRPC) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *JSONRPC) SetScope(scope String, target Object) {
	cargs := []gdc.ConstTypePtr{scope.AsCTypePtr(), target.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnSetScope), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *JSONRPC) ProcessAction(action Variant, recurse bool) Variant {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&recurse)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&recurse)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnProcessAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JSONRPC) ProcessString(action String) String {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnProcessString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JSONRPC) MakeRequest(method String, params Variant, id Variant) Dictionary {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), params.AsCTypePtr(), id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnMakeRequest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JSONRPC) MakeResponse(result Variant, id Variant) Dictionary {
	cargs := []gdc.ConstTypePtr{result.AsCTypePtr(), id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnMakeResponse), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JSONRPC) MakeNotification(method String, params Variant) Dictionary {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), params.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnMakeNotification), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JSONRPC) MakeResponseError(code int64, message String, id Variant) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code), message.AsCTypePtr(), id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&code)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJSONRPC.fnMakeResponseError), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
