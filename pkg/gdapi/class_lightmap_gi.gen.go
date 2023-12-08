// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightmapGI struct {
  obj gdc.ObjectPtr
}

func (me *LightmapGI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightmapGI) BaseClass() string {
  return "LightmapGI"
}

type LightmapGIBakeQuality int
const (
  LightmapGIBakeQualityLow LightmapGIBakeQuality = 0
  LightmapGIBakeQualityMedium LightmapGIBakeQuality = 1
  LightmapGIBakeQualityHigh LightmapGIBakeQuality = 2
  LightmapGIBakeQualityUltra LightmapGIBakeQuality = 3
)

type LightmapGIGenerateProbes int
const (
  LightmapGIGenerateProbesDisabled LightmapGIGenerateProbes = 0
  LightmapGIGenerateProbesSubdiv4 LightmapGIGenerateProbes = 1
  LightmapGIGenerateProbesSubdiv8 LightmapGIGenerateProbes = 2
  LightmapGIGenerateProbesSubdiv16 LightmapGIGenerateProbes = 3
  LightmapGIGenerateProbesSubdiv32 LightmapGIGenerateProbes = 4
)

type LightmapGIBakeError int
const (
  LightmapGIBakeErrorOk LightmapGIBakeError = 0
  LightmapGIBakeErrorNoSceneRoot LightmapGIBakeError = 1
  LightmapGIBakeErrorForeignData LightmapGIBakeError = 2
  LightmapGIBakeErrorNoLightmapper LightmapGIBakeError = 3
  LightmapGIBakeErrorNoSavePath LightmapGIBakeError = 4
  LightmapGIBakeErrorNoMeshes LightmapGIBakeError = 5
  LightmapGIBakeErrorMeshesInvalid LightmapGIBakeError = 6
  LightmapGIBakeErrorCantCreateImage LightmapGIBakeError = 7
  LightmapGIBakeErrorUserAborted LightmapGIBakeError = 8
  LightmapGIBakeErrorTextureSizeTooSmall LightmapGIBakeError = 9
)

type LightmapGIEnvironmentMode int
const (
  LightmapGIEnvironmentModeDisabled LightmapGIEnvironmentMode = 0
  LightmapGIEnvironmentModeScene LightmapGIEnvironmentMode = 1
  LightmapGIEnvironmentModeCustomSky LightmapGIEnvironmentMode = 2
  LightmapGIEnvironmentModeCustomColor LightmapGIEnvironmentMode = 3
)
