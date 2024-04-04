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

type Object struct {
  obj gdc.ObjectPtr
}

func (me *Object) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Object) BaseClass() string {
  return "Object"
}

func NewObject() *Object {
  str := StringNameFromStr("Object") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Object{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  ObjectNotificationPostinitialize = "0" // TODO: construct correctly
  ObjectNotificationPredelete = "1" // TODO: construct correctly
)

// Enums

type ObjectConnectFlags int
const (
  ObjectConnectFlagsConnectDeferred ObjectConnectFlags = 1
  ObjectConnectFlagsConnectPersist ObjectConnectFlags = 2
  ObjectConnectFlagsConnectOneShot ObjectConnectFlags = 4
  ObjectConnectFlagsConnectReferenceCounted ObjectConnectFlags = 8
)

func (me *Object) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Object) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Object) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Object) GetClass() String {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) IsClass(class String, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{class.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) Set(property StringName, value Variant, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) Get(property StringName, ) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) SetIndexed(property_path NodePath, value Variant, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indexed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3500910842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property_path.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) GetIndexed(property_path NodePath, ) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_indexed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4006125091) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) GetPropertyList() []Dictionary {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_property_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Object) GetMethodList() []Dictionary {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_method_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Object) PropertyCanRevert(property StringName, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_can_revert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) PropertyGetRevert(property StringName, ) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_revert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) Notification(what int64, reversed bool, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4023243586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what) , gdc.ConstTypePtr(&reversed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) ToString() String {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) GetInstanceId() int64 {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) SetScript(script Variant, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_script")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1114965689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{script.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) GetScript() Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_script")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214101251) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) SetMeta(name StringName, value Variant, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) RemoveMeta(name StringName, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) GetMeta(name StringName, default_ Variant, ) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3990617847) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), default_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) HasMeta(name StringName, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) GetMetaList() []StringName {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_meta_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[StringName](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Object) AddUserSignal(signal String, arguments Array, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_user_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 85656714) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), arguments.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) HasUserSignal(signal StringName, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_user_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) EmitSignal(signal StringName, varargs ...Variant) Error {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("emit_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4047867050) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 1 + len(varargs))
  var0 := signal.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  ret := NewVariant()
  defer ret.Destroy()
  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return Error(-1)
  }
  retInt, err := ret.AsInt()
  if err != nil {
    log.Printf("Error converting return value to int enum: %v", err) // FIXME: bad logging
    return Error(-1)
  }
  return Error(retInt.Get())
}

func  (me *Object) Call(method StringName, varargs ...Variant) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3400424181) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 1 + len(varargs))
  var0 := method.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  ret := NewVariant()
  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return *ret
  }
  return *ret
}

func  (me *Object) CallDeferred(method StringName, varargs ...Variant) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("call_deferred")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3400424181) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 1 + len(varargs))
  var0 := method.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  ret := NewVariant()
  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return *ret
  }
  return *ret
}

func  (me *Object) SetDeferred(property StringName, value Variant, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deferred")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) Callv(method StringName, arg_array Array, ) Variant {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("callv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1260104456) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), arg_array.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) HasMethod(method StringName, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) HasSignal(signal StringName, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_signal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) GetSignalList() []Dictionary {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_signal_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Object) GetSignalConnectionList(signal StringName, ) []Dictionary {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_signal_connection_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3147814860) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Object) GetIncomingConnections() []Dictionary {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_incoming_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Object) Connect(signal StringName, callable Callable, flags int64, ) Error {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1518946055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), callable.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Object) Disconnect(signal StringName, callable Callable, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1874754934) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) IsConnected(signal StringName, callable Callable, ) bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 768136979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{signal.AsCTypePtr(), callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) SetBlockSignals(enable bool, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_block_signals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) IsBlockingSignals() bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_blocking_signals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) NotifyPropertyListChanged()  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("notify_property_list_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) SetMessageTranslation(enable bool, )  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_message_translation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Object) CanTranslateMessages() bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_translate_messages")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) Tr(message StringName, context StringName, ) String {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1195764410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) TrN(message StringName, plural_message StringName, n int64, context StringName, ) String {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tr_n")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 162698058) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), plural_message.AsCTypePtr(), gdc.ConstTypePtr(&n) , context.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&n)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Object) IsQueuedForDeletion() bool {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_queued_for_deletion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Object) CancelFree()  {
  classNameV := StringNameFromStr("Object")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cancel_free")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type ObjectScriptChangedSignalFn func()

func (me *Object) ConnectScriptChanged(subs SignalSubscribers, fn ObjectScriptChangedSignalFn) {
  sig := StringNameFromStr("script_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Object) DisconnectScriptChanged(subs SignalSubscribers, fn ObjectScriptChangedSignalFn) {
  sig := StringNameFromStr("script_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ObjectPropertyListChangedSignalFn func()

func (me *Object) ConnectPropertyListChanged(subs SignalSubscribers, fn ObjectPropertyListChangedSignalFn) {
  sig := StringNameFromStr("property_list_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Object) DisconnectPropertyListChanged(subs SignalSubscribers, fn ObjectPropertyListChangedSignalFn) {
  sig := StringNameFromStr("property_list_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
