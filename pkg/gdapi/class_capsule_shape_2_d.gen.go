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

type ptrsForCapsuleShape2DList struct {
	fnSetRadius gdc.MethodBindPtr
	fnGetRadius gdc.MethodBindPtr
	fnSetHeight gdc.MethodBindPtr
	fnGetHeight gdc.MethodBindPtr
}

var ptrsForCapsuleShape2D ptrsForCapsuleShape2DList

func initCapsuleShape2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CapsuleShape2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForCapsuleShape2D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForCapsuleShape2D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_height")
		defer methodName.Destroy()
		ptrsForCapsuleShape2D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_height")
		defer methodName.Destroy()
		ptrsForCapsuleShape2D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type CapsuleShape2D struct {
	Shape2D
}

func (me *CapsuleShape2D) BaseClass() string {
	return "CapsuleShape2D"
}

func NewCapsuleShape2D() *CapsuleShape2D {
	str := StringNameFromStr("CapsuleShape2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CapsuleShape2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CapsuleShape2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CapsuleShape2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CapsuleShape2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CapsuleShape2D) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleShape2D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CapsuleShape2D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleShape2D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CapsuleShape2D) SetHeight(height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleShape2D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CapsuleShape2D) GetHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCapsuleShape2D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
