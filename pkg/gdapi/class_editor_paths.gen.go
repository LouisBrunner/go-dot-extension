// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorPaths struct {
  obj gdc.ObjectPtr
}

func (me *EditorPaths) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorPaths) BaseClass() string {
  return "EditorPaths"
}



// Enums

func (me *EditorPaths) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorPaths) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorPaths) GetDataDir()  {
  panic("TODO: implement")
}

func  (me *EditorPaths) GetConfigDir()  {
  panic("TODO: implement")
}

func  (me *EditorPaths) GetCacheDir()  {
  panic("TODO: implement")
}

func  (me *EditorPaths) IsSelfContained()  {
  panic("TODO: implement")
}

func  (me *EditorPaths) GetSelfContainedFile()  {
  panic("TODO: implement")
}

func  (me *EditorPaths) GetProjectSettingsDir()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
