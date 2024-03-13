// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ZIPPacker struct {
  obj gdc.ObjectPtr
}

func (me *ZIPPacker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ZIPPacker) BaseClass() string {
  return "ZIPPacker"
}



// Enums

type ZIPPackerZipAppend int
const (
  ZIPPackerZipAppendAppendCreate ZIPPackerZipAppend = 0
  ZIPPackerZipAppendAppendCreateafter ZIPPackerZipAppend = 1
  ZIPPackerZipAppendAppendAddinzip ZIPPackerZipAppend = 2
)

func (me *ZIPPacker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ZIPPacker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ZIPPacker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ZIPPacker) Open(path String, append ZIPPackerZipAppend, ) Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1936816515) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&append), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ZIPPacker) StartFile(path String, ) Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ZIPPacker) WriteFile(data PackedByteArray, ) Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("write_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ZIPPacker) CloseFile() Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ZIPPacker) Close() Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
