// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SceneReplicationConfig struct {
  obj gdc.ObjectPtr
}

func (me *SceneReplicationConfig) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneReplicationConfig) BaseClass() string {
  return "SceneReplicationConfig"
}



// Enums

type SceneReplicationConfigReplicationMode int
const (
  SceneReplicationConfigReplicationModeReplicationModeNever SceneReplicationConfigReplicationMode = 0
  SceneReplicationConfigReplicationModeReplicationModeAlways SceneReplicationConfigReplicationMode = 1
  SceneReplicationConfigReplicationModeReplicationModeOnChange SceneReplicationConfigReplicationMode = 2
)

func (me *SceneReplicationConfig) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SceneReplicationConfig) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneReplicationConfig) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SceneReplicationConfig) GetProperties() NodePath {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) AddProperty(path NodePath, index int, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4094619021) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneReplicationConfig) HasProperty(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 861721659) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) RemoveProperty(path NodePath, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneReplicationConfig) PropertyGetIndex(path NodePath, ) int {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1382022557) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) PropertyGetSpawn(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_spawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3456846888) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) PropertySetSpawn(path NodePath, enabled bool, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_spawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneReplicationConfig) PropertyGetReplicationMode(path NodePath, ) SceneReplicationConfigReplicationMode {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_replication_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2870606336) // FIXME: should cache?
  var ret SceneReplicationConfigReplicationMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) PropertySetReplicationMode(path NodePath, mode SceneReplicationConfigReplicationMode, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_replication_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200083865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneReplicationConfig) PropertyGetSync(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3456846888) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) PropertySetSync(path NodePath, enabled bool, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneReplicationConfig) PropertyGetWatch(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_watch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3456846888) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SceneReplicationConfig) PropertySetWatch(path NodePath, enabled bool, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_watch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
