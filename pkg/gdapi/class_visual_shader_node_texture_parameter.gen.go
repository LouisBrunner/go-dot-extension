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

type ptrsForVisualShaderNodeTextureParameterList struct {
	fnSetTextureType   gdc.MethodBindPtr
	fnGetTextureType   gdc.MethodBindPtr
	fnSetColorDefault  gdc.MethodBindPtr
	fnGetColorDefault  gdc.MethodBindPtr
	fnSetTextureFilter gdc.MethodBindPtr
	fnGetTextureFilter gdc.MethodBindPtr
	fnSetTextureRepeat gdc.MethodBindPtr
	fnGetTextureRepeat gdc.MethodBindPtr
	fnSetTextureSource gdc.MethodBindPtr
	fnGetTextureSource gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTextureParameter ptrsForVisualShaderNodeTextureParameterList

func initVisualShaderNodeTextureParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTextureParameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnSetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2227296876))
	}
	{
		methodName := StringNameFromStr("get_texture_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnGetTextureType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 367922070))
	}
	{
		methodName := StringNameFromStr("set_color_default")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnSetColorDefault = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4217624432))
	}
	{
		methodName := StringNameFromStr("get_color_default")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnGetColorDefault = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3837060134))
	}
	{
		methodName := StringNameFromStr("set_texture_filter")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2147684752))
	}
	{
		methodName := StringNameFromStr("get_texture_filter")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnGetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4184490817))
	}
	{
		methodName := StringNameFromStr("set_texture_repeat")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnSetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2036143070))
	}
	{
		methodName := StringNameFromStr("get_texture_repeat")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnGetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1690132794))
	}
	{
		methodName := StringNameFromStr("set_texture_source")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnSetTextureSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1212687372))
	}
	{
		methodName := StringNameFromStr("get_texture_source")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTextureParameter.fnGetTextureSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2039092262))
	}
}

type VisualShaderNodeTextureParameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeTextureParameter) BaseClass() string {
	return "VisualShaderNodeTextureParameter"
}

func NewVisualShaderNodeTextureParameter() *VisualShaderNodeTextureParameter {
	str := StringNameFromStr("VisualShaderNodeTextureParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTextureParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeTextureParameterTextureType int

const (
	VisualShaderNodeTextureParameterTextureTypeTypeData       VisualShaderNodeTextureParameterTextureType = 0
	VisualShaderNodeTextureParameterTextureTypeTypeColor      VisualShaderNodeTextureParameterTextureType = 1
	VisualShaderNodeTextureParameterTextureTypeTypeNormalMap  VisualShaderNodeTextureParameterTextureType = 2
	VisualShaderNodeTextureParameterTextureTypeTypeAnisotropy VisualShaderNodeTextureParameterTextureType = 3
	VisualShaderNodeTextureParameterTextureTypeTypeMax        VisualShaderNodeTextureParameterTextureType = 4
)

type VisualShaderNodeTextureParameterColorDefault int

const (
	VisualShaderNodeTextureParameterColorDefaultColorDefaultWhite       VisualShaderNodeTextureParameterColorDefault = 0
	VisualShaderNodeTextureParameterColorDefaultColorDefaultBlack       VisualShaderNodeTextureParameterColorDefault = 1
	VisualShaderNodeTextureParameterColorDefaultColorDefaultTransparent VisualShaderNodeTextureParameterColorDefault = 2
	VisualShaderNodeTextureParameterColorDefaultColorDefaultMax         VisualShaderNodeTextureParameterColorDefault = 3
)

type VisualShaderNodeTextureParameterTextureFilter int

const (
	VisualShaderNodeTextureParameterTextureFilterFilterDefault                  VisualShaderNodeTextureParameterTextureFilter = 0
	VisualShaderNodeTextureParameterTextureFilterFilterNearest                  VisualShaderNodeTextureParameterTextureFilter = 1
	VisualShaderNodeTextureParameterTextureFilterFilterLinear                   VisualShaderNodeTextureParameterTextureFilter = 2
	VisualShaderNodeTextureParameterTextureFilterFilterNearestMipmap            VisualShaderNodeTextureParameterTextureFilter = 3
	VisualShaderNodeTextureParameterTextureFilterFilterLinearMipmap             VisualShaderNodeTextureParameterTextureFilter = 4
	VisualShaderNodeTextureParameterTextureFilterFilterNearestMipmapAnisotropic VisualShaderNodeTextureParameterTextureFilter = 5
	VisualShaderNodeTextureParameterTextureFilterFilterLinearMipmapAnisotropic  VisualShaderNodeTextureParameterTextureFilter = 6
	VisualShaderNodeTextureParameterTextureFilterFilterMax                      VisualShaderNodeTextureParameterTextureFilter = 7
)

type VisualShaderNodeTextureParameterTextureRepeat int

const (
	VisualShaderNodeTextureParameterTextureRepeatRepeatDefault  VisualShaderNodeTextureParameterTextureRepeat = 0
	VisualShaderNodeTextureParameterTextureRepeatRepeatEnabled  VisualShaderNodeTextureParameterTextureRepeat = 1
	VisualShaderNodeTextureParameterTextureRepeatRepeatDisabled VisualShaderNodeTextureParameterTextureRepeat = 2
	VisualShaderNodeTextureParameterTextureRepeatRepeatMax      VisualShaderNodeTextureParameterTextureRepeat = 3
)

type VisualShaderNodeTextureParameterTextureSource int

const (
	VisualShaderNodeTextureParameterTextureSourceSourceNone            VisualShaderNodeTextureParameterTextureSource = 0
	VisualShaderNodeTextureParameterTextureSourceSourceScreen          VisualShaderNodeTextureParameterTextureSource = 1
	VisualShaderNodeTextureParameterTextureSourceSourceDepth           VisualShaderNodeTextureParameterTextureSource = 2
	VisualShaderNodeTextureParameterTextureSourceSourceNormalRoughness VisualShaderNodeTextureParameterTextureSource = 3
	VisualShaderNodeTextureParameterTextureSourceSourceMax             VisualShaderNodeTextureParameterTextureSource = 4
)

func (me *VisualShaderNodeTextureParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeTextureParameter) SetTextureType(type_ VisualShaderNodeTextureParameterTextureType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnSetTextureType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTextureParameter) GetTextureType() VisualShaderNodeTextureParameterTextureType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureParameterTextureType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnGetTextureType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeTextureParameter) SetColorDefault(color VisualShaderNodeTextureParameterColorDefault) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&color)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnSetColorDefault), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTextureParameter) GetColorDefault() VisualShaderNodeTextureParameterColorDefault {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureParameterColorDefault

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnGetColorDefault), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeTextureParameter) SetTextureFilter(filter VisualShaderNodeTextureParameterTextureFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTextureParameter) GetTextureFilter() VisualShaderNodeTextureParameterTextureFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureParameterTextureFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnGetTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeTextureParameter) SetTextureRepeat(repeat VisualShaderNodeTextureParameterTextureRepeat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnSetTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTextureParameter) GetTextureRepeat() VisualShaderNodeTextureParameterTextureRepeat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureParameterTextureRepeat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnGetTextureRepeat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeTextureParameter) SetTextureSource(source VisualShaderNodeTextureParameterTextureSource) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnSetTextureSource), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTextureParameter) GetTextureSource() VisualShaderNodeTextureParameterTextureSource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeTextureParameterTextureSource

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTextureParameter.fnGetTextureSource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
