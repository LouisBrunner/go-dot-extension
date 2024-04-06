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

type ptrsForXMLParserList struct {
	fnRead                       gdc.MethodBindPtr
	fnGetNodeType                gdc.MethodBindPtr
	fnGetNodeName                gdc.MethodBindPtr
	fnGetNodeData                gdc.MethodBindPtr
	fnGetNodeOffset              gdc.MethodBindPtr
	fnGetAttributeCount          gdc.MethodBindPtr
	fnGetAttributeName           gdc.MethodBindPtr
	fnGetAttributeValue          gdc.MethodBindPtr
	fnHasAttribute               gdc.MethodBindPtr
	fnGetNamedAttributeValue     gdc.MethodBindPtr
	fnGetNamedAttributeValueSafe gdc.MethodBindPtr
	fnIsEmpty                    gdc.MethodBindPtr
	fnGetCurrentLine             gdc.MethodBindPtr
	fnSkipSection                gdc.MethodBindPtr
	fnSeek                       gdc.MethodBindPtr
	fnOpen                       gdc.MethodBindPtr
	fnOpenBuffer                 gdc.MethodBindPtr
}

var ptrsForXMLParser ptrsForXMLParserList

func initXMLParserPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XMLParser")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("read")
		defer methodName.Destroy()
		ptrsForXMLParser.fnRead = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("get_node_type")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetNodeType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2984359541))
	}
	{
		methodName := StringNameFromStr("get_node_name")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetNodeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_node_data")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetNodeData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_node_offset")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetNodeOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_attribute_count")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetAttributeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_attribute_name")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetAttributeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_attribute_value")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetAttributeValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("has_attribute")
		defer methodName.Destroy()
		ptrsForXMLParser.fnHasAttribute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}
	{
		methodName := StringNameFromStr("get_named_attribute_value")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetNamedAttributeValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
	}
	{
		methodName := StringNameFromStr("get_named_attribute_value_safe")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetNamedAttributeValueSafe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForXMLParser.fnIsEmpty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_current_line")
		defer methodName.Destroy()
		ptrsForXMLParser.fnGetCurrentLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("skip_section")
		defer methodName.Destroy()
		ptrsForXMLParser.fnSkipSection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("seek")
		defer methodName.Destroy()
		ptrsForXMLParser.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844576869))
	}
	{
		methodName := StringNameFromStr("open")
		defer methodName.Destroy()
		ptrsForXMLParser.fnOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("open_buffer")
		defer methodName.Destroy()
		ptrsForXMLParser.fnOpenBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
	}

}

type XMLParser struct {
	RefCounted
}

func (me *XMLParser) BaseClass() string {
	return "XMLParser"
}

func NewXMLParser() *XMLParser {
	str := StringNameFromStr("XMLParser") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XMLParser{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XMLParserNodeType int

const (
	XMLParserNodeTypeNodeNone       XMLParserNodeType = 0
	XMLParserNodeTypeNodeElement    XMLParserNodeType = 1
	XMLParserNodeTypeNodeElementEnd XMLParserNodeType = 2
	XMLParserNodeTypeNodeText       XMLParserNodeType = 3
	XMLParserNodeTypeNodeComment    XMLParserNodeType = 4
	XMLParserNodeTypeNodeCdata      XMLParserNodeType = 5
	XMLParserNodeTypeNodeUnknown    XMLParserNodeType = 6
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

func (me *XMLParser) Read() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnRead), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XMLParser) GetNodeType() XMLParserNodeType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XMLParserNodeType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetNodeType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XMLParser) GetNodeName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetNodeName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XMLParser) GetNodeData() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetNodeData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XMLParser) GetNodeOffset() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetNodeOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XMLParser) GetAttributeCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetAttributeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XMLParser) GetAttributeName(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetAttributeName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XMLParser) GetAttributeValue(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetAttributeValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XMLParser) HasAttribute(name String) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnHasAttribute), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XMLParser) GetNamedAttributeValue(name String) String {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetNamedAttributeValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XMLParser) GetNamedAttributeValueSafe(name String) String {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetNamedAttributeValueSafe), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XMLParser) IsEmpty() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnIsEmpty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XMLParser) GetCurrentLine() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnGetCurrentLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XMLParser) SkipSection() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnSkipSection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XMLParser) Seek(position int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&position)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnSeek), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XMLParser) Open(file String) Error {
	cargs := []gdc.ConstTypePtr{file.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnOpen), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XMLParser) OpenBuffer(buffer PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXMLParser.fnOpenBuffer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
