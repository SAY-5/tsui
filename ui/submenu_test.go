package ui

import "testing"

func TestSubmenuRenderWithNoRoom(t *testing.T) {
	// The main menu previews the selected item's submenu, so a terminal too
	// short to fit any item still renders it.
	var submenu Submenu
	submenu.SetItems([]SubmenuItem{
		&ToggleableSubmenuItem{LabeledSubmenuItem: LabeledSubmenuItem{Label: "None"}, IsActive: true},
		&DividerSubmenuItem{},
	})

	if rendered := submenu.Render(false, 0, 40); rendered == "" {
		t.Fatal("expected an overflow indicator")
	}
}
