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

type ptrsForPhysicsServer3DRenderingServerHandlerList struct {
	fnXSetVertex gdc.MethodBindPtr
	fnXSetNormal gdc.MethodBindPtr
	fnXSetAabb   gdc.MethodBindPtr
	fnSetVertex  gdc.MethodBindPtr
	fnSetNormal  gdc.MethodBindPtr
	fnSetAabb    gdc.MethodBindPtr
}

var ptrsForPhysicsServer3DRenderingServerHandler ptrsForPhysicsServer3DRenderingServerHandlerList

func initPhysicsServer3DRenderingServerHandlerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer3DRenderingServerHandler")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_vertex")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DRenderingServerHandler.fnSetVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("set_normal")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DRenderingServerHandler.fnSetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("set_aabb")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DRenderingServerHandler.fnSetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
	}

}

type PhysicsServer3DRenderingServerHandler struct {
	Object
}

func (me *PhysicsServer3DRenderingServerHandler) BaseClass() string {
	return "PhysicsServer3DRenderingServerHandler"
}

func NewPhysicsServer3DRenderingServerHandler() *PhysicsServer3DRenderingServerHandler {
	str := StringNameFromStr("PhysicsServer3DRenderingServerHandler") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer3DRenderingServerHandler{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsServer3DRenderingServerHandler) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsServer3DRenderingServerHandler) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3DRenderingServerHandler) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsServer3DRenderingServerHandler) SetVertex(vertex_id int64, vertex Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_id), vertex.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DRenderingServerHandler.fnSetVertex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3DRenderingServerHandler) SetNormal(vertex_id int64, normal Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_id), normal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DRenderingServerHandler.fnSetNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3DRenderingServerHandler) SetAabb(aabb AABB) {
	cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DRenderingServerHandler.fnSetAabb), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
