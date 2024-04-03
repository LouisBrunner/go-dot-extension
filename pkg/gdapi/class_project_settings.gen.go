// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ProjectSettings struct {
  Object
}

func (me *ProjectSettings) BaseClass() string {
  return "ProjectSettings"
}

func NewProjectSettings() *ProjectSettings {
  str := StringNameFromStr("ProjectSettings") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ProjectSettings{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ProjectSettings) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ProjectSettings) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ProjectSettings) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ProjectSettings) HasSetting(name String, ) bool {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProjectSettings) SetSetting(name String, value Variant, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402577236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) GetSetting(name String, default_value Variant, ) Variant {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_setting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 223050753) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(default_value.AsCTypePtr()), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProjectSettings) GetSettingWithOverride(name StringName, ) Variant {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_setting_with_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProjectSettings) GetGlobalClassList() []Dictionary {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_class_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Dictionary](ret)
}

func  (me *ProjectSettings) SetOrder(name String, position int64, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&position), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) GetOrder(name String, ) int64 {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_order")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProjectSettings) SetInitialValue(name String, value Variant, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_initial_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402577236) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) SetAsBasic(name String, basic bool, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_basic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&basic), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) SetAsInternal(name String, internal bool, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_internal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&internal), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) AddPropertyInfo(hint Dictionary, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_property_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(hint.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) SetRestartIfChanged(name String, restart bool, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_restart_if_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2678287736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&restart), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) Clear(name String, )  {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ProjectSettings) LocalizePath(path String, ) String {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("localize_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProjectSettings) GlobalizePath(path String, ) String {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("globalize_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ProjectSettings) Save() Error {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ProjectSettings) LoadResourcePack(pack String, replace_files bool, offset int64, ) bool {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_resource_pack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 708980503) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pack.AsCTypePtr()), gdc.ConstTypePtr(&replace_files), gdc.ConstTypePtr(&offset), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ProjectSettings) SaveCustom(file String, ) Error {
  classNameV := StringNameFromStr("ProjectSettings")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_custom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals

type ProjectSettingsSettingsChangedSignalFn func()

func (me *ProjectSettings) ConnectSettingsChanged(subs SignalSubscribers, fn ProjectSettingsSettingsChangedSignalFn) {
  sig := StringNameFromStr("settings_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ProjectSettings) DisconnectSettingsChanged(subs SignalSubscribers, fn ProjectSettingsSettingsChangedSignalFn) {
  sig := StringNameFromStr("settings_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
