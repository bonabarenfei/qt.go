package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQuickItem struct {
	*qtcore.QObject
	*qtqml.QQmlParserStatus
}

func (this *QQuickItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickItem) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QQmlParserStatus = qtqml.NewQQmlParserStatusFromPointer(cthis)
}
func NewQQuickItemFromPointer(cthis unsafe.Pointer) *QQuickItem {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := qtqml.NewQQmlParserStatusFromPointer(cthis)
	return &QQuickItem{bcthis0, bcthis1}
}
func (*QQuickItem) NewFromPointer(cthis unsafe.Pointer) *QQuickItem {
	return NewQQuickItemFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickitem.h:98
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:199
// index:0
// Public
// void QQuickItem(QQuickItem *)
func NewQQuickItem(parent *QQuickItem /*777 QQuickItem **/) *QQuickItem {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItemC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickitem.h:200
// index:0
// Public virtual
// void ~QQuickItem()
func DeleteQQuickItem(*QQuickItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:202
// index:0
// Public
// QQuickWindow * window()
func (this *QQuickItem) Window() *QQuickWindow /*777 QQuickWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:203
// index:0
// Public
// QQuickItem * parentItem()
func (this *QQuickItem) ParentItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem10parentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:204
// index:0
// Public
// void setParentItem(QQuickItem *)
func (this *QQuickItem) SetParentItem(parent *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13setParentItemEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:205
// index:0
// Public
// void stackBefore(const QQuickItem *)
func (this *QQuickItem) StackBefore(arg0 *QQuickItem /*777 const QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11stackBeforeEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:206
// index:0
// Public
// void stackAfter(const QQuickItem *)
func (this *QQuickItem) StackAfter(arg0 *QQuickItem /*777 const QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10stackAfterEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:208
// index:0
// Public
// QRectF childrenRect()
func (this *QQuickItem) ChildrenRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12childrenRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:211
// index:0
// Public
// bool clip()
func (this *QQuickItem) Clip() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem4clipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:212
// index:0
// Public
// void setClip(_Bool)
func (this *QQuickItem) SetClip(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem7setClipEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:214
// index:0
// Public
// QString state()
func (this *QQuickItem) State() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem5stateEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:215
// index:0
// Public
// void setState(const QString &)
func (this *QQuickItem) SetState(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8setStateERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:217
// index:0
// Public
// qreal baselineOffset()
func (this *QQuickItem) BaselineOffset() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem14baselineOffsetEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:218
// index:0
// Public
// void setBaselineOffset(qreal)
func (this *QQuickItem) SetBaselineOffset(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem17setBaselineOffsetEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:222
// index:0
// Public
// qreal x()
func (this *QQuickItem) X() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem1xEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:223
// index:0
// Public
// qreal y()
func (this *QQuickItem) Y() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem1yEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:224
// index:0
// Public
// QPointF position()
func (this *QQuickItem) Position() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem8positionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:225
// index:0
// Public
// void setX(qreal)
func (this *QQuickItem) SetX(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem4setXEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:226
// index:0
// Public
// void setY(qreal)
func (this *QQuickItem) SetY(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem4setYEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:227
// index:0
// Public
// void setPosition(const QPointF &)
func (this *QQuickItem) SetPosition(arg0 *qtcore.QPointF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11setPositionERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:229
// index:0
// Public
// qreal width()
func (this *QQuickItem) Width() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem5widthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:230
// index:0
// Public
// void setWidth(qreal)
func (this *QQuickItem) SetWidth(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8setWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:231
// index:0
// Public
// void resetWidth()
func (this *QQuickItem) ResetWidth() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10resetWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:232
// index:0
// Public
// void setImplicitWidth(qreal)
func (this *QQuickItem) SetImplicitWidth(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16setImplicitWidthEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:233
// index:0
// Public
// qreal implicitWidth()
func (this *QQuickItem) ImplicitWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem13implicitWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:235
// index:0
// Public
// qreal height()
func (this *QQuickItem) Height() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem6heightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:236
// index:0
// Public
// void setHeight(qreal)
func (this *QQuickItem) SetHeight(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem9setHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:237
// index:0
// Public
// void resetHeight()
func (this *QQuickItem) ResetHeight() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11resetHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:238
// index:0
// Public
// void setImplicitHeight(qreal)
func (this *QQuickItem) SetImplicitHeight(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem17setImplicitHeightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:239
// index:0
// Public
// qreal implicitHeight()
func (this *QQuickItem) ImplicitHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem14implicitHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:241
// index:0
// Public
// QSizeF size()
func (this *QQuickItem) Size() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK10QQuickItem4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:242
// index:0
// Public
// void setSize(const QSizeF &)
func (this *QQuickItem) SetSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem7setSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:244
// index:0
// Public
// QQuickItem::TransformOrigin transformOrigin()
func (this *QQuickItem) TransformOrigin() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem15transformOriginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:245
// index:0
// Public
// void setTransformOrigin(enum QQuickItem::TransformOrigin)
func (this *QQuickItem) SetTransformOrigin(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem18setTransformOriginENS_15TransformOriginE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:246
// index:0
// Public
// QPointF transformOriginPoint()
func (this *QQuickItem) TransformOriginPoint() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem20transformOriginPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:247
// index:0
// Public
// void setTransformOriginPoint(const QPointF &)
func (this *QQuickItem) SetTransformOriginPoint(arg0 *qtcore.QPointF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem23setTransformOriginPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:249
// index:0
// Public
// qreal z()
func (this *QQuickItem) Z() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem1zEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:250
// index:0
// Public
// void setZ(qreal)
func (this *QQuickItem) SetZ(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem4setZEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:252
// index:0
// Public
// qreal rotation()
func (this *QQuickItem) Rotation() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem8rotationEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:253
// index:0
// Public
// void setRotation(qreal)
func (this *QQuickItem) SetRotation(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11setRotationEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:254
// index:0
// Public
// qreal scale()
func (this *QQuickItem) Scale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem5scaleEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:255
// index:0
// Public
// void setScale(qreal)
func (this *QQuickItem) SetScale(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8setScaleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:257
// index:0
// Public
// qreal opacity()
func (this *QQuickItem) Opacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem7opacityEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:258
// index:0
// Public
// void setOpacity(qreal)
func (this *QQuickItem) SetOpacity(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10setOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:260
// index:0
// Public
// bool isVisible()
func (this *QQuickItem) IsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:261
// index:0
// Public
// void setVisible(_Bool)
func (this *QQuickItem) SetVisible(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:263
// index:0
// Public
// bool isEnabled()
func (this *QQuickItem) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:264
// index:0
// Public
// void setEnabled(_Bool)
func (this *QQuickItem) SetEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:266
// index:0
// Public
// bool smooth()
func (this *QQuickItem) Smooth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem6smoothEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:267
// index:0
// Public
// void setSmooth(_Bool)
func (this *QQuickItem) SetSmooth(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem9setSmoothEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:269
// index:0
// Public
// bool activeFocusOnTab()
func (this *QQuickItem) ActiveFocusOnTab() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem16activeFocusOnTabEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:270
// index:0
// Public
// void setActiveFocusOnTab(_Bool)
func (this *QQuickItem) SetActiveFocusOnTab(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem19setActiveFocusOnTabEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:272
// index:0
// Public
// bool antialiasing()
func (this *QQuickItem) Antialiasing() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem12antialiasingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:273
// index:0
// Public
// void setAntialiasing(_Bool)
func (this *QQuickItem) SetAntialiasing(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15setAntialiasingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:274
// index:0
// Public
// void resetAntialiasing()
func (this *QQuickItem) ResetAntialiasing() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem17resetAntialiasingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:276
// index:0
// Public
// QQuickItem::Flags flags()
func (this *QQuickItem) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:277
// index:0
// Public
// void setFlag(enum QQuickItem::Flag, _Bool)
func (this *QQuickItem) SetFlag(flag int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem7setFlagENS_4FlagEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:278
// index:0
// Public
// void setFlags(QQuickItem::Flags)
func (this *QQuickItem) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8setFlagsE6QFlagsINS_4FlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:280
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QQuickItem) BoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:281
// index:0
// Public virtual
// QRectF clipRect()
func (this *QQuickItem) ClipRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem8clipRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:283
// index:0
// Public
// bool hasActiveFocus()
func (this *QQuickItem) HasActiveFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem14hasActiveFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:284
// index:0
// Public
// bool hasFocus()
func (this *QQuickItem) HasFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem8hasFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:285
// index:0
// Public
// void setFocus(_Bool)
func (this *QQuickItem) SetFocus(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8setFocusEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:286
// index:1
// Public
// void setFocus(_Bool, Qt::FocusReason)
func (this *QQuickItem) SetFocus_1(focus bool, reason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8setFocusEbN2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), focus, reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:287
// index:0
// Public
// bool isFocusScope()
func (this *QQuickItem) IsFocusScope() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem12isFocusScopeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:288
// index:0
// Public
// QQuickItem * scopedFocusItem()
func (this *QQuickItem) ScopedFocusItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem15scopedFocusItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:290
// index:0
// Public
// bool isAncestorOf(const QQuickItem *)
func (this *QQuickItem) IsAncestorOf(child *QQuickItem /*777 const QQuickItem **/) bool {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem12isAncestorOfEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:292
// index:0
// Public
// Qt::MouseButtons acceptedMouseButtons()
func (this *QQuickItem) AcceptedMouseButtons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem20acceptedMouseButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:294
// index:0
// Public
// bool acceptHoverEvents()
func (this *QQuickItem) AcceptHoverEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem17acceptHoverEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:295
// index:0
// Public
// void setAcceptHoverEvents(_Bool)
func (this *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem20setAcceptHoverEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:296
// index:0
// Public
// bool acceptTouchEvents()
func (this *QQuickItem) AcceptTouchEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem17acceptTouchEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:297
// index:0
// Public
// void setAcceptTouchEvents(_Bool)
func (this *QQuickItem) SetAcceptTouchEvents(accept bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem20setAcceptTouchEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), accept)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:300
// index:0
// Public
// QCursor cursor()
func (this *QQuickItem) Cursor() *qtgui.QCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem6cursorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:301
// index:0
// Public
// void setCursor(const QCursor &)
func (this *QQuickItem) SetCursor(cursor *qtgui.QCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem9setCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:302
// index:0
// Public
// void unsetCursor()
func (this *QQuickItem) UnsetCursor() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11unsetCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:305
// index:0
// Public
// bool isUnderMouse()
func (this *QQuickItem) IsUnderMouse() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem12isUnderMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:306
// index:0
// Public
// void grabMouse()
func (this *QQuickItem) GrabMouse() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem9grabMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:307
// index:0
// Public
// void ungrabMouse()
func (this *QQuickItem) UngrabMouse() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11ungrabMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:308
// index:0
// Public
// bool keepMouseGrab()
func (this *QQuickItem) KeepMouseGrab() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem13keepMouseGrabEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:309
// index:0
// Public
// void setKeepMouseGrab(_Bool)
func (this *QQuickItem) SetKeepMouseGrab(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16setKeepMouseGrabEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:310
// index:0
// Public
// bool filtersChildMouseEvents()
func (this *QQuickItem) FiltersChildMouseEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem23filtersChildMouseEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:311
// index:0
// Public
// void setFiltersChildMouseEvents(_Bool)
func (this *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem26setFiltersChildMouseEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:314
// index:0
// Public
// void ungrabTouchPoints()
func (this *QQuickItem) UngrabTouchPoints() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem17ungrabTouchPointsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:315
// index:0
// Public
// bool keepTouchGrab()
func (this *QQuickItem) KeepTouchGrab() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem13keepTouchGrabEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:316
// index:0
// Public
// void setKeepTouchGrab(_Bool)
func (this *QQuickItem) SetKeepTouchGrab(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16setKeepTouchGrabEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:319
// index:0
// Public
// bool grabToImage(const QJSValue &, const QSize &)
func (this *QQuickItem) GrabToImage(callback *qtqml.QJSValue, targetSize *qtcore.QSize) bool {
	var convArg0 = callback.GetCthis()
	var convArg1 = targetSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11grabToImageERK8QJSValueRK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:322
// index:0
// Public virtual
// bool contains(const QPointF &)
func (this *QQuickItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:324
// index:0
// Public
// QTransform itemTransform(QQuickItem *, _Bool *)
func (this *QQuickItem) ItemTransform(arg0 *QQuickItem /*777 QQuickItem **/, arg1 unsafe.Pointer /*666*/) *qtgui.QTransform /*123*/ {
	var convArg0 = arg0.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem13itemTransformEPS_Pb", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:325
// index:0
// Public
// QPointF mapToItem(const QQuickItem *, const QPointF &)
func (this *QQuickItem) MapToItem(item *QQuickItem /*777 const QQuickItem **/, point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem9mapToItemEPKS_RK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:326
// index:0
// Public
// QPointF mapToScene(const QPointF &)
func (this *QQuickItem) MapToScene(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem10mapToSceneERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:327
// index:0
// Public
// QPointF mapToGlobal(const QPointF &)
func (this *QQuickItem) MapToGlobal(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem11mapToGlobalERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:328
// index:0
// Public
// QRectF mapRectToItem(const QQuickItem *, const QRectF &)
func (this *QQuickItem) MapRectToItem(item *QQuickItem /*777 const QQuickItem **/, rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem13mapRectToItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:329
// index:0
// Public
// QRectF mapRectToScene(const QRectF &)
func (this *QQuickItem) MapRectToScene(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem14mapRectToSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:330
// index:0
// Public
// QPointF mapFromItem(const QQuickItem *, const QPointF &)
func (this *QQuickItem) MapFromItem(item *QQuickItem /*777 const QQuickItem **/, point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem11mapFromItemEPKS_RK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:331
// index:0
// Public
// QPointF mapFromScene(const QPointF &)
func (this *QQuickItem) MapFromScene(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem12mapFromSceneERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:332
// index:0
// Public
// QPointF mapFromGlobal(const QPointF &)
func (this *QQuickItem) MapFromGlobal(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem13mapFromGlobalERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:333
// index:0
// Public
// QRectF mapRectFromItem(const QQuickItem *, const QRectF &)
func (this *QQuickItem) MapRectFromItem(item *QQuickItem /*777 const QQuickItem **/, rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem15mapRectFromItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:334
// index:0
// Public
// QRectF mapRectFromScene(const QRectF &)
func (this *QQuickItem) MapRectFromScene(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem16mapRectFromSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:336
// index:0
// Public
// void polish()
func (this *QQuickItem) Polish() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem6polishEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:342
// index:0
// Public
// void forceActiveFocus()
func (this *QQuickItem) ForceActiveFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16forceActiveFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:343
// index:1
// Public
// void forceActiveFocus(Qt::FocusReason)
func (this *QQuickItem) ForceActiveFocus_1(reason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16forceActiveFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:344
// index:0
// Public
// QQuickItem * nextItemInFocusChain(_Bool)
func (this *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem20nextItemInFocusChainEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), forward)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:345
// index:0
// Public
// QQuickItem * childAt(qreal, qreal)
func (this *QQuickItem) ChildAt(x float64, y float64) *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem7childAtEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:348
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QQuickItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:358
// index:0
// Public virtual
// bool isTextureProvider()
func (this *QQuickItem) IsTextureProvider() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem17isTextureProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:359
// index:0
// Public virtual
// QSGTextureProvider * textureProvider()
func (this *QQuickItem) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem15textureProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:362
// index:0
// Public
// void update()
func (this *QQuickItem) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:365
// index:0
// Public
// void childrenRectChanged(const QRectF &)
func (this *QQuickItem) ChildrenRectChanged(arg0 *qtcore.QRectF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem19childrenRectChangedERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:366
// index:0
// Public
// void baselineOffsetChanged(qreal)
func (this *QQuickItem) BaselineOffsetChanged(arg0 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem21baselineOffsetChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:367
// index:0
// Public
// void stateChanged(const QString &)
func (this *QQuickItem) StateChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12stateChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:368
// index:0
// Public
// void focusChanged(_Bool)
func (this *QQuickItem) FocusChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12focusChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:369
// index:0
// Public
// void activeFocusChanged(_Bool)
func (this *QQuickItem) ActiveFocusChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem18activeFocusChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:370
// index:0
// Public
// void activeFocusOnTabChanged(_Bool)
func (this *QQuickItem) ActiveFocusOnTabChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem23activeFocusOnTabChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:371
// index:0
// Public
// void parentChanged(QQuickItem *)
func (this *QQuickItem) ParentChanged(arg0 *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13parentChangedEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:372
// index:0
// Public
// void transformOriginChanged(enum QQuickItem::TransformOrigin)
func (this *QQuickItem) TransformOriginChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem22transformOriginChangedENS_15TransformOriginE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:373
// index:0
// Public
// void smoothChanged(_Bool)
func (this *QQuickItem) SmoothChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13smoothChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:374
// index:0
// Public
// void antialiasingChanged(_Bool)
func (this *QQuickItem) AntialiasingChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem19antialiasingChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:375
// index:0
// Public
// void clipChanged(_Bool)
func (this *QQuickItem) ClipChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem11clipChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:376
// index:0
// Public
// void windowChanged(QQuickWindow *)
func (this *QQuickItem) WindowChanged(window *QQuickWindow /*777 QQuickWindow **/) {
	var convArg0 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13windowChangedEP12QQuickWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:378
// index:0
// Public
// void childrenChanged()
func (this *QQuickItem) ChildrenChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15childrenChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:379
// index:0
// Public
// void opacityChanged()
func (this *QQuickItem) OpacityChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14opacityChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:380
// index:0
// Public
// void enabledChanged()
func (this *QQuickItem) EnabledChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14enabledChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:381
// index:0
// Public
// void visibleChanged()
func (this *QQuickItem) VisibleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14visibleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:382
// index:0
// Public
// void visibleChildrenChanged()
func (this *QQuickItem) VisibleChildrenChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem22visibleChildrenChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:383
// index:0
// Public
// void rotationChanged()
func (this *QQuickItem) RotationChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15rotationChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:384
// index:0
// Public
// void scaleChanged()
func (this *QQuickItem) ScaleChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12scaleChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:386
// index:0
// Public
// void xChanged()
func (this *QQuickItem) XChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8xChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:387
// index:0
// Public
// void yChanged()
func (this *QQuickItem) YChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8yChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:388
// index:0
// Public
// void widthChanged()
func (this *QQuickItem) WidthChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12widthChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:389
// index:0
// Public
// void heightChanged()
func (this *QQuickItem) HeightChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13heightChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:390
// index:0
// Public
// void zChanged()
func (this *QQuickItem) ZChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem8zChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:391
// index:0
// Public
// void implicitWidthChanged()
func (this *QQuickItem) ImplicitWidthChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem20implicitWidthChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:392
// index:0
// Public
// void implicitHeightChanged()
func (this *QQuickItem) ImplicitHeightChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem21implicitHeightChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:395
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QQuickItem) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:397
// index:0
// Protected
// bool isComponentComplete()
func (this *QQuickItem) IsComponentComplete() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem19isComponentCompleteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:404
// index:0
// Protected
// bool widthValid()
func (this *QQuickItem) WidthValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem10widthValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:405
// index:0
// Protected
// bool heightValid()
func (this *QQuickItem) HeightValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickItem11heightValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:406
// index:0
// Protected
// void setImplicitSize(qreal, qreal)
func (this *QQuickItem) SetImplicitSize(arg0 float64, arg1 float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15setImplicitSizeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:408
// index:0
// Protected virtual
// void classBegin()
func (this *QQuickItem) ClassBegin() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10classBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:409
// index:0
// Protected virtual
// void componentComplete()
func (this *QQuickItem) ComponentComplete() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem17componentCompleteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:411
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QQuickItem) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:412
// index:0
// Protected virtual
// void keyReleaseEvent(QKeyEvent *)
func (this *QQuickItem) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:414
// index:0
// Protected virtual
// void inputMethodEvent(QInputMethodEvent *)
func (this *QQuickItem) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:416
// index:0
// Protected virtual
// void focusInEvent(QFocusEvent *)
func (this *QQuickItem) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:417
// index:0
// Protected virtual
// void focusOutEvent(QFocusEvent *)
func (this *QQuickItem) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:418
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QQuickItem) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:419
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QQuickItem) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:420
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickItem) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:421
// index:0
// Protected virtual
// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickItem) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:422
// index:0
// Protected virtual
// void mouseUngrabEvent()
func (this *QQuickItem) MouseUngrabEvent() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16mouseUngrabEventEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:423
// index:0
// Protected virtual
// void touchUngrabEvent()
func (this *QQuickItem) TouchUngrabEvent() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16touchUngrabEventEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:425
// index:0
// Protected virtual
// void wheelEvent(QWheelEvent *)
func (this *QQuickItem) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:427
// index:0
// Protected virtual
// void touchEvent(QTouchEvent *)
func (this *QQuickItem) TouchEvent(event *qtgui.QTouchEvent /*777 QTouchEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem10touchEventEP11QTouchEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:428
// index:0
// Protected virtual
// void hoverEnterEvent(QHoverEvent *)
func (this *QQuickItem) HoverEnterEvent(event *qtgui.QHoverEvent /*777 QHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15hoverEnterEventEP11QHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:429
// index:0
// Protected virtual
// void hoverMoveEvent(QHoverEvent *)
func (this *QQuickItem) HoverMoveEvent(event *qtgui.QHoverEvent /*777 QHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14hoverMoveEventEP11QHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:430
// index:0
// Protected virtual
// void hoverLeaveEvent(QHoverEvent *)
func (this *QQuickItem) HoverLeaveEvent(event *qtgui.QHoverEvent /*777 QHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15hoverLeaveEventEP11QHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:432
// index:0
// Protected virtual
// void dragEnterEvent(QDragEnterEvent *)
func (this *QQuickItem) DragEnterEvent(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:433
// index:0
// Protected virtual
// void dragMoveEvent(QDragMoveEvent *)
func (this *QQuickItem) DragMoveEvent(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:434
// index:0
// Protected virtual
// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QQuickItem) DragLeaveEvent(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:435
// index:0
// Protected virtual
// void dropEvent(QDropEvent *)
func (this *QQuickItem) DropEvent(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:437
// index:0
// Protected virtual
// bool childMouseEventFilter(QQuickItem *, QEvent *)
func (this *QQuickItem) ChildMouseEventFilter(arg0 *QQuickItem /*777 QQuickItem **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem21childMouseEventFilterEPS_P6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:438
// index:0
// Protected virtual
// void windowDeactivateEvent()
func (this *QQuickItem) WindowDeactivateEvent() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem21windowDeactivateEventEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:440
// index:0
// Protected virtual
// void geometryChanged(const QRectF &, const QRectF &)
func (this *QQuickItem) GeometryChanged(newGeometry *qtcore.QRectF, oldGeometry *qtcore.QRectF) {
	var convArg0 = newGeometry.GetCthis()
	var convArg1 = oldGeometry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem15geometryChangedERK6QRectFS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:444
// index:0
// Protected virtual
// void releaseResources()
func (this *QQuickItem) ReleaseResources() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem16releaseResourcesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:445
// index:0
// Protected virtual
// void updatePolish()
func (this *QQuickItem) UpdatePolish() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickItem12updatePolishEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QQuickItem__Flag = int

const QQuickItem__ItemClipsChildrenToShape QQuickItem__Flag = 1
const QQuickItem__ItemAcceptsInputMethod QQuickItem__Flag = 2
const QQuickItem__ItemIsFocusScope QQuickItem__Flag = 4
const QQuickItem__ItemHasContents QQuickItem__Flag = 8
const QQuickItem__ItemAcceptsDrops QQuickItem__Flag = 16

type QQuickItem__ItemChange = int

const QQuickItem__ItemChildAddedChange QQuickItem__ItemChange = 0
const QQuickItem__ItemChildRemovedChange QQuickItem__ItemChange = 1
const QQuickItem__ItemSceneChange QQuickItem__ItemChange = 2
const QQuickItem__ItemVisibleHasChanged QQuickItem__ItemChange = 3
const QQuickItem__ItemParentHasChanged QQuickItem__ItemChange = 4
const QQuickItem__ItemOpacityHasChanged QQuickItem__ItemChange = 5
const QQuickItem__ItemActiveFocusHasChanged QQuickItem__ItemChange = 6
const QQuickItem__ItemRotationHasChanged QQuickItem__ItemChange = 7
const QQuickItem__ItemAntialiasingHasChanged QQuickItem__ItemChange = 8
const QQuickItem__ItemDevicePixelRatioHasChanged QQuickItem__ItemChange = 9
const QQuickItem__ItemEnabledHasChanged QQuickItem__ItemChange = 10

type QQuickItem__TransformOrigin = int

const QQuickItem__TopLeft QQuickItem__TransformOrigin = 0
const QQuickItem__Top QQuickItem__TransformOrigin = 1
const QQuickItem__TopRight QQuickItem__TransformOrigin = 2
const QQuickItem__Left QQuickItem__TransformOrigin = 3
const QQuickItem__Center QQuickItem__TransformOrigin = 4
const QQuickItem__Right QQuickItem__TransformOrigin = 5
const QQuickItem__BottomLeft QQuickItem__TransformOrigin = 6
const QQuickItem__Bottom QQuickItem__TransformOrigin = 7
const QQuickItem__BottomRight QQuickItem__TransformOrigin = 8

//  body block end
