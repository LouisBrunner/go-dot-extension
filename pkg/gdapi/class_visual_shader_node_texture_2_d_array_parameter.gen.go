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

type ptrsForVisualShaderNodeTexture2DArrayParameterList struct {
}

var ptrsForVisualShaderNodeTexture2DArrayParameter ptrsForVisualShaderNodeTexture2DArrayParameterList

func initVisualShaderNodeTexture2DArrayParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTexture2DArrayParameter")
	defer className.Destroy()
}

type VisualShaderNodeTexture2DArrayParameter struct {
	VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeTexture2DArrayParameter) BaseClass() string {
	return "VisualShaderNodeTexture2DArrayParameter"
}

func NewVisualShaderNodeTexture2DArrayParameter() *VisualShaderNodeTexture2DArrayParameter {
	str := StringNameFromStr("VisualShaderNodeTexture2DArrayParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTexture2DArrayParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTexture2DArrayParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture2DArrayParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture2DArrayParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
