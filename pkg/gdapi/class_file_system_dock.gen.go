// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FileSystemDock struct {
  obj gdc.ObjectPtr
}

func (me *FileSystemDock) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FileSystemDock) BaseClass() string {
  return "FileSystemDock"
}



// Enums

func (me *FileSystemDock) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *FileSystemDock) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *FileSystemDock) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *FileSystemDock) NavigateToPath(path String, )  {
  panic("TODO: implement")
}

func  (me *FileSystemDock) AddResourceTooltipPlugin(plugin EditorResourceTooltipPlugin, )  {
  panic("TODO: implement")
}

func  (me *FileSystemDock) RemoveResourceTooltipPlugin(plugin EditorResourceTooltipPlugin, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
