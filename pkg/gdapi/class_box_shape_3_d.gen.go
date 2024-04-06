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

type ptrsForBoxShape3DList struct {
	fnSetSize gdc.MethodBindPtr
	fnGetSize gdc.MethodBindPtr
}

var ptrsForBoxShape3D ptrsForBoxShape3DList

func initBoxShape3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BoxShape3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForBoxShape3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForBoxShape3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
}

type BoxShape3D struct {
	Shape3D
}

func (me *BoxShape3D) BaseClass() string {
	return "BoxShape3D"
}

func NewBoxShape3D() *BoxShape3D {
	str := StringNameFromStr("BoxShape3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &BoxShape3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *BoxShape3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *BoxShape3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *BoxShape3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *BoxShape3D) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxShape3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoxShape3D) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxShape3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
