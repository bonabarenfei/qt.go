package qtwidgets

// /usr/include/qt/QtWidgets/qfileiconprovider.h
// #include <qfileiconprovider.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QFileIconProvider struct {
	*qtrt.CObject
}
type QFileIconProvider_ITF interface {
	QFileIconProvider_PTR() *QFileIconProvider
}

func (ptr *QFileIconProvider) QFileIconProvider_PTR() *QFileIconProvider { return ptr }

func (this *QFileIconProvider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFileIconProvider) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFileIconProviderFromPointer(cthis unsafe.Pointer) *QFileIconProvider {
	return &QFileIconProvider{&qtrt.CObject{cthis}}
}
func (*QFileIconProvider) NewFromPointer(cthis unsafe.Pointer) *QFileIconProvider {
	return NewQFileIconProviderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileIconProvider()

/*
Constructs a file icon provider.
*/
func (*QFileIconProvider) NewForInherit() *QFileIconProvider {
	return NewQFileIconProvider()
}
func NewQFileIconProvider() *QFileIconProvider {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QFileIconProviderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFileIconProvider)
	return gothis
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileIconProvider()

/*

 */
func DeleteQFileIconProvider(this *QFileIconProvider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QFileIconProviderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QIcon icon(QFileIconProvider::IconType) const

/*
Returns an icon set for the given type.
*/
func (this *QFileIconProvider) Icon(type_ int) *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QFileIconProvider4iconENS_8IconTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:66
// index:1
// Public virtual Visibility=Default Availability=Available
// [8] QIcon icon(const QFileInfo &) const

/*
Returns an icon set for the given type.
*/
func (this *QFileIconProvider) Icon1(info qtcore.QFileInfo_ITF) *qtgui.QIcon /*123*/ {
	var convArg0 unsafe.Pointer
	if info != nil && info.QFileInfo_PTR() != nil {
		convArg0 = info.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QFileIconProvider4iconERK9QFileInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString type(const QFileInfo &) const

/*
Returns the type of the file described by info.
*/
func (this *QFileIconProvider) Type(info qtcore.QFileInfo_ITF) string {
	var convArg0 unsafe.Pointer
	if info != nil && info.QFileInfo_PTR() != nil {
		convArg0 = info.QFileInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QFileIconProvider4typeERK9QFileInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QFileIconProvider::Options)

/*
Sets options that affect the icon provider.

This function was introduced in  Qt 5.2.

See also options().
*/
func (this *QFileIconProvider) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QFileIconProvider10setOptionsE6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfileiconprovider.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileIconProvider::Options options() const

/*
Returns all the options that affect the icon provider. By default, all options are disabled.

This function was introduced in  Qt 5.2.

See also setOptions().
*/
func (this *QFileIconProvider) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QFileIconProvider7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*
ConstantValue
QFileIconProvider::Computer0
QFileIconProvider::Desktop1
QFileIconProvider::Trashcan2
QFileIconProvider::Network3
QFileIconProvider::Drive4
QFileIconProvider::Folder5
QFileIconProvider::File6

*/
type QFileIconProvider__IconType = int

//
const QFileIconProvider__Computer QFileIconProvider__IconType = 0

//
const QFileIconProvider__Desktop QFileIconProvider__IconType = 1

//
const QFileIconProvider__Trashcan QFileIconProvider__IconType = 2

//
const QFileIconProvider__Network QFileIconProvider__IconType = 3

//
const QFileIconProvider__Drive QFileIconProvider__IconType = 4

//
const QFileIconProvider__Folder QFileIconProvider__IconType = 5

//
const QFileIconProvider__File QFileIconProvider__IconType = 6

func (this *QFileIconProvider) IconTypeItemName(val int) string {
	switch val {
	case QFileIconProvider__Computer: // 0
		return "Computer"
	case QFileIconProvider__Desktop: // 1
		return "Desktop"
	case QFileIconProvider__Trashcan: // 2
		return "Trashcan"
	case QFileIconProvider__Network: // 3
		return "Network"
	case QFileIconProvider__Drive: // 4
		return "Drive"
	case QFileIconProvider__Folder: // 5
		return "Folder"
	case QFileIconProvider__File: // 6
		return "File"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFileIconProvider_IconTypeItemName(val int) string {
	var nilthis *QFileIconProvider
	return nilthis.IconTypeItemName(val)
}

/*


 */
type QFileIconProvider__Option = int

//
const QFileIconProvider__DontUseCustomDirectoryIcons QFileIconProvider__Option = 1

func (this *QFileIconProvider) OptionItemName(val int) string {
	switch val {
	case QFileIconProvider__DontUseCustomDirectoryIcons: // 1
		return "DontUseCustomDirectoryIcons"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFileIconProvider_OptionItemName(val int) string {
	var nilthis *QFileIconProvider
	return nilthis.OptionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11127() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
