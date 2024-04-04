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

type ZIPPacker struct {
  RefCounted
}

func (me *ZIPPacker) BaseClass() string {
  return "ZIPPacker"
}

func NewZIPPacker() *ZIPPacker {
  str := StringNameFromStr("ZIPPacker") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ZIPPacker{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&append) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&append)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPPacker) StartFile(path String, ) Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPPacker) WriteFile(data PackedByteArray, ) Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("write_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPPacker) CloseFile() Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPPacker) Close() Error {
  classNameV := StringNameFromStr("ZIPPacker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
