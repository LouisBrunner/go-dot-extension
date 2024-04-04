// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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

func  GLTFPhysicsBodyFromNode(body_node CollisionObject3D, ) GLTFPhysicsBody {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 420544174) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGLTFPhysicsBody()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) ToNode() CollisionObject3D {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3224013656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCollisionObject3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  GLTFPhysicsBodyFromDictionary(dictionary Dictionary, ) GLTFPhysicsBody {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1177544336) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{dictionary.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGLTFPhysicsBody()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) ToDictionary() Dictionary {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) GetBodyType() String {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_body_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) SetBodyType(body_type String, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_body_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{body_type.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFPhysicsBody) GetMass() float64 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFPhysicsBody) SetMass(mass float64, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFPhysicsBody) GetLinearVelocity() Vector3 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) SetLinearVelocity(linear_velocity Vector3, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{linear_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFPhysicsBody) GetAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) SetAngularVelocity(angular_velocity Vector3, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{angular_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFPhysicsBody) GetCenterOfMass() Vector3 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) SetCenterOfMass(center_of_mass Vector3, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_center_of_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{center_of_mass.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFPhysicsBody) GetInertiaTensor() Basis {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inertia_tensor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2716978435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBasis()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFPhysicsBody) SetInertiaTensor(inertia_tensor Basis, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inertia_tensor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1055510324) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{inertia_tensor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
