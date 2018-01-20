//  header block begin
// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QTreeWidgetItem struct {
	*qtrt.CObject
}

func (this *QTreeWidgetItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTreeWidgetItemFromPointer(cthis unsafe.Pointer) *QTreeWidgetItem {
	return &QTreeWidgetItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qtreewidget.h:67
// index:0
// Public
// void QTreeWidgetItem(int)
func NewQTreeWidgetItem(type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:68
// index:1
// Public
// void QTreeWidgetItem(const class QStringList &, int)
func NewQTreeWidgetItem_1(strings *qtcore.QStringList, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg0 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2ERK11QStringListi", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:69
// index:2
// Public
// void QTreeWidgetItem(class QTreeWidget *, int)
func NewQTreeWidgetItem_2(view unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgeti", ffiqt.FFI_TYPE_VOID, cthis, view, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:70
// index:3
// Public
// void QTreeWidgetItem(class QTreeWidget *, const class QStringList &, int)
func NewQTreeWidgetItem_3(view unsafe.Pointer, strings *qtcore.QStringList, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg1 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi", ffiqt.FFI_TYPE_VOID, cthis, view, convArg1, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:71
// index:4
// Public
// void QTreeWidgetItem(class QTreeWidget *, class QTreeWidgetItem *, int)
func NewQTreeWidgetItem_4(view unsafe.Pointer, after unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i", ffiqt.FFI_TYPE_VOID, cthis, view, after, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:72
// index:5
// Public
// void QTreeWidgetItem(class QTreeWidgetItem *, int)
func NewQTreeWidgetItem_5(parent unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_i", ffiqt.FFI_TYPE_VOID, cthis, parent, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:73
// index:6
// Public
// void QTreeWidgetItem(class QTreeWidgetItem *, const class QStringList &, int)
func NewQTreeWidgetItem_6(parent unsafe.Pointer, strings *qtcore.QStringList, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	var convArg1 = strings.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_RK11QStringListi", ffiqt.FFI_TYPE_VOID, cthis, parent, convArg1, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:74
// index:7
// Public
// void QTreeWidgetItem(class QTreeWidgetItem *, class QTreeWidgetItem *, int)
func NewQTreeWidgetItem_7(parent unsafe.Pointer, after unsafe.Pointer, type_ int) *QTreeWidgetItem {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemC2EPS_S0_i", ffiqt.FFI_TYPE_VOID, cthis, parent, after, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTreeWidgetItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtreewidget.h:76
// index:0
// Public virtual
// void ~QTreeWidgetItem()
func DeleteQTreeWidgetItem(*QTreeWidgetItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:78
// index:0
// Public virtual
// QTreeWidgetItem * clone()
func (this *QTreeWidgetItem) Clone() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5cloneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:80
// index:0
// Public inline
// QTreeWidget * treeWidget()
func (this *QTreeWidgetItem) TreeWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10treeWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:82
// index:0
// Public inline
// void setSelected(_Bool)
func (this *QTreeWidgetItem) SetSelected(select_ bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setSelectedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &select_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:83
// index:0
// Public inline
// bool isSelected()
func (this *QTreeWidgetItem) IsSelected() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:85
// index:0
// Public inline
// void setHidden(_Bool)
func (this *QTreeWidgetItem) SetHidden(hide bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem9setHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &hide)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:86
// index:0
// Public inline
// bool isHidden()
func (this *QTreeWidgetItem) IsHidden() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem8isHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:88
// index:0
// Public inline
// void setExpanded(_Bool)
func (this *QTreeWidgetItem) SetExpanded(expand bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setExpandedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &expand)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:89
// index:0
// Public inline
// bool isExpanded()
func (this *QTreeWidgetItem) IsExpanded() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isExpandedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:91
// index:0
// Public inline
// void setFirstColumnSpanned(_Bool)
func (this *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem21setFirstColumnSpannedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &span)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:92
// index:0
// Public inline
// bool isFirstColumnSpanned()
func (this *QTreeWidgetItem) IsFirstColumnSpanned() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem20isFirstColumnSpannedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:94
// index:0
// Public inline
// void setDisabled(_Bool)
func (this *QTreeWidgetItem) SetDisabled(disabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setDisabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &disabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:95
// index:0
// Public inline
// bool isDisabled()
func (this *QTreeWidgetItem) IsDisabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10isDisabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:98
// index:0
// Public
// void setChildIndicatorPolicy(class QTreeWidgetItem::ChildIndicatorPolicy)
func (this *QTreeWidgetItem) SetChildIndicatorPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem23setChildIndicatorPolicyENS_20ChildIndicatorPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:99
// index:0
// Public
// QTreeWidgetItem::ChildIndicatorPolicy childIndicatorPolicy()
func (this *QTreeWidgetItem) ChildIndicatorPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem20childIndicatorPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:101
// index:0
// Public
// Qt::ItemFlags flags()
func (this *QTreeWidgetItem) Flags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:104
// index:0
// Public inline
// QString text(int)
func (this *QTreeWidgetItem) Text(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4textEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:106
// index:0
// Public inline
// void setText(int, const class QString &)
func (this *QTreeWidgetItem) SetText(column int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:108
// index:0
// Public inline
// QIcon icon(int)
func (this *QTreeWidgetItem) Icon(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4iconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:110
// index:0
// Public inline
// void setIcon(int, const class QIcon &)
func (this *QTreeWidgetItem) SetIcon(column int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:112
// index:0
// Public inline
// QString statusTip(int)
func (this *QTreeWidgetItem) StatusTip(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9statusTipEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:114
// index:0
// Public inline
// void setStatusTip(int, const class QString &)
func (this *QTreeWidgetItem) SetStatusTip(column int, statusTip *qtcore.QString) {
	var convArg1 = statusTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setStatusTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:117
// index:0
// Public inline
// QString toolTip(int)
func (this *QTreeWidgetItem) ToolTip(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem7toolTipEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:119
// index:0
// Public inline
// void setToolTip(int, const class QString &)
func (this *QTreeWidgetItem) SetToolTip(column int, toolTip *qtcore.QString) {
	var convArg1 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem10setToolTipEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:123
// index:0
// Public inline
// QString whatsThis(int)
func (this *QTreeWidgetItem) WhatsThis(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9whatsThisEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:125
// index:0
// Public inline
// void setWhatsThis(int, const class QString &)
func (this *QTreeWidgetItem) SetWhatsThis(column int, whatsThis *qtcore.QString) {
	var convArg1 = whatsThis.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:128
// index:0
// Public inline
// QFont font(int)
func (this *QTreeWidgetItem) Font(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4fontEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:130
// index:0
// Public inline
// void setFont(int, const class QFont &)
func (this *QTreeWidgetItem) SetFont(column int, font *qtgui.QFont) {
	var convArg1 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setFontEiRK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:132
// index:0
// Public inline
// int textAlignment(int)
func (this *QTreeWidgetItem) TextAlignment(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem13textAlignmentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:134
// index:0
// Public inline
// void setTextAlignment(int, int)
func (this *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem16setTextAlignmentEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:137
// index:0
// Public inline
// QColor backgroundColor(int)
func (this *QTreeWidgetItem) BackgroundColor(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem15backgroundColorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:139
// index:0
// Public inline
// void setBackgroundColor(int, const class QColor &)
func (this *QTreeWidgetItem) SetBackgroundColor(column int, color *qtgui.QColor) {
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:142
// index:0
// Public inline
// QBrush background(int)
func (this *QTreeWidgetItem) Background(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10backgroundEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:144
// index:0
// Public inline
// void setBackground(int, const class QBrush &)
func (this *QTreeWidgetItem) SetBackground(column int, brush *qtgui.QBrush) {
	var convArg1 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:147
// index:0
// Public inline
// QColor textColor(int)
func (this *QTreeWidgetItem) TextColor(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9textColorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:149
// index:0
// Public inline
// void setTextColor(int, const class QColor &)
func (this *QTreeWidgetItem) SetTextColor(column int, color *qtgui.QColor) {
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setTextColorEiRK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:152
// index:0
// Public inline
// QBrush foreground(int)
func (this *QTreeWidgetItem) Foreground(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10foregroundEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:154
// index:0
// Public inline
// void setForeground(int, const class QBrush &)
func (this *QTreeWidgetItem) SetForeground(column int, brush *qtgui.QBrush) {
	var convArg1 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setForegroundEiRK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:157
// index:0
// Public inline
// Qt::CheckState checkState(int)
func (this *QTreeWidgetItem) CheckState(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10checkStateEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:159
// index:0
// Public inline
// void setCheckState(int, Qt::CheckState)
func (this *QTreeWidgetItem) SetCheckState(column int, state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem13setCheckStateEiN2Qt10CheckStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:162
// index:0
// Public inline
// QSize sizeHint(int)
func (this *QTreeWidgetItem) SizeHint(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem8sizeHintEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:164
// index:0
// Public inline
// void setSizeHint(int, const class QSize &)
func (this *QTreeWidgetItem) SetSizeHint(column int, size *qtcore.QSize) {
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11setSizeHintEiRK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:167
// index:0
// Public virtual
// QVariant data(int, int)
func (this *QTreeWidgetItem) Data(column int, role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4dataEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:168
// index:0
// Public virtual
// void setData(int, int, const class QVariant &)
func (this *QTreeWidgetItem) SetData(column int, role int, value *qtcore.QVariant) {
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem7setDataEiiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &role, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:173
// index:0
// Public virtual
// void read(class QDataStream &)
func (this *QTreeWidgetItem) Read(in *qtcore.QDataStream) {
	var convArg0 = in.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem4readER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:174
// index:0
// Public virtual
// void write(class QDataStream &)
func (this *QTreeWidgetItem) Write(out *qtcore.QDataStream) {
	var convArg0 = out.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5writeER11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:178
// index:0
// Public inline
// QTreeWidgetItem * parent()
func (this *QTreeWidgetItem) Parent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:179
// index:0
// Public inline
// QTreeWidgetItem * child(int)
func (this *QTreeWidgetItem) Child(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:185
// index:0
// Public inline
// int childCount()
func (this *QTreeWidgetItem) ChildCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:186
// index:0
// Public inline
// int columnCount()
func (this *QTreeWidgetItem) ColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:187
// index:0
// Public inline
// int indexOfChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) IndexOfChild(child unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem12indexOfChildEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:189
// index:0
// Public
// void addChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) AddChild(child unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem8addChildEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:190
// index:0
// Public
// void insertChild(int, class QTreeWidgetItem *)
func (this *QTreeWidgetItem) InsertChild(index int, child unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11insertChildEiPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:191
// index:0
// Public
// void removeChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) RemoveChild(child unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem11removeChildEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:192
// index:0
// Public
// QTreeWidgetItem * takeChild(int)
func (this *QTreeWidgetItem) TakeChild(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem9takeChildEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:196
// index:0
// Public
// QList<QTreeWidgetItem *> takeChildren()
func (this *QTreeWidgetItem) TakeChildren() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12takeChildrenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:198
// index:0
// Public inline
// int type()
func (this *QTreeWidgetItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTreeWidgetItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtreewidget.h:199
// index:0
// Public inline
// void sortChildren(int, Qt::SortOrder)
func (this *QTreeWidgetItem) SortChildren(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem12sortChildrenEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtreewidget.h:203
// index:0
// Protected
// void emitDataChanged()
func (this *QTreeWidgetItem) EmitDataChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTreeWidgetItem15emitDataChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
