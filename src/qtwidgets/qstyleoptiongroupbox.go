//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleOptionGroupBox struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionGroupBox) GetCthis() unsafe.Pointer {
	return this.QStyleOptionComplex.GetCthis()
}
func NewQStyleOptionGroupBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionGroupBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionGroupBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:639
// index:0
// Public
// void QStyleOptionGroupBox()
func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {
	cthis := qtrt.Calloc(1, 256) // 120
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGroupBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:642
// index:1
// Protected
// void QStyleOptionGroupBox(int)
func NewQStyleOptionGroupBox_1(version int) *QStyleOptionGroupBox {
	cthis := qtrt.Calloc(1, 256) // 120
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionGroupBoxC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGroupBoxFromPointer(cthis)
	return gothis
}

//  body block end
