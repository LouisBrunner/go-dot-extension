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

type ptrsForWorld3DList struct {
	fnGetSpace               gdc.MethodBindPtr
	fnGetNavigationMap       gdc.MethodBindPtr
	fnGetScenario            gdc.MethodBindPtr
	fnSetEnvironment         gdc.MethodBindPtr
	fnGetEnvironment         gdc.MethodBindPtr
	fnSetFallbackEnvironment gdc.MethodBindPtr
	fnGetFallbackEnvironment gdc.MethodBindPtr
	fnSetCameraAttributes    gdc.MethodBindPtr
	fnGetCameraAttributes    gdc.MethodBindPtr
	fnGetDirectSpaceState    gdc.MethodBindPtr
}

var ptrsForWorld3D ptrsForWorld3DList

func initWorld3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("World3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_space")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_scenario")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetScenario = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_environment")
		defer methodName.Destroy()
		ptrsForWorld3D.fnSetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4143518816))
	}
	{
		methodName := StringNameFromStr("get_environment")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3082064660))
	}
	{
		methodName := StringNameFromStr("set_fallback_environment")
		defer methodName.Destroy()
		ptrsForWorld3D.fnSetFallbackEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4143518816))
	}
	{
		methodName := StringNameFromStr("get_fallback_environment")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetFallbackEnvironment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3082064660))
	}
	{
		methodName := StringNameFromStr("set_camera_attributes")
		defer methodName.Destroy()
		ptrsForWorld3D.fnSetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2817810567))
	}
	{
		methodName := StringNameFromStr("get_camera_attributes")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetCameraAttributes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3921283215))
	}
	{
		methodName := StringNameFromStr("get_direct_space_state")
		defer methodName.Destroy()
		ptrsForWorld3D.fnGetDirectSpaceState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2069328350))
	}

}

type World3D struct {
	Resource
}

func (me *World3D) BaseClass() string {
	return "World3D"
}

func NewWorld3D() *World3D {
	str := StringNameFromStr("World3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &World3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *World3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *World3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *World3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *World3D) GetSpace() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World3D) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World3D) GetScenario() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetScenario), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World3D) SetEnvironment(env Environment) {
	cargs := []gdc.ConstTypePtr{env.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnSetEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *World3D) GetEnvironment() Environment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEnvironment()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World3D) SetFallbackEnvironment(env Environment) {
	cargs := []gdc.ConstTypePtr{env.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnSetFallbackEnvironment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *World3D) GetFallbackEnvironment() Environment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEnvironment()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetFallbackEnvironment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World3D) SetCameraAttributes(attributes CameraAttributes) {
	cargs := []gdc.ConstTypePtr{attributes.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnSetCameraAttributes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *World3D) GetCameraAttributes() CameraAttributes {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCameraAttributes()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetCameraAttributes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *World3D) GetDirectSpaceState() PhysicsDirectSpaceState3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectSpaceState3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorld3D.fnGetDirectSpaceState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
