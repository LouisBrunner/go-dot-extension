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

type VisualShader struct {
  Shader
}

func (me *VisualShader) BaseClass() string {
  return "VisualShader"
}

func NewVisualShader() *VisualShader {
  str := StringNameFromStr("VisualShader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShader{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) AddNode(type_ VisualShaderType, node VisualShaderNode, position Vector2, id int64, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1560769431) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , node.AsCTypePtr(), position.AsCTypePtr(), gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) GetNode(type_ VisualShaderType, id int64, ) VisualShaderNode {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3784670312) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVisualShaderNode()
  pinner.Pin(&type_)
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShader) SetNodePosition(type_ VisualShaderType, id int64, position Vector2, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726660721) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&id) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) GetNodePosition(type_ VisualShaderType, id int64, ) Vector2 {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2175036082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&type_)
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShader) GetNodeList(type_ VisualShaderType, ) PackedInt32Array {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2370592410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShader) GetValidNodeId(type_ VisualShaderType, ) int64 {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_valid_node_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 629467342) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShader) RemoveNode(type_ VisualShaderType, id int64, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844050912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) ReplaceNode(type_ VisualShaderType, id int64, new_class StringName, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("replace_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3144735253) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&id) , new_class.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) IsNodeConnection(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64, ) bool {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_node_connection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3922381898) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&from_node) , gdc.ConstTypePtr(&from_port) , gdc.ConstTypePtr(&to_node) , gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&type_)
  pinner.Pin(&from_node)
  pinner.Pin(&from_port)
  pinner.Pin(&to_node)
  pinner.Pin(&to_port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShader) CanConnectNodes(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64, ) bool {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_connect_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3922381898) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&from_node) , gdc.ConstTypePtr(&from_port) , gdc.ConstTypePtr(&to_node) , gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&type_)
  pinner.Pin(&from_node)
  pinner.Pin(&from_port)
  pinner.Pin(&to_node)
  pinner.Pin(&to_port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShader) ConnectNodes(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64, ) Error {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3081049573) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&from_node) , gdc.ConstTypePtr(&from_port) , gdc.ConstTypePtr(&to_node) , gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&type_)
  pinner.Pin(&from_node)
  pinner.Pin(&from_port)
  pinner.Pin(&to_node)
  pinner.Pin(&to_port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShader) DisconnectNodes(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("disconnect_nodes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2268060358) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&from_node) , gdc.ConstTypePtr(&from_port) , gdc.ConstTypePtr(&to_node) , gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) ConnectNodesForced(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_nodes_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2268060358) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&from_node) , gdc.ConstTypePtr(&from_port) , gdc.ConstTypePtr(&to_node) , gdc.ConstTypePtr(&to_port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) GetNodeConnections(type_ VisualShaderType, ) []Dictionary {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1441964831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *VisualShader) SetGraphOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) GetGraphOffset() Vector2 {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_graph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShader) AddVarying(name String, mode VisualShaderVaryingMode, type_ VisualShaderVaryingType, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_varying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2084110726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&mode) , gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) RemoveVarying(name String, )  {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_varying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShader) HasVarying(name String, ) bool {
  classNameV := StringNameFromStr("VisualShader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_varying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
