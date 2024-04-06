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

type ptrsForVisualShaderNodeResizableBaseList struct {
	fnSetSize gdc.MethodBindPtr
	fnGetSize gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeResizableBase ptrsForVisualShaderNodeResizableBaseList

func initVisualShaderNodeResizableBasePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeResizableBase")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeResizableBase.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeResizableBase.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}

}

type VisualShaderNodeResizableBase struct {
	VisualShaderNode
}

func (me *VisualShaderNodeResizableBase) BaseClass() string {
	return "VisualShaderNodeResizableBase"
}

func NewVisualShaderNodeResizableBase() *VisualShaderNodeResizableBase {
	str := StringNameFromStr("VisualShaderNodeResizableBase") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeResizableBase{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeResizableBase) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeResizableBase) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeResizableBase) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeResizableBase) SetSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeResizableBase.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeResizableBase) GetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeResizableBase.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
