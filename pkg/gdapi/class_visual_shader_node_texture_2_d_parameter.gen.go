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

type ptrsForVisualShaderNodeTexture2DParameterList struct {
}

var ptrsForVisualShaderNodeTexture2DParameter ptrsForVisualShaderNodeTexture2DParameterList

func initVisualShaderNodeTexture2DParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTexture2DParameter")
	defer className.Destroy()

}

type VisualShaderNodeTexture2DParameter struct {
	VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeTexture2DParameter) BaseClass() string {
	return "VisualShaderNodeTexture2DParameter"
}

func NewVisualShaderNodeTexture2DParameter() *VisualShaderNodeTexture2DParameter {
	str := StringNameFromStr("VisualShaderNodeTexture2DParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTexture2DParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTexture2DParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture2DParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture2DParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
