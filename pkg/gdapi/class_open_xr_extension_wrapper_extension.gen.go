// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRExtensionWrapperExtension struct {
  Object
}

func (me *OpenXRExtensionWrapperExtension) BaseClass() string {
  return "OpenXRExtensionWrapperExtension"
}



// Enums

func (me *OpenXRExtensionWrapperExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRExtensionWrapperExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRExtensionWrapperExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRExtensionWrapperExtension) GetOpenxrApi() OpenXRAPIExtension {
  classNameV := StringNameFromStr("OpenXRExtensionWrapperExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_openxr_api")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1637791613) // FIXME: should cache?
  var ret OpenXRAPIExtension
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRExtensionWrapperExtension) RegisterExtensionWrapper()  {
  classNameV := StringNameFromStr("OpenXRExtensionWrapperExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_extension_wrapper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
