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

type ptrsForPlaceholderTexture3DList struct {
	fnSetSize gdc.MethodBindPtr
	fnGetSize gdc.MethodBindPtr
}

var ptrsForPlaceholderTexture3D ptrsForPlaceholderTexture3DList

func initPlaceholderTexture3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PlaceholderTexture3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForPlaceholderTexture3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 560364750))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForPlaceholderTexture3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2785653706))
	}

}

type PlaceholderTexture3D struct {
	Texture3D
}

func (me *PlaceholderTexture3D) BaseClass() string {
	return "PlaceholderTexture3D"
}

func NewPlaceholderTexture3D() *PlaceholderTexture3D {
	str := StringNameFromStr("PlaceholderTexture3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PlaceholderTexture3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PlaceholderTexture3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PlaceholderTexture3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PlaceholderTexture3D) SetSize(size Vector3i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderTexture3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PlaceholderTexture3D) GetSize() Vector3i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderTexture3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
