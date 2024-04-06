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

type ptrsForMultiMeshInstance2DList struct {
	fnSetMultimesh gdc.MethodBindPtr
	fnGetMultimesh gdc.MethodBindPtr
	fnSetTexture   gdc.MethodBindPtr
	fnGetTexture   gdc.MethodBindPtr
}

var ptrsForMultiMeshInstance2D ptrsForMultiMeshInstance2DList

func initMultiMeshInstance2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MultiMeshInstance2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_multimesh")
		defer methodName.Destroy()
		ptrsForMultiMeshInstance2D.fnSetMultimesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2246127404))
	}
	{
		methodName := StringNameFromStr("get_multimesh")
		defer methodName.Destroy()
		ptrsForMultiMeshInstance2D.fnGetMultimesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1385450523))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForMultiMeshInstance2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForMultiMeshInstance2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
}

type MultiMeshInstance2D struct {
	Node2D
}

func (me *MultiMeshInstance2D) BaseClass() string {
	return "MultiMeshInstance2D"
}

func NewMultiMeshInstance2D() *MultiMeshInstance2D {
	str := StringNameFromStr("MultiMeshInstance2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MultiMeshInstance2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MultiMeshInstance2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MultiMeshInstance2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MultiMeshInstance2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MultiMeshInstance2D) SetMultimesh(multimesh MultiMesh) {
	cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMeshInstance2D.fnSetMultimesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMeshInstance2D) GetMultimesh() MultiMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMultiMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMeshInstance2D.fnGetMultimesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MultiMeshInstance2D) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMeshInstance2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMeshInstance2D) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMeshInstance2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type MultiMeshInstance2DTextureChangedSignalFn func()

func (me *MultiMeshInstance2D) ConnectTextureChanged(subs SignalSubscribers, fn MultiMeshInstance2DTextureChangedSignalFn) {
	sig := StringNameFromStr("texture_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *MultiMeshInstance2D) DisconnectTextureChanged(subs SignalSubscribers, fn MultiMeshInstance2DTextureChangedSignalFn) {
	sig := StringNameFromStr("texture_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
