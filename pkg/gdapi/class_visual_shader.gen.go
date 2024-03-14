// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShader struct {
  Shader
}

func (me *VisualShader) BaseClass() string {
  return "VisualShader"
}



// Constants

var (
  VisualShaderNodeIdInvalid = "-1" // TODO: construct correctly
  VisualShaderNodeIdOutput = "0" // TODO: construct correctly
)

// Enums

type VisualShaderType int
const (
  VisualShaderTypeTypeVertex VisualShaderType = 0
  VisualShaderTypeTypeFragment VisualShaderType = 1
  VisualShaderTypeTypeLight VisualShaderType = 2
  VisualShaderTypeTypeStart VisualShaderType = 3
  VisualShaderTypeTypeProcess VisualShaderType = 4
  VisualShaderTypeTypeCollide VisualShaderType = 5
  VisualShaderTypeTypeStartCustom VisualShaderType = 6
  VisualShaderTypeTypeProcessCustom VisualShaderType = 7
  VisualShaderTypeTypeSky VisualShaderType = 8
  VisualShaderTypeTypeFog VisualShaderType = 9
  VisualShaderTypeTypeMax VisualShaderType = 10
)

type VisualShaderVaryingMode int
const (
  VisualShaderVaryingModeVaryingModeVertexToFragLight VisualShaderVaryingMode = 0
  VisualShaderVaryingModeVaryingModeFragToLight VisualShaderVaryingMode = 1
  VisualShaderVaryingModeVaryingModeMax VisualShaderVaryingMode = 2
)

type VisualShaderVaryingType int
const (
  VisualShaderVaryingTypeVaryingTypeFloat VisualShaderVaryingType = 0
  VisualShaderVaryingTypeVaryingTypeInt VisualShaderVaryingType = 1
  VisualShaderVaryingTypeVaryingTypeUint VisualShaderVaryingType = 2
  VisualShaderVaryingTypeVaryingTypeVector2D VisualShaderVaryingType = 3
  VisualShaderVaryingTypeVaryingTypeVector3D VisualShaderVaryingType = 4
  VisualShaderVaryingTypeVaryingTypeVector4D VisualShaderVaryingType = 5
  VisualShaderVaryingTypeVaryingTypeBoolean VisualShaderVaryingType = 6
  VisualShaderVaryingTypeVaryingTypeTransform VisualShaderVaryingType = 7
  VisualShaderVaryingTypeVaryingTypeMax VisualShaderVaryingType = 8
)

func (me *VisualShader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShader) SetMode(mode ShaderMode, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3978014962) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) AddNode(type_ VisualShaderType, node VisualShaderNode, position Vector2, id int, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1560769431) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) GetNode(type_ VisualShaderType, id int, ) VisualShaderNode {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3784670312) // FIXME: should cache?
  var ret VisualShaderNode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) SetNodePosition(type_ VisualShaderType, id int, position Vector2, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726660721) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) GetNodePosition(type_ VisualShaderType, id int, ) Vector2 {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2175036082) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) GetNodeList(type_ VisualShaderType, ) PackedInt32Array {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2370592410) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) GetValidNodeId(type_ VisualShaderType, ) int {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_valid_node_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 629467342) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) RemoveNode(type_ VisualShaderType, id int, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844050912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) ReplaceNode(type_ VisualShaderType, id int, new_class StringName, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("replace_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3144735253) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(new_class.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) IsNodeConnection(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, ) bool {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_node_connection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3922381898) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) CanConnectNodes(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, ) bool {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_connect_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3922381898) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) ConnectNodes(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, ) Error {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3081049573) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) DisconnectNodes(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2268060358) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) ConnectNodesForced(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_nodes_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2268060358) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) GetNodeConnections(type_ VisualShaderType, ) Dictionary {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1441964831) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) SetGraphOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) GetGraphOffset() Vector2 {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShader) AddVarying(name String, mode VisualShaderVaryingMode, type_ VisualShaderVaryingType, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_varying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2084110726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) RemoveVarying(name String, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_varying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShader) HasVarying(name String, ) bool {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_varying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
