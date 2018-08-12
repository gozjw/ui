// 12 december 2015

package ui

import (
	"unsafe"
)

// #include "ui.h"
// extern void doRadioButtonsOnSelected(uiRadioButtons *, void *);
import "C"

// RadioButtons is a Control that represents a set of checkable
// buttons from which exactly one may be chosen by the user.
type RadioButtons struct {
	ControlBase
	r	*C.uiRadioButtons
	onSelected	func(*RadioButtons)
}

// NewRadioButtons creates a new RadioButtons.
func NewRadioButtons() *RadioButtons {
	r := new(RadioButtons)

	r.r = C.uiNewRadioButtons()

	C.uiRadioButtonsOnSelected(r.r, C.doRadioButtonsOnSelected, nil)

	r.ControlBase = NewControlBase(r, uintptr(unsafe.Pointer(r.r)))
	return r
}

// Append adds the named button to the end of the RadioButtons.
// If this button is the first button, it is automatically selected.
func (r *RadioButtons) Append(text string) {
	ctext := C.CString(text)
	C.uiRadioButtonsAppend(r.r, ctext)
	freestr(ctext)
}

// Selected returns the index of the currently selected option in the
// RadioButtons, or -1 if there are no items.
func (r *RadioButtons) Selected() int {
	return int(C.uiRadioButtonsSelected(r.r))
}

// SetSelected sets the currently selected option in the RadioButtons
// to index.
func (r *RadioButtons) SetSelected(index int) {
	C.uiRadioButtonsSetSelected(r.r, C.int(index))
}

// OnSelected registers f to be run when the user selects an option in
// the RadioButtons. Only one function can be registered at a time.
func (r *RadioButtons) OnSelected(f func(*RadioButtons)) {
	r.onSelected = f
}

//export doRadioButtonsOnSelected
func doRadioButtonsOnSelected(rr *C.uiRadioButtons, data unsafe.Pointer) {
	r := ControlFromLibui(uintptr(unsafe.Pointer(rr))).(*RadioButtons)
	if r.onSelected != nil {
		r.onSelected(r)
	}
}
