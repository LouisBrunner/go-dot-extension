// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LightmapGIData struct {
  obj gdc.ObjectPtr
}

func (me *LightmapGIData) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightmapGIData) BaseClass() string {
  return "LightmapGIData"
}



// Enums

func (me *LightmapGIData) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LightmapGIData) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightmapGIData) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *LightmapGIData) SetLightTexture(light_texture TextureLayered, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1278366092) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(light_texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LightmapGIData) GetLightTexture() TextureLayered {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3984243839) // FIXME: should cache?
  var ret TextureLayered
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LightmapGIData) SetUsesSphericalHarmonics(uses_spherical_harmonics bool, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uses_spherical_harmonics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uses_spherical_harmonics), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LightmapGIData) IsUsingSphericalHarmonics() bool {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_spherical_harmonics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LightmapGIData) AddUser(path NodePath, uv_scale Rect2, slice_index int, sub_instance int, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_user")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4272570515) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(uv_scale.AsCTypePtr()), gdc.ConstTypePtr(&slice_index), gdc.ConstTypePtr(&sub_instance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *LightmapGIData) GetUserCount() int {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_user_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LightmapGIData) GetUserPath(user_idx int, ) NodePath {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_user_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&user_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *LightmapGIData) ClearUsers()  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_users")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *LightmapGIData) GetPropLightTexture() TextureLayered {
  panic("TODO: implement")
}

func (me *LightmapGIData) SetPropLightTexture(value TextureLayered) {
  panic("TODO: implement")
}

func (me *LightmapGIData) GetPropLightTextures() Array {
  panic("TODO: implement")
}

func (me *LightmapGIData) SetPropLightTextures(value Array) {
  panic("TODO: implement")
}

func (me *LightmapGIData) GetPropUsesSphericalHarmonics() bool {
  panic("TODO: implement")
}

func (me *LightmapGIData) SetPropUsesSphericalHarmonics(value bool) {
  panic("TODO: implement")
}

func (me *LightmapGIData) GetPropUserData() Array {
  panic("TODO: implement")
}

func (me *LightmapGIData) SetPropUserData(value Array) {
  panic("TODO: implement")
}

func (me *LightmapGIData) GetPropProbeData() Dictionary {
  panic("TODO: implement")
}

func (me *LightmapGIData) SetPropProbeData(value Dictionary) {
  panic("TODO: implement")
}