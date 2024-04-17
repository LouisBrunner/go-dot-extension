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

type ptrsForPrimitiveMeshList struct {
	fnXCreateMeshArray gdc.MethodBindPtr
	fnSetMaterial      gdc.MethodBindPtr
	fnGetMaterial      gdc.MethodBindPtr
	fnGetMeshArrays    gdc.MethodBindPtr
	fnSetCustomAabb    gdc.MethodBindPtr
	fnGetCustomAabb    gdc.MethodBindPtr
	fnSetFlipFaces     gdc.MethodBindPtr
	fnGetFlipFaces     gdc.MethodBindPtr
	fnSetAddUv2        gdc.MethodBindPtr
	fnGetAddUv2        gdc.MethodBindPtr
	fnSetUv2Padding    gdc.MethodBindPtr
	fnGetUv2Padding    gdc.MethodBindPtr
	fnRequestUpdate    gdc.MethodBindPtr
}

var ptrsForPrimitiveMesh ptrsForPrimitiveMeshList

func initPrimitiveMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PrimitiveMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("get_mesh_arrays")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnGetMeshArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_custom_aabb")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnSetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
	}
	{
		methodName := StringNameFromStr("get_custom_aabb")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnGetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
	}
	{
		methodName := StringNameFromStr("set_flip_faces")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnSetFlipFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_flip_faces")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnGetFlipFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_add_uv2")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnSetAddUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_add_uv2")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnGetAddUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_uv2_padding")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnSetUv2Padding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_uv2_padding")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnGetUv2Padding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("request_update")
		defer methodName.Destroy()
		ptrsForPrimitiveMesh.fnRequestUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type PrimitiveMesh struct {
	Mesh
}

func (me *PrimitiveMesh) BaseClass() string {
	return "PrimitiveMesh"
}

func NewPrimitiveMesh() *PrimitiveMesh {
	str := StringNameFromStr("PrimitiveMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PrimitiveMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PrimitiveMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PrimitiveMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PrimitiveMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PrimitiveMesh) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrimitiveMesh) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PrimitiveMesh) GetMeshArrays() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnGetMeshArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PrimitiveMesh) SetCustomAabb(aabb AABB) {
	cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnSetCustomAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrimitiveMesh) GetCustomAabb() AABB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnGetCustomAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PrimitiveMesh) SetFlipFaces(flip_faces bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_faces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnSetFlipFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrimitiveMesh) GetFlipFaces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnGetFlipFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PrimitiveMesh) SetAddUv2(add_uv2 bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&add_uv2)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnSetAddUv2), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrimitiveMesh) GetAddUv2() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnGetAddUv2), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PrimitiveMesh) SetUv2Padding(uv2_padding float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uv2_padding)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnSetUv2Padding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PrimitiveMesh) GetUv2Padding() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnGetUv2Padding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PrimitiveMesh) RequestUpdate() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPrimitiveMesh.fnRequestUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
