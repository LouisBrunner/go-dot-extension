// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShader struct {
  obj gdc.ObjectPtr
}

func (me *VisualShader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *VisualShader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShader) SetMode(mode ShaderMode, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) AddNode(type_ VisualShaderType, node VisualShaderNode, position Vector2, id int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) GetNode(type_ VisualShaderType, id int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) SetNodePosition(type_ VisualShaderType, id int, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) GetNodePosition(type_ VisualShaderType, id int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) GetNodeList(type_ VisualShaderType, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) GetValidNodeId(type_ VisualShaderType, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) RemoveNode(type_ VisualShaderType, id int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) ReplaceNode(type_ VisualShaderType, id int, new_class StringName, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) IsNodeConnection(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) CanConnectNodes(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) ConnectNodes(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) DisconnectNodes(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) ConnectNodesForced(type_ VisualShaderType, from_node int, from_port int, to_node int, to_port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) GetNodeConnections(type_ VisualShaderType, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) SetGraphOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) GetGraphOffset()  {
  panic("TODO: implement")
}

func  (me *VisualShader) AddVarying(name String, mode VisualShaderVaryingMode, type_ VisualShaderVaryingType, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) RemoveVarying(name String, )  {
  panic("TODO: implement")
}

func  (me *VisualShader) HasVarying(name String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
