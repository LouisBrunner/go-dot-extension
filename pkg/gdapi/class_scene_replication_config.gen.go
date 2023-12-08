// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneReplicationConfig struct {
  obj gdc.ObjectPtr
}

func createSceneReplicationConfig(obj gdc.ObjectPtr) *SceneReplicationConfig {
  return &SceneReplicationConfig{
    obj: obj,
  }
}

func (me *SceneReplicationConfig) BaseClass() string {
  return "SceneReplicationConfig"
}
