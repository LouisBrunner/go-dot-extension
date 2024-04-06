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

type ptrsForVisualShaderNodeTexture3DParameterList struct {
}

var ptrsForVisualShaderNodeTexture3DParameter ptrsForVisualShaderNodeTexture3DParameterList

func initVisualShaderNodeTexture3DParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTexture3DParameter")
	defer className.Destroy()
}

type VisualShaderNodeTexture3DParameter struct {
	VisualShaderNodeTextureParameter
}

func (me *VisualShaderNodeTexture3DParameter) BaseClass() string {
	return "VisualShaderNodeTexture3DParameter"
}

func NewVisualShaderNodeTexture3DParameter() *VisualShaderNodeTexture3DParameter {
	str := StringNameFromStr("VisualShaderNodeTexture3DParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTexture3DParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTexture3DParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture3DParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture3DParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
