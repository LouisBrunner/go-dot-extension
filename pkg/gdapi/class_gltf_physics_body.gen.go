// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFPhysicsBody struct {
  obj gdc.ObjectPtr
}

func (me *GLTFPhysicsBody) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFPhysicsBody) BaseClass() string {
  return "GLTFPhysicsBody"
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
  var ret GLTFPhysicsBody
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) ToNode() CollisionObject3D {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3224013656) // FIXME: should cache?
  var ret CollisionObject3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  GLTFPhysicsBodyFromDictionary(dictionary Dictionary, ) GLTFPhysicsBody {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1177544336) // FIXME: should cache?
  var ret GLTFPhysicsBody
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dictionary.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) ToDictionary() Dictionary {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) GetBodyType() String {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_body_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) SetBodyType(body_type String, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_body_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsBody) GetMass() float32 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) SetMass(mass float32, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsBody) GetLinearVelocity() Vector3 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) SetLinearVelocity(linear_velocity Vector3, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(linear_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsBody) GetAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) SetAngularVelocity(angular_velocity Vector3, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(angular_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsBody) GetInertiaTensor() Basis {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inertia_tensor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2716978435) // FIXME: should cache?
  var ret Basis
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsBody) SetInertiaTensor(inertia_tensor Basis, )  {
  classNameV := StringNameFromStr("GLTFPhysicsBody")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inertia_tensor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1055510324) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(inertia_tensor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
