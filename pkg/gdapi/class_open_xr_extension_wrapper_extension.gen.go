// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type OpenXRExtensionWrapperExtension struct {
  Object
}

func (me *OpenXRExtensionWrapperExtension) BaseClass() string {
  return "OpenXRExtensionWrapperExtension"
}

func NewOpenXRExtensionWrapperExtension() *OpenXRExtensionWrapperExtension {
  str := StringNameFromStr("OpenXRExtensionWrapperExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OpenXRExtensionWrapperExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewOpenXRAPIExtension()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OpenXRExtensionWrapperExtension) RegisterExtensionWrapper()  {
  classNameV := StringNameFromStr("OpenXRExtensionWrapperExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_extension_wrapper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
