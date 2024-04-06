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

type ptrsForSkeletonModification2DStackHolderList struct {
	fnSetHeldModificationStack gdc.MethodBindPtr
	fnGetHeldModificationStack gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DStackHolder ptrsForSkeletonModification2DStackHolderList

func initSkeletonModification2DStackHolderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonModification2DStackHolder")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_held_modification_stack")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DStackHolder.fnSetHeldModificationStack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3907307132))
	}
	{
		methodName := StringNameFromStr("get_held_modification_stack")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DStackHolder.fnGetHeldModificationStack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2107508396))
	}

}

type SkeletonModification2DStackHolder struct {
	SkeletonModification2D
}

func (me *SkeletonModification2DStackHolder) BaseClass() string {
	return "SkeletonModification2DStackHolder"
}

func NewSkeletonModification2DStackHolder() *SkeletonModification2DStackHolder {
	str := StringNameFromStr("SkeletonModification2DStackHolder") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonModification2DStackHolder{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonModification2DStackHolder) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonModification2DStackHolder) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DStackHolder) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonModification2DStackHolder) SetHeldModificationStack(held_modification_stack SkeletonModificationStack2D) {
	cargs := []gdc.ConstTypePtr{held_modification_stack.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DStackHolder.fnSetHeldModificationStack), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DStackHolder) GetHeldModificationStack() SkeletonModificationStack2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkeletonModificationStack2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DStackHolder.fnGetHeldModificationStack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
