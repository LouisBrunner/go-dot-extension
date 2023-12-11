// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ReflectionProbe struct {
  obj gdc.ObjectPtr
}

func (me *ReflectionProbe) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ReflectionProbe) BaseClass() string {
  return "ReflectionProbe"
}



// Enums

type ReflectionProbeUpdateMode int
const (
  ReflectionProbeUpdateModeUpdateOnce ReflectionProbeUpdateMode = 0
  ReflectionProbeUpdateModeUpdateAlways ReflectionProbeUpdateMode = 1
)

type ReflectionProbeAmbientMode int
const (
  ReflectionProbeAmbientModeAmbientDisabled ReflectionProbeAmbientMode = 0
  ReflectionProbeAmbientModeAmbientEnvironment ReflectionProbeAmbientMode = 1
  ReflectionProbeAmbientModeAmbientColor ReflectionProbeAmbientMode = 2
)

func (me *ReflectionProbe) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ReflectionProbe) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ReflectionProbe) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ReflectionProbe) SetIntensity(intensity float32, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetIntensity() float32 {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetAmbientMode(ambient ReflectionProbeAmbientMode, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1748981278) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ambient), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetAmbientMode() ReflectionProbeAmbientMode {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1014607621) // FIXME: should cache?
  var ret ReflectionProbeAmbientMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetAmbientColor(ambient Color, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(ambient.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetAmbientColor() Color {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetAmbientColorEnergy(ambient_energy float32, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ambient_color_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ambient_energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetAmbientColorEnergy() float32 {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ambient_color_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetMaxDistance(max_distance float32, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetMaxDistance() float32 {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetMeshLodThreshold(ratio float32, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh_lod_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetMeshLodThreshold() float32 {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh_lod_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetSize() Vector3 {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetOriginOffset(origin_offset Vector3, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_origin_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(origin_offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetOriginOffset() Vector3 {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_origin_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetAsInterior(enable bool, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) IsSetAsInterior() bool {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_set_as_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetEnableBoxProjection(enable bool, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_box_projection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) IsBoxProjectionEnabled() bool {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_box_projection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetEnableShadows(enable bool, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_shadows")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) AreShadowsEnabled() bool {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_shadows_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetCullMask(layers int, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetCullMask() int {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReflectionProbe) SetUpdateMode(mode ReflectionProbeUpdateMode, )  {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4090221187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReflectionProbe) GetUpdateMode() ReflectionProbeUpdateMode {
  classNameV := StringNameFromStr("ReflectionProbe")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2367550552) // FIXME: should cache?
  var ret ReflectionProbeUpdateMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
