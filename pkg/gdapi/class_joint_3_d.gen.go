// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Joint3D struct {
  obj gdc.ObjectPtr
}

func (me *Joint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Joint3D) BaseClass() string {
  return "Joint3D"
}



// Enums

func (me *Joint3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Joint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Joint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Joint3D) SetNodeA(node NodePath, )  {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint3D) GetNodeA() NodePath {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Joint3D) SetNodeB(node NodePath, )  {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint3D) GetNodeB() NodePath {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Joint3D) SetSolverPriority(priority int, )  {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_solver_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint3D) GetSolverPriority() int {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_solver_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Joint3D) SetExcludeNodesFromCollision(enable bool, )  {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude_nodes_from_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Joint3D) GetExcludeNodesFromCollision() bool {
  classNameV := StringNameFromStr("Joint3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude_nodes_from_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Joint3D) GetPropNodeA() NodePath {
  panic("TODO: implement")
}

func (me *Joint3D) SetPropNodeA(value NodePath) {
  panic("TODO: implement")
}

func (me *Joint3D) GetPropNodeB() NodePath {
  panic("TODO: implement")
}

func (me *Joint3D) SetPropNodeB(value NodePath) {
  panic("TODO: implement")
}

func (me *Joint3D) GetPropSolverPriority() int {
  panic("TODO: implement")
}

func (me *Joint3D) SetPropSolverPriority(value int) {
  panic("TODO: implement")
}

func (me *Joint3D) GetPropExcludeNodesFromCollision() bool {
  panic("TODO: implement")
}

func (me *Joint3D) SetPropExcludeNodesFromCollision(value bool) {
  panic("TODO: implement")
}