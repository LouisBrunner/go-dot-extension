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

type ptrsForVisibleOnScreenEnabler3DList struct {
	fnSetEnableMode     gdc.MethodBindPtr
	fnGetEnableMode     gdc.MethodBindPtr
	fnSetEnableNodePath gdc.MethodBindPtr
	fnGetEnableNodePath gdc.MethodBindPtr
}

var ptrsForVisibleOnScreenEnabler3D ptrsForVisibleOnScreenEnabler3DList

func initVisibleOnScreenEnabler3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisibleOnScreenEnabler3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_enable_mode")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler3D.fnSetEnableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 320303646))
	}
	{
		methodName := StringNameFromStr("get_enable_mode")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler3D.fnGetEnableMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3352990031))
	}
	{
		methodName := StringNameFromStr("set_enable_node_path")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler3D.fnSetEnableNodePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_enable_node_path")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenEnabler3D.fnGetEnableNodePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 277076166))
	}

}

type VisibleOnScreenEnabler3D struct {
	VisibleOnScreenNotifier3D
}

func (me *VisibleOnScreenEnabler3D) BaseClass() string {
	return "VisibleOnScreenEnabler3D"
}

func NewVisibleOnScreenEnabler3D() *VisibleOnScreenEnabler3D {
	str := StringNameFromStr("VisibleOnScreenEnabler3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisibleOnScreenEnabler3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisibleOnScreenEnabler3DEnableMode int

const (
	VisibleOnScreenEnabler3DEnableModeEnableModeInherit    VisibleOnScreenEnabler3DEnableMode = 0
	VisibleOnScreenEnabler3DEnableModeEnableModeAlways     VisibleOnScreenEnabler3DEnableMode = 1
	VisibleOnScreenEnabler3DEnableModeEnableModeWhenPaused VisibleOnScreenEnabler3DEnableMode = 2
)

func (me *VisibleOnScreenEnabler3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisibleOnScreenEnabler3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenEnabler3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisibleOnScreenEnabler3D) SetEnableMode(mode VisibleOnScreenEnabler3DEnableMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler3D.fnSetEnableMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisibleOnScreenEnabler3D) GetEnableMode() VisibleOnScreenEnabler3DEnableMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisibleOnScreenEnabler3DEnableMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler3D.fnGetEnableMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisibleOnScreenEnabler3D) SetEnableNodePath(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler3D.fnSetEnableNodePath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisibleOnScreenEnabler3D) GetEnableNodePath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenEnabler3D.fnGetEnableNodePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
