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

type ptrsForVisualShaderNodeTextureList struct {
	fnSetSource      gdc.MethodBindPtr
	fnGetSource      gdc.MethodBindPtr
	fnSetTexture     gdc.MethodBindPtr
	fnGetTexture     gdc.MethodBindPtr
	fnSetTextureType gdc.MethodBindPtr
	fnGetTextureType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTexture ptrsForVisualShaderNodeTextureList

func initVisualShaderNodeTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_source")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture.fnSetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 905262939))
	}
	{
		methodName := StringNameFromStr("get_source")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture.fnGetSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2896297444))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_texture_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture.fnSetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 986314081))
	}
	{
		methodName := StringNameFromStr("get_texture_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTexture.fnGetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3290430153))
	}

}

type VisualShaderNodeTexture struct {
	VisualShaderNode
}

func (me *VisualShaderNodeTexture) BaseClass() string {
	return "VisualShaderNodeTexture"
}

func NewVisualShaderNodeTexture() *VisualShaderNodeTexture {
	str := StringNameFromStr("VisualShaderNodeTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeTextureSource int

const (
	VisualShaderNodeTextureSourceSourceTexture   VisualShaderNodeTextureSource = 0
	VisualShaderNodeTextureSourceSourceScreen    VisualShaderNodeTextureSource = 1
	VisualShaderNodeTextureSourceSource2DTexture VisualShaderNodeTextureSource = 2
	VisualShaderNodeTextureSourceSource2DNormal  VisualShaderNodeTextureSource = 3
	VisualShaderNodeTextureSourceSourceDepth     VisualShaderNodeTextureSource = 4
	VisualShaderNodeTextureSourceSourcePort      VisualShaderNodeTextureSource = 5
	VisualShaderNodeTextureSourceSource3DNormal  VisualShaderNodeTextureSource = 6
	VisualShaderNodeTextureSourceSourceRoughness VisualShaderNodeTextureSource = 7
	VisualShaderNodeTextureSourceSourceMax       VisualShaderNodeTextureSource = 8
)

type VisualShaderNodeTextureTextureType int

const (
	VisualShaderNodeTextureTextureTypeTypeData      VisualShaderNodeTextureTextureType = 0
	VisualShaderNodeTextureTextureTypeTypeColor     VisualShaderNodeTextureTextureType = 1
	VisualShaderNodeTextureTextureTypeTypeNormalMap VisualShaderNodeTextureTextureType = 2
	VisualShaderNodeTextureTextureTypeTypeMax       VisualShaderNodeTextureTextureType = 3
)

func (me *VisualShaderNodeTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeTexture) SetSource(value VisualShaderNodeTextureSource) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture.fnSetSource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTexture) GetSource() VisualShaderNodeTextureSource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureSource

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture.fnGetSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeTexture) SetTexture(value Texture2D) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTexture) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeTexture) SetTextureType(value VisualShaderNodeTextureTextureType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture.fnSetTextureType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTexture) GetTextureType() VisualShaderNodeTextureTextureType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureTextureType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTexture.fnGetTextureType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
