// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ThemeDB struct {
  obj gdc.ObjectPtr
}

func createThemeDB(obj gdc.ObjectPtr) *ThemeDB {
  return &ThemeDB{
    obj: obj,
  }
}

func (me *ThemeDB) BaseClass() string {
  return "ThemeDB"
}
