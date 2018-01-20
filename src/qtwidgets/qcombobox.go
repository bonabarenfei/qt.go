//  header block begin
// /usr/include/qt/QtWidgets/qcombobox.h
// #include <qcombobox.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
type QComboBox struct {
	*QWidget
}

func (this *QComboBox) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}
func NewQComboBoxFromPointer(cthis unsafe.Pointer) *QComboBox {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QComboBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qcombobox.h:61
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QComboBox) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:85
// index:0
// Public
// void QComboBox(class QWidget *)
func NewQComboBox(parent unsafe.Pointer) *QComboBox {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQComboBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcombobox.h:86
// index:0
// Public virtual
// void ~QComboBox()
func DeleteQComboBox(*QComboBox) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBoxD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:88
// index:0
// Public
// int maxVisibleItems()
func (this *QComboBox) MaxVisibleItems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15maxVisibleItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:89
// index:0
// Public
// void setMaxVisibleItems(int)
func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18setMaxVisibleItemsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxItems)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:91
// index:0
// Public
// int count()
func (this *QComboBox) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:92
// index:0
// Public
// void setMaxCount(int)
func (this *QComboBox) SetMaxCount(max int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setMaxCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &max)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:93
// index:0
// Public
// int maxCount()
func (this *QComboBox) MaxCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8maxCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:96
// index:0
// Public
// bool autoCompletion()
func (this *QComboBox) AutoCompletion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14autoCompletionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:97
// index:0
// Public
// void setAutoCompletion(_Bool)
func (this *QComboBox) SetAutoCompletion(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setAutoCompletionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:99
// index:0
// Public
// Qt::CaseSensitivity autoCompletionCaseSensitivity()
func (this *QComboBox) AutoCompletionCaseSensitivity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox29autoCompletionCaseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:100
// index:0
// Public
// void setAutoCompletionCaseSensitivity(Qt::CaseSensitivity)
func (this *QComboBox) SetAutoCompletionCaseSensitivity(sensitivity int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox32setAutoCompletionCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &sensitivity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:103
// index:0
// Public
// bool duplicatesEnabled()
func (this *QComboBox) DuplicatesEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox17duplicatesEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:104
// index:0
// Public
// void setDuplicatesEnabled(_Bool)
func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox20setDuplicatesEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:106
// index:0
// Public
// void setFrame(_Bool)
func (this *QComboBox) SetFrame(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setFrameEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:107
// index:0
// Public
// bool hasFrame()
func (this *QComboBox) HasFrame() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8hasFrameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:126
// index:0
// Public
// QComboBox::InsertPolicy insertPolicy()
func (this *QComboBox) InsertPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12insertPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:127
// index:0
// Public
// void setInsertPolicy(enum QComboBox::InsertPolicy)
func (this *QComboBox) SetInsertPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setInsertPolicyENS_12InsertPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:137
// index:0
// Public
// QComboBox::SizeAdjustPolicy sizeAdjustPolicy()
func (this *QComboBox) SizeAdjustPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16sizeAdjustPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:138
// index:0
// Public
// void setSizeAdjustPolicy(enum QComboBox::SizeAdjustPolicy)
func (this *QComboBox) SetSizeAdjustPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:139
// index:0
// Public
// int minimumContentsLength()
func (this *QComboBox) MinimumContentsLength() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox21minimumContentsLengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:140
// index:0
// Public
// void setMinimumContentsLength(int)
func (this *QComboBox) SetMinimumContentsLength(characters int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox24setMinimumContentsLengthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &characters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:141
// index:0
// Public
// QSize iconSize()
func (this *QComboBox) IconSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:142
// index:0
// Public
// void setIconSize(const class QSize &)
func (this *QComboBox) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:144
// index:0
// Public
// bool isEditable()
func (this *QComboBox) IsEditable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox10isEditableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:145
// index:0
// Public
// void setEditable(_Bool)
func (this *QComboBox) SetEditable(editable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &editable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:146
// index:0
// Public
// void setLineEdit(class QLineEdit *)
func (this *QComboBox) SetLineEdit(edit unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setLineEditEP9QLineEdit", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), edit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:147
// index:0
// Public
// QLineEdit * lineEdit()
func (this *QComboBox) LineEdit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8lineEditEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:149
// index:0
// Public
// void setValidator(const class QValidator *)
func (this *QComboBox) SetValidator(v unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setValidatorEPK10QValidator", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:150
// index:0
// Public
// const QValidator * validator()
func (this *QComboBox) Validator() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9validatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:154
// index:0
// Public
// void setCompleter(class QCompleter *)
func (this *QComboBox) SetCompleter(c unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12setCompleterEP10QCompleter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:155
// index:0
// Public
// QCompleter * completer()
func (this *QComboBox) Completer() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox9completerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:158
// index:0
// Public
// QAbstractItemDelegate * itemDelegate()
func (this *QComboBox) ItemDelegate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12itemDelegateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:159
// index:0
// Public
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QComboBox) SetItemDelegate(delegate unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:161
// index:0
// Public
// QAbstractItemModel * model()
func (this *QComboBox) Model() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:162
// index:0
// Public
// void setModel(class QAbstractItemModel *)
func (this *QComboBox) SetModel(model unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:164
// index:0
// Public
// QModelIndex rootModelIndex()
func (this *QComboBox) RootModelIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox14rootModelIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:165
// index:0
// Public
// void setRootModelIndex(const class QModelIndex &)
func (this *QComboBox) SetRootModelIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17setRootModelIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:167
// index:0
// Public
// int modelColumn()
func (this *QComboBox) ModelColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11modelColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:168
// index:0
// Public
// void setModelColumn(int)
func (this *QComboBox) SetModelColumn(visibleColumn int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setModelColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visibleColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:170
// index:0
// Public
// int currentIndex()
func (this *QComboBox) CurrentIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:171
// index:0
// Public
// QString currentText()
func (this *QComboBox) CurrentText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:172
// index:0
// Public
// QVariant currentData(int)
func (this *QComboBox) CurrentData(role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox11currentDataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:174
// index:0
// Public
// QString itemText(int)
func (this *QComboBox) ItemText(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemTextEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:175
// index:0
// Public
// QIcon itemIcon(int)
func (this *QComboBox) ItemIcon(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemIconEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:176
// index:0
// Public
// QVariant itemData(int, int)
func (this *QComboBox) ItemData(index int, role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8itemDataEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:178
// index:0
// Public inline
// void addItem(const class QString &, const class QVariant &)
func (this *QComboBox) AddItem(text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg0 = text.GetCthis()
	var convArg1 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:179
// index:1
// Public inline
// void addItem(const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) AddItem_1(icon *qtgui.QIcon, text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:181
// index:0
// Public inline
// void addItems(const class QStringList &)
func (this *QComboBox) AddItems(texts *qtcore.QStringList) {
	var convArg0 = texts.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox8addItemsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:184
// index:0
// Public inline
// void insertItem(int, const class QString &, const class QVariant &)
func (this *QComboBox) InsertItem(index int, text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg1 = text.GetCthis()
	var convArg2 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:185
// index:1
// Public
// void insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) InsertItem_1(index int, icon *qtgui.QIcon, text *qtcore.QString, userData *qtcore.QVariant) {
	var convArg1 = icon.GetCthis()
	var convArg2 = text.GetCthis()
	var convArg3 = userData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:187
// index:0
// Public
// void insertItems(int, const class QStringList &)
func (this *QComboBox) InsertItems(index int, texts *qtcore.QStringList) {
	var convArg1 = texts.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11insertItemsEiRK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:188
// index:0
// Public
// void insertSeparator(int)
func (this *QComboBox) InsertSeparator(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15insertSeparatorEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:190
// index:0
// Public
// void removeItem(int)
func (this *QComboBox) RemoveItem(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10removeItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:192
// index:0
// Public
// void setItemText(int, const class QString &)
func (this *QComboBox) SetItemText(index int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:193
// index:0
// Public
// void setItemIcon(int, const class QIcon &)
func (this *QComboBox) SetItemIcon(index int, icon *qtgui.QIcon) {
	var convArg1 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemIconEiRK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:194
// index:0
// Public
// void setItemData(int, const class QVariant &, int)
func (this *QComboBox) SetItemData(index int, value *qtcore.QVariant, role int) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setItemDataEiRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:196
// index:0
// Public
// QAbstractItemView * view()
func (this *QComboBox) View() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox4viewEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:197
// index:0
// Public
// void setView(class QAbstractItemView *)
func (this *QComboBox) SetView(itemView unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox7setViewEP17QAbstractItemView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), itemView)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:199
// index:0
// Public virtual
// QSize sizeHint()
func (this *QComboBox) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:200
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QComboBox) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:202
// index:0
// Public virtual
// void showPopup()
func (this *QComboBox) ShowPopup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showPopupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:203
// index:0
// Public virtual
// void hidePopup()
func (this *QComboBox) HidePopup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hidePopupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:205
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QComboBox) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:206
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QComboBox) InputMethodQuery(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:207
// index:1
// Public
// QVariant inputMethodQuery(Qt::InputMethodQuery, const class QVariant &)
func (this *QComboBox) InputMethodQuery_1(query int, argument *qtcore.QVariant) interface{} {
	var convArg1 = argument.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox16inputMethodQueryEN2Qt16InputMethodQueryERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &query, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qcombobox.h:210
// index:0
// Public
// void clear()
func (this *QComboBox) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:211
// index:0
// Public
// void clearEditText()
func (this *QComboBox) ClearEditText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13clearEditTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:212
// index:0
// Public
// void setEditText(const class QString &)
func (this *QComboBox) SetEditText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11setEditTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:213
// index:0
// Public
// void setCurrentIndex(int)
func (this *QComboBox) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:214
// index:0
// Public
// void setCurrentText(const class QString &)
func (this *QComboBox) SetCurrentText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox14setCurrentTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:217
// index:0
// Public
// void editTextChanged(const class QString &)
func (this *QComboBox) EditTextChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15editTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:218
// index:0
// Public
// void activated(int)
func (this *QComboBox) Activated(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:219
// index:1
// Public
// void activated(const class QString &)
func (this *QComboBox) Activated_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9activatedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:220
// index:0
// Public
// void highlighted(int)
func (this *QComboBox) Highlighted(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:221
// index:1
// Public
// void highlighted(const class QString &)
func (this *QComboBox) Highlighted_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11highlightedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:222
// index:0
// Public
// void currentIndexChanged(int)
func (this *QComboBox) CurrentIndexChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:223
// index:1
// Public
// void currentIndexChanged(const class QString &)
func (this *QComboBox) CurrentIndexChanged_1(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox19currentIndexChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:224
// index:0
// Public
// void currentTextChanged(const class QString &)
func (this *QComboBox) CurrentTextChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox18currentTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:227
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QComboBox) FocusInEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:228
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QComboBox) FocusOutEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:229
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QComboBox) ChangeEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:230
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QComboBox) ResizeEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:231
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QComboBox) PaintEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:232
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QComboBox) ShowEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:233
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QComboBox) HideEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:234
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QComboBox) MousePressEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:235
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QComboBox) MouseReleaseEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:236
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QComboBox) KeyPressEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:237
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QComboBox) KeyReleaseEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:239
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QComboBox) WheelEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:242
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QComboBox) ContextMenuEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:244
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QComboBox) InputMethodEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcombobox.h:245
// index:0
// Protected
// void initStyleOption(class QStyleOptionComboBox *)
func (this *QComboBox) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
