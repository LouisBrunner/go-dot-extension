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

type ptrsForGLTFPhysicsBodyList struct {
	fnFromNode           gdc.MethodBindPtr
	fnToNode             gdc.MethodBindPtr
	fnFromDictionary     gdc.MethodBindPtr
	fnToDictionary       gdc.MethodBindPtr
	fnGetBodyType        gdc.MethodBindPtr
	fnSetBodyType        gdc.MethodBindPtr
	fnGetMass            gdc.MethodBindPtr
	fnSetMass            gdc.MethodBindPtr
	fnGetLinearVelocity  gdc.MethodBindPtr
	fnSetLinearVelocity  gdc.MethodBindPtr
	fnGetAngularVelocity gdc.MethodBindPtr
	fnSetAngularVelocity gdc.MethodBindPtr
	fnGetCenterOfMass    gdc.MethodBindPtr
	fnSetCenterOfMass    gdc.MethodBindPtr
	fnGetInertiaTensor   gdc.MethodBindPtr
	fnSetInertiaTensor   gdc.MethodBindPtr
}

var ptrsForGLTFPhysicsBody ptrsForGLTFPhysicsBodyList

func initGLTFPhysicsBodyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFPhysicsBody")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("from_node")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnFromNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 420544174))
	}
	{
		methodName := StringNameFromStr("to_node")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnToNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3224013656))
	}
	{
		methodName := StringNameFromStr("from_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnFromDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1177544336))
	}
	{
		methodName := StringNameFromStr("to_dictionary")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnToDictionary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_body_type")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnGetBodyType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_body_type")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnSetBodyType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_mass")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnGetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_mass")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnSetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_linear_velocity")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_linear_velocity")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_angular_velocity")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_angular_velocity")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_center_of_mass")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnGetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_center_of_mass")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnSetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_inertia_tensor")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnGetInertiaTensor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2716978435))
	}
	{
		methodName := StringNameFromStr("set_inertia_tensor")
		defer methodName.Destroy()
		ptrsForGLTFPhysicsBody.fnSetInertiaTensor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1055510324))
	}
}

type GLTFPhysicsBody struct {
	Resource
}

func (me *GLTFPhysicsBody) BaseClass() string {
	return "GLTFPhysicsBody"
}

func NewGLTFPhysicsBody() *GLTFPhysicsBody {
	str := StringNameFromStr("GLTFPhysicsBody") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GLTFPhysicsBody{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GLTFPhysicsBody) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GLTFPhysicsBody) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GLTFPhysicsBody) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func GLTFPhysicsBodyFromNode(body_node CollisionObject3D) GLTFPhysicsBody {
	cargs := []gdc.ConstTypePtr{body_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFPhysicsBody()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnFromNode), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) ToNode() CollisionObject3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCollisionObject3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnToNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func GLTFPhysicsBodyFromDictionary(dictionary Dictionary) GLTFPhysicsBody {
	cargs := []gdc.ConstTypePtr{dictionary.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewGLTFPhysicsBody()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnFromDictionary), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) ToDictionary() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnToDictionary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) GetBodyType() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnGetBodyType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) SetBodyType(body_type String) {
	cargs := []gdc.ConstTypePtr{body_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnSetBodyType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsBody) GetMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnGetMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GLTFPhysicsBody) SetMass(mass float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnSetMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsBody) GetLinearVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) SetLinearVelocity(linear_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{linear_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsBody) GetAngularVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) SetAngularVelocity(angular_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{angular_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsBody) GetCenterOfMass() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnGetCenterOfMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) SetCenterOfMass(center_of_mass Vector3) {
	cargs := []gdc.ConstTypePtr{center_of_mass.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnSetCenterOfMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GLTFPhysicsBody) GetInertiaTensor() Basis {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBasis()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnGetInertiaTensor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GLTFPhysicsBody) SetInertiaTensor(inertia_tensor Basis) {
	cargs := []gdc.ConstTypePtr{inertia_tensor.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFPhysicsBody.fnSetInertiaTensor), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
