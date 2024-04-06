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

type ptrsForVisualShaderList struct {
	fnSetMode            gdc.MethodBindPtr
	fnAddNode            gdc.MethodBindPtr
	fnGetNode            gdc.MethodBindPtr
	fnSetNodePosition    gdc.MethodBindPtr
	fnGetNodePosition    gdc.MethodBindPtr
	fnGetNodeList        gdc.MethodBindPtr
	fnGetValidNodeId     gdc.MethodBindPtr
	fnRemoveNode         gdc.MethodBindPtr
	fnReplaceNode        gdc.MethodBindPtr
	fnIsNodeConnection   gdc.MethodBindPtr
	fnCanConnectNodes    gdc.MethodBindPtr
	fnConnectNodes       gdc.MethodBindPtr
	fnDisconnectNodes    gdc.MethodBindPtr
	fnConnectNodesForced gdc.MethodBindPtr
	fnGetNodeConnections gdc.MethodBindPtr
	fnSetGraphOffset     gdc.MethodBindPtr
	fnGetGraphOffset     gdc.MethodBindPtr
	fnAddVarying         gdc.MethodBindPtr
	fnRemoveVarying      gdc.MethodBindPtr
	fnHasVarying         gdc.MethodBindPtr
}

var ptrsForVisualShader ptrsForVisualShaderList

func initVisualShaderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShader")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mode")
		defer methodName.Destroy()
		ptrsForVisualShader.fnSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3978014962))
	}
	{
		methodName := StringNameFromStr("add_node")
		defer methodName.Destroy()
		ptrsForVisualShader.fnAddNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1560769431))
	}
	{
		methodName := StringNameFromStr("get_node")
		defer methodName.Destroy()
		ptrsForVisualShader.fnGetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3784670312))
	}
	{
		methodName := StringNameFromStr("set_node_position")
		defer methodName.Destroy()
		ptrsForVisualShader.fnSetNodePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726660721))
	}
	{
		methodName := StringNameFromStr("get_node_position")
		defer methodName.Destroy()
		ptrsForVisualShader.fnGetNodePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2175036082))
	}
	{
		methodName := StringNameFromStr("get_node_list")
		defer methodName.Destroy()
		ptrsForVisualShader.fnGetNodeList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2370592410))
	}
	{
		methodName := StringNameFromStr("get_valid_node_id")
		defer methodName.Destroy()
		ptrsForVisualShader.fnGetValidNodeId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 629467342))
	}
	{
		methodName := StringNameFromStr("remove_node")
		defer methodName.Destroy()
		ptrsForVisualShader.fnRemoveNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844050912))
	}
	{
		methodName := StringNameFromStr("replace_node")
		defer methodName.Destroy()
		ptrsForVisualShader.fnReplaceNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3144735253))
	}
	{
		methodName := StringNameFromStr("is_node_connection")
		defer methodName.Destroy()
		ptrsForVisualShader.fnIsNodeConnection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3922381898))
	}
	{
		methodName := StringNameFromStr("can_connect_nodes")
		defer methodName.Destroy()
		ptrsForVisualShader.fnCanConnectNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3922381898))
	}
	{
		methodName := StringNameFromStr("connect_nodes")
		defer methodName.Destroy()
		ptrsForVisualShader.fnConnectNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3081049573))
	}
	{
		methodName := StringNameFromStr("disconnect_nodes")
		defer methodName.Destroy()
		ptrsForVisualShader.fnDisconnectNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2268060358))
	}
	{
		methodName := StringNameFromStr("connect_nodes_forced")
		defer methodName.Destroy()
		ptrsForVisualShader.fnConnectNodesForced = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2268060358))
	}
	{
		methodName := StringNameFromStr("get_node_connections")
		defer methodName.Destroy()
		ptrsForVisualShader.fnGetNodeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1441964831))
	}
	{
		methodName := StringNameFromStr("set_graph_offset")
		defer methodName.Destroy()
		ptrsForVisualShader.fnSetGraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_graph_offset")
		defer methodName.Destroy()
		ptrsForVisualShader.fnGetGraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("add_varying")
		defer methodName.Destroy()
		ptrsForVisualShader.fnAddVarying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2084110726))
	}
	{
		methodName := StringNameFromStr("remove_varying")
		defer methodName.Destroy()
		ptrsForVisualShader.fnRemoveVarying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("has_varying")
		defer methodName.Destroy()
		ptrsForVisualShader.fnHasVarying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}

}

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
	VisualShaderNodeIdInvalid = -1
	VisualShaderNodeIdOutput  = 0
)

// Enums

type VisualShaderType int

const (
	VisualShaderTypeTypeVertex        VisualShaderType = 0
	VisualShaderTypeTypeFragment      VisualShaderType = 1
	VisualShaderTypeTypeLight         VisualShaderType = 2
	VisualShaderTypeTypeStart         VisualShaderType = 3
	VisualShaderTypeTypeProcess       VisualShaderType = 4
	VisualShaderTypeTypeCollide       VisualShaderType = 5
	VisualShaderTypeTypeStartCustom   VisualShaderType = 6
	VisualShaderTypeTypeProcessCustom VisualShaderType = 7
	VisualShaderTypeTypeSky           VisualShaderType = 8
	VisualShaderTypeTypeFog           VisualShaderType = 9
	VisualShaderTypeTypeMax           VisualShaderType = 10
)

type VisualShaderVaryingMode int

const (
	VisualShaderVaryingModeVaryingModeVertexToFragLight VisualShaderVaryingMode = 0
	VisualShaderVaryingModeVaryingModeFragToLight       VisualShaderVaryingMode = 1
	VisualShaderVaryingModeVaryingModeMax               VisualShaderVaryingMode = 2
)

type VisualShaderVaryingType int

const (
	VisualShaderVaryingTypeVaryingTypeFloat     VisualShaderVaryingType = 0
	VisualShaderVaryingTypeVaryingTypeInt       VisualShaderVaryingType = 1
	VisualShaderVaryingTypeVaryingTypeUint      VisualShaderVaryingType = 2
	VisualShaderVaryingTypeVaryingTypeVector2D  VisualShaderVaryingType = 3
	VisualShaderVaryingTypeVaryingTypeVector3D  VisualShaderVaryingType = 4
	VisualShaderVaryingTypeVaryingTypeVector4D  VisualShaderVaryingType = 5
	VisualShaderVaryingTypeVaryingTypeBoolean   VisualShaderVaryingType = 6
	VisualShaderVaryingTypeVaryingTypeTransform VisualShaderVaryingType = 7
	VisualShaderVaryingTypeVaryingTypeMax       VisualShaderVaryingType = 8
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

func (me *VisualShader) SetMode(mode ShaderMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) AddNode(type_ VisualShaderType, node VisualShaderNode, position Vector2, id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), node.AsCTypePtr(), position.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnAddNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) GetNode(type_ VisualShaderType, id int64) VisualShaderNode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVisualShaderNode()
	pinner.Pin(&type_)
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnGetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShader) SetNodePosition(type_ VisualShaderType, id int64, position Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnSetNodePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) GetNodePosition(type_ VisualShaderType, id int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&type_)
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnGetNodePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShader) GetNodeList(type_ VisualShaderType) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&type_)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnGetNodeList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShader) GetValidNodeId(type_ VisualShaderType) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&type_)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnGetValidNodeId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShader) RemoveNode(type_ VisualShaderType, id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnRemoveNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) ReplaceNode(type_ VisualShaderType, id int64, new_class StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&id), new_class.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnReplaceNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) IsNodeConnection(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&type_)
	pinner.Pin(&from_node)
	pinner.Pin(&from_port)
	pinner.Pin(&to_node)
	pinner.Pin(&to_port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnIsNodeConnection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShader) CanConnectNodes(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&type_)
	pinner.Pin(&from_node)
	pinner.Pin(&from_port)
	pinner.Pin(&to_node)
	pinner.Pin(&to_port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnCanConnectNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShader) ConnectNodes(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&type_)
	pinner.Pin(&from_node)
	pinner.Pin(&from_port)
	pinner.Pin(&to_node)
	pinner.Pin(&to_port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnConnectNodes), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShader) DisconnectNodes(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnDisconnectNodes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) ConnectNodesForced(type_ VisualShaderType, from_node int64, from_port int64, to_node int64, to_port int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&from_node), gdc.ConstTypePtr(&from_port), gdc.ConstTypePtr(&to_node), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnConnectNodesForced), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) GetNodeConnections(type_ VisualShaderType) []Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&type_)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnGetNodeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *VisualShader) SetGraphOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnSetGraphOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) GetGraphOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnGetGraphOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShader) AddVarying(name String, mode VisualShaderVaryingMode, type_ VisualShaderVaryingType) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnAddVarying), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) RemoveVarying(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnRemoveVarying), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShader) HasVarying(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShader.fnHasVarying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
