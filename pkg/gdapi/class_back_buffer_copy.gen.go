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

type ptrsForBackBufferCopyList struct {
	fnSetRect     gdc.MethodBindPtr
	fnGetRect     gdc.MethodBindPtr
	fnSetCopyMode gdc.MethodBindPtr
	fnGetCopyMode gdc.MethodBindPtr
}

var ptrsForBackBufferCopy ptrsForBackBufferCopyList

func initBackBufferCopyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BackBufferCopy")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_rect")
		defer methodName.Destroy()
		ptrsForBackBufferCopy.fnSetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("get_rect")
		defer methodName.Destroy()
		ptrsForBackBufferCopy.fnGetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_copy_mode")
		defer methodName.Destroy()
		ptrsForBackBufferCopy.fnSetCopyMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1713538590))
	}
	{
		methodName := StringNameFromStr("get_copy_mode")
		defer methodName.Destroy()
		ptrsForBackBufferCopy.fnGetCopyMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3271169440))
	}

}

type BackBufferCopy struct {
	Node2D
}

func (me *BackBufferCopy) BaseClass() string {
	return "BackBufferCopy"
}

func NewBackBufferCopy() *BackBufferCopy {
	str := StringNameFromStr("BackBufferCopy") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &BackBufferCopy{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type BackBufferCopyCopyMode int

const (
	BackBufferCopyCopyModeCopyModeDisabled BackBufferCopyCopyMode = 0
	BackBufferCopyCopyModeCopyModeRect     BackBufferCopyCopyMode = 1
	BackBufferCopyCopyModeCopyModeViewport BackBufferCopyCopyMode = 2
)

func (me *BackBufferCopy) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *BackBufferCopy) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *BackBufferCopy) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *BackBufferCopy) SetRect(rect Rect2) {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBackBufferCopy.fnSetRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BackBufferCopy) GetRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBackBufferCopy.fnGetRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BackBufferCopy) SetCopyMode(copy_mode BackBufferCopyCopyMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&copy_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBackBufferCopy.fnSetCopyMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BackBufferCopy) GetCopyMode() BackBufferCopyCopyMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BackBufferCopyCopyMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBackBufferCopy.fnGetCopyMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
