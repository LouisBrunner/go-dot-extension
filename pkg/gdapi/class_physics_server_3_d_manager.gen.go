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

type ptrsForPhysicsServer3DManagerList struct {
	fnRegisterServer   gdc.MethodBindPtr
	fnSetDefaultServer gdc.MethodBindPtr
}

var ptrsForPhysicsServer3DManager ptrsForPhysicsServer3DManagerList

func initPhysicsServer3DManagerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer3DManager")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("register_server")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DManager.fnRegisterServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2137474292))
	}
	{
		methodName := StringNameFromStr("set_default_server")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DManager.fnSetDefaultServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
	}
}

type PhysicsServer3DManager struct {
	Object
}

func (me *PhysicsServer3DManager) BaseClass() string {
	return "PhysicsServer3DManager"
}

func NewPhysicsServer3DManager() *PhysicsServer3DManager {
	str := StringNameFromStr("PhysicsServer3DManager") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer3DManager{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsServer3DManager) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsServer3DManager) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3DManager) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsServer3DManager) RegisterServer(name String, create_callback Callable) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), create_callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DManager.fnRegisterServer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3DManager) SetDefaultServer(name String, priority int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DManager.fnSetDefaultServer), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
