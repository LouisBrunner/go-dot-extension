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

type ptrsForGrooveJoint2DList struct {
	fnSetLength        gdc.MethodBindPtr
	fnGetLength        gdc.MethodBindPtr
	fnSetInitialOffset gdc.MethodBindPtr
	fnGetInitialOffset gdc.MethodBindPtr
}

var ptrsForGrooveJoint2D ptrsForGrooveJoint2DList

func initGrooveJoint2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GrooveJoint2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_length")
		defer methodName.Destroy()
		ptrsForGrooveJoint2D.fnSetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_length")
		defer methodName.Destroy()
		ptrsForGrooveJoint2D.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_initial_offset")
		defer methodName.Destroy()
		ptrsForGrooveJoint2D.fnSetInitialOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_initial_offset")
		defer methodName.Destroy()
		ptrsForGrooveJoint2D.fnGetInitialOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type GrooveJoint2D struct {
	Joint2D
}

func (me *GrooveJoint2D) BaseClass() string {
	return "GrooveJoint2D"
}

func NewGrooveJoint2D() *GrooveJoint2D {
	str := StringNameFromStr("GrooveJoint2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GrooveJoint2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GrooveJoint2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GrooveJoint2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GrooveJoint2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GrooveJoint2D) SetLength(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGrooveJoint2D.fnSetLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GrooveJoint2D) GetLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGrooveJoint2D.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GrooveJoint2D) SetInitialOffset(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGrooveJoint2D.fnSetInitialOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GrooveJoint2D) GetInitialOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGrooveJoint2D.fnGetInitialOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
