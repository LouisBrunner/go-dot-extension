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

type ptrsForMeshInstance2DList struct {
	fnSetMesh    gdc.MethodBindPtr
	fnGetMesh    gdc.MethodBindPtr
	fnSetTexture gdc.MethodBindPtr
	fnGetTexture gdc.MethodBindPtr
}

var ptrsForMeshInstance2D ptrsForMeshInstance2DList

func initMeshInstance2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MeshInstance2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mesh")
		defer methodName.Destroy()
		ptrsForMeshInstance2D.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForMeshInstance2D.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808005922))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForMeshInstance2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForMeshInstance2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
}

type MeshInstance2D struct {
	Node2D
}

func (me *MeshInstance2D) BaseClass() string {
	return "MeshInstance2D"
}

func NewMeshInstance2D() *MeshInstance2D {
	str := StringNameFromStr("MeshInstance2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MeshInstance2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MeshInstance2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MeshInstance2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MeshInstance2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MeshInstance2D) SetMesh(mesh Mesh) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance2D.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshInstance2D) GetMesh() Mesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance2D.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshInstance2D) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshInstance2D) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MeshInstance2DTextureChangedSignalFn func()

func (me *MeshInstance2D) ConnectTextureChanged(subs SignalSubscribers, fn MeshInstance2DTextureChangedSignalFn) {
	sig := StringNameFromStr("texture_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *MeshInstance2D) DisconnectTextureChanged(subs SignalSubscribers, fn MeshInstance2DTextureChangedSignalFn) {
	sig := StringNameFromStr("texture_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
