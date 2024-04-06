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

type ptrsForWorldBoundaryShape2DList struct {
	fnSetNormal   gdc.MethodBindPtr
	fnGetNormal   gdc.MethodBindPtr
	fnSetDistance gdc.MethodBindPtr
	fnGetDistance gdc.MethodBindPtr
}

var ptrsForWorldBoundaryShape2D ptrsForWorldBoundaryShape2DList

func initWorldBoundaryShape2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WorldBoundaryShape2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_normal")
		defer methodName.Destroy()
		ptrsForWorldBoundaryShape2D.fnSetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_normal")
		defer methodName.Destroy()
		ptrsForWorldBoundaryShape2D.fnGetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_distance")
		defer methodName.Destroy()
		ptrsForWorldBoundaryShape2D.fnSetDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_distance")
		defer methodName.Destroy()
		ptrsForWorldBoundaryShape2D.fnGetDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type WorldBoundaryShape2D struct {
	Shape2D
}

func (me *WorldBoundaryShape2D) BaseClass() string {
	return "WorldBoundaryShape2D"
}

func NewWorldBoundaryShape2D() *WorldBoundaryShape2D {
	str := StringNameFromStr("WorldBoundaryShape2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WorldBoundaryShape2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *WorldBoundaryShape2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *WorldBoundaryShape2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *WorldBoundaryShape2D) SetNormal(normal Vector2) {
	cargs := []gdc.ConstTypePtr{normal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldBoundaryShape2D.fnSetNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WorldBoundaryShape2D) GetNormal() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldBoundaryShape2D.fnGetNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WorldBoundaryShape2D) SetDistance(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldBoundaryShape2D.fnSetDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WorldBoundaryShape2D) GetDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorldBoundaryShape2D.fnGetDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
