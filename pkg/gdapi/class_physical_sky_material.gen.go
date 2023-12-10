// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicalSkyMaterial struct {
  obj gdc.ObjectPtr
}

func (me *PhysicalSkyMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicalSkyMaterial) BaseClass() string {
  return "PhysicalSkyMaterial"
}



// Enums

func (me *PhysicalSkyMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicalSkyMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicalSkyMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicalSkyMaterial) SetRayleighCoefficient(rayleigh float32, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rayleigh_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rayleigh), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetRayleighCoefficient() float32 {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rayleigh_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetRayleighColor(color Color, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rayleigh_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetRayleighColor() Color {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rayleigh_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetMieCoefficient(mie float32, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mie_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mie), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetMieCoefficient() float32 {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mie_coefficient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetMieEccentricity(eccentricity float32, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mie_eccentricity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&eccentricity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetMieEccentricity() float32 {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mie_eccentricity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetMieColor(color Color, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mie_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetMieColor() Color {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mie_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetTurbidity(turbidity float32, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_turbidity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&turbidity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetTurbidity() float32 {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_turbidity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetSunDiskScale(scale float32, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sun_disk_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetSunDiskScale() float32 {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sun_disk_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetGroundColor(color Color, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ground_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetGroundColor() Color {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ground_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetEnergyMultiplier(multiplier float32, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetEnergyMultiplier() float32 {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_energy_multiplier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetUseDebanding(use_debanding bool, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_debanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_debanding), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetUseDebanding() bool {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_debanding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalSkyMaterial) SetNightSky(night_sky Texture2D, )  {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_night_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(night_sky.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalSkyMaterial) GetNightSky() Texture2D {
  classNameV := StringNameFromStr("PhysicalSkyMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_night_sky")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PhysicalSkyMaterial) GetPropRayleighCoefficient() float32 {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropRayleighCoefficient(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropRayleighColor() Color {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropRayleighColor(value Color) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropMieCoefficient() float32 {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropMieCoefficient(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropMieEccentricity() float32 {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropMieEccentricity(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropMieColor() Color {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropMieColor(value Color) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropTurbidity() float32 {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropTurbidity(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropSunDiskScale() float32 {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropSunDiskScale(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropGroundColor() Color {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropGroundColor(value Color) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropEnergyMultiplier() float32 {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropEnergyMultiplier(value float32) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropUseDebanding() bool {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropUseDebanding(value bool) {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) GetPropNightSky() Texture2D {
  panic("TODO: implement")
}

func (me *PhysicalSkyMaterial) SetPropNightSky(value Texture2D) {
  panic("TODO: implement")
}