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

type ptrsForSkinList struct {
	fnSetBindCount gdc.MethodBindPtr
	fnGetBindCount gdc.MethodBindPtr
	fnAddBind      gdc.MethodBindPtr
	fnAddNamedBind gdc.MethodBindPtr
	fnSetBindPose  gdc.MethodBindPtr
	fnGetBindPose  gdc.MethodBindPtr
	fnSetBindName  gdc.MethodBindPtr
	fnGetBindName  gdc.MethodBindPtr
	fnSetBindBone  gdc.MethodBindPtr
	fnGetBindBone  gdc.MethodBindPtr
	fnClearBinds   gdc.MethodBindPtr
}

var ptrsForSkin ptrsForSkinList

func initSkinPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Skin")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_bind_count")
		defer methodName.Destroy()
		ptrsForSkin.fnSetBindCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bind_count")
		defer methodName.Destroy()
		ptrsForSkin.fnGetBindCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_bind")
		defer methodName.Destroy()
		ptrsForSkin.fnAddBind = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("add_named_bind")
		defer methodName.Destroy()
		ptrsForSkin.fnAddNamedBind = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3154712474))
	}
	{
		methodName := StringNameFromStr("set_bind_pose")
		defer methodName.Destroy()
		ptrsForSkin.fnSetBindPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("get_bind_pose")
		defer methodName.Destroy()
		ptrsForSkin.fnGetBindPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("set_bind_name")
		defer methodName.Destroy()
		ptrsForSkin.fnSetBindName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
	{
		methodName := StringNameFromStr("get_bind_name")
		defer methodName.Destroy()
		ptrsForSkin.fnGetBindName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_bind_bone")
		defer methodName.Destroy()
		ptrsForSkin.fnSetBindBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_bind_bone")
		defer methodName.Destroy()
		ptrsForSkin.fnGetBindBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("clear_binds")
		defer methodName.Destroy()
		ptrsForSkin.fnClearBinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type Skin struct {
	Resource
}

func (me *Skin) BaseClass() string {
	return "Skin"
}

func NewSkin() *Skin {
	str := StringNameFromStr("Skin") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Skin{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Skin) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Skin) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Skin) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Skin) SetBindCount(bind_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnSetBindCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skin) GetBindCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnGetBindCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skin) AddBind(bone int64, pose Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone), pose.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnAddBind), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skin) AddNamedBind(name String, pose Transform3D) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), pose.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnAddNamedBind), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skin) SetBindPose(bind_index int64, pose Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), pose.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnSetBindPose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skin) GetBindPose(bind_index int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bind_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnGetBindPose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skin) SetBindName(bind_index int64, name StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnSetBindName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skin) GetBindName(bind_index int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&bind_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnGetBindName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skin) SetBindBone(bind_index int64, bone int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), gdc.ConstTypePtr(&bone)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnSetBindBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skin) GetBindBone(bind_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&bind_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnGetBindBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skin) ClearBinds() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkin.fnClearBinds), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
