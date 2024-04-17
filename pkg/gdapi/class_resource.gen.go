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

type ptrsForResourceList struct {
	fnXSetupLocalToScene    gdc.MethodBindPtr
	fnSetPath               gdc.MethodBindPtr
	fnTakeOverPath          gdc.MethodBindPtr
	fnGetPath               gdc.MethodBindPtr
	fnSetName               gdc.MethodBindPtr
	fnGetName               gdc.MethodBindPtr
	fnGetRid                gdc.MethodBindPtr
	fnSetLocalToScene       gdc.MethodBindPtr
	fnIsLocalToScene        gdc.MethodBindPtr
	fnGetLocalScene         gdc.MethodBindPtr
	fnSetupLocalToScene     gdc.MethodBindPtr
	fnGenerateSceneUniqueId gdc.MethodBindPtr
	fnSetSceneUniqueId      gdc.MethodBindPtr
	fnGetSceneUniqueId      gdc.MethodBindPtr
	fnEmitChanged           gdc.MethodBindPtr
	fnDuplicate             gdc.MethodBindPtr
}

var ptrsForResource ptrsForResourceList

func initResourcePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Resource")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_path")
		defer methodName.Destroy()
		ptrsForResource.fnSetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("take_over_path")
		defer methodName.Destroy()
		ptrsForResource.fnTakeOverPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_path")
		defer methodName.Destroy()
		ptrsForResource.fnGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_name")
		defer methodName.Destroy()
		ptrsForResource.fnSetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForResource.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForResource.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_local_to_scene")
		defer methodName.Destroy()
		ptrsForResource.fnSetLocalToScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_local_to_scene")
		defer methodName.Destroy()
		ptrsForResource.fnIsLocalToScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_local_scene")
		defer methodName.Destroy()
		ptrsForResource.fnGetLocalScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
	}
	{
		methodName := StringNameFromStr("setup_local_to_scene")
		defer methodName.Destroy()
		ptrsForResource.fnSetupLocalToScene = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("generate_scene_unique_id")
		defer methodName.Destroy()
		ptrsForResource.fnGenerateSceneUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("set_scene_unique_id")
		defer methodName.Destroy()
		ptrsForResource.fnSetSceneUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_scene_unique_id")
		defer methodName.Destroy()
		ptrsForResource.fnGetSceneUniqueId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("emit_changed")
		defer methodName.Destroy()
		ptrsForResource.fnEmitChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForResource.fnDuplicate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 482882304))
	}

}

type Resource struct {
	RefCounted
}

func (me *Resource) BaseClass() string {
	return "Resource"
}

func NewResource() *Resource {
	str := StringNameFromStr("Resource") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Resource{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Resource) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Resource) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Resource) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Resource) SetPath(path String) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnSetPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Resource) TakeOverPath(path String) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnTakeOverPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Resource) GetPath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Resource) SetName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnSetName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Resource) GetName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Resource) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Resource) SetLocalToScene(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnSetLocalToScene), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Resource) IsLocalToScene() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnIsLocalToScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Resource) GetLocalScene() Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnGetLocalScene), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Resource) SetupLocalToScene() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnSetupLocalToScene), me.obj, unsafe.SliceData(cargs), nil)

}

func ResourceGenerateSceneUniqueId() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnGenerateSceneUniqueId), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Resource) SetSceneUniqueId(id String) {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnSetSceneUniqueId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Resource) GetSceneUniqueId() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnGetSceneUniqueId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Resource) EmitChanged() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnEmitChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Resource) Duplicate(subresources bool) Resource {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subresources)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()
	pinner.Pin(&subresources)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResource.fnDuplicate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ResourceChangedSignalFn func()

func (me *Resource) ConnectChanged(subs SignalSubscribers, fn ResourceChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Resource) DisconnectChanged(subs SignalSubscribers, fn ResourceChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ResourceSetupLocalToSceneRequestedSignalFn func()

func (me *Resource) ConnectSetupLocalToSceneRequested(subs SignalSubscribers, fn ResourceSetupLocalToSceneRequestedSignalFn) {
	sig := StringNameFromStr("setup_local_to_scene_requested")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Resource) DisconnectSetupLocalToSceneRequested(subs SignalSubscribers, fn ResourceSetupLocalToSceneRequestedSignalFn) {
	sig := StringNameFromStr("setup_local_to_scene_requested")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
