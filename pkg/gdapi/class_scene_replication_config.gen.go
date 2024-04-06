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

type ptrsForSceneReplicationConfigList struct {
	fnGetProperties              gdc.MethodBindPtr
	fnAddProperty                gdc.MethodBindPtr
	fnHasProperty                gdc.MethodBindPtr
	fnRemoveProperty             gdc.MethodBindPtr
	fnPropertyGetIndex           gdc.MethodBindPtr
	fnPropertyGetSpawn           gdc.MethodBindPtr
	fnPropertySetSpawn           gdc.MethodBindPtr
	fnPropertyGetReplicationMode gdc.MethodBindPtr
	fnPropertySetReplicationMode gdc.MethodBindPtr
	fnPropertyGetSync            gdc.MethodBindPtr
	fnPropertySetSync            gdc.MethodBindPtr
	fnPropertyGetWatch           gdc.MethodBindPtr
	fnPropertySetWatch           gdc.MethodBindPtr
}

var ptrsForSceneReplicationConfig ptrsForSceneReplicationConfigList

func initSceneReplicationConfigPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SceneReplicationConfig")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_properties")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnGetProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("add_property")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnAddProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4094619021))
	}
	{
		methodName := StringNameFromStr("has_property")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnHasProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 861721659))
	}
	{
		methodName := StringNameFromStr("remove_property")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnRemoveProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("property_get_index")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertyGetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1382022557))
	}
	{
		methodName := StringNameFromStr("property_get_spawn")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertyGetSpawn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3456846888))
	}
	{
		methodName := StringNameFromStr("property_set_spawn")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertySetSpawn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3868023870))
	}
	{
		methodName := StringNameFromStr("property_get_replication_mode")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertyGetReplicationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2870606336))
	}
	{
		methodName := StringNameFromStr("property_set_replication_mode")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertySetReplicationMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3200083865))
	}
	{
		methodName := StringNameFromStr("property_get_sync")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertyGetSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3456846888))
	}
	{
		methodName := StringNameFromStr("property_set_sync")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertySetSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3868023870))
	}
	{
		methodName := StringNameFromStr("property_get_watch")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertyGetWatch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3456846888))
	}
	{
		methodName := StringNameFromStr("property_set_watch")
		defer methodName.Destroy()
		ptrsForSceneReplicationConfig.fnPropertySetWatch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3868023870))
	}
}

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
	SceneReplicationConfigReplicationModeReplicationModeNever    SceneReplicationConfigReplicationMode = 0
	SceneReplicationConfigReplicationModeReplicationModeAlways   SceneReplicationConfigReplicationMode = 1
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

func (me *SceneReplicationConfig) GetProperties() []NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnGetProperties), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[NodePath](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *SceneReplicationConfig) AddProperty(path NodePath, index int64) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnAddProperty), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneReplicationConfig) HasProperty(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnHasProperty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneReplicationConfig) RemoveProperty(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnRemoveProperty), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneReplicationConfig) PropertyGetIndex(path NodePath) int64 {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertyGetIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneReplicationConfig) PropertyGetSpawn(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertyGetSpawn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneReplicationConfig) PropertySetSpawn(path NodePath, enabled bool) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertySetSpawn), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneReplicationConfig) PropertyGetReplicationMode(path NodePath) SceneReplicationConfigReplicationMode {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret SceneReplicationConfigReplicationMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertyGetReplicationMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SceneReplicationConfig) PropertySetReplicationMode(path NodePath, mode SceneReplicationConfigReplicationMode) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertySetReplicationMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneReplicationConfig) PropertyGetSync(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertyGetSync), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneReplicationConfig) PropertySetSync(path NodePath, enabled bool) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertySetSync), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneReplicationConfig) PropertyGetWatch(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertyGetWatch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SceneReplicationConfig) PropertySetWatch(path NodePath, enabled bool) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneReplicationConfig.fnPropertySetWatch), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
