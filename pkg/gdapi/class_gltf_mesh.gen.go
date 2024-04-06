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

type ptrsForGLTFMeshList struct {
	fnGetMesh              gdc.MethodBindPtr
	fnSetMesh              gdc.MethodBindPtr
	fnGetBlendWeights      gdc.MethodBindPtr
	fnSetBlendWeights      gdc.MethodBindPtr
	fnGetInstanceMaterials gdc.MethodBindPtr
	fnSetInstanceMaterials gdc.MethodBindPtr
}

var ptrsForGLTFMesh ptrsForGLTFMeshList

func initGLTFMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForGLTFMesh.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3754628756))
	}
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForGLTFMesh.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2255166972))
	}
	{
		methodName := StringNameFromStr("get_blend_weights")
		defer methodName.Destroy()
		ptrsForGLTFMesh.fnGetBlendWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2445143706))
	}
	{
		methodName := StringNameFromStr("set_blend_weights")
		defer methodName.Destroy()
		ptrsForGLTFMesh.fnSetBlendWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
	{
		methodName := StringNameFromStr("get_instance_materials")
		defer methodName.Destroy()
		ptrsForGLTFMesh.fnGetInstanceMaterials = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("set_instance_materials")
		defer methodName.Destroy()
		ptrsForGLTFMesh.fnSetInstanceMaterials = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
}

type GLTFMesh struct {
	Resource
}

func (me *GLTFMesh) BaseClass() string {
	return "GLTFMesh"
}

func NewGLTFMesh() *GLTFMesh {
	str := StringNameFromStr("GLTFMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFMesh{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFMesh) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFMesh) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFMesh) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GLTFMesh) GetMesh() ImporterMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImporterMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFMesh.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFMesh) SetMesh(mesh ImporterMesh) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFMesh.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFMesh) GetBlendWeights() PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFMesh.fnGetBlendWeights), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFMesh) SetBlendWeights(blend_weights PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{blend_weights.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFMesh.fnSetBlendWeights), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFMesh) GetInstanceMaterials() []Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFMesh.fnGetInstanceMaterials), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Material](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *GLTFMesh) SetInstanceMaterials(instance_materials []Material) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance_materials)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFMesh.fnSetInstanceMaterials), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
