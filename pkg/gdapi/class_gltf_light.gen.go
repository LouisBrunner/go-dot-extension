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

type ptrsForGLTFLightList struct {
	fnFromNode          gdc.MethodBindPtr
	fnToNode            gdc.MethodBindPtr
	fnFromDictionary    gdc.MethodBindPtr
	fnToDictionary      gdc.MethodBindPtr
	fnGetColor          gdc.MethodBindPtr
	fnSetColor          gdc.MethodBindPtr
	fnGetIntensity      gdc.MethodBindPtr
	fnSetIntensity      gdc.MethodBindPtr
	fnGetLightType      gdc.MethodBindPtr
	fnSetLightType      gdc.MethodBindPtr
	fnGetRange          gdc.MethodBindPtr
	fnSetRange          gdc.MethodBindPtr
	fnGetInnerConeAngle gdc.MethodBindPtr
	fnSetInnerConeAngle gdc.MethodBindPtr
	fnGetOuterConeAngle gdc.MethodBindPtr
	fnSetOuterConeAngle gdc.MethodBindPtr
}

var ptrsForGLTFLight ptrsForGLTFLightList

func initGLTFLightPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFLight")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("from_node")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnFromNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3907677874))
	}
	{
		methodName := StringNameFromStr("to_node")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnToNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2040811672))
	}
	{
		methodName := StringNameFromStr("from_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnFromDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4057087208))
	}
	{
		methodName := StringNameFromStr("to_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnToDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200896285))
	}
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_intensity")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnGetIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_intensity")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnSetIntensity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_light_type")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnGetLightType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("set_light_type")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnSetLightType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_range")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnGetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_range")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnSetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_inner_cone_angle")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnGetInnerConeAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_inner_cone_angle")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnSetInnerConeAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_outer_cone_angle")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnGetOuterConeAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_outer_cone_angle")
		defer methodName.Destroy()
		ptrsForGLTFLight.fnSetOuterConeAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}

}

type GLTFLight struct {
	Resource
}

func (me *GLTFLight) BaseClass() string {
	return "GLTFLight"
}

func NewGLTFLight() *GLTFLight {
	str := StringNameFromStr("GLTFLight") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFLight{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFLight) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFLight) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFLight) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func GLTFLightFromNode(light_node Light3D) GLTFLight {
	cargs := []gdc.ConstTypePtr{light_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFLight()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnFromNode), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFLight) ToNode() Light3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewLight3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnToNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func GLTFLightFromDictionary(dictionary Dictionary) GLTFLight {
	cargs := []gdc.ConstTypePtr{dictionary.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFLight()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnFromDictionary), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFLight) ToDictionary() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnToDictionary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFLight) GetColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFLight) SetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFLight) GetIntensity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnGetIntensity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFLight) SetIntensity(intensity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnSetIntensity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFLight) GetLightType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnGetLightType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFLight) SetLightType(light_type String) {
	cargs := []gdc.ConstTypePtr{light_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnSetLightType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFLight) GetRange() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnGetRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFLight) SetRange(range_ float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&range_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnSetRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFLight) GetInnerConeAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnGetInnerConeAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFLight) SetInnerConeAngle(inner_cone_angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_cone_angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnSetInnerConeAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFLight) GetOuterConeAngle() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnGetOuterConeAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFLight) SetOuterConeAngle(outer_cone_angle float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outer_cone_angle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFLight.fnSetOuterConeAngle), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
