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

type ptrsForDampedSpringJoint2DList struct {
	fnSetLength     gdc.MethodBindPtr
	fnGetLength     gdc.MethodBindPtr
	fnSetRestLength gdc.MethodBindPtr
	fnGetRestLength gdc.MethodBindPtr
	fnSetStiffness  gdc.MethodBindPtr
	fnGetStiffness  gdc.MethodBindPtr
	fnSetDamping    gdc.MethodBindPtr
	fnGetDamping    gdc.MethodBindPtr
}

var ptrsForDampedSpringJoint2D ptrsForDampedSpringJoint2DList

func initDampedSpringJoint2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("DampedSpringJoint2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_length")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnSetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_length")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rest_length")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnSetRestLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_rest_length")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnGetRestLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_stiffness")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnSetStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_stiffness")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnGetStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_damping")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnSetDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_damping")
		defer methodName.Destroy()
		ptrsForDampedSpringJoint2D.fnGetDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type DampedSpringJoint2D struct {
	Joint2D
}

func (me *DampedSpringJoint2D) BaseClass() string {
	return "DampedSpringJoint2D"
}

func NewDampedSpringJoint2D() *DampedSpringJoint2D {
	str := StringNameFromStr("DampedSpringJoint2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &DampedSpringJoint2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *DampedSpringJoint2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *DampedSpringJoint2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *DampedSpringJoint2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *DampedSpringJoint2D) SetLength(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnSetLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DampedSpringJoint2D) GetLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DampedSpringJoint2D) SetRestLength(rest_length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rest_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnSetRestLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DampedSpringJoint2D) GetRestLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnGetRestLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DampedSpringJoint2D) SetStiffness(stiffness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stiffness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnSetStiffness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DampedSpringJoint2D) GetStiffness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnGetStiffness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DampedSpringJoint2D) SetDamping(damping float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnSetDamping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DampedSpringJoint2D) GetDamping() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDampedSpringJoint2D.fnGetDamping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
