//  header block begin
// /usr/include/qt/QtWidgets/qstatusbar.h
// #include <qstatusbar.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QStatusBar struct {
	*QWidget
}

func (this *QStatusBar) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQStatusBarFromPointer(cthis unsafe.Pointer) *QStatusBar {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QStatusBar{bcthis0}
}

// /usr/include/qt/QtWidgets/qstatusbar.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStatusBar) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStatusBar10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstatusbar.h:59
// index:0
// Public
// void QStatusBar(class QWidget *)
func NewQStatusBar(parent unsafe.Pointer) *QStatusBar {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBarC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStatusBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstatusbar.h:60
// index:0
// Public virtual
// void ~QStatusBar()
func DeleteQStatusBar(*QStatusBar) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBarD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:62
// index:0
// Public
// void addWidget(class QWidget *, int)
func (this *QStatusBar) AddWidget(widget unsafe.Pointer, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar9addWidgetEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:63
// index:0
// Public
// int insertWidget(int, class QWidget *, int)
func (this *QStatusBar) InsertWidget(index int, widget unsafe.Pointer, stretch int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar12insertWidgetEiP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget, &stretch)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstatusbar.h:64
// index:0
// Public
// void addPermanentWidget(class QWidget *, int)
func (this *QStatusBar) AddPermanentWidget(widget unsafe.Pointer, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:65
// index:0
// Public
// int insertPermanentWidget(int, class QWidget *, int)
func (this *QStatusBar) InsertPermanentWidget(index int, widget unsafe.Pointer, stretch int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget, &stretch)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstatusbar.h:66
// index:0
// Public
// void removeWidget(class QWidget *)
func (this *QStatusBar) RemoveWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar12removeWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:68
// index:0
// Public
// void setSizeGripEnabled(_Bool)
func (this *QStatusBar) SetSizeGripEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar18setSizeGripEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:69
// index:0
// Public
// bool isSizeGripEnabled()
func (this *QStatusBar) IsSizeGripEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStatusBar17isSizeGripEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstatusbar.h:71
// index:0
// Public
// QString currentMessage()
func (this *QStatusBar) CurrentMessage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QStatusBar14currentMessageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstatusbar.h:74
// index:0
// Public
// void showMessage(const class QString &, int)
func (this *QStatusBar) ShowMessage(text *qtcore.QString, timeout int) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar11showMessageERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &timeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:75
// index:0
// Public
// void clearMessage()
func (this *QStatusBar) ClearMessage() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar12clearMessageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:79
// index:0
// Public
// void messageChanged(const class QString &)
func (this *QStatusBar) MessageChanged(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar14messageChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:82
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QStatusBar) ShowEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:83
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QStatusBar) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:84
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QStatusBar) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:87
// index:0
// Protected
// void reformat()
func (this *QStatusBar) Reformat() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar8reformatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:88
// index:0
// Protected
// void hideOrShow()
func (this *QStatusBar) HideOrShow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar10hideOrShowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstatusbar.h:89
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QStatusBar) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QStatusBar5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
