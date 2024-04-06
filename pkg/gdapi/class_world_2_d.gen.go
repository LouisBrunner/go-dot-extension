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

type ptrsForWorld2DList struct {
	fnGetCanvas           gdc.MethodBindPtr
	fnGetSpace            gdc.MethodBindPtr
	fnGetNavigationMap    gdc.MethodBindPtr
	fnGetDirectSpaceState gdc.MethodBindPtr
}

var ptrsForWorld2D ptrsForWorld2DList

func initWorld2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("World2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_canvas")
		defer methodName.Destroy()
		ptrsForWorld2D.fnGetCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_space")
		defer methodName.Destroy()
		ptrsForWorld2D.fnGetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForWorld2D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_direct_space_state")
		defer methodName.Destroy()
		ptrsForWorld2D.fnGetDirectSpaceState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2506717822))
	}
}

type World2D struct {
	Resource
}

func (me *World2D) BaseClass() string {
	return "World2D"
}

func NewWorld2D() *World2D {
	str := StringNameFromStr("World2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &World2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *World2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *World2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *World2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *World2D) GetCanvas() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld2D.fnGetCanvas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World2D) GetSpace() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld2D.fnGetSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World2D) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld2D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World2D) GetDirectSpaceState() PhysicsDirectSpaceState2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectSpaceState2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld2D.fnGetDirectSpaceState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
