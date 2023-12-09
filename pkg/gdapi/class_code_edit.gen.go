// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *CodeEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CodeEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CodeEdit) XConfirmCodeCompletion(replace bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) XRequestCodeCompletion(force bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) XFilterCodeCompletionCandidates(candidates Dictionary, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetIndentSize(size int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetIndentSize()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetIndentUsingSpaces(use_spaces bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsIndentUsingSpaces()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetAutoIndentEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsAutoIndentEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetAutoIndentPrefixes(prefixes String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetAutoIndentPrefixes()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) DoIndent()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IndentLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) UnindentLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ConvertIndent(from_line int, to_line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetAutoBraceCompletionEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsAutoBraceCompletionEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetHighlightMatchingBracesEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsHighlightMatchingBracesEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) AddAutoBraceCompletionPair(start_key String, end_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetAutoBraceCompletionPairs(pairs Dictionary, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetAutoBraceCompletionPairs()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) HasAutoBraceCompletionOpenKey(open_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) HasAutoBraceCompletionCloseKey(close_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetAutoBraceCompletionCloseKey(open_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetDrawBreakpointsGutter(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsDrawingBreakpointsGutter()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetDrawBookmarksGutter(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsDrawingBookmarksGutter()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetDrawExecutingLinesGutter(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsDrawingExecutingLinesGutter()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetLineAsBreakpoint(line int, breakpointed bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsLineBreakpointed(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ClearBreakpointedLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetBreakpointedLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetLineAsBookmarked(line int, bookmarked bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsLineBookmarked(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ClearBookmarkedLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetBookmarkedLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetLineAsExecuting(line int, executing bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsLineExecuting(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ClearExecutingLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetExecutingLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetDrawLineNumbers(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsDrawLineNumbersEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetLineNumbersZeroPadded(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsLineNumbersZeroPadded()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetDrawFoldGutter(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsDrawingFoldGutter()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetLineFoldingEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsLineFoldingEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) CanFoldLine(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) FoldLine(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) UnfoldLine(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) FoldAllLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) UnfoldAllLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ToggleFoldableLine(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsLineFolded(line int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetFoldedLines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) AddStringDelimiter(start_key String, end_key String, line_only bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) RemoveStringDelimiter(start_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) HasStringDelimiter(start_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetStringDelimiters(string_delimiters String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ClearStringDelimiters()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetStringDelimiters()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsInString(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) AddCommentDelimiter(start_key String, end_key String, line_only bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) RemoveCommentDelimiter(start_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) HasCommentDelimiter(start_key String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetCommentDelimiters(comment_delimiters String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ClearCommentDelimiters()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetCommentDelimiters()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsInComment(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetDelimiterStartKey(delimiter_index int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetDelimiterEndKey(delimiter_index int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetDelimiterStartPosition(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetDelimiterEndPosition(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetCodeHint(code_hint String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetCodeHintDrawBelow(draw_below bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetTextForCodeCompletion()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) RequestCodeCompletion(force bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) AddCodeCompletionOption(type_ CodeEditCodeCompletionKind, display_text String, insert_text String, text_color Color, icon Resource, value Variant, location int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) UpdateCodeCompletionOptions(force bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetCodeCompletionOptions()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetCodeCompletionOption(index int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetCodeCompletionSelectedIndex()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetCodeCompletionSelectedIndex(index int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) ConfirmCodeCompletion(replace bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) CancelCodeCompletion()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetCodeCompletionEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsCodeCompletionEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetCodeCompletionPrefixes(prefixes String, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetCodeCompletionPrefixes()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetLineLengthGuidelines(guideline_columns int, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetLineLengthGuidelines()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetSymbolLookupOnClickEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *CodeEdit) IsSymbolLookupOnClickEnabled()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) GetTextForSymbolLookup()  {
  panic("TODO: implement")
}

func  (me *CodeEdit) SetSymbolLookupWordAsValid(valid bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
