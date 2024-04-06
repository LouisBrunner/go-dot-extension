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

type ptrsForShape3DList struct {
	fnSetCustomSolverBias gdc.MethodBindPtr
	fnGetCustomSolverBias gdc.MethodBindPtr
	fnSetMargin           gdc.MethodBindPtr
	fnGetMargin           gdc.MethodBindPtr
	fnGetDebugMesh        gdc.MethodBindPtr
}

var ptrsForShape3D ptrsForShape3DList

func initShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Shape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_custom_solver_bias")
		defer methodName.Destroy()
		ptrsForShape3D.fnSetCustomSolverBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_custom_solver_bias")
		defer methodName.Destroy()
		ptrsForShape3D.fnGetCustomSolverBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_margin")
		defer methodName.Destroy()
		ptrsForShape3D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_margin")
		defer methodName.Destroy()
		ptrsForShape3D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_debug_mesh")
		defer methodName.Destroy()
		ptrsForShape3D.fnGetDebugMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1605880883))
	}

}

type Shape3D struct {
	Resource
}

func (me *Shape3D) BaseClass() string {
	return "Shape3D"
}

func NewShape3D() *Shape3D {
	str := StringNameFromStr("Shape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Shape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Shape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Shape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Shape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Shape3D) SetCustomSolverBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape3D.fnSetCustomSolverBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shape3D) GetCustomSolverBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape3D.fnGetCustomSolverBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shape3D) SetMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape3D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shape3D) GetMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape3D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shape3D) GetDebugMesh() ArrayMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArrayMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape3D.fnGetDebugMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
