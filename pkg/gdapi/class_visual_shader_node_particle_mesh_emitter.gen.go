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

type ptrsForVisualShaderNodeParticleMeshEmitterList struct {
	fnSetMesh           gdc.MethodBindPtr
	fnGetMesh           gdc.MethodBindPtr
	fnSetUseAllSurfaces gdc.MethodBindPtr
	fnIsUseAllSurfaces  gdc.MethodBindPtr
	fnSetSurfaceIndex   gdc.MethodBindPtr
	fnGetSurfaceIndex   gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParticleMeshEmitter ptrsForVisualShaderNodeParticleMeshEmitterList

func initVisualShaderNodeParticleMeshEmitterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMeshEmitter.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMeshEmitter.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808005922))
	}
	{
		methodName := StringNameFromStr("set_use_all_surfaces")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMeshEmitter.fnSetUseAllSurfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_use_all_surfaces")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMeshEmitter.fnIsUseAllSurfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_surface_index")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMeshEmitter.fnSetSurfaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_surface_index")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParticleMeshEmitter.fnGetSurfaceIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type VisualShaderNodeParticleMeshEmitter struct {
	VisualShaderNodeParticleEmitter
}

func (me *VisualShaderNodeParticleMeshEmitter) BaseClass() string {
	return "VisualShaderNodeParticleMeshEmitter"
}

func NewVisualShaderNodeParticleMeshEmitter() *VisualShaderNodeParticleMeshEmitter {
	str := StringNameFromStr("VisualShaderNodeParticleMeshEmitter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParticleMeshEmitter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeParticleMeshEmitter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleMeshEmitter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleMeshEmitter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParticleMeshEmitter) SetMesh(mesh Mesh) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMeshEmitter.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleMeshEmitter) GetMesh() Mesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMeshEmitter.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeParticleMeshEmitter) SetUseAllSurfaces(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMeshEmitter.fnSetUseAllSurfaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleMeshEmitter) IsUseAllSurfaces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMeshEmitter.fnIsUseAllSurfaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeParticleMeshEmitter) SetSurfaceIndex(surface_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMeshEmitter.fnSetSurfaceIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParticleMeshEmitter) GetSurfaceIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleMeshEmitter.fnGetSurfaceIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
