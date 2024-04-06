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

type ptrsForPhysicsServer2DManagerList struct {
	fnRegisterServer   gdc.MethodBindPtr
	fnSetDefaultServer gdc.MethodBindPtr
}

var ptrsForPhysicsServer2DManager ptrsForPhysicsServer2DManagerList

func initPhysicsServer2DManagerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer2DManager")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("register_server")
		defer methodName.Destroy()
		ptrsForPhysicsServer2DManager.fnRegisterServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2137474292))
	}
	{
		methodName := StringNameFromStr("set_default_server")
		defer methodName.Destroy()
		ptrsForPhysicsServer2DManager.fnSetDefaultServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
	}

}

type PhysicsServer2DManager struct {
	Object
}

func (me *PhysicsServer2DManager) BaseClass() string {
	return "PhysicsServer2DManager"
}

func NewPhysicsServer2DManager() *PhysicsServer2DManager {
	str := StringNameFromStr("PhysicsServer2DManager") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer2DManager{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsServer2DManager) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsServer2DManager) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2DManager) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsServer2DManager) RegisterServer(name String, create_callback Callable) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), create_callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2DManager.fnRegisterServer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2DManager) SetDefaultServer(name String, priority int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2DManager.fnSetDefaultServer), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
