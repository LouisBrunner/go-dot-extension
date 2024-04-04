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

type Marshalls struct {
  Object
}

func (me *Marshalls) BaseClass() string {
  return "Marshalls"
}

func NewMarshalls() *Marshalls {
  str := StringNameFromStr("Marshalls") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Marshalls{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Marshalls) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Marshalls) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Marshalls) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Marshalls) VariantToBase64(variant Variant, full_objects bool, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("variant_to_base64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3876248563) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{variant.AsCTypePtr(), gdc.ConstTypePtr(&full_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&full_objects)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Marshalls) Base64ToVariant(base64_str String, allow_objects bool, ) Variant {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("base64_to_variant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 218087648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{base64_str.AsCTypePtr(), gdc.ConstTypePtr(&allow_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&allow_objects)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Marshalls) RawToBase64(array PackedByteArray, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("raw_to_base64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3999417757) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Marshalls) Base64ToRaw(base64_str String, ) PackedByteArray {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("base64_to_raw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659035735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{base64_str.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Marshalls) Utf8ToBase64(utf8_str String, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("utf8_to_base64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{utf8_str.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Marshalls) Base64ToUtf8(base64_str String, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("base64_to_utf8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{base64_str.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
