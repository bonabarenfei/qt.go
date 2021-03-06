package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 266
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QCharRef struct {
	*qtrt.CObject
}
type QCharRef_ITF interface {
	QCharRef_PTR() *QCharRef
}

func (ptr *QCharRef) QCharRef_PTR() *QCharRef { return ptr }

func (this *QCharRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCharRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCharRefFromPointer(cthis unsafe.Pointer) *QCharRef {
	return &QCharRef{&qtrt.CObject{cthis}}
}
func (*QCharRef) NewFromPointer(cthis unsafe.Pointer) *QCharRef {
	return NewQCharRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstring.h:1052
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(QChar)

/*

 */
func (this *QCharRef) Operator_equal(c QChar_ITF /*123*/) *QCharRef {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1058
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(char)

/*

 */
func (this *QCharRef) Operator_equal1(c byte) *QCharRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1060
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(uchar)

/*

 */
func (this *QCharRef) Operator_equal2(c byte) *QCharRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1063
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(const QCharRef &)

/*

 */
func (this *QCharRef) Operator_equal3(c QCharRef_ITF) *QCharRef {
	var convArg0 unsafe.Pointer
	if c != nil && c.QCharRef_PTR() != nil {
		convArg0 = c.QCharRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1064
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(ushort)

/*

 */
func (this *QCharRef) Operator_equal4(rc uint16) *QCharRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1065
// index:5
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(short)

/*

 */
func (this *QCharRef) Operator_equal5(rc int16) *QCharRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSEs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1066
// index:6
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(uint)

/*

 */
func (this *QCharRef) Operator_equal6(rc uint) *QCharRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1067
// index:7
// Public inline Visibility=Default Availability=Available
// [16] QCharRef & operator=(int)

/*

 */
func (this *QCharRef) Operator_equal7(rc int) *QCharRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefaSEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rc)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCharRef)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1070
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if this string is null; otherwise returns false.

Example:


  QString().isNull();             // returns true
  QString("").isNull();           // returns false
  QString("abc").isNull();        // returns false



Qt makes a distinction between null strings and empty strings for historical reasons. For most applications, what matters is whether or not a string contains any data, and this can be determined using the isEmpty() function.

See also isEmpty().
*/
func (this *QCharRef) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1071
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPrint() const

/*

 */
func (this *QCharRef) IsPrint() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isPrintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1072
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isPunct() const

/*

 */
func (this *QCharRef) IsPunct() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isPunctEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1073
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSpace() const

/*

 */
func (this *QCharRef) IsSpace() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1074
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isMark() const

/*

 */
func (this *QCharRef) IsMark() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef6isMarkEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1075
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLetter() const

/*

 */
func (this *QCharRef) IsLetter() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8isLetterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1076
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNumber() const

/*

 */
func (this *QCharRef) IsNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8isNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1077
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLetterOrNumber()

/*

 */
func (this *QCharRef) IsLetterOrNumber() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef16isLetterOrNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1078
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDigit() const

/*

 */
func (this *QCharRef) IsDigit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isDigitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1079
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLower() const

/*
Returns true if the string only contains lowercase letters, otherwise returns false.

This function was introduced in  Qt 5.12.

See also QChar::isLower() and isUpper().
*/
func (this *QCharRef) IsLower() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1080
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isUpper() const

/*
Returns true if the string only contains uppercase letters, otherwise returns false.

This function was introduced in  Qt 5.12.

See also QChar::isUpper() and isLower().
*/
func (this *QCharRef) IsUpper() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7isUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1081
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTitleCase() const

/*

 */
func (this *QCharRef) IsTitleCase() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11isTitleCaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1083
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int digitValue() const

/*

 */
func (this *QCharRef) DigitValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef10digitValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1084
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toLower() const

/*
Returns a lowercase copy of the string.


  QString str = "The Qt PROJECT";
  str = str.toLower();        // str == "the qt project"



The case conversion will always happen in the 'C' locale. For locale dependent case folding use QLocale::toLower()

See also toUpper() and QLocale::toLower().
*/
func (this *QCharRef) ToLower() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7toLowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1085
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toUpper() const

/*
Returns an uppercase copy of the string.


  QString str = "TeXt";
  str = str.toUpper();        // str == "TEXT"



The case conversion will always happen in the 'C' locale. For locale dependent case folding use QLocale::toUpper()

See also toLower() and QLocale::toLower().
*/
func (this *QCharRef) ToUpper() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7toUpperEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1086
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar toTitleCase() const

/*

 */
func (this *QCharRef) ToTitleCase() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11toTitleCaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1088
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Category category() const

/*

 */
func (this *QCharRef) Category() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8categoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1089
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Direction direction() const

/*

 */
func (this *QCharRef) Direction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef9directionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1090
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::JoiningType joiningType() const

/*

 */
func (this *QCharRef) JoiningType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11joiningTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1092
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Joining joining() const

/*

 */
func (this *QCharRef) Joining() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7joiningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1105
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasMirrored() const

/*

 */
func (this *QCharRef) HasMirrored() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef11hasMirroredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstring.h:1106
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar mirroredChar() const

/*

 */
func (this *QCharRef) MirroredChar() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef12mirroredCharEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstring.h:1107
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString decomposition() const

/*

 */
func (this *QCharRef) Decomposition() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef13decompositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstring.h:1108
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Decomposition decompositionTag() const

/*

 */
func (this *QCharRef) DecompositionTag() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef16decompositionTagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1109
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar combiningClass() const

/*

 */
func (this *QCharRef) CombiningClass() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef14combiningClassEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1111
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::Script script() const

/*

 */
func (this *QCharRef) Script() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef6scriptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1113
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QChar::UnicodeVersion unicodeVersion() const

/*

 */
func (this *QCharRef) UnicodeVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef14unicodeVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstring.h:1115
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar cell() const

/*

 */
func (this *QCharRef) Cell() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef4cellEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1116
// index:0
// Public inline Visibility=Default Availability=Available
// [1] uchar row() const

/*

 */
func (this *QCharRef) Row() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return byte(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCell(uchar)

/*

 */
func (this *QCharRef) SetCell(cell byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef7setCellEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cell)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1118
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRow(uchar)

/*

 */
func (this *QCharRef) SetRow(row byte) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef6setRowEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstring.h:1123
// index:0
// Public inline Visibility=Default Availability=Available
// [1] char toLatin1() const

/*
Returns a Latin-1 representation of the string as a QByteArray.

The returned byte array is undefined if the string contains non-Latin1 characters. Those characters may be suppressed or replaced with a question mark.

See also fromLatin1(), toUtf8(), toLocal8Bit(), and QTextCodec.
*/
func (this *QCharRef) ToLatin1() byte {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("byte", rv).(byte) // 1111
}

// /usr/include/qt/QtCore/qstring.h:1124
// index:0
// Public inline Visibility=Default Availability=Available
// [2] ushort unicode() const

/*
Returns a Unicode representation of the string. The result remains valid until the string is modified.

Note: The returned string may not be '\0'-terminated. Use size() to determine the length of the array.

See also setUnicode(), utf16(), and fromRawData().
*/
func (this *QCharRef) Unicode() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QCharRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qstring.h:1125
// index:1
// Public inline Visibility=Default Availability=Available
// [2] ushort & unicode()

/*
Returns a Unicode representation of the string. The result remains valid until the string is modified.

Note: The returned string may not be '\0'-terminated. Use size() to determine the length of the array.

See also setUnicode(), utf16(), and fromRawData().
*/
func (this *QCharRef) Unicode1() uint16 {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRef7unicodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint16(rv)
}

func DeleteQCharRef(this *QCharRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QCharRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10191() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
