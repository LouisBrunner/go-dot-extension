// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Light3D struct {
  obj gdc.ObjectPtr
}

func (me *Light3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light3D) BaseClass() string {
  return "Light3D"
}



// Enums

type Light3DParam int
const (
  Light3DParamParamEnergy Light3DParam = 0
  Light3DParamParamIndirectEnergy Light3DParam = 1
  Light3DParamParamVolumetricFogEnergy Light3DParam = 2
  Light3DParamParamSpecular Light3DParam = 3
  Light3DParamParamRange Light3DParam = 4
  Light3DParamParamSize Light3DParam = 5
  Light3DParamParamAttenuation Light3DParam = 6
  Light3DParamParamSpotAngle Light3DParam = 7
  Light3DParamParamSpotAttenuation Light3DParam = 8
  Light3DParamParamShadowMaxDistance Light3DParam = 9
  Light3DParamParamShadowSplit1Offset Light3DParam = 10
  Light3DParamParamShadowSplit2Offset Light3DParam = 11
  Light3DParamParamShadowSplit3Offset Light3DParam = 12
  Light3DParamParamShadowFadeStart Light3DParam = 13
  Light3DParamParamShadowNormalBias Light3DParam = 14
  Light3DParamParamShadowBias Light3DParam = 15
  Light3DParamParamShadowPancakeSize Light3DParam = 16
  Light3DParamParamShadowOpacity Light3DParam = 17
  Light3DParamParamShadowBlur Light3DParam = 18
  Light3DParamParamTransmittanceBias Light3DParam = 19
  Light3DParamParamIntensity Light3DParam = 20
  Light3DParamParamMax Light3DParam = 21
)

type Light3DBakeMode int
const (
  Light3DBakeModeBakeDisabled Light3DBakeMode = 0
  Light3DBakeModeBakeStatic Light3DBakeMode = 1
  Light3DBakeModeBakeDynamic Light3DBakeMode = 2
)

func (me *Light3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Light3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Light3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Light3D) SetEditorOnly(editor_only bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editor_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) IsEditorOnly() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetParam(param Light3DParam, value float32, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1722734213) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetParam(param Light3DParam, ) float32 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1844084987) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetShadow(enabled bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) HasShadow() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetNegative(enabled bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_negative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) IsNegative() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_negative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetCullMask(cull_mask int, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetCullMask() int {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetEnableDistanceFade(enable bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_distance_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) IsDistanceFadeEnabled() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_distance_fade_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetDistanceFadeBegin(distance float32, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetDistanceFadeBegin() float32 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetDistanceFadeShadow(distance float32, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetDistanceFadeShadow() float32 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_shadow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetDistanceFadeLength(distance float32, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetDistanceFadeLength() float32 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetColor() Color {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetShadowReverseCullFace(enable bool, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_reverse_cull_face")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetShadowReverseCullFace() bool {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_reverse_cull_face")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetBakeMode(bake_mode Light3DBakeMode, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 37739303) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bake_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetBakeMode() Light3DBakeMode {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 371737608) // FIXME: should cache?
  var ret Light3DBakeMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetProjector(projector Texture2D, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_projector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(projector.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetProjector() Texture2D {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_projector")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) SetTemperature(temperature float32, )  {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_temperature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&temperature), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light3D) GetTemperature() float32 {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_temperature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light3D) GetCorrelatedColor() Color {
  classNameV := StringNameFromStr("Light3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_correlated_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Light3D) GetPropLightIntensityLumens() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightIntensityLumens(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightIntensityLux() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightIntensityLux(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightTemperature() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightTemperature(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightColor() Color {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightColor(value Color) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightEnergy() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightIndirectEnergy() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightIndirectEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightVolumetricFogEnergy() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightVolumetricFogEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightProjector() Texture2D {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightProjector(value Texture2D) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightSize() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightSize(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightAngularDistance() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightAngularDistance(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightNegative() bool {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightNegative(value bool) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightSpecular() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightSpecular(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightBakeMode() int {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightBakeMode(value int) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropLightCullMask() int {
  panic("TODO: implement")
}

func (me *Light3D) SetPropLightCullMask(value int) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowEnabled() bool {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowBias() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowBias(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowNormalBias() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowNormalBias(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowReverseCullFace() bool {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowReverseCullFace(value bool) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowTransmittanceBias() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowTransmittanceBias(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowOpacity() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowOpacity(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropShadowBlur() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropShadowBlur(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropDistanceFadeEnabled() bool {
  panic("TODO: implement")
}

func (me *Light3D) SetPropDistanceFadeEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropDistanceFadeBegin() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropDistanceFadeBegin(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropDistanceFadeShadow() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropDistanceFadeShadow(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropDistanceFadeLength() float32 {
  panic("TODO: implement")
}

func (me *Light3D) SetPropDistanceFadeLength(value float32) {
  panic("TODO: implement")
}

func (me *Light3D) GetPropEditorOnly() bool {
  panic("TODO: implement")
}

func (me *Light3D) SetPropEditorOnly(value bool) {
  panic("TODO: implement")
}