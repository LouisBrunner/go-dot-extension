// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Sky struct {
  Resource
}

func (me *Sky) BaseClass() string {
  return "Sky"
}

func NewSky() *Sky {
  str := StringNameFromStr("Sky") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Sky{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type SkyRadianceSize int
const (
  SkyRadianceSizeRadianceSize32 SkyRadianceSize = 0
  SkyRadianceSizeRadianceSize64 SkyRadianceSize = 1
  SkyRadianceSizeRadianceSize128 SkyRadianceSize = 2
  SkyRadianceSizeRadianceSize256 SkyRadianceSize = 3
  SkyRadianceSizeRadianceSize512 SkyRadianceSize = 4
  SkyRadianceSizeRadianceSize1024 SkyRadianceSize = 5
  SkyRadianceSizeRadianceSize2048 SkyRadianceSize = 6
  SkyRadianceSizeRadianceSizeMax SkyRadianceSize = 7
)

type SkyProcessMode int
const (
  SkyProcessModeProcessModeAutomatic SkyProcessMode = 0
  SkyProcessModeProcessModeQuality SkyProcessMode = 1
  SkyProcessModeProcessModeIncremental SkyProcessMode = 2
  SkyProcessModeProcessModeRealtime SkyProcessMode = 3
)

func (me *Sky) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Sky) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Sky) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Sky) SetRadianceSize(size SkyRadianceSize, )  {
  classNameV := StringNameFromStr("Sky")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radiance_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1512957179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sky) GetRadianceSize() SkyRadianceSize {
  classNameV := StringNameFromStr("Sky")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radiance_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2708733976) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret SkyRadianceSize

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Sky) SetProcessMode(mode SkyProcessMode, )  {
  classNameV := StringNameFromStr("Sky")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 875986769) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sky) GetProcessMode() SkyProcessMode {
  classNameV := StringNameFromStr("Sky")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 731245043) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret SkyProcessMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Sky) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("Sky")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sky) GetMaterial() Material {
  classNameV := StringNameFromStr("Sky")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
