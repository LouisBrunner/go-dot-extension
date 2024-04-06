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

type ptrsForCSGBox3DList struct {
	fnSetSize     gdc.MethodBindPtr
	fnGetSize     gdc.MethodBindPtr
	fnSetMaterial gdc.MethodBindPtr
	fnGetMaterial gdc.MethodBindPtr
}

var ptrsForCSGBox3D ptrsForCSGBox3DList

func initCSGBox3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CSGBox3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForCSGBox3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForCSGBox3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForCSGBox3D.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForCSGBox3D.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
}

type CSGBox3D struct {
	CSGPrimitive3D
}

func (me *CSGBox3D) BaseClass() string {
	return "CSGBox3D"
}

func NewCSGBox3D() *CSGBox3D {
	str := StringNameFromStr("CSGBox3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CSGBox3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CSGBox3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CSGBox3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CSGBox3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CSGBox3D) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGBox3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGBox3D) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGBox3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CSGBox3D) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGBox3D.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CSGBox3D) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGBox3D.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
