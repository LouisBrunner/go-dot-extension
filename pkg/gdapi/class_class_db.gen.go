// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ClassDB struct {
  obj gdc.ObjectPtr
}

func (me *ClassDB) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ClassDB) BaseClass() string {
  return "ClassDB"
}



// Enums

func (me *ClassDB) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ClassDB) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ClassDB) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ClassDB) GetClassList() PackedStringArray {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_class_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) GetInheritersFromClass(class StringName, ) PackedStringArray {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inheriters_from_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1761182771) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) GetParentClass(class StringName, ) StringName {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parent_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassExists(class StringName, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_exists")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) IsParentClass(class StringName, inherits StringName, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_parent_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(inherits.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) CanInstantiate(class StringName, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_instantiate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) Instantiate(class StringName, ) Variant {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("instantiate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassHasSignal(class StringName, signal StringName, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_has_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(signal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetSignal(class StringName, signal StringName, ) Dictionary {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3061114238) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(signal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetSignalList(class StringName, no_inheritance bool, ) Dictionary {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_signal_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3504980660) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetPropertyList(class StringName, no_inheritance bool, ) Dictionary {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_property_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3504980660) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetProperty(object Object, property StringName, ) Variant {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2498641674) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassSetProperty(object Object, property StringName, value Variant, ) Error {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_set_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1690314931) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassHasMethod(class StringName, method StringName, no_inheritance bool, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_has_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3860701026) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetMethodList(class StringName, no_inheritance bool, ) Dictionary {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_method_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3504980660) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetIntegerConstantList(class StringName, no_inheritance bool, ) PackedStringArray {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_integer_constant_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3031669221) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassHasIntegerConstant(class StringName, name StringName, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_has_integer_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 471820014) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetIntegerConstant(class StringName, name StringName, ) int {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_integer_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2419549490) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassHasEnum(class StringName, name StringName, no_inheritance bool, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_has_enum")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3860701026) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetEnumList(class StringName, no_inheritance bool, ) PackedStringArray {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_enum_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3031669221) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetEnumConstants(class StringName, enum StringName, no_inheritance bool, ) PackedStringArray {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_enum_constants")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 661528303) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(enum.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) ClassGetIntegerConstantEnum(class StringName, name StringName, no_inheritance bool, ) StringName {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("class_get_integer_constant_enum")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2457504236) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&no_inheritance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ClassDB) IsClassEnabled(class StringName, ) bool {
  classNameV := StringNameFromStr("ClassDB")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_class_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
