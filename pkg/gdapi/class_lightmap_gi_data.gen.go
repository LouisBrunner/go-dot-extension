// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type LightmapGIData struct {
  Resource
}

func (me *LightmapGIData) BaseClass() string {
  return "LightmapGIData"
}

func NewLightmapGIData() *LightmapGIData {
  str := StringNameFromStr("LightmapGIData") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &LightmapGIData{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *LightmapGIData) SetLightmapTextures(light_textures []TextureLayered, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lightmap_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_textures) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) GetLightmapTextures() []TextureLayered {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lightmap_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[TextureLayered](ret)
}

func  (me *LightmapGIData) SetUsesSphericalHarmonics(uses_spherical_harmonics bool, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uses_spherical_harmonics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uses_spherical_harmonics) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) IsUsingSphericalHarmonics() bool {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_spherical_harmonics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGIData) AddUser(path NodePath, uv_scale Rect2, slice_index int64, sub_instance int64, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_user")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4272570515) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), uv_scale.AsCTypePtr(), gdc.ConstTypePtr(&slice_index) , gdc.ConstTypePtr(&sub_instance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) GetUserCount() int64 {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_user_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGIData) GetUserPath(user_idx int64, ) NodePath {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_user_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&user_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&user_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LightmapGIData) ClearUsers()  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_users")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) SetLightTexture(light_texture TextureLayered, )  {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1278366092) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{light_texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) GetLightTexture() TextureLayered {
  classNameV := StringNameFromStr("LightmapGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3984243839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTextureLayered()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
