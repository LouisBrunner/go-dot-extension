// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CodeEdit struct {
  obj gdc.ObjectPtr
}

func (me *CodeEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CodeEdit) BaseClass() string {
  return "CodeEdit"
}



// Enums

type CodeEditCodeCompletionKind int
const (
  CodeEditCodeCompletionKindKindClass CodeEditCodeCompletionKind = 0
  CodeEditCodeCompletionKindKindFunction CodeEditCodeCompletionKind = 1
  CodeEditCodeCompletionKindKindSignal CodeEditCodeCompletionKind = 2
  CodeEditCodeCompletionKindKindVariable CodeEditCodeCompletionKind = 3
  CodeEditCodeCompletionKindKindMember CodeEditCodeCompletionKind = 4
  CodeEditCodeCompletionKindKindEnum CodeEditCodeCompletionKind = 5
  CodeEditCodeCompletionKindKindConstant CodeEditCodeCompletionKind = 6
  CodeEditCodeCompletionKindKindNodePath CodeEditCodeCompletionKind = 7
  CodeEditCodeCompletionKindKindFilePath CodeEditCodeCompletionKind = 8
  CodeEditCodeCompletionKindKindPlainText CodeEditCodeCompletionKind = 9
)

type CodeEditCodeCompletionLocation int
const (
  CodeEditCodeCompletionLocationLocationLocal CodeEditCodeCompletionLocation = 0
  CodeEditCodeCompletionLocationLocationParentMask CodeEditCodeCompletionLocation = 256
  CodeEditCodeCompletionLocationLocationOtherUserCode CodeEditCodeCompletionLocation = 512
  CodeEditCodeCompletionLocationLocationOther CodeEditCodeCompletionLocation = 1024
)

func (me *CodeEdit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CodeEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CodeEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CodeEdit) SetIndentSize(size int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indent_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetIndentSize() int {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_indent_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetIndentUsingSpaces(use_spaces bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indent_using_spaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_spaces), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsIndentUsingSpaces() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_indent_using_spaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetAutoIndentEnabled(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_indent_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsAutoIndentEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_indent_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetAutoIndentPrefixes(prefixes String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_indent_prefixes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(prefixes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetAutoIndentPrefixes() String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_indent_prefixes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) DoIndent()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("do_indent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IndentLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("indent_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) UnindentLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unindent_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) ConvertIndent(from_line int, to_line int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convert_indent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 423910286) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&to_line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) SetAutoBraceCompletionEnabled(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_brace_completion_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsAutoBraceCompletionEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_auto_brace_completion_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetHighlightMatchingBracesEnabled(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_highlight_matching_braces_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsHighlightMatchingBracesEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_highlight_matching_braces_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) AddAutoBraceCompletionPair(start_key String, end_key String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_auto_brace_completion_pair")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3186203200) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), gdc.ConstTypePtr(end_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) SetAutoBraceCompletionPairs(pairs Dictionary, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_brace_completion_pairs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155329257) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(pairs.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetAutoBraceCompletionPairs() Dictionary {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_brace_completion_pairs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) HasAutoBraceCompletionOpenKey(open_key String, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_auto_brace_completion_open_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(open_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) HasAutoBraceCompletionCloseKey(close_key String, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_auto_brace_completion_close_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(close_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetAutoBraceCompletionCloseKey(open_key String, ) String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_brace_completion_close_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(open_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetDrawBreakpointsGutter(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_breakpoints_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsDrawingBreakpointsGutter() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_breakpoints_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetDrawBookmarksGutter(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_bookmarks_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsDrawingBookmarksGutter() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_bookmarks_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetDrawExecutingLinesGutter(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_executing_lines_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsDrawingExecutingLinesGutter() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_executing_lines_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetLineAsBreakpoint(line int, breakpointed bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_as_breakpoint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&breakpointed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsLineBreakpointed(line int, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_breakpointed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) ClearBreakpointedLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_breakpointed_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetBreakpointedLines() PackedInt32Array {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_breakpointed_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetLineAsBookmarked(line int, bookmarked bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_as_bookmarked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&bookmarked), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsLineBookmarked(line int, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_bookmarked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) ClearBookmarkedLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_bookmarked_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetBookmarkedLines() PackedInt32Array {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bookmarked_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetLineAsExecuting(line int, executing bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_as_executing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&executing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsLineExecuting(line int, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_executing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) ClearExecutingLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_executing_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetExecutingLines() PackedInt32Array {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_executing_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetDrawLineNumbers(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_line_numbers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsDrawLineNumbersEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_line_numbers_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetLineNumbersZeroPadded(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_numbers_zero_padded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsLineNumbersZeroPadded() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_numbers_zero_padded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetDrawFoldGutter(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_fold_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsDrawingFoldGutter() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_drawing_fold_gutter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetLineFoldingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_folding_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsLineFoldingEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_folding_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) CanFoldLine(line int, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_fold_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) FoldLine(line int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fold_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) UnfoldLine(line int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unfold_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) FoldAllLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fold_all_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) UnfoldAllLines()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unfold_all_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) ToggleFoldableLine(line int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("toggle_foldable_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsLineFolded(line int, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_line_folded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetFoldedLines() int {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_folded_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) AddStringDelimiter(start_key String, end_key String, line_only bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_string_delimiter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3146098955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), gdc.ConstTypePtr(end_key.AsCTypePtr()), gdc.ConstTypePtr(&line_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) RemoveStringDelimiter(start_key String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_string_delimiter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) HasStringDelimiter(start_key String, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_string_delimiter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetStringDelimiters(string_delimiters String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_string_delimiters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_delimiters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) ClearStringDelimiters()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_string_delimiters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetStringDelimiters() String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_string_delimiters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) IsInString(line int, column int, ) int {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3294126239) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) AddCommentDelimiter(start_key String, end_key String, line_only bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_comment_delimiter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3146098955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), gdc.ConstTypePtr(end_key.AsCTypePtr()), gdc.ConstTypePtr(&line_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) RemoveCommentDelimiter(start_key String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_comment_delimiter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) HasCommentDelimiter(start_key String, ) bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_comment_delimiter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(start_key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetCommentDelimiters(comment_delimiters String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_comment_delimiters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(comment_delimiters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) ClearCommentDelimiters()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_comment_delimiters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetCommentDelimiters() String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_comment_delimiters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) IsInComment(line int, column int, ) int {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_in_comment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3294126239) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetDelimiterStartKey(delimiter_index int, ) String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delimiter_start_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delimiter_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetDelimiterEndKey(delimiter_index int, ) String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delimiter_end_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delimiter_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetDelimiterStartPosition(line int, column int, ) Vector2 {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delimiter_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3016396712) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetDelimiterEndPosition(line int, column int, ) Vector2 {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_delimiter_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3016396712) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetCodeHint(code_hint String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(code_hint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) SetCodeHintDrawBelow(draw_below bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code_hint_draw_below")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_below), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetTextForCodeCompletion() String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_for_code_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) RequestCodeCompletion(force bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("request_code_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107499316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) AddCodeCompletionOption(type_ CodeEditCodeCompletionKind, display_text String, insert_text String, text_color Color, icon Resource, value Variant, location int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_code_completion_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1629240608) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(display_text.AsCTypePtr()), gdc.ConstTypePtr(insert_text.AsCTypePtr()), gdc.ConstTypePtr(text_color.AsCTypePtr()), gdc.ConstTypePtr(icon.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), gdc.ConstTypePtr(&location), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) UpdateCodeCompletionOptions(force bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_code_completion_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetCodeCompletionOptions() Dictionary {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code_completion_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetCodeCompletionOption(index int, ) Dictionary {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code_completion_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3485342025) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetCodeCompletionSelectedIndex() int {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code_completion_selected_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetCodeCompletionSelectedIndex(index int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code_completion_selected_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) ConfirmCodeCompletion(replace bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("confirm_code_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107499316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&replace), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) CancelCodeCompletion()  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cancel_code_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) SetCodeCompletionEnabled(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code_completion_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsCodeCompletionEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_code_completion_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetCodeCompletionPrefixes(prefixes String, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code_completion_prefixes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(prefixes.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetCodeCompletionPrefixes() String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code_completion_prefixes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetLineLengthGuidelines(guideline_columns int, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_length_guidelines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&guideline_columns), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) GetLineLengthGuidelines() int {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_length_guidelines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetSymbolLookupOnClickEnabled(enable bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_symbol_lookup_on_click_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CodeEdit) IsSymbolLookupOnClickEnabled() bool {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_symbol_lookup_on_click_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) GetTextForSymbolLookup() String {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_for_symbol_lookup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CodeEdit) SetSymbolLookupWordAsValid(valid bool, )  {
  classNameV := StringNameFromStr("CodeEdit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_symbol_lookup_word_as_valid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&valid), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CodeEditBreakpointToggledSignalFn func(line int, )

func (me *CodeEdit) ConnectBreakpointToggled(subs SignalSubscribers, fn CodeEditBreakpointToggledSignalFn) {
  sig := StringNameFromStr("breakpoint_toggled")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectBreakpointToggled(subs SignalSubscribers, fn CodeEditBreakpointToggledSignalFn) {
  sig := StringNameFromStr("breakpoint_toggled")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CodeEditCodeCompletionRequestedSignalFn func()

func (me *CodeEdit) ConnectCodeCompletionRequested(subs SignalSubscribers, fn CodeEditCodeCompletionRequestedSignalFn) {
  sig := StringNameFromStr("code_completion_requested")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectCodeCompletionRequested(subs SignalSubscribers, fn CodeEditCodeCompletionRequestedSignalFn) {
  sig := StringNameFromStr("code_completion_requested")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CodeEditSymbolLookupSignalFn func(symbol String, line int, column int, )

func (me *CodeEdit) ConnectSymbolLookup(subs SignalSubscribers, fn CodeEditSymbolLookupSignalFn) {
  sig := StringNameFromStr("symbol_lookup")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectSymbolLookup(subs SignalSubscribers, fn CodeEditSymbolLookupSignalFn) {
  sig := StringNameFromStr("symbol_lookup")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type CodeEditSymbolValidateSignalFn func(symbol String, )

func (me *CodeEdit) ConnectSymbolValidate(subs SignalSubscribers, fn CodeEditSymbolValidateSignalFn) {
  sig := StringNameFromStr("symbol_validate")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectSymbolValidate(subs SignalSubscribers, fn CodeEditSymbolValidateSignalFn) {
  sig := StringNameFromStr("symbol_validate")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
