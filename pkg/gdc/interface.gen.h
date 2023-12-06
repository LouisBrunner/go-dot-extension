// Code generated by cgen. DO NOT EDIT.
#pragma once

#include <gdextension_interface.h>

void* callGetProcAddress(GDExtensionInterfaceGetProcAddress getProcAddr, const char* procName);


GDExtensionVariantPtr callArrayOperatorIndex(GDExtensionInterfaceArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionVariantPtr callArrayOperatorIndexConst(GDExtensionInterfaceArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

void callArrayRef(GDExtensionInterfaceArrayRef fn, GDExtensionTypePtr pSelf, GDExtensionConstTypePtr pFrom);

void callArraySetTyped(GDExtensionInterfaceArraySetTyped fn, GDExtensionTypePtr pSelf, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pClassName, GDExtensionConstVariantPtr pScript);

GDExtensionObjectPtr callClassdbConstructObject(GDExtensionInterfaceClassdbConstructObject fn, GDExtensionConstStringNamePtr pClassname);

void* callClassdbGetClassTag(GDExtensionInterfaceClassdbGetClassTag fn, GDExtensionConstStringNamePtr pClassname);

const GDExtensionMethodBindPtr callClassdbGetMethodBind(GDExtensionInterfaceClassdbGetMethodBind fn, GDExtensionConstStringNamePtr pClassname, GDExtensionConstStringNamePtr pMethodname, GDExtensionInt pHash);

void callClassdbRegisterExtensionClass(GDExtensionInterfaceClassdbRegisterExtensionClass fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionConstStringNamePtr pParentClassName, GDExtensionClassCreationInfo* pExtensionFuncs);

void callClassdbRegisterExtensionClassIntegerConstant(GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionConstStringNamePtr pEnumName, GDExtensionConstStringNamePtr pConstantName, GDExtensionInt pConstantValue, GDExtensionBool pIsBitfield);

void callClassdbRegisterExtensionClassMethod(GDExtensionInterfaceClassdbRegisterExtensionClassMethod fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionClassMethodInfo* pMethodInfo);

void callClassdbRegisterExtensionClassProperty(GDExtensionInterfaceClassdbRegisterExtensionClassProperty fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionPropertyInfo* pInfo, GDExtensionConstStringNamePtr pSetter, GDExtensionConstStringNamePtr pGetter);

void callClassdbRegisterExtensionClassPropertyGroup(GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionConstStringPtr pGroupName, GDExtensionConstStringPtr pPrefix);

void callClassdbRegisterExtensionClassPropertySubgroup(GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionConstStringPtr pSubgroupName, GDExtensionConstStringPtr pPrefix);

void callClassdbRegisterExtensionClassSignal(GDExtensionInterfaceClassdbRegisterExtensionClassSignal fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName, GDExtensionConstStringNamePtr pSignalName, GDExtensionPropertyInfo* pArgumentInfo, GDExtensionInt pArgumentCount);

void callClassdbUnregisterExtensionClass(GDExtensionInterfaceClassdbUnregisterExtensionClass fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionConstStringNamePtr pClassName);

GDExtensionVariantPtr callDictionaryOperatorIndex(GDExtensionInterfaceDictionaryOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionConstVariantPtr pKey);

GDExtensionVariantPtr callDictionaryOperatorIndexConst(GDExtensionInterfaceDictionaryOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionConstVariantPtr pKey);

void callEditorAddPlugin(GDExtensionInterfaceEditorAddPlugin fn, GDExtensionConstStringNamePtr pClassName);

void callEditorRemovePlugin(GDExtensionInterfaceEditorRemovePlugin fn, GDExtensionConstStringNamePtr pClassName);

uint64_t callFileAccessGetBuffer(GDExtensionInterfaceFileAccessGetBuffer fn, GDExtensionConstObjectPtr pInstance, uint8_t* pDst, uint64_t pLength);

void callFileAccessStoreBuffer(GDExtensionInterfaceFileAccessStoreBuffer fn, GDExtensionObjectPtr pInstance, const uint8_t* pSrc, uint64_t pLength);

void callGetGodotVersion(GDExtensionInterfaceGetGodotVersion fn, GDExtensionGodotVersion* rGodotVersion);

void callGetLibraryPath(GDExtensionInterfaceGetLibraryPath fn, GDExtensionClassLibraryPtr pLibrary, GDExtensionUninitializedStringPtr rPath);

uint64_t callGetNativeStructSize(GDExtensionInterfaceGetNativeStructSize fn, GDExtensionConstStringNamePtr pName);

GDExtensionVariantFromTypeConstructorFunc callGetVariantFromTypeConstructor(GDExtensionInterfaceGetVariantFromTypeConstructor fn, GDExtensionVariantType pType);

void callCallVariantFromTypeConstructorFunc(GDExtensionVariantFromTypeConstructorFunc fn, GDExtensionUninitializedVariantPtr arg0, GDExtensionTypePtr arg1);

GDExtensionTypeFromVariantConstructorFunc callGetVariantToTypeConstructor(GDExtensionInterfaceGetVariantToTypeConstructor fn, GDExtensionVariantType pType);

void callCallTypeFromVariantConstructorFunc(GDExtensionTypeFromVariantConstructorFunc fn, GDExtensionUninitializedTypePtr arg0, GDExtensionVariantPtr arg1);

GDExtensionObjectPtr callGlobalGetSingleton(GDExtensionInterfaceGlobalGetSingleton fn, GDExtensionConstStringNamePtr pName);

void* callMemAlloc(GDExtensionInterfaceMemAlloc fn, size_t pBytes);

void callMemFree(GDExtensionInterfaceMemFree fn, void* pPtr);

void* callMemRealloc(GDExtensionInterfaceMemRealloc fn, void* pPtr, size_t pBytes);

GDExtensionObjectPtr callObjectCastTo(GDExtensionInterfaceObjectCastTo fn, GDExtensionConstObjectPtr pObject, void* pClassTag);

void callObjectDestroy(GDExtensionInterfaceObjectDestroy fn, GDExtensionObjectPtr pO);

GDExtensionBool callObjectGetClassName(GDExtensionInterfaceObjectGetClassName fn, GDExtensionConstObjectPtr pObject, GDExtensionClassLibraryPtr pLibrary, GDExtensionUninitializedStringNamePtr rClassName);

void* callObjectGetInstanceBinding(GDExtensionInterfaceObjectGetInstanceBinding fn, GDExtensionObjectPtr pO, void* pToken, GDExtensionInstanceBindingCallbacks* pCallbacks);

GDExtensionObjectPtr callObjectGetInstanceFromId(GDExtensionInterfaceObjectGetInstanceFromId fn, GDObjectInstanceID pInstanceId);

GDObjectInstanceID callObjectGetInstanceId(GDExtensionInterfaceObjectGetInstanceId fn, GDExtensionConstObjectPtr pObject);

void callObjectMethodBindCall(GDExtensionInterfaceObjectMethodBindCall fn, GDExtensionMethodBindPtr pMethodBind, GDExtensionObjectPtr pInstance, const GDExtensionConstVariantPtr* pArgs, GDExtensionInt pArgCount, GDExtensionUninitializedVariantPtr rRet, GDExtensionCallError* rError);

void callObjectMethodBindPtrcall(GDExtensionInterfaceObjectMethodBindPtrcall fn, GDExtensionMethodBindPtr pMethodBind, GDExtensionObjectPtr pInstance, const GDExtensionConstTypePtr* pArgs, GDExtensionTypePtr rRet);

void callObjectSetInstance(GDExtensionInterfaceObjectSetInstance fn, GDExtensionObjectPtr pO, GDExtensionConstStringNamePtr pClassname, GDExtensionClassInstancePtr pInstance);

void callObjectSetInstanceBinding(GDExtensionInterfaceObjectSetInstanceBinding fn, GDExtensionObjectPtr pO, void* pToken, void* pBinding, GDExtensionInstanceBindingCallbacks* pCallbacks);

uint8_t* callPackedByteArrayOperatorIndex(GDExtensionInterfacePackedByteArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

const uint8_t* callPackedByteArrayOperatorIndexConst(GDExtensionInterfacePackedByteArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionTypePtr callPackedColorArrayOperatorIndex(GDExtensionInterfacePackedColorArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionTypePtr callPackedColorArrayOperatorIndexConst(GDExtensionInterfacePackedColorArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

float* callPackedFloat32ArrayOperatorIndex(GDExtensionInterfacePackedFloat32ArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

const float* callPackedFloat32ArrayOperatorIndexConst(GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

double* callPackedFloat64ArrayOperatorIndex(GDExtensionInterfacePackedFloat64ArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

const double* callPackedFloat64ArrayOperatorIndexConst(GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

int* callPackedInt32ArrayOperatorIndex(GDExtensionInterfacePackedInt32ArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

const int* callPackedInt32ArrayOperatorIndexConst(GDExtensionInterfacePackedInt32ArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

long long* callPackedInt64ArrayOperatorIndex(GDExtensionInterfacePackedInt64ArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

const long long* callPackedInt64ArrayOperatorIndexConst(GDExtensionInterfacePackedInt64ArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionStringPtr callPackedStringArrayOperatorIndex(GDExtensionInterfacePackedStringArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionStringPtr callPackedStringArrayOperatorIndexConst(GDExtensionInterfacePackedStringArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionTypePtr callPackedVector2ArrayOperatorIndex(GDExtensionInterfacePackedVector2ArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionTypePtr callPackedVector2ArrayOperatorIndexConst(GDExtensionInterfacePackedVector2ArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionTypePtr callPackedVector3ArrayOperatorIndex(GDExtensionInterfacePackedVector3ArrayOperatorIndex fn, GDExtensionTypePtr pSelf, GDExtensionInt pIndex);

GDExtensionTypePtr callPackedVector3ArrayOperatorIndexConst(GDExtensionInterfacePackedVector3ArrayOperatorIndexConst fn, GDExtensionConstTypePtr pSelf, GDExtensionInt pIndex);

void callPrintError(GDExtensionInterfacePrintError fn, const char* pDescription, const char* pFunction, const char* pFile, int32_t pLine, GDExtensionBool pEditorNotify);

void callPrintErrorWithMessage(GDExtensionInterfacePrintErrorWithMessage fn, const char* pDescription, const char* pMessage, const char* pFunction, const char* pFile, int32_t pLine, GDExtensionBool pEditorNotify);

void callPrintScriptError(GDExtensionInterfacePrintScriptError fn, const char* pDescription, const char* pFunction, const char* pFile, int32_t pLine, GDExtensionBool pEditorNotify);

void callPrintScriptErrorWithMessage(GDExtensionInterfacePrintScriptErrorWithMessage fn, const char* pDescription, const char* pMessage, const char* pFunction, const char* pFile, int32_t pLine, GDExtensionBool pEditorNotify);

void callPrintWarning(GDExtensionInterfacePrintWarning fn, const char* pDescription, const char* pFunction, const char* pFile, int32_t pLine, GDExtensionBool pEditorNotify);

void callPrintWarningWithMessage(GDExtensionInterfacePrintWarningWithMessage fn, const char* pDescription, const char* pMessage, const char* pFunction, const char* pFile, int32_t pLine, GDExtensionBool pEditorNotify);

GDExtensionObjectPtr callRefGetObject(GDExtensionInterfaceRefGetObject fn, GDExtensionConstRefPtr pRef);

void callRefSetObject(GDExtensionInterfaceRefSetObject fn, GDExtensionRefPtr pRef, GDExtensionObjectPtr pObject);

GDExtensionScriptInstancePtr callScriptInstanceCreate(GDExtensionInterfaceScriptInstanceCreate fn, GDExtensionScriptInstanceInfo* pInfo, GDExtensionScriptInstanceDataPtr pInstanceData);

void callStringNewWithLatin1Chars(GDExtensionInterfaceStringNewWithLatin1Chars fn, GDExtensionUninitializedStringPtr rDest, const char* pContents);

void callStringNewWithLatin1CharsAndLen(GDExtensionInterfaceStringNewWithLatin1CharsAndLen fn, GDExtensionUninitializedStringPtr rDest, const char* pContents, GDExtensionInt pSize);

void callStringNewWithUtf16Chars(GDExtensionInterfaceStringNewWithUtf16Chars fn, GDExtensionUninitializedStringPtr rDest, const uint16_t* pContents);

void callStringNewWithUtf16CharsAndLen(GDExtensionInterfaceStringNewWithUtf16CharsAndLen fn, GDExtensionUninitializedStringPtr rDest, const uint16_t* pContents, GDExtensionInt pSize);

void callStringNewWithUtf32Chars(GDExtensionInterfaceStringNewWithUtf32Chars fn, GDExtensionUninitializedStringPtr rDest, const unsigned* pContents);

void callStringNewWithUtf32CharsAndLen(GDExtensionInterfaceStringNewWithUtf32CharsAndLen fn, GDExtensionUninitializedStringPtr rDest, const unsigned* pContents, GDExtensionInt pSize);

void callStringNewWithUtf8Chars(GDExtensionInterfaceStringNewWithUtf8Chars fn, GDExtensionUninitializedStringPtr rDest, const char* pContents);

void callStringNewWithUtf8CharsAndLen(GDExtensionInterfaceStringNewWithUtf8CharsAndLen fn, GDExtensionUninitializedStringPtr rDest, const char* pContents, GDExtensionInt pSize);

void callStringNewWithWideChars(GDExtensionInterfaceStringNewWithWideChars fn, GDExtensionUninitializedStringPtr rDest, const int* pContents);

void callStringNewWithWideCharsAndLen(GDExtensionInterfaceStringNewWithWideCharsAndLen fn, GDExtensionUninitializedStringPtr rDest, const int* pContents, GDExtensionInt pSize);

unsigned* callStringOperatorIndex(GDExtensionInterfaceStringOperatorIndex fn, GDExtensionStringPtr pSelf, GDExtensionInt pIndex);

const unsigned* callStringOperatorIndexConst(GDExtensionInterfaceStringOperatorIndexConst fn, GDExtensionConstStringPtr pSelf, GDExtensionInt pIndex);

void callStringOperatorPlusEqC32Str(GDExtensionInterfaceStringOperatorPlusEqC32str fn, GDExtensionStringPtr pSelf, const unsigned* pB);

void callStringOperatorPlusEqChar(GDExtensionInterfaceStringOperatorPlusEqChar fn, GDExtensionStringPtr pSelf, char32_t pB);

void callStringOperatorPlusEqCstr(GDExtensionInterfaceStringOperatorPlusEqCstr fn, GDExtensionStringPtr pSelf, const char* pB);

void callStringOperatorPlusEqString(GDExtensionInterfaceStringOperatorPlusEqString fn, GDExtensionStringPtr pSelf, GDExtensionConstStringPtr pB);

void callStringOperatorPlusEqWcstr(GDExtensionInterfaceStringOperatorPlusEqWcstr fn, GDExtensionStringPtr pSelf, const int* pB);

GDExtensionInt callStringToLatin1Chars(GDExtensionInterfaceStringToLatin1Chars fn, GDExtensionConstStringPtr pSelf, char* rText, GDExtensionInt pMaxWriteLength);

GDExtensionInt callStringToUtf16Chars(GDExtensionInterfaceStringToUtf16Chars fn, GDExtensionConstStringPtr pSelf, uint16_t* rText, GDExtensionInt pMaxWriteLength);

GDExtensionInt callStringToUtf32Chars(GDExtensionInterfaceStringToUtf32Chars fn, GDExtensionConstStringPtr pSelf, unsigned* rText, GDExtensionInt pMaxWriteLength);

GDExtensionInt callStringToUtf8Chars(GDExtensionInterfaceStringToUtf8Chars fn, GDExtensionConstStringPtr pSelf, char* rText, GDExtensionInt pMaxWriteLength);

GDExtensionInt callStringToWideChars(GDExtensionInterfaceStringToWideChars fn, GDExtensionConstStringPtr pSelf, int* rText, GDExtensionInt pMaxWriteLength);

GDExtensionBool callVariantBooleanize(GDExtensionInterfaceVariantBooleanize fn, GDExtensionConstVariantPtr pSelf);

void callVariantCall(GDExtensionInterfaceVariantCall fn, GDExtensionVariantPtr pSelf, GDExtensionConstStringNamePtr pMethod, const GDExtensionConstVariantPtr* pArgs, GDExtensionInt pArgumentCount, GDExtensionUninitializedVariantPtr rReturn, GDExtensionCallError* rError);

void callVariantCallStatic(GDExtensionInterfaceVariantCallStatic fn, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pMethod, const GDExtensionConstVariantPtr* pArgs, GDExtensionInt pArgumentCount, GDExtensionUninitializedVariantPtr rReturn, GDExtensionCallError* rError);

GDExtensionBool callVariantCanConvert(GDExtensionInterfaceVariantCanConvert fn, GDExtensionVariantType pFrom, GDExtensionVariantType pTo);

GDExtensionBool callVariantCanConvertStrict(GDExtensionInterfaceVariantCanConvertStrict fn, GDExtensionVariantType pFrom, GDExtensionVariantType pTo);

void callVariantConstruct(GDExtensionInterfaceVariantConstruct fn, GDExtensionVariantType pType, GDExtensionUninitializedVariantPtr rBase, const GDExtensionConstVariantPtr* pArgs, int32_t pArgumentCount, GDExtensionCallError* rError);

void callVariantDestroy(GDExtensionInterfaceVariantDestroy fn, GDExtensionVariantPtr pSelf);

void callVariantDuplicate(GDExtensionInterfaceVariantDuplicate fn, GDExtensionConstVariantPtr pSelf, GDExtensionVariantPtr rRet, GDExtensionBool pDeep);

void callVariantEvaluate(GDExtensionInterfaceVariantEvaluate fn, GDExtensionVariantOperator pOp, GDExtensionConstVariantPtr pA, GDExtensionConstVariantPtr pB, GDExtensionUninitializedVariantPtr rReturn, uint8_t* rValid);

void callVariantGet(GDExtensionInterfaceVariantGet fn, GDExtensionConstVariantPtr pSelf, GDExtensionConstVariantPtr pKey, GDExtensionUninitializedVariantPtr rRet, uint8_t* rValid);

void callVariantGetConstantValue(GDExtensionInterfaceVariantGetConstantValue fn, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pConstant, GDExtensionUninitializedVariantPtr rRet);

void callVariantGetIndexed(GDExtensionInterfaceVariantGetIndexed fn, GDExtensionConstVariantPtr pSelf, GDExtensionInt pIndex, GDExtensionUninitializedVariantPtr rRet, uint8_t* rValid, uint8_t* rOob);

void callVariantGetKeyed(GDExtensionInterfaceVariantGetKeyed fn, GDExtensionConstVariantPtr pSelf, GDExtensionConstVariantPtr pKey, GDExtensionUninitializedVariantPtr rRet, uint8_t* rValid);

void callVariantGetNamed(GDExtensionInterfaceVariantGetNamed fn, GDExtensionConstVariantPtr pSelf, GDExtensionConstStringNamePtr pKey, GDExtensionUninitializedVariantPtr rRet, uint8_t* rValid);

GDExtensionPtrBuiltInMethod callVariantGetPtrBuiltinMethod(GDExtensionInterfaceVariantGetPtrBuiltinMethod fn, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pMethod, GDExtensionInt pHash);

void callCallPtrBuiltInMethod(GDExtensionPtrBuiltInMethod fn, GDExtensionTypePtr pBase, const GDExtensionConstTypePtr* pArgs, GDExtensionTypePtr rReturn, int pArgumentCount);

GDExtensionPtrConstructor callVariantGetPtrConstructor(GDExtensionInterfaceVariantGetPtrConstructor fn, GDExtensionVariantType pType, int32_t pConstructor);

void callCallPtrConstructor(GDExtensionPtrConstructor fn, GDExtensionUninitializedTypePtr pBase, const GDExtensionConstTypePtr* pArgs);

GDExtensionPtrDestructor callVariantGetPtrDestructor(GDExtensionInterfaceVariantGetPtrDestructor fn, GDExtensionVariantType pType);

void callCallPtrDestructor(GDExtensionPtrDestructor fn, GDExtensionTypePtr pBase);

GDExtensionPtrGetter callVariantGetPtrGetter(GDExtensionInterfaceVariantGetPtrGetter fn, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pMember);

void callCallPtrGetter(GDExtensionPtrGetter fn, GDExtensionConstTypePtr pBase, GDExtensionTypePtr rValue);

GDExtensionPtrIndexedGetter callVariantGetPtrIndexedGetter(GDExtensionInterfaceVariantGetPtrIndexedGetter fn, GDExtensionVariantType pType);

void callCallPtrIndexedGetter(GDExtensionPtrIndexedGetter fn, GDExtensionConstTypePtr pBase, GDExtensionInt pIndex, GDExtensionTypePtr rValue);

GDExtensionPtrIndexedSetter callVariantGetPtrIndexedSetter(GDExtensionInterfaceVariantGetPtrIndexedSetter fn, GDExtensionVariantType pType);

void callCallPtrIndexedSetter(GDExtensionPtrIndexedSetter fn, GDExtensionTypePtr pBase, GDExtensionInt pIndex, GDExtensionConstTypePtr pValue);

GDExtensionPtrKeyedChecker callVariantGetPtrKeyedChecker(GDExtensionInterfaceVariantGetPtrKeyedChecker fn, GDExtensionVariantType pType);

uint32_t callCallPtrKeyedChecker(GDExtensionPtrKeyedChecker fn, GDExtensionConstVariantPtr pBase, GDExtensionConstVariantPtr pKey);

GDExtensionPtrKeyedGetter callVariantGetPtrKeyedGetter(GDExtensionInterfaceVariantGetPtrKeyedGetter fn, GDExtensionVariantType pType);

void callCallPtrKeyedGetter(GDExtensionPtrKeyedGetter fn, GDExtensionConstTypePtr pBase, GDExtensionConstTypePtr pKey, GDExtensionTypePtr rValue);

GDExtensionPtrKeyedSetter callVariantGetPtrKeyedSetter(GDExtensionInterfaceVariantGetPtrKeyedSetter fn, GDExtensionVariantType pType);

void callCallPtrKeyedSetter(GDExtensionPtrKeyedSetter fn, GDExtensionTypePtr pBase, GDExtensionConstTypePtr pKey, GDExtensionConstTypePtr pValue);

GDExtensionPtrOperatorEvaluator callVariantGetPtrOperatorEvaluator(GDExtensionInterfaceVariantGetPtrOperatorEvaluator fn, GDExtensionVariantOperator pOperator, GDExtensionVariantType pTypeA, GDExtensionVariantType pTypeB);

void callCallPtrOperatorEvaluator(GDExtensionPtrOperatorEvaluator fn, GDExtensionConstTypePtr pLeft, GDExtensionConstTypePtr pRight, GDExtensionTypePtr rResult);

GDExtensionPtrSetter callVariantGetPtrSetter(GDExtensionInterfaceVariantGetPtrSetter fn, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pMember);

void callCallPtrSetter(GDExtensionPtrSetter fn, GDExtensionTypePtr pBase, GDExtensionConstTypePtr pValue);

GDExtensionPtrUtilityFunction callVariantGetPtrUtilityFunction(GDExtensionInterfaceVariantGetPtrUtilityFunction fn, GDExtensionConstStringNamePtr pFunction, GDExtensionInt pHash);

void callCallPtrUtilityFunction(GDExtensionPtrUtilityFunction fn, GDExtensionTypePtr rReturn, const GDExtensionConstTypePtr* pArgs, int pArgumentCount);

GDExtensionVariantType callVariantGetType(GDExtensionInterfaceVariantGetType fn, GDExtensionConstVariantPtr pSelf);

void callVariantGetTypeName(GDExtensionInterfaceVariantGetTypeName fn, GDExtensionVariantType pType, GDExtensionUninitializedStringPtr rName);

GDExtensionBool callVariantHasKey(GDExtensionInterfaceVariantHasKey fn, GDExtensionConstVariantPtr pSelf, GDExtensionConstVariantPtr pKey, uint8_t* rValid);

GDExtensionBool callVariantHasMember(GDExtensionInterfaceVariantHasMember fn, GDExtensionVariantType pType, GDExtensionConstStringNamePtr pMember);

GDExtensionBool callVariantHasMethod(GDExtensionInterfaceVariantHasMethod fn, GDExtensionConstVariantPtr pSelf, GDExtensionConstStringNamePtr pMethod);

GDExtensionInt callVariantHash(GDExtensionInterfaceVariantHash fn, GDExtensionConstVariantPtr pSelf);

GDExtensionBool callVariantHashCompare(GDExtensionInterfaceVariantHashCompare fn, GDExtensionConstVariantPtr pSelf, GDExtensionConstVariantPtr pOther);

void callVariantIterGet(GDExtensionInterfaceVariantIterGet fn, GDExtensionConstVariantPtr pSelf, GDExtensionVariantPtr rIter, GDExtensionUninitializedVariantPtr rRet, uint8_t* rValid);

GDExtensionBool callVariantIterInit(GDExtensionInterfaceVariantIterInit fn, GDExtensionConstVariantPtr pSelf, GDExtensionUninitializedVariantPtr rIter, uint8_t* rValid);

GDExtensionBool callVariantIterNext(GDExtensionInterfaceVariantIterNext fn, GDExtensionConstVariantPtr pSelf, GDExtensionVariantPtr rIter, uint8_t* rValid);

void callVariantNewCopy(GDExtensionInterfaceVariantNewCopy fn, GDExtensionUninitializedVariantPtr rDest, GDExtensionConstVariantPtr pSrc);

void callVariantNewNil(GDExtensionInterfaceVariantNewNil fn, GDExtensionUninitializedVariantPtr rDest);

GDExtensionInt callVariantRecursiveHash(GDExtensionInterfaceVariantRecursiveHash fn, GDExtensionConstVariantPtr pSelf, GDExtensionInt pRecursionCount);

void callVariantSet(GDExtensionInterfaceVariantSet fn, GDExtensionVariantPtr pSelf, GDExtensionConstVariantPtr pKey, GDExtensionConstVariantPtr pValue, uint8_t* rValid);

void callVariantSetIndexed(GDExtensionInterfaceVariantSetIndexed fn, GDExtensionVariantPtr pSelf, GDExtensionInt pIndex, GDExtensionConstVariantPtr pValue, uint8_t* rValid, uint8_t* rOob);

void callVariantSetKeyed(GDExtensionInterfaceVariantSetKeyed fn, GDExtensionVariantPtr pSelf, GDExtensionConstVariantPtr pKey, GDExtensionConstVariantPtr pValue, uint8_t* rValid);

void callVariantSetNamed(GDExtensionInterfaceVariantSetNamed fn, GDExtensionVariantPtr pSelf, GDExtensionConstStringNamePtr pKey, GDExtensionConstVariantPtr pValue, uint8_t* rValid);

void callVariantStringify(GDExtensionInterfaceVariantStringify fn, GDExtensionConstVariantPtr pSelf, GDExtensionStringPtr rRet);

int64_t callWorkerThreadPoolAddNativeGroupTask(GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask fn, GDExtensionObjectPtr pInstance, void* pFunc, void* pUserdata, int pElements, int pTasks, GDExtensionBool pHighPriority, GDExtensionConstStringPtr pDescription);

int64_t callWorkerThreadPoolAddNativeTask(GDExtensionInterfaceWorkerThreadPoolAddNativeTask fn, GDExtensionObjectPtr pInstance, void* pFunc, void* pUserdata, GDExtensionBool pHighPriority, GDExtensionConstStringPtr pDescription);

GDExtensionInt callXmlParserOpenBuffer(GDExtensionInterfaceXmlParserOpenBuffer fn, GDExtensionObjectPtr pInstance, const uint8_t* pBuffer, size_t pSize);

