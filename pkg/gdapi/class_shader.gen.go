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

type ptrsForShaderList struct {
	fnGetMode                    gdc.MethodBindPtr
	fnSetCode                    gdc.MethodBindPtr
	fnGetCode                    gdc.MethodBindPtr
	fnSetDefaultTextureParameter gdc.MethodBindPtr
	fnGetDefaultTextureParameter gdc.MethodBindPtr
	fnGetShaderUniformList       gdc.MethodBindPtr
}

var ptrsForShader ptrsForShaderList

func initShaderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Shader")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_mode")
		defer methodName.Destroy()
		ptrsForShader.fnGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3392948163))
	}
	{
		methodName := StringNameFromStr("set_code")
		defer methodName.Destroy()
		ptrsForShader.fnSetCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_code")
		defer methodName.Destroy()
		ptrsForShader.fnGetCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_default_texture_parameter")
		defer methodName.Destroy()
		ptrsForShader.fnSetDefaultTextureParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2750740428))
	}
	{
		methodName := StringNameFromStr("get_default_texture_parameter")
		defer methodName.Destroy()
		ptrsForShader.fnGetDefaultTextureParameter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3090538643))
	}
	{
		methodName := StringNameFromStr("get_shader_uniform_list")
		defer methodName.Destroy()
		ptrsForShader.fnGetShaderUniformList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1230511656))
	}
}

type Shader struct {
	Resource
}

func (me *Shader) BaseClass() string {
	return "Shader"
}

func NewShader() *Shader {
	str := StringNameFromStr("Shader") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Shader{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ShaderMode int

const (
	ShaderModeModeSpatial    ShaderMode = 0
	ShaderModeModeCanvasItem ShaderMode = 1
	ShaderModeModeParticles  ShaderMode = 2
	ShaderModeModeSky        ShaderMode = 3
	ShaderModeModeFog        ShaderMode = 4
)

func (me *Shader) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Shader) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Shader) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Shader) GetMode() ShaderMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ShaderMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShader.fnGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Shader) SetCode(code String) {
	cargs := []gdc.ConstTypePtr{code.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShader.fnSetCode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shader) GetCode() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShader.fnGetCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Shader) SetDefaultTextureParameter(name StringName, texture Texture2D, index int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShader.fnSetDefaultTextureParameter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Shader) GetDefaultTextureParameter(name StringName, index int64) Texture2D {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShader.fnGetDefaultTextureParameter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Shader) GetShaderUniformList(get_groups bool) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&get_groups)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&get_groups)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShader.fnGetShaderUniformList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
