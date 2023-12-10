// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFPhysicsShape struct {
  obj gdc.ObjectPtr
}

func (me *GLTFPhysicsShape) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFPhysicsShape) BaseClass() string {
  return "GLTFPhysicsShape"
}



// Enums

func (me *GLTFPhysicsShape) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFPhysicsShape) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFPhysicsShape) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  GLTFPhysicsShapeFromNode(shape_node CollisionShape3D, ) GLTFPhysicsShape {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3613751275) // FIXME: should cache?
  var ret GLTFPhysicsShape
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) ToNode(cache_shapes bool, ) CollisionShape3D {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 563689933) // FIXME: should cache?
  var ret CollisionShape3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cache_shapes), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  GLTFPhysicsShapeFromDictionary(dictionary Dictionary, ) GLTFPhysicsShape {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2390691823) // FIXME: should cache?
  var ret GLTFPhysicsShape
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dictionary.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) ToDictionary() Dictionary {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) GetShapeType() String {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetShapeType(shape_type String, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsShape) GetSize() Vector3 {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsShape) GetRadius() float32 {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetRadius(radius float32, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsShape) GetHeight() float32 {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetHeight(height float32, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsShape) GetIsTrigger() bool {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_trigger")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetIsTrigger(is_trigger bool, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_is_trigger")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_trigger), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsShape) GetMeshIndex() int {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetMeshIndex(mesh_index int, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mesh_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFPhysicsShape) GetImporterMesh() ImporterMesh {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_importer_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3161779525) // FIXME: should cache?
  var ret ImporterMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFPhysicsShape) SetImporterMesh(importer_mesh ImporterMesh, )  {
  classNameV := StringNameFromStr("GLTFPhysicsShape")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_importer_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2255166972) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(importer_mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *GLTFPhysicsShape) GetPropShapeType() String {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropShapeType(value String) {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropSize(value Vector3) {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) GetPropRadius() float32 {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropRadius(value float32) {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) GetPropHeight() float32 {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropHeight(value float32) {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) GetPropIsTrigger() bool {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropIsTrigger(value bool) {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) GetPropMeshIndex() int {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropMeshIndex(value int) {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) GetPropImporterMesh() ImporterMesh {
  panic("TODO: implement")
}

func (me *GLTFPhysicsShape) SetPropImporterMesh(value ImporterMesh) {
  panic("TODO: implement")
}