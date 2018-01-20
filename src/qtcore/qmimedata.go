//  header block begin
// /usr/include/qt/QtCore/qmimedata.h
// #include <qmimedata.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QMimeData struct {
	*QObject
}

func (this *QMimeData) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQMimeDataFromPointer(cthis unsafe.Pointer) *QMimeData {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QMimeData{bcthis0}
}

// /usr/include/qt/QtCore/qmimedata.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMimeData) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:56
// index:0
// Public
// void QMimeData()
func NewQMimeData() *QMimeData {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeDataC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMimeDataFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmimedata.h:57
// index:0
// Public virtual
// void ~QMimeData()
func DeleteQMimeData(*QMimeData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:59
// index:0
// Public
// QList<QUrl> urls()
func (this *QMimeData) Urls() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4urlsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:61
// index:0
// Public
// bool hasUrls()
func (this *QMimeData) HasUrls() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasUrlsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:63
// index:0
// Public
// QString text()
func (this *QMimeData) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:64
// index:0
// Public
// void setText(const class QString &)
func (this *QMimeData) SetText(text *QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:65
// index:0
// Public
// bool hasText()
func (this *QMimeData) HasText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:67
// index:0
// Public
// QString html()
func (this *QMimeData) Html() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4htmlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:68
// index:0
// Public
// void setHtml(const class QString &)
func (this *QMimeData) SetHtml(html *QString) {
	var convArg0 = html.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:69
// index:0
// Public
// bool hasHtml()
func (this *QMimeData) HasHtml() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7hasHtmlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:71
// index:0
// Public
// QVariant imageData()
func (this *QMimeData) ImageData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9imageDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:72
// index:0
// Public
// void setImageData(const class QVariant &)
func (this *QMimeData) SetImageData(image *QVariant) {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12setImageDataERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:73
// index:0
// Public
// bool hasImage()
func (this *QMimeData) HasImage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData8hasImageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:75
// index:0
// Public
// QVariant colorData()
func (this *QMimeData) ColorData() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9colorDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:76
// index:0
// Public
// void setColorData(const class QVariant &)
func (this *QMimeData) SetColorData(color *QVariant) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12setColorDataERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:77
// index:0
// Public
// bool hasColor()
func (this *QMimeData) HasColor() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData8hasColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:79
// index:0
// Public
// QByteArray data(const class QString &)
func (this *QMimeData) Data(mimetype *QString) interface{} {
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData4dataERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:80
// index:0
// Public
// void setData(const class QString &, const class QByteArray &)
func (this *QMimeData) SetData(mimetype *QString, data *QByteArray) {
	var convArg0 = mimetype.GetCthis()
	var convArg1 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData7setDataERK7QStringRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:81
// index:0
// Public
// void removeFormat(const class QString &)
func (this *QMimeData) RemoveFormat(mimetype *QString) {
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData12removeFormatERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:83
// index:0
// Public virtual
// bool hasFormat(const class QString &)
func (this *QMimeData) HasFormat(mimetype *QString) interface{} {
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData9hasFormatERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:84
// index:0
// Public virtual
// QStringList formats()
func (this *QMimeData) Formats() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData7formatsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmimedata.h:86
// index:0
// Public
// void clear()
func (this *QMimeData) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMimeData5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmimedata.h:88
// index:0
// Protected virtual
// QVariant retrieveData(const class QString &, class QVariant::Type)
func (this *QMimeData) RetrieveData(mimetype *QString, preferredType int) interface{} {
	var convArg0 = mimetype.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &preferredType)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
