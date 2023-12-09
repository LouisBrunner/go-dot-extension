// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *LightmapGIData) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightmapGIData) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *LightmapGIData) SetLightTexture(light_texture TextureLayered, )  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) GetLightTexture()  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) SetUsesSphericalHarmonics(uses_spherical_harmonics bool, )  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) IsUsingSphericalHarmonics()  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) AddUser(path NodePath, uv_scale Rect2, slice_index int, sub_instance int, )  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) GetUserCount()  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) GetUserPath(user_idx int, )  {
  panic("TODO: implement")
}

func  (me *LightmapGIData) ClearUsers()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
