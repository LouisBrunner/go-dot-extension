// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Marshalls struct {
  obj gdc.ObjectPtr
}

func (me *Marshalls) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Marshalls) BaseClass() string {
  return "Marshalls"
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
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(variant.AsCTypePtr()), gdc.ConstTypePtr(&full_objects), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Marshalls) Base64ToVariant(base64_str String, allow_objects bool, ) Variant {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("base64_to_variant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 218087648) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base64_str.AsCTypePtr()), gdc.ConstTypePtr(&allow_objects), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Marshalls) RawToBase64(array PackedByteArray, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("raw_to_base64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3999417757) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(array.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Marshalls) Base64ToRaw(base64_str String, ) PackedByteArray {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("base64_to_raw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659035735) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base64_str.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Marshalls) Utf8ToBase64(utf8_str String, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("utf8_to_base64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(utf8_str.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Marshalls) Base64ToUtf8(base64_str String, ) String {
  classNameV := StringNameFromStr("Marshalls")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("base64_to_utf8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1703090593) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base64_str.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
