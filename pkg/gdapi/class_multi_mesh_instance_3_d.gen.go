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

type ptrsForMultiMeshInstance3DList struct {
	fnSetMultimesh gdc.MethodBindPtr
	fnGetMultimesh gdc.MethodBindPtr
}

var ptrsForMultiMeshInstance3D ptrsForMultiMeshInstance3DList

func initMultiMeshInstance3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MultiMeshInstance3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_multimesh")
		defer methodName.Destroy()
		ptrsForMultiMeshInstance3D.fnSetMultimesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2246127404))
	}
	{
		methodName := StringNameFromStr("get_multimesh")
		defer methodName.Destroy()
		ptrsForMultiMeshInstance3D.fnGetMultimesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1385450523))
	}
}

type MultiMeshInstance3D struct {
	GeometryInstance3D
}

func (me *MultiMeshInstance3D) BaseClass() string {
	return "MultiMeshInstance3D"
}

func NewMultiMeshInstance3D() *MultiMeshInstance3D {
	str := StringNameFromStr("MultiMeshInstance3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MultiMeshInstance3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MultiMeshInstance3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MultiMeshInstance3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MultiMeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MultiMeshInstance3D) SetMultimesh(multimesh MultiMesh) {
	cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMeshInstance3D.fnSetMultimesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MultiMeshInstance3D) GetMultimesh() MultiMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMultiMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMultiMeshInstance3D.fnGetMultimesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
