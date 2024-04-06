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

type ptrsForJavaScriptBridgeList struct {
	fnEval           gdc.MethodBindPtr
	fnGetInterface   gdc.MethodBindPtr
	fnCreateCallback gdc.MethodBindPtr
	fnCreateObject   gdc.MethodBindPtr
	fnDownloadBuffer gdc.MethodBindPtr
	fnPwaNeedsUpdate gdc.MethodBindPtr
	fnPwaUpdate      gdc.MethodBindPtr
	fnForceFsSync    gdc.MethodBindPtr
}

var ptrsForJavaScriptBridge ptrsForJavaScriptBridgeList

func initJavaScriptBridgePtrs(iface gdc.Interface) {

	className := StringNameFromStr("JavaScriptBridge")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("eval")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnEval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 218087648))
	}
	{
		methodName := StringNameFromStr("get_interface")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnGetInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1355533281))
	}
	{
		methodName := StringNameFromStr("create_callback")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnCreateCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 422818440))
	}
	{
		methodName := StringNameFromStr("create_object")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnCreateObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093893586))
	}
	{
		methodName := StringNameFromStr("download_buffer")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnDownloadBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3352272093))
	}
	{
		methodName := StringNameFromStr("pwa_needs_update")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnPwaNeedsUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("pwa_update")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnPwaUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("force_fs_sync")
		defer methodName.Destroy()
		ptrsForJavaScriptBridge.fnForceFsSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type JavaScriptBridge struct {
	Object
}

func (me *JavaScriptBridge) BaseClass() string {
	return "JavaScriptBridge"
}

func NewJavaScriptBridge() *JavaScriptBridge {
	str := StringNameFromStr("JavaScriptBridge") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &JavaScriptBridge{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *JavaScriptBridge) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *JavaScriptBridge) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *JavaScriptBridge) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *JavaScriptBridge) Eval(code String, use_global_execution_context bool) Variant {
	cargs := []gdc.ConstTypePtr{code.AsCTypePtr(), gdc.ConstTypePtr(&use_global_execution_context)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&use_global_execution_context)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnEval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JavaScriptBridge) GetInterface(interface_ String) JavaScriptObject {
	cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewJavaScriptObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnGetInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JavaScriptBridge) CreateCallback(callable Callable) JavaScriptObject {
	cargs := []gdc.ConstTypePtr{callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewJavaScriptObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnCreateCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *JavaScriptBridge) CreateObject(object String, varargs ...Variant) Variant {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := object.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForJavaScriptBridge.fnCreateObject), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return *ret
	}
	return *ret
}

func (me *JavaScriptBridge) DownloadBuffer(buffer PackedByteArray, name String, mime String) {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), name.AsCTypePtr(), mime.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnDownloadBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *JavaScriptBridge) PwaNeedsUpdate() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnPwaNeedsUpdate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *JavaScriptBridge) PwaUpdate() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnPwaUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *JavaScriptBridge) ForceFsSync() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForJavaScriptBridge.fnForceFsSync), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type JavaScriptBridgePwaUpdateAvailableSignalFn func()

func (me *JavaScriptBridge) ConnectPwaUpdateAvailable(subs SignalSubscribers, fn JavaScriptBridgePwaUpdateAvailableSignalFn) {
	sig := StringNameFromStr("pwa_update_available")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *JavaScriptBridge) DisconnectPwaUpdateAvailable(subs SignalSubscribers, fn JavaScriptBridgePwaUpdateAvailableSignalFn) {
	sig := StringNameFromStr("pwa_update_available")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
