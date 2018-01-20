//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QAccessibleTextInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleTextInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAccessibleTextInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleTextInterface {
	return &QAccessibleTextInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:523
// index:0
// Public virtual
// void ~QAccessibleTextInterface()
func DeleteQAccessibleTextInterface(*QAccessibleTextInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAccessibleTextInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:525
// index:0
// Public pure virtual
// void selection(int, int *, int *)
func (this *QAccessibleTextInterface) Selection(selectionIndex int, startOffset unsafe.Pointer, endOffset unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface9selectionEiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &selectionIndex, startOffset, endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:526
// index:0
// Public pure virtual
// int selectionCount()
func (this *QAccessibleTextInterface) SelectionCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface14selectionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:527
// index:0
// Public pure virtual
// void addSelection(int, int)
func (this *QAccessibleTextInterface) AddSelection(startOffset int, endOffset int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAccessibleTextInterface12addSelectionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:528
// index:0
// Public pure virtual
// void removeSelection(int)
func (this *QAccessibleTextInterface) RemoveSelection(selectionIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAccessibleTextInterface15removeSelectionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &selectionIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:529
// index:0
// Public pure virtual
// void setSelection(int, int, int)
func (this *QAccessibleTextInterface) SetSelection(selectionIndex int, startOffset int, endOffset int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAccessibleTextInterface12setSelectionEiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &selectionIndex, &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:532
// index:0
// Public pure virtual
// int cursorPosition()
func (this *QAccessibleTextInterface) CursorPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface14cursorPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:533
// index:0
// Public pure virtual
// void setCursorPosition(int)
func (this *QAccessibleTextInterface) SetCursorPosition(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAccessibleTextInterface17setCursorPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:536
// index:0
// Public pure virtual
// QString text(int, int)
func (this *QAccessibleTextInterface) Text(startOffset int, endOffset int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface4textEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &startOffset, &endOffset)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:537
// index:0
// Public virtual
// QString textBeforeOffset(int, class QAccessible::TextBoundaryType, int *, int *)
func (this *QAccessibleTextInterface) TextBeforeOffset(offset int, boundaryType int, startOffset unsafe.Pointer, endOffset unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface16textBeforeOffsetEiN11QAccessible16TextBoundaryTypeEPiS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offset, &boundaryType, startOffset, endOffset)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:539
// index:0
// Public virtual
// QString textAfterOffset(int, class QAccessible::TextBoundaryType, int *, int *)
func (this *QAccessibleTextInterface) TextAfterOffset(offset int, boundaryType int, startOffset unsafe.Pointer, endOffset unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface15textAfterOffsetEiN11QAccessible16TextBoundaryTypeEPiS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offset, &boundaryType, startOffset, endOffset)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:541
// index:0
// Public virtual
// QString textAtOffset(int, class QAccessible::TextBoundaryType, int *, int *)
func (this *QAccessibleTextInterface) TextAtOffset(offset int, boundaryType int, startOffset unsafe.Pointer, endOffset unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface12textAtOffsetEiN11QAccessible16TextBoundaryTypeEPiS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offset, &boundaryType, startOffset, endOffset)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:543
// index:0
// Public pure virtual
// int characterCount()
func (this *QAccessibleTextInterface) CharacterCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface14characterCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:546
// index:0
// Public pure virtual
// QRect characterRect(int)
func (this *QAccessibleTextInterface) CharacterRect(offset int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface13characterRectEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offset)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:547
// index:0
// Public pure virtual
// int offsetAtPoint(const class QPoint &)
func (this *QAccessibleTextInterface) OffsetAtPoint(point *qtcore.QPoint) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface13offsetAtPointERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:549
// index:0
// Public pure virtual
// void scrollToSubstring(int, int)
func (this *QAccessibleTextInterface) ScrollToSubstring(startIndex int, endIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAccessibleTextInterface17scrollToSubstringEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &startIndex, &endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:550
// index:0
// Public pure virtual
// QString attributes(int, int *, int *)
func (this *QAccessibleTextInterface) Attributes(offset int, startOffset unsafe.Pointer, endOffset unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAccessibleTextInterface10attributesEiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &offset, startOffset, endOffset)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
