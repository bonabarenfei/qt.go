//  header block begin
// /usr/include/qt/QtWidgets/qtoolbox.h
// #include <qtoolbox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
type QToolBox struct {
	*QFrame
}

func (this *QToolBox) GetCthis() unsafe.Pointer {
	return this.QFrame.GetCthis()
}
func NewQToolBoxFromPointer(cthis unsafe.Pointer) *QToolBox {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QToolBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qtoolbox.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QToolBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:61
// index:0
// Public virtual
// void ~QToolBox()
func DeleteQToolBox(*QToolBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:63
// index:0
// Public
// int addItem(class QWidget *, const class QString &)
func (this *QToolBox) AddItem(widget unsafe.Pointer, text *qtcore.QString) interface{} {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:64
// index:1
// Public
// int addItem(class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) AddItem_1(widget unsafe.Pointer, icon *qtgui.QIcon, text *qtcore.QString) interface{} {
	var convArg1 = icon.GetCthis()
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:65
// index:0
// Public
// int insertItem(int, class QWidget *, const class QString &)
func (this *QToolBox) InsertItem(index int, widget unsafe.Pointer, text *qtcore.QString) interface{} {
	var convArg2 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:66
// index:1
// Public
// int insertItem(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) InsertItem_1(index int, widget unsafe.Pointer, icon *qtgui.QIcon, text *qtcore.QString) interface{} {
	var convArg2 = icon.GetCthis()
	var convArg3 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:68
// index:0
// Public
// void removeItem(int)
func (this *QToolBox) RemoveItem(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox10removeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:70
// index:0
// Public
// void setItemEnabled(int, _Bool)
func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:71
// index:0
// Public
// bool isItemEnabled(int)
func (this *QToolBox) IsItemEnabled(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13isItemEnabledEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:73
// index:0
// Public
// void setItemText(int, const class QString &)
func (this *QToolBox) SetItemText(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:74
// index:0
// Public
// QString itemText(int)
func (this *QToolBox) ItemText(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:76
// index:0
// Public
// void setItemIcon(int, const class QIcon &)
func (this *QToolBox) SetItemIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:77
// index:0
// Public
// QIcon itemIcon(int)
func (this *QToolBox) ItemIcon(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox8itemIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:80
// index:0
// Public
// void setItemToolTip(int, const class QString &)
func (this *QToolBox) SetItemToolTip(index int, toolTip *qtcore.QString) {
	var convArg1 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14setItemToolTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:81
// index:0
// Public
// QString itemToolTip(int)
func (this *QToolBox) ItemToolTip(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox11itemToolTipEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:84
// index:0
// Public
// int currentIndex()
func (this *QToolBox) CurrentIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:85
// index:0
// Public
// QWidget * currentWidget()
func (this *QToolBox) CurrentWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:86
// index:0
// Public
// QWidget * widget(int)
func (this *QToolBox) Widget(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:87
// index:0
// Public
// int indexOf(class QWidget *)
func (this *QToolBox) IndexOf(widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:88
// index:0
// Public
// int count()
func (this *QToolBox) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QToolBox5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:91
// index:0
// Public
// void setCurrentIndex(int)
func (this *QToolBox) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:92
// index:0
// Public
// void setCurrentWidget(class QWidget *)
func (this *QToolBox) SetCurrentWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:95
// index:0
// Public
// void currentChanged(int)
func (this *QToolBox) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:98
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QToolBox) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtoolbox.h:99
// index:0
// Protected virtual
// void itemInserted(int)
func (this *QToolBox) ItemInserted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox12itemInsertedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:100
// index:0
// Protected virtual
// void itemRemoved(int)
func (this *QToolBox) ItemRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11itemRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:101
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QToolBox) ShowEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtoolbox.h:102
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QToolBox) ChangeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QToolBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
