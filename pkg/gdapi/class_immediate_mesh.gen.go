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

type ptrsForImmediateMeshList struct {
	fnSurfaceBegin       gdc.MethodBindPtr
	fnSurfaceSetColor    gdc.MethodBindPtr
	fnSurfaceSetNormal   gdc.MethodBindPtr
	fnSurfaceSetTangent  gdc.MethodBindPtr
	fnSurfaceSetUv       gdc.MethodBindPtr
	fnSurfaceSetUv2      gdc.MethodBindPtr
	fnSurfaceAddVertex   gdc.MethodBindPtr
	fnSurfaceAddVertex2D gdc.MethodBindPtr
	fnSurfaceEnd         gdc.MethodBindPtr
	fnClearSurfaces      gdc.MethodBindPtr
}

var ptrsForImmediateMesh ptrsForImmediateMeshList

func initImmediateMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ImmediateMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("surface_begin")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2794442543))
	}
	{
		methodName := StringNameFromStr("surface_set_color")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("surface_set_normal")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceSetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("surface_set_tangent")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceSetTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3505987427))
	}
	{
		methodName := StringNameFromStr("surface_set_uv")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceSetUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("surface_set_uv2")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceSetUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("surface_add_vertex")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceAddVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("surface_add_vertex_2d")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceAddVertex2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("surface_end")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnSurfaceEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_surfaces")
		defer methodName.Destroy()
		ptrsForImmediateMesh.fnClearSurfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type ImmediateMesh struct {
	Mesh
}

func (me *ImmediateMesh) BaseClass() string {
	return "ImmediateMesh"
}

func NewImmediateMesh() *ImmediateMesh {
	str := StringNameFromStr("ImmediateMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ImmediateMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ImmediateMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ImmediateMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ImmediateMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ImmediateMesh) SurfaceBegin(primitive MeshPrimitiveType, material Material) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive), material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceSetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceSetNormal(normal Vector3) {
	cargs := []gdc.ConstTypePtr{normal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceSetNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceSetTangent(tangent Plane) {
	cargs := []gdc.ConstTypePtr{tangent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceSetTangent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceSetUv(uv Vector2) {
	cargs := []gdc.ConstTypePtr{uv.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceSetUv), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceSetUv2(uv2 Vector2) {
	cargs := []gdc.ConstTypePtr{uv2.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceSetUv2), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceAddVertex(vertex Vector3) {
	cargs := []gdc.ConstTypePtr{vertex.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceAddVertex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceAddVertex2D(vertex Vector2) {
	cargs := []gdc.ConstTypePtr{vertex.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceAddVertex2D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) SurfaceEnd() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnSurfaceEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImmediateMesh) ClearSurfaces() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImmediateMesh.fnClearSurfaces), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
