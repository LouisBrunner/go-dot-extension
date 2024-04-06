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

type ptrsForShape2DList struct {
	fnSetCustomSolverBias             gdc.MethodBindPtr
	fnGetCustomSolverBias             gdc.MethodBindPtr
	fnCollide                         gdc.MethodBindPtr
	fnCollideWithMotion               gdc.MethodBindPtr
	fnCollideAndGetContacts           gdc.MethodBindPtr
	fnCollideWithMotionAndGetContacts gdc.MethodBindPtr
	fnDraw                            gdc.MethodBindPtr
	fnGetRect                         gdc.MethodBindPtr
}

var ptrsForShape2D ptrsForShape2DList

func initShape2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Shape2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_custom_solver_bias")
		defer methodName.Destroy()
		ptrsForShape2D.fnSetCustomSolverBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_custom_solver_bias")
		defer methodName.Destroy()
		ptrsForShape2D.fnGetCustomSolverBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("collide")
		defer methodName.Destroy()
		ptrsForShape2D.fnCollide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3709843132))
	}
	{
		methodName := StringNameFromStr("collide_with_motion")
		defer methodName.Destroy()
		ptrsForShape2D.fnCollideWithMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869556801))
	}
	{
		methodName := StringNameFromStr("collide_and_get_contacts")
		defer methodName.Destroy()
		ptrsForShape2D.fnCollideAndGetContacts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3056932662))
	}
	{
		methodName := StringNameFromStr("collide_with_motion_and_get_contacts")
		defer methodName.Destroy()
		ptrsForShape2D.fnCollideWithMotionAndGetContacts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3620351573))
	}
	{
		methodName := StringNameFromStr("draw")
		defer methodName.Destroy()
		ptrsForShape2D.fnDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2948539648))
	}
	{
		methodName := StringNameFromStr("get_rect")
		defer methodName.Destroy()
		ptrsForShape2D.fnGetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}

}

type Shape2D struct {
	Resource
}

func (me *Shape2D) BaseClass() string {
	return "Shape2D"
}

func NewShape2D() *Shape2D {
	str := StringNameFromStr("Shape2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Shape2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Shape2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Shape2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Shape2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Shape2D) SetCustomSolverBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnSetCustomSolverBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shape2D) GetCustomSolverBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnGetCustomSolverBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shape2D) Collide(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D) bool {
	cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnCollide), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shape2D) CollideWithMotion(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2) bool {
	cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), local_motion.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr(), shape_motion.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnCollideWithMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Shape2D) CollideAndGetContacts(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnCollideAndGetContacts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Shape2D) CollideWithMotionAndGetContacts(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), local_motion.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr(), shape_motion.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnCollideWithMotionAndGetContacts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Shape2D) Draw(canvas_item RID, color Color) {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shape2D) GetRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShape2D.fnGetRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
