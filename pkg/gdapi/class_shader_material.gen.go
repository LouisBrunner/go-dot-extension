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

type ptrsForShaderMaterialList struct {
	fnSetShader          gdc.MethodBindPtr
	fnGetShader          gdc.MethodBindPtr
	fnSetShaderParameter gdc.MethodBindPtr
	fnGetShaderParameter gdc.MethodBindPtr
}

var ptrsForShaderMaterial ptrsForShaderMaterialList

func initShaderMaterialPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ShaderMaterial")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_shader")
		defer methodName.Destroy()
		ptrsForShaderMaterial.fnSetShader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341921675))
	}
	{
		methodName := StringNameFromStr("get_shader")
		defer methodName.Destroy()
		ptrsForShaderMaterial.fnGetShader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2078273437))
	}
	{
		methodName := StringNameFromStr("set_shader_parameter")
		defer methodName.Destroy()
		ptrsForShaderMaterial.fnSetShaderParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("get_shader_parameter")
		defer methodName.Destroy()
		ptrsForShaderMaterial.fnGetShaderParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2760726917))
	}
}

type ShaderMaterial struct {
	Material
}

func (me *ShaderMaterial) BaseClass() string {
	return "ShaderMaterial"
}

func NewShaderMaterial() *ShaderMaterial {
	str := StringNameFromStr("ShaderMaterial") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ShaderMaterial{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ShaderMaterial) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ShaderMaterial) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ShaderMaterial) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ShaderMaterial) SetShader(shader Shader) {
	cargs := []gdc.ConstTypePtr{shader.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShaderMaterial.fnSetShader), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ShaderMaterial) GetShader() Shader {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShader()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShaderMaterial.fnGetShader), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ShaderMaterial) SetShaderParameter(param StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{param.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShaderMaterial.fnSetShaderParameter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ShaderMaterial) GetShaderParameter(param StringName) Variant {
	cargs := []gdc.ConstTypePtr{param.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShaderMaterial.fnGetShaderParameter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
