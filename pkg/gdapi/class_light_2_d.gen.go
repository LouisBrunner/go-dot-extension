// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Light2D struct {
  obj gdc.ObjectPtr
}

func (me *Light2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Light2D) BaseClass() string {
  return "Light2D"
}



// Enums

type Light2DShadowFilter int
const (
  Light2DShadowFilterShadowFilterNone Light2DShadowFilter = 0
  Light2DShadowFilterShadowFilterPcf5 Light2DShadowFilter = 1
  Light2DShadowFilterShadowFilterPcf13 Light2DShadowFilter = 2
)

type Light2DBlendMode int
const (
  Light2DBlendModeBlendModeAdd Light2DBlendMode = 0
  Light2DBlendModeBlendModeSub Light2DBlendMode = 1
  Light2DBlendModeBlendModeMix Light2DBlendMode = 2
)

func (me *Light2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Light2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Light2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Light2D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) IsEnabled() bool {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetEditorOnly(editor_only bool, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editor_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) IsEditorOnly() bool {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetColor(color Color, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetColor() Color {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetEnergy(energy float32, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetEnergy() float32 {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetZRangeMin(z int, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_range_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetZRangeMin() int {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_z_range_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetZRangeMax(z int, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_range_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetZRangeMax() int {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_z_range_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetLayerRangeMin(layer int, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_range_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetLayerRangeMin() int {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_range_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetLayerRangeMax(layer int, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer_range_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetLayerRangeMax() int {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_range_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetItemCullMask(item_cull_mask int, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&item_cull_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetItemCullMask() int {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetItemShadowCullMask(item_shadow_cull_mask int, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_shadow_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&item_shadow_cull_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetItemShadowCullMask() int {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_shadow_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetShadowEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) IsShadowEnabled() bool {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shadow_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetShadowSmooth(smooth float32, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_smooth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetShadowSmooth() float32 {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_smooth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetShadowFilter(filter Light2DShadowFilter, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3209356555) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetShadowFilter() Light2DShadowFilter {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1973619177) // FIXME: should cache?
  var ret Light2DShadowFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetShadowColor(shadow_color Color, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shadow_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetShadowColor() Color {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetBlendMode(mode Light2DBlendMode, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2916638796) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetBlendMode() Light2DBlendMode {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 936255250) // FIXME: should cache?
  var ret Light2DBlendMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Light2D) SetHeight(height float32, )  {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Light2D) GetHeight() float32 {
  classNameV := StringNameFromStr("Light2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Light2D) GetPropEnabled() bool {
  panic("TODO: implement")
}

func (me *Light2D) SetPropEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropEditorOnly() bool {
  panic("TODO: implement")
}

func (me *Light2D) SetPropEditorOnly(value bool) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *Light2D) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropEnergy() float32 {
  panic("TODO: implement")
}

func (me *Light2D) SetPropEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropBlendMode() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropBlendMode(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropRangeZMin() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropRangeZMin(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropRangeZMax() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropRangeZMax(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropRangeLayerMin() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropRangeLayerMin(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropRangeLayerMax() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropRangeLayerMax(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropRangeItemCullMask() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropRangeItemCullMask(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropShadowEnabled() bool {
  panic("TODO: implement")
}

func (me *Light2D) SetPropShadowEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropShadowColor() Color {
  panic("TODO: implement")
}

func (me *Light2D) SetPropShadowColor(value Color) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropShadowFilter() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropShadowFilter(value int) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropShadowFilterSmooth() float32 {
  panic("TODO: implement")
}

func (me *Light2D) SetPropShadowFilterSmooth(value float32) {
  panic("TODO: implement")
}

func (me *Light2D) GetPropShadowItemCullMask() int {
  panic("TODO: implement")
}

func (me *Light2D) SetPropShadowItemCullMask(value int) {
  panic("TODO: implement")
}