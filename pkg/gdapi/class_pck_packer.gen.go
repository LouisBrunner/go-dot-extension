// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PCKPacker struct {
  RefCounted
}

func (me *PCKPacker) BaseClass() string {
  return "PCKPacker"
}



// Enums

func (me *PCKPacker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PCKPacker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PCKPacker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PCKPacker) PckStart(pck_name String, alignment int, key String, encrypt_directory bool, ) Error {
  classNameV := StringNameFromStr("PCKPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pck_start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 508410629) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pck_name.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(&encrypt_directory), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PCKPacker) AddFile(pck_path String, source_path String, encrypt bool, ) Error {
  classNameV := StringNameFromStr("PCKPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2215643711) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pck_path.AsCTypePtr()), gdc.ConstTypePtr(source_path.AsCTypePtr()), gdc.ConstTypePtr(&encrypt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PCKPacker) Flush(verbose bool, ) Error {
  classNameV := StringNameFromStr("PCKPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flush")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1633102583) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&verbose), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
