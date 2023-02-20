package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createTextDialog(title, message string, icon fyne.Resource, parent fyne.Window) Dialog {
	d := newDialog(title, message, icon, nil, parent)

	d.dismiss = &widget.Button{Text: "OK",
		OnTapped: d.Hide,
	}
	d.create(newButtonList(d.dismiss))

	return d
}

// NewInformation creates a dialog over the specified window for user information.
// The title is used for the dialog window and message is the content.
// After creation you should call Show().
func NewInformation(title, message string, parent fyne.Window) Dialog {
	return createTextDialog(title, message, theme.InfoIcon(), parent)
}

// ShowInformation shows a dialog over the specified window for user information.
// The title is used for the dialog window and message is the content.
func ShowInformation(title, message string, parent fyne.Window) {
	NewInformation(title, message, parent).Show()
}

// NewError creates a dialog over the specified window for an application error.
// The message is extracted from the provided error (should not be nil).
// After creation you should call Show().
func NewError(err error, parent fyne.Window) Dialog {
	return createTextDialog("Error", err.Error(), theme.ErrorIcon(), parent)
}

// ShowError shows a dialog over the specified window for an application error.
// The message is extracted from the provided error (should not be nil).
func ShowError(err error, parent fyne.Window) {
	NewError(err, parent).Show()
}

// NewCustomError creates a dialog over the specified window for an application error using custom
// content.
// The button will have the dismiss text.
// The message is extracted from the provided error (should not be nil).
// The response callback is called on user action.
func NewCustomError(title, dismiss, message string, callback func(bool), parent fyne.Window) *ConfirmDialog {

	d := newDialog(title, message, theme.ErrorIcon(), callback, parent)

	confirm := &widget.Button{Text: dismiss, Icon: theme.ErrorIcon(), Importance: widget.DangerImportance,
		OnTapped: func() {
			d.hideWithResponse(true)
		},
	}
	d.create(newButtonList(confirm))

	return &ConfirmDialog{d, confirm}
}

// ShowCustomError shows a dialog over the specified window for an application error.
// The message is extracted from the provided error (should not be nil).
// The callback is executed when the user decides.
func ShowCustomError(title, dismiss, message string, callback func(bool), parent fyne.Window) {
	NewCustomError(title, dismiss, message, callback, parent).Show()
}
