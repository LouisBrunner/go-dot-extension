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

type ptrsForSyntaxHighlighterList struct {
	fnXGetLineSyntaxHighlighting gdc.MethodBindPtr
	fnXClearHighlightingCache    gdc.MethodBindPtr
	fnXUpdateCache               gdc.MethodBindPtr
	fnGetLineSyntaxHighlighting  gdc.MethodBindPtr
	fnUpdateCache                gdc.MethodBindPtr
	fnClearHighlightingCache     gdc.MethodBindPtr
	fnGetTextEdit                gdc.MethodBindPtr
}

var ptrsForSyntaxHighlighter ptrsForSyntaxHighlighterList

func initSyntaxHighlighterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SyntaxHighlighter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_line_syntax_highlighting")
		defer methodName.Destroy()
		ptrsForSyntaxHighlighter.fnGetLineSyntaxHighlighting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3554694381))
	}
	{
		methodName := StringNameFromStr("update_cache")
		defer methodName.Destroy()
		ptrsForSyntaxHighlighter.fnUpdateCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_highlighting_cache")
		defer methodName.Destroy()
		ptrsForSyntaxHighlighter.fnClearHighlightingCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_text_edit")
		defer methodName.Destroy()
		ptrsForSyntaxHighlighter.fnGetTextEdit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1893027089))
	}

}

type SyntaxHighlighter struct {
	Resource
}

func (me *SyntaxHighlighter) BaseClass() string {
	return "SyntaxHighlighter"
}

func NewSyntaxHighlighter() *SyntaxHighlighter {
	str := StringNameFromStr("SyntaxHighlighter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SyntaxHighlighter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SyntaxHighlighter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SyntaxHighlighter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SyntaxHighlighter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SyntaxHighlighter) GetLineSyntaxHighlighting(line int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSyntaxHighlighter.fnGetLineSyntaxHighlighting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SyntaxHighlighter) UpdateCache() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSyntaxHighlighter.fnUpdateCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SyntaxHighlighter) ClearHighlightingCache() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSyntaxHighlighter.fnClearHighlightingCache), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SyntaxHighlighter) GetTextEdit() TextEdit {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTextEdit()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSyntaxHighlighter.fnGetTextEdit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
