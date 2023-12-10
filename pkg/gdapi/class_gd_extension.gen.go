// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GDExtension struct {
  obj gdc.ObjectPtr
}

func (me *GDExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDExtension) BaseClass() string {
  return "GDExtension"
}



// Enums

type GDExtensionInitializationLevel int
const (
  GDExtensionInitializationLevelInitializationLevelCore GDExtensionInitializationLevel = 0
  GDExtensionInitializationLevelInitializationLevelServers GDExtensionInitializationLevel = 1
  GDExtensionInitializationLevelInitializationLevelScene GDExtensionInitializationLevel = 2
  GDExtensionInitializationLevelInitializationLevelEditor GDExtensionInitializationLevel = 3
)

func (me *GDExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GDExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GDExtension) OpenLibrary(path String, entry_symbol String, ) Error {
  classNameV := StringNameFromStr("GDExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(entry_symbol.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GDExtension) CloseLibrary()  {
  classNameV := StringNameFromStr("GDExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GDExtension) IsLibraryOpen() bool {
  classNameV := StringNameFromStr("GDExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_library_open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GDExtension) GetMinimumLibraryInitializationLevel() GDExtensionInitializationLevel {
  classNameV := StringNameFromStr("GDExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_minimum_library_initialization_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 964858755) // FIXME: should cache?
  var ret GDExtensionInitializationLevel
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GDExtension) InitializeLibrary(level GDExtensionInitializationLevel, )  {
  classNameV := StringNameFromStr("GDExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("initialize_library")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3409922941) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&level), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties