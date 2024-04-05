// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForMissingResourceList struct {
  fnSetOriginalClass gdc.MethodBindPtr
  fnGetOriginalClass gdc.MethodBindPtr
  fnSetRecordingProperties gdc.MethodBindPtr
  fnIsRecordingProperties gdc.MethodBindPtr
}

var ptrsForMissingResource ptrsForMissingResourceList

func initMissingResourcePtrs(iface gdc.Interface) {

  className := StringNameFromStr("MissingResource")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_original_class")
    defer methodName.Destroy()
    ptrsForMissingResource.fnSetOriginalClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_original_class")
    defer methodName.Destroy()
    ptrsForMissingResource.fnGetOriginalClass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_recording_properties")
    defer methodName.Destroy()
    ptrsForMissingResource.fnSetRecordingProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_recording_properties")
    defer methodName.Destroy()
    ptrsForMissingResource.fnIsRecordingProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type MissingResource struct {
  Resource
}

func (me *MissingResource) BaseClass() string {
  return "MissingResource"
}

func NewMissingResource() *MissingResource {
  str := StringNameFromStr("MissingResource") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MissingResource{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MissingResource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MissingResource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MissingResource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MissingResource) SetOriginalClass(name String, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingResource.fnSetOriginalClass), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MissingResource) GetOriginalClass() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingResource.fnGetOriginalClass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MissingResource) SetRecordingProperties(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingResource.fnSetRecordingProperties), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MissingResource) IsRecordingProperties() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMissingResource.fnIsRecordingProperties), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
