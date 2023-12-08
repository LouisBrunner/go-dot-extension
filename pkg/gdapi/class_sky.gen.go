// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type SkyRadianceSize int
const (
  SkyRadianceSize32 SkyRadianceSize = 0
  SkyRadianceSize64 SkyRadianceSize = 1
  SkyRadianceSize128 SkyRadianceSize = 2
  SkyRadianceSize256 SkyRadianceSize = 3
  SkyRadianceSize512 SkyRadianceSize = 4
  SkyRadianceSize1024 SkyRadianceSize = 5
  SkyRadianceSize2048 SkyRadianceSize = 6
  SkyRadianceSizeMax SkyRadianceSize = 7
)

type SkyProcessMode int
const (
  SkyProcessModeAutomatic SkyProcessMode = 0
  SkyProcessModeQuality SkyProcessMode = 1
  SkyProcessModeIncremental SkyProcessMode = 2
  SkyProcessModeRealtime SkyProcessMode = 3
)
