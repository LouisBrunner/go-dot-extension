// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("FileSystemDock")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("navigate_to_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileSystemDock) AddResourceTooltipPlugin(plugin EditorResourceTooltipPlugin, )  {
  classNameV := StringNameFromStr("FileSystemDock")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_resource_tooltip_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2258356838) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plugin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *FileSystemDock) RemoveResourceTooltipPlugin(plugin EditorResourceTooltipPlugin, )  {
  classNameV := StringNameFromStr("FileSystemDock")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource_tooltip_plugin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2258356838) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plugin.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API