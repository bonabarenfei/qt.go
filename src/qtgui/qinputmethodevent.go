//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QInputMethodEvent struct {
	*qtcore.QEvent
}

func (this *QInputMethodEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}
func NewQInputMethodEventFromPointer(cthis unsafe.Pointer) *QInputMethodEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputMethodEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:555
// index:0
// Public
// void QInputMethodEvent()
func NewQInputMethodEvent() *QInputMethodEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QInputMethodEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQInputMethodEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:557
// index:0
// Public virtual
// void ~QInputMethodEvent()
func DeleteQInputMethodEvent(*QInputMethodEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QInputMethodEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:559
// index:0
// Public
// void setCommitString(const class QString &, int, int)
func (this *QInputMethodEvent) SetCommitString(commitString *qtcore.QString, replaceFrom int, replaceLength int) {
	var convArg0 = commitString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QInputMethodEvent15setCommitStringERK7QStringii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &replaceFrom, &replaceLength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:560
// index:0
// Public inline
// const QList<QInputMethodEvent::Attribute> & attributes()
func (this *QInputMethodEvent) Attributes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent10attributesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:561
// index:0
// Public inline
// const QString & preeditString()
func (this *QInputMethodEvent) PreeditString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent13preeditStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:563
// index:0
// Public inline
// const QString & commitString()
func (this *QInputMethodEvent) CommitString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent12commitStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:564
// index:0
// Public inline
// int replacementStart()
func (this *QInputMethodEvent) ReplacementStart() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent16replacementStartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:565
// index:0
// Public inline
// int replacementLength()
func (this *QInputMethodEvent) ReplacementLength() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QInputMethodEvent17replacementLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
