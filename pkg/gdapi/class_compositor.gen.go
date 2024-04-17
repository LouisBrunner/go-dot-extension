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

type ptrsForCompositorList struct {
	fnSetCompositorEffects gdc.MethodBindPtr
	fnGetCompositorEffects gdc.MethodBindPtr
}

var ptrsForCompositor ptrsForCompositorList

func initCompositorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Compositor")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_compositor_effects")
		defer methodName.Destroy()
		ptrsForCompositor.fnSetCompositorEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_compositor_effects")
		defer methodName.Destroy()
		ptrsForCompositor.fnGetCompositorEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}

}

type Compositor struct {
	Resource
}

func (me *Compositor) BaseClass() string {
	return "Compositor"
}

func NewCompositor() *Compositor {
	str := StringNameFromStr("Compositor") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Compositor{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Compositor) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Compositor) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Compositor) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Compositor) SetCompositorEffects(compositor_effects []CompositorEffect) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compositor_effects)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositor.fnSetCompositorEffects), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Compositor) GetCompositorEffects() []CompositorEffect {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCompositor.fnGetCompositorEffects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[CompositorEffect](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
