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

type ptrsForVisualShaderNodeTexture3DList struct {
	fnSetTexture gdc.MethodBindPtr
	fnGetTexture gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTexture3D ptrsForVisualShaderNodeTexture3DList

func initVisualShaderNodeTexture3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTexture3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture3D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1188404210))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture3D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373985333))
	}

}

type VisualShaderNodeTexture3D struct {
	VisualShaderNodeSample3D
}

func (me *VisualShaderNodeTexture3D) BaseClass() string {
	return "VisualShaderNodeTexture3D"
}

func NewVisualShaderNodeTexture3D() *VisualShaderNodeTexture3D {
	str := StringNameFromStr("VisualShaderNodeTexture3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTexture3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTexture3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeTexture3D) SetTexture(value Texture3D) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture3D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTexture3D) GetTexture() Texture3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture3D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
