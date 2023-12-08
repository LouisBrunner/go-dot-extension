// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConfigFile struct {
  obj gdc.ObjectPtr
}

func createConfigFile(obj gdc.ObjectPtr) *ConfigFile {
  return &ConfigFile{
    obj: obj,
  }
}

func (me *ConfigFile) BaseClass() string {
  return "ConfigFile"
}
