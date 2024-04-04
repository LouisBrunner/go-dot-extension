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

type ZIPReader struct {
  RefCounted
}

func (me *ZIPReader) BaseClass() string {
  return "ZIPReader"
}

func NewZIPReader() *ZIPReader {
  str := StringNameFromStr("ZIPReader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ZIPReader{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ZIPReader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ZIPReader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ZIPReader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ZIPReader) Open(path String, ) Error {
  classNameV := StringNameFromStr("ZIPReader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ZIPReader) Close() Error {
  classNameV := StringNameFromStr("ZIPReader")
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

func  (me *ZIPReader) GetFiles() PackedStringArray {
  classNameV := StringNameFromStr("ZIPReader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_files")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ZIPReader) ReadFile(path String, case_sensitive bool, ) PackedByteArray {
  classNameV := StringNameFromStr("ZIPReader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("read_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 740857591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&case_sensitive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&case_sensitive)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ZIPReader) FileExists(path String, case_sensitive bool, ) bool {
  classNameV := StringNameFromStr("ZIPReader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("file_exists")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 35364943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&case_sensitive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&case_sensitive)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
