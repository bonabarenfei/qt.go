//  header block begin
// /usr/include/qt/QtWidgets/qscrollbar.h
// #include <qscrollbar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QScrollBar struct {
	*QAbstractSlider
}

func (this *QScrollBar) GetCthis() unsafe.Pointer {
	return this.QAbstractSlider.GetCthis()
}
func NewQScrollBarFromPointer(cthis unsafe.Pointer) *QScrollBar {
	bcthis0 := NewQAbstractSliderFromPointer(cthis)
	return &QScrollBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qscrollbar.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QScrollBar) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscrollbar.h:59
// index:0
// Public
// void QScrollBar(class QWidget *)
func NewQScrollBar(parent unsafe.Pointer) *QScrollBar {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:60
// index:1
// Public
// void QScrollBar(Qt::Orientation, class QWidget *)
func NewQScrollBar_1(arg0 int, parent unsafe.Pointer) *QScrollBar {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollbar.h:61
// index:0
// Public virtual
// void ~QScrollBar()
func DeleteQScrollBar(*QScrollBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:63
// index:0
// Public virtual
// QSize sizeHint()
func (this *QScrollBar) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscrollbar.h:64
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QScrollBar) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qscrollbar.h:68
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QScrollBar) WheelEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:70
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QScrollBar) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:71
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QScrollBar) MousePressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:72
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QScrollBar) MouseReleaseEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:73
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QScrollBar) MouseMoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:74
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QScrollBar) HideEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:75
// index:0
// Protected virtual
// void sliderChange(enum QAbstractSlider::SliderChange)
func (this *QScrollBar) SliderChange(change int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &change)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:77
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QScrollBar) ContextMenuEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollbar.h:79
// index:0
// Protected
// void initStyleOption(class QStyleOptionSlider *)
func (this *QScrollBar) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
