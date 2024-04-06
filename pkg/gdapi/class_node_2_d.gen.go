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

type ptrsForNode2DList struct {
	fnSetPosition                  gdc.MethodBindPtr
	fnSetRotation                  gdc.MethodBindPtr
	fnSetRotationDegrees           gdc.MethodBindPtr
	fnSetSkew                      gdc.MethodBindPtr
	fnSetScale                     gdc.MethodBindPtr
	fnGetPosition                  gdc.MethodBindPtr
	fnGetRotation                  gdc.MethodBindPtr
	fnGetRotationDegrees           gdc.MethodBindPtr
	fnGetSkew                      gdc.MethodBindPtr
	fnGetScale                     gdc.MethodBindPtr
	fnRotate                       gdc.MethodBindPtr
	fnMoveLocalX                   gdc.MethodBindPtr
	fnMoveLocalY                   gdc.MethodBindPtr
	fnTranslate                    gdc.MethodBindPtr
	fnGlobalTranslate              gdc.MethodBindPtr
	fnApplyScale                   gdc.MethodBindPtr
	fnSetGlobalPosition            gdc.MethodBindPtr
	fnGetGlobalPosition            gdc.MethodBindPtr
	fnSetGlobalRotation            gdc.MethodBindPtr
	fnSetGlobalRotationDegrees     gdc.MethodBindPtr
	fnGetGlobalRotation            gdc.MethodBindPtr
	fnGetGlobalRotationDegrees     gdc.MethodBindPtr
	fnSetGlobalSkew                gdc.MethodBindPtr
	fnGetGlobalSkew                gdc.MethodBindPtr
	fnSetGlobalScale               gdc.MethodBindPtr
	fnGetGlobalScale               gdc.MethodBindPtr
	fnSetTransform                 gdc.MethodBindPtr
	fnSetGlobalTransform           gdc.MethodBindPtr
	fnLookAt                       gdc.MethodBindPtr
	fnGetAngleTo                   gdc.MethodBindPtr
	fnToLocal                      gdc.MethodBindPtr
	fnToGlobal                     gdc.MethodBindPtr
	fnGetRelativeTransformToParent gdc.MethodBindPtr
}

var ptrsForNode2D ptrsForNode2DList

func initNode2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Node2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_rotation")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_rotation_degrees")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_skew")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetSkew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_scale")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_rotation")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_rotation_degrees")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_skew")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetSkew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_scale")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("rotate")
		defer methodName.Destroy()
		ptrsForNode2D.fnRotate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("move_local_x")
		defer methodName.Destroy()
		ptrsForNode2D.fnMoveLocalX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2087892650))
	}
	{
		methodName := StringNameFromStr("move_local_y")
		defer methodName.Destroy()
		ptrsForNode2D.fnMoveLocalY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2087892650))
	}
	{
		methodName := StringNameFromStr("translate")
		defer methodName.Destroy()
		ptrsForNode2D.fnTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("global_translate")
		defer methodName.Destroy()
		ptrsForNode2D.fnGlobalTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("apply_scale")
		defer methodName.Destroy()
		ptrsForNode2D.fnApplyScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_global_position")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_global_position")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_global_rotation")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetGlobalRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_global_rotation_degrees")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetGlobalRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_global_rotation")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetGlobalRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_global_rotation_degrees")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetGlobalRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_global_skew")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetGlobalSkew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_global_skew")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetGlobalSkew = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_global_scale")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetGlobalScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_global_scale")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetGlobalScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_transform")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("set_global_transform")
		defer methodName.Destroy()
		ptrsForNode2D.fnSetGlobalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("look_at")
		defer methodName.Destroy()
		ptrsForNode2D.fnLookAt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_angle_to")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetAngleTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2276447920))
	}
	{
		methodName := StringNameFromStr("to_local")
		defer methodName.Destroy()
		ptrsForNode2D.fnToLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
	}
	{
		methodName := StringNameFromStr("to_global")
		defer methodName.Destroy()
		ptrsForNode2D.fnToGlobal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
	}
	{
		methodName := StringNameFromStr("get_relative_transform_to_parent")
		defer methodName.Destroy()
		ptrsForNode2D.fnGetRelativeTransformToParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 904556875))
	}
}

type Node2D struct {
	CanvasItem
}

func (me *Node2D) BaseClass() string {
	return "Node2D"
}

func NewNode2D() *Node2D {
	str := StringNameFromStr("Node2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Node2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Node2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Node2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Node2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Node2D) SetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetRotation(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetRotationDegrees(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetRotationDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetSkew(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetSkew), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetScale(scale Vector2) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node2D) GetRotation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) GetRotationDegrees() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetRotationDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) GetSkew() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetSkew), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) GetScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node2D) Rotate(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnRotate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) MoveLocalX(delta float64, scaled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&scaled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnMoveLocalX), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) MoveLocalY(delta float64, scaled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&scaled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnMoveLocalY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) Translate(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnTranslate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GlobalTranslate(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGlobalTranslate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) ApplyScale(ratio Vector2) {
	cargs := []gdc.ConstTypePtr{ratio.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnApplyScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetGlobalPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetGlobalPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GetGlobalPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetGlobalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node2D) SetGlobalRotation(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetGlobalRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetGlobalRotationDegrees(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetGlobalRotationDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GetGlobalRotation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetGlobalRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) GetGlobalRotationDegrees() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetGlobalRotationDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) SetGlobalSkew(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetGlobalSkew), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GetGlobalSkew() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetGlobalSkew), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) SetGlobalScale(scale Vector2) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetGlobalScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GetGlobalScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetGlobalScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node2D) SetTransform(xform Transform2D) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) SetGlobalTransform(xform Transform2D) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnSetGlobalTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) LookAt(point Vector2) {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnLookAt), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node2D) GetAngleTo(point Vector2) float64 {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetAngleTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node2D) ToLocal(global_point Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{global_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnToLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node2D) ToGlobal(local_point Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnToGlobal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node2D) GetRelativeTransformToParent(parent Node) Transform2D {
	cargs := []gdc.ConstTypePtr{parent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode2D.fnGetRelativeTransformToParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
