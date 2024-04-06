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

type ptrsForSkeletonModificationStack2DList struct {
	fnSetup                  gdc.MethodBindPtr
	fnExecute                gdc.MethodBindPtr
	fnEnableAllModifications gdc.MethodBindPtr
	fnGetModification        gdc.MethodBindPtr
	fnAddModification        gdc.MethodBindPtr
	fnDeleteModification     gdc.MethodBindPtr
	fnSetModification        gdc.MethodBindPtr
	fnSetModificationCount   gdc.MethodBindPtr
	fnGetModificationCount   gdc.MethodBindPtr
	fnGetIsSetup             gdc.MethodBindPtr
	fnSetEnabled             gdc.MethodBindPtr
	fnGetEnabled             gdc.MethodBindPtr
	fnSetStrength            gdc.MethodBindPtr
	fnGetStrength            gdc.MethodBindPtr
	fnGetSkeleton            gdc.MethodBindPtr
}

var ptrsForSkeletonModificationStack2D ptrsForSkeletonModificationStack2DList

func initSkeletonModificationStack2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonModificationStack2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("setup")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("execute")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnExecute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1005356550))
	}
	{
		methodName := StringNameFromStr("enable_all_modifications")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnEnableAllModifications = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_modification")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnGetModification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2570274329))
	}
	{
		methodName := StringNameFromStr("add_modification")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnAddModification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 354162120))
	}
	{
		methodName := StringNameFromStr("delete_modification")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnDeleteModification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_modification")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnSetModification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1098262544))
	}
	{
		methodName := StringNameFromStr("set_modification_count")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnSetModificationCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_modification_count")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnGetModificationCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_is_setup")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnGetIsSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enabled")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_strength")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnSetStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_strength")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnGetStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_skeleton")
		defer methodName.Destroy()
		ptrsForSkeletonModificationStack2D.fnGetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1697361217))
	}

}

type SkeletonModificationStack2D struct {
	Resource
}

func (me *SkeletonModificationStack2D) BaseClass() string {
	return "SkeletonModificationStack2D"
}

func NewSkeletonModificationStack2D() *SkeletonModificationStack2D {
	str := StringNameFromStr("SkeletonModificationStack2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonModificationStack2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonModificationStack2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonModificationStack2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonModificationStack2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonModificationStack2D) Setup() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnSetup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) Execute(delta float64, execution_mode int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&execution_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnExecute), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) EnableAllModifications(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnEnableAllModifications), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) GetModification(mod_idx int64) SkeletonModification2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkeletonModification2D()
	pinner.Pin(&mod_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnGetModification), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModificationStack2D) AddModification(modification SkeletonModification2D) {
	cargs := []gdc.ConstTypePtr{modification.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnAddModification), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) DeleteModification(mod_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnDeleteModification), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) SetModification(mod_idx int64, modification SkeletonModification2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx), modification.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnSetModification), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) SetModificationCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnSetModificationCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) GetModificationCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnGetModificationCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModificationStack2D) GetIsSetup() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnGetIsSetup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModificationStack2D) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) GetEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModificationStack2D) SetStrength(strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnSetStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModificationStack2D) GetStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnGetStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModificationStack2D) GetSkeleton() Skeleton2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkeleton2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModificationStack2D.fnGetSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
