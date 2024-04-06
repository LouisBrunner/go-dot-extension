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

type ptrsForConcavePolygonShape3DList struct {
	fnSetFaces                    gdc.MethodBindPtr
	fnGetFaces                    gdc.MethodBindPtr
	fnSetBackfaceCollisionEnabled gdc.MethodBindPtr
	fnIsBackfaceCollisionEnabled  gdc.MethodBindPtr
}

var ptrsForConcavePolygonShape3D ptrsForConcavePolygonShape3DList

func initConcavePolygonShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ConcavePolygonShape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_faces")
		defer methodName.Destroy()
		ptrsForConcavePolygonShape3D.fnSetFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
	}
	{
		methodName := StringNameFromStr("get_faces")
		defer methodName.Destroy()
		ptrsForConcavePolygonShape3D.fnGetFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
	{
		methodName := StringNameFromStr("set_backface_collision_enabled")
		defer methodName.Destroy()
		ptrsForConcavePolygonShape3D.fnSetBackfaceCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_backface_collision_enabled")
		defer methodName.Destroy()
		ptrsForConcavePolygonShape3D.fnIsBackfaceCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type ConcavePolygonShape3D struct {
	Shape3D
}

func (me *ConcavePolygonShape3D) BaseClass() string {
	return "ConcavePolygonShape3D"
}

func NewConcavePolygonShape3D() *ConcavePolygonShape3D {
	str := StringNameFromStr("ConcavePolygonShape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ConcavePolygonShape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ConcavePolygonShape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ConcavePolygonShape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ConcavePolygonShape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ConcavePolygonShape3D) SetFaces(faces PackedVector3Array) {
	cargs := []gdc.ConstTypePtr{faces.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConcavePolygonShape3D.fnSetFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConcavePolygonShape3D) GetFaces() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConcavePolygonShape3D.fnGetFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ConcavePolygonShape3D) SetBackfaceCollisionEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConcavePolygonShape3D.fnSetBackfaceCollisionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ConcavePolygonShape3D) IsBackfaceCollisionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForConcavePolygonShape3D.fnIsBackfaceCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
