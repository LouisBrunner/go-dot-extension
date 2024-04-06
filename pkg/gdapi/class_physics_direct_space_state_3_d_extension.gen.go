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

type ptrsForPhysicsDirectSpaceState3DExtensionList struct {
	fnXIntersectRay                  gdc.MethodBindPtr
	fnXIntersectPoint                gdc.MethodBindPtr
	fnXIntersectShape                gdc.MethodBindPtr
	fnXCastMotion                    gdc.MethodBindPtr
	fnXCollideShape                  gdc.MethodBindPtr
	fnXRestInfo                      gdc.MethodBindPtr
	fnXGetClosestPointToObjectVolume gdc.MethodBindPtr
	fnIsBodyExcludedFromQuery        gdc.MethodBindPtr
}

var ptrsForPhysicsDirectSpaceState3DExtension ptrsForPhysicsDirectSpaceState3DExtensionList

func initPhysicsDirectSpaceState3DExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsDirectSpaceState3DExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_body_excluded_from_query")
		defer methodName.Destroy()
		ptrsForPhysicsDirectSpaceState3DExtension.fnIsBodyExcludedFromQuery = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}

}

type PhysicsDirectSpaceState3DExtension struct {
	PhysicsDirectSpaceState3D
}

func (me *PhysicsDirectSpaceState3DExtension) BaseClass() string {
	return "PhysicsDirectSpaceState3DExtension"
}

func NewPhysicsDirectSpaceState3DExtension() *PhysicsDirectSpaceState3DExtension {
	str := StringNameFromStr("PhysicsDirectSpaceState3DExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsDirectSpaceState3DExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsDirectSpaceState3DExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState3DExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState3DExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsDirectSpaceState3DExtension) IsBodyExcludedFromQuery(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3DExtension.fnIsBodyExcludedFromQuery), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
