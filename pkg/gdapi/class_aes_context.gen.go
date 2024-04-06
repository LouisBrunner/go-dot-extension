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

type ptrsForAESContextList struct {
	fnStart      gdc.MethodBindPtr
	fnUpdate     gdc.MethodBindPtr
	fnGetIvState gdc.MethodBindPtr
	fnFinish     gdc.MethodBindPtr
}

var ptrsForAESContext ptrsForAESContextList

func initAESContextPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AESContext")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("start")
		defer methodName.Destroy()
		ptrsForAESContext.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3122411423))
	}
	{
		methodName := StringNameFromStr("update")
		defer methodName.Destroy()
		ptrsForAESContext.fnUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 527836100))
	}
	{
		methodName := StringNameFromStr("get_iv_state")
		defer methodName.Destroy()
		ptrsForAESContext.fnGetIvState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2115431945))
	}
	{
		methodName := StringNameFromStr("finish")
		defer methodName.Destroy()
		ptrsForAESContext.fnFinish = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type AESContext struct {
	RefCounted
}

func (me *AESContext) BaseClass() string {
	return "AESContext"
}

func NewAESContext() *AESContext {
	str := StringNameFromStr("AESContext") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AESContext{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AESContextMode int

const (
	AESContextModeModeEcbEncrypt AESContextMode = 0
	AESContextModeModeEcbDecrypt AESContextMode = 1
	AESContextModeModeCbcEncrypt AESContextMode = 2
	AESContextModeModeCbcDecrypt AESContextMode = 3
	AESContextModeModeMax        AESContextMode = 4
)

func (me *AESContext) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AESContext) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AESContext) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AESContext) Start(mode AESContextMode, key PackedByteArray, iv PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), key.AsCTypePtr(), iv.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAESContext.fnStart), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AESContext) Update(src PackedByteArray) PackedByteArray {
	cargs := []gdc.ConstTypePtr{src.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAESContext.fnUpdate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AESContext) GetIvState() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAESContext.fnGetIvState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AESContext) Finish() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAESContext.fnFinish), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
