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

type ptrsForVisualShaderNodeTexture2DArrayList struct {
	fnSetTextureArray gdc.MethodBindPtr
	fnGetTextureArray gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTexture2DArray ptrsForVisualShaderNodeTexture2DArrayList

func initVisualShaderNodeTexture2DArrayPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTexture2DArray")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture_array")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture2DArray.fnSetTextureArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2206200446))
	}
	{
		methodName := StringNameFromStr("get_texture_array")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture2DArray.fnGetTextureArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 146117123))
	}

}

type VisualShaderNodeTexture2DArray struct {
	VisualShaderNodeSample3D
}

func (me *VisualShaderNodeTexture2DArray) BaseClass() string {
	return "VisualShaderNodeTexture2DArray"
}

func NewVisualShaderNodeTexture2DArray() *VisualShaderNodeTexture2DArray {
	str := StringNameFromStr("VisualShaderNodeTexture2DArray") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTexture2DArray{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTexture2DArray) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture2DArray) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture2DArray) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeTexture2DArray) SetTextureArray(value Texture2DArray) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture2DArray.fnSetTextureArray), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTexture2DArray) GetTextureArray() Texture2DArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2DArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture2DArray.fnGetTextureArray), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
