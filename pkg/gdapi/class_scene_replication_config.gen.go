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

type SceneReplicationConfig struct {
  Resource
}

func (me *SceneReplicationConfig) BaseClass() string {
  return "SceneReplicationConfig"
}

func NewSceneReplicationConfig() *SceneReplicationConfig {
  str := StringNameFromStr("SceneReplicationConfig") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SceneReplicationConfig{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *SceneReplicationConfig) GetProperties() []NodePath {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[NodePath](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *SceneReplicationConfig) AddProperty(path NodePath, index int64, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4094619021) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneReplicationConfig) HasProperty(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 861721659) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneReplicationConfig) RemoveProperty(path NodePath, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneReplicationConfig) PropertyGetIndex(path NodePath, ) int64 {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1382022557) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneReplicationConfig) PropertyGetSpawn(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_spawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3456846888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneReplicationConfig) PropertySetSpawn(path NodePath, enabled bool, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_spawn")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneReplicationConfig) PropertyGetReplicationMode(path NodePath, ) SceneReplicationConfigReplicationMode {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_replication_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2870606336) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SceneReplicationConfigReplicationMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SceneReplicationConfig) PropertySetReplicationMode(path NodePath, mode SceneReplicationConfigReplicationMode, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_replication_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200083865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneReplicationConfig) PropertyGetSync(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3456846888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneReplicationConfig) PropertySetSync(path NodePath, enabled bool, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneReplicationConfig) PropertyGetWatch(path NodePath, ) bool {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_get_watch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3456846888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SceneReplicationConfig) PropertySetWatch(path NodePath, enabled bool, )  {
  classNameV := StringNameFromStr("SceneReplicationConfig")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("property_set_watch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3868023870) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
