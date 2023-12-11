// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XMLParser struct {
  obj gdc.ObjectPtr
}

func (me *XMLParser) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XMLParser) BaseClass() string {
  return "XMLParser"
}



// Enums

type XMLParserNodeType int
const (
  XMLParserNodeTypeNodeNone XMLParserNodeType = 0
  XMLParserNodeTypeNodeElement XMLParserNodeType = 1
  XMLParserNodeTypeNodeElementEnd XMLParserNodeType = 2
  XMLParserNodeTypeNodeText XMLParserNodeType = 3
  XMLParserNodeTypeNodeComment XMLParserNodeType = 4
  XMLParserNodeTypeNodeCdata XMLParserNodeType = 5
  XMLParserNodeTypeNodeUnknown XMLParserNodeType = 6
)

func (me *XMLParser) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XMLParser) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XMLParser) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XMLParser) Read() Error {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("read")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetNodeType() XMLParserNodeType {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2984359541) // FIXME: should cache?
  var ret XMLParserNodeType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetNodeName() String {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetNodeData() String {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetNodeOffset() int {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_node_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetAttributeCount() int {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attribute_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetAttributeName(idx int, ) String {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attribute_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetAttributeValue(idx int, ) String {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attribute_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) HasAttribute(name String, ) bool {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_attribute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetNamedAttributeValue(name String, ) String {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_named_attribute_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetNamedAttributeValueSafe(name String, ) String {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_named_attribute_value_safe")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) IsEmpty() bool {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_empty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) GetCurrentLine() int {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) SkipSection()  {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("skip_section")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XMLParser) Seek(position int, ) Error {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) Open(file String, ) Error {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(file.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XMLParser) OpenBuffer(buffer PackedByteArray, ) Error {
  classNameV := StringNameFromStr("XMLParser")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
