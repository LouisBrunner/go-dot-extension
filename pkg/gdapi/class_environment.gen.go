// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Environment struct {
  obj gdc.ObjectPtr
}

func (me *Environment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Environment) BaseClass() string {
  return "Environment"
}

type EnvironmentBGMode int
const (
  EnvironmentBgClearColor EnvironmentBGMode = 0
  EnvironmentBgColor EnvironmentBGMode = 1
  EnvironmentBgSky EnvironmentBGMode = 2
  EnvironmentBgCanvas EnvironmentBGMode = 3
  EnvironmentBgKeep EnvironmentBGMode = 4
  EnvironmentBgCameraFeed EnvironmentBGMode = 5
  EnvironmentBgMax EnvironmentBGMode = 6
)

type EnvironmentAmbientSource int
const (
  EnvironmentAmbientSourceBg EnvironmentAmbientSource = 0
  EnvironmentAmbientSourceDisabled EnvironmentAmbientSource = 1
  EnvironmentAmbientSourceColor EnvironmentAmbientSource = 2
  EnvironmentAmbientSourceSky EnvironmentAmbientSource = 3
)

type EnvironmentReflectionSource int
const (
  EnvironmentReflectionSourceBg EnvironmentReflectionSource = 0
  EnvironmentReflectionSourceDisabled EnvironmentReflectionSource = 1
  EnvironmentReflectionSourceSky EnvironmentReflectionSource = 2
)

type EnvironmentToneMapper int
const (
  EnvironmentToneMapperLinear EnvironmentToneMapper = 0
  EnvironmentToneMapperReinhardt EnvironmentToneMapper = 1
  EnvironmentToneMapperFilmic EnvironmentToneMapper = 2
  EnvironmentToneMapperAces EnvironmentToneMapper = 3
)

type EnvironmentGlowBlendMode int
const (
  EnvironmentGlowBlendModeAdditive EnvironmentGlowBlendMode = 0
  EnvironmentGlowBlendModeScreen EnvironmentGlowBlendMode = 1
  EnvironmentGlowBlendModeSoftlight EnvironmentGlowBlendMode = 2
  EnvironmentGlowBlendModeReplace EnvironmentGlowBlendMode = 3
  EnvironmentGlowBlendModeMix EnvironmentGlowBlendMode = 4
)

type EnvironmentSDFGIYScale int
const (
  EnvironmentSdfgiYScale50Percent EnvironmentSDFGIYScale = 0
  EnvironmentSdfgiYScale75Percent EnvironmentSDFGIYScale = 1
  EnvironmentSdfgiYScale100Percent EnvironmentSDFGIYScale = 2
)
