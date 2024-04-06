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

type ptrsForCircleShape2DList struct {
	fnSetRadius gdc.MethodBindPtr
	fnGetRadius gdc.MethodBindPtr
}

var ptrsForCircleShape2D ptrsForCircleShape2DList

func initCircleShape2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CircleShape2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForCircleShape2D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForCircleShape2D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type CircleShape2D struct {
	Shape2D
}

func (me *CircleShape2D) BaseClass() string {
	return "CircleShape2D"
}

func NewCircleShape2D() *CircleShape2D {
	str := StringNameFromStr("CircleShape2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CircleShape2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CircleShape2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CircleShape2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CircleShape2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CircleShape2D) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCircleShape2D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CircleShape2D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCircleShape2D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
