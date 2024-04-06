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

type ptrsForVisibleOnScreenEnabler2DList struct {
	fnSetEnableMode     gdc.MethodBindPtr
	fnGetEnableMode     gdc.MethodBindPtr
	fnSetEnableNodePath gdc.MethodBindPtr
	fnGetEnableNodePath gdc.MethodBindPtr
}

var ptrsForVisibleOnScreenEnabler2D ptrsForVisibleOnScreenEnabler2DList

func initVisibleOnScreenEnabler2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisibleOnScreenEnabler2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enable_mode")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler2D.fnSetEnableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961788752))
	}
	{
		methodName := StringNameFromStr("get_enable_mode")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler2D.fnGetEnableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2650445576))
	}
	{
		methodName := StringNameFromStr("set_enable_node_path")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler2D.fnSetEnableNodePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_enable_node_path")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler2D.fnGetEnableNodePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 277076166))
	}

}

type VisibleOnScreenEnabler2D struct {
	VisibleOnScreenNotifier2D
}

func (me *VisibleOnScreenEnabler2D) BaseClass() string {
	return "VisibleOnScreenEnabler2D"
}

func NewVisibleOnScreenEnabler2D() *VisibleOnScreenEnabler2D {
	str := StringNameFromStr("VisibleOnScreenEnabler2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisibleOnScreenEnabler2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisibleOnScreenEnabler2DEnableMode int

const (
	VisibleOnScreenEnabler2DEnableModeEnableModeInherit    VisibleOnScreenEnabler2DEnableMode = 0
	VisibleOnScreenEnabler2DEnableModeEnableModeAlways     VisibleOnScreenEnabler2DEnableMode = 1
	VisibleOnScreenEnabler2DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler2DEnableMode = 2
)

func (me *VisibleOnScreenEnabler2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisibleOnScreenEnabler2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenEnabler2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisibleOnScreenEnabler2D) SetEnableMode(mode VisibleOnScreenEnabler2DEnableMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler2D.fnSetEnableMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisibleOnScreenEnabler2D) GetEnableMode() VisibleOnScreenEnabler2DEnableMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisibleOnScreenEnabler2DEnableMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler2D.fnGetEnableMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisibleOnScreenEnabler2D) SetEnableNodePath(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler2D.fnSetEnableNodePath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisibleOnScreenEnabler2D) GetEnableNodePath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler2D.fnGetEnableNodePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
