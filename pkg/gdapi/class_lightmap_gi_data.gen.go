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

type ptrsForLightmapGIDataList struct {
  fnSetLightmapTextures gdc.MethodBindPtr
  fnGetLightmapTextures gdc.MethodBindPtr
  fnSetUsesSphericalHarmonics gdc.MethodBindPtr
  fnIsUsingSphericalHarmonics gdc.MethodBindPtr
  fnAddUser gdc.MethodBindPtr
  fnGetUserCount gdc.MethodBindPtr
  fnGetUserPath gdc.MethodBindPtr
  fnClearUsers gdc.MethodBindPtr
  fnSetLightTexture gdc.MethodBindPtr
  fnGetLightTexture gdc.MethodBindPtr
}

var ptrsForLightmapGIData ptrsForLightmapGIDataList

func initLightmapGIDataPtrs(iface gdc.Interface) {

  className := StringNameFromStr("LightmapGIData")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_lightmap_textures")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnSetLightmapTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_lightmap_textures")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnGetLightmapTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_uses_spherical_harmonics")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnSetUsesSphericalHarmonics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_spherical_harmonics")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnIsUsingSphericalHarmonics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("add_user")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnAddUser = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4272570515))
  }
  {
    methodName := StringNameFromStr("get_user_count")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnGetUserCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_user_path")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnGetUserPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
  }
  {
    methodName := StringNameFromStr("clear_users")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnClearUsers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_light_texture")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnSetLightTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1278366092))
  }
  {
    methodName := StringNameFromStr("get_light_texture")
    defer methodName.Destroy()
    ptrsForLightmapGIData.fnGetLightTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3984243839))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_textures) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnSetLightmapTextures), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) GetLightmapTextures() []TextureLayered {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnGetLightmapTextures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[TextureLayered](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *LightmapGIData) SetUsesSphericalHarmonics(uses_spherical_harmonics bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uses_spherical_harmonics) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnSetUsesSphericalHarmonics), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) IsUsingSphericalHarmonics() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnIsUsingSphericalHarmonics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGIData) AddUser(path NodePath, uv_scale Rect2, slice_index int64, sub_instance int64, )  {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), uv_scale.AsCTypePtr(), gdc.ConstTypePtr(&slice_index) , gdc.ConstTypePtr(&sub_instance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnAddUser), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) GetUserCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnGetUserCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LightmapGIData) GetUserPath(user_idx int64, ) NodePath {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&user_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&user_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnGetUserPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LightmapGIData) ClearUsers()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnClearUsers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) SetLightTexture(light_texture TextureLayered, )  {
  cargs := []gdc.ConstTypePtr{light_texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnSetLightTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LightmapGIData) GetLightTexture() TextureLayered {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTextureLayered()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightmapGIData.fnGetLightTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
