//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 104
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
type QGraphicsSceneEvent struct {
	*qtcore.QEvent
}

func (this *QGraphicsSceneEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}
func NewQGraphicsSceneEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QGraphicsSceneEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:67
// index:0
// Public
// void QGraphicsSceneEvent(enum QEvent::Type)
func NewQGraphicsSceneEvent(type_ int) *QGraphicsSceneEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:68
// index:0
// Public virtual
// void ~QGraphicsSceneEvent()
func DeleteQGraphicsSceneEvent(*QGraphicsSceneEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:70
// index:0
// Public
// QWidget * widget()
func (this *QGraphicsSceneEvent) Widget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsSceneEvent6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:71
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QGraphicsSceneEvent) SetWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEvent9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:76
// index:0
// Protected inline
// QGraphicsSceneEventPrivate * d_func()
func (this *QGraphicsSceneEvent) D_func() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsSceneEvent6d_funcEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:76
// index:1
// Protected inline
// const QGraphicsSceneEventPrivate * d_func()
func (this *QGraphicsSceneEvent) D_func_1() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsSceneEvent6d_funcEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
