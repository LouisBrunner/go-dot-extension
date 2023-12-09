// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Sky struct {
  obj gdc.ObjectPtr
}

func (me *Sky) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Sky) BaseClass() string {
  return "Sky"
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

func (me *Sky) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Sky) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Sky) SetRadianceSize(size SkyRadianceSize, )  {
  panic("TODO: implement")
}

func  (me *Sky) GetRadianceSize()  {
  panic("TODO: implement")
}

func  (me *Sky) SetProcessMode(mode SkyProcessMode, )  {
  panic("TODO: implement")
}

func  (me *Sky) GetProcessMode()  {
  panic("TODO: implement")
}

func  (me *Sky) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *Sky) GetMaterial()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
