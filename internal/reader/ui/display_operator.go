// Copyright (c) 2023 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package ui

import (
	"github.com/rivo/tview"

	"github.com/bow/neon/internal/reader/backend"
)

type DisplayOperator struct {
	focusStack tview.Primitive
}

//nolint:revive
func NewDisplayOperator() *DisplayOperator {
	do := DisplayOperator{}
	return &do
}

//nolint:revive
func (do *DisplayOperator) ClearStatusBar(d *Display) {
	panic("ClearStatusBar is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) FocusFeedsPane(d *Display) {
	panic("FocusFeedsPane is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) FocusEntriesPane(d *Display) {
	panic("FocusEntriesPane is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) FocusNextPane(d *Display) {
	panic("FocusNextPane is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) FocusPreviousPane(d *Display) {
	panic("FocusPreviousPane is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) FocusReadingPane(d *Display) {
	panic("FocusReadingPane is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) HideIntroPopup(d *Display) {
	panic("HideIntroPopup is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) NotifyInfof(d *Display, text string, a ...any) {
	panic("NotifyInfof is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) NotifyErr(d *Display, err error) {
	panic("NotifyErr is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) NotifyErrf(d *Display, text string, a ...any) {
	panic("NotifyErrf is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) NotifyWarnf(d *Display, text string, a ...any) {
	panic("NotifyWarnf is unimplemented")
}

func (do *DisplayOperator) ToggleAboutPopup(d *Display, b backend.Backend) {
	if name := do.frontPageName(d); name == aboutPageName {
		do.hidePopup(d, name)
	} else if name != introPageName {
		d.setAboutPopupText(b)
		do.showPopup(d, aboutPageName, name)
	}
}

//nolint:revive
func (do *DisplayOperator) ToggleFeedsInPane(d *Display, b backend.Backend) {
	panic("ToggleFeedsInPane is unimplemented")
}

func (do *DisplayOperator) ToggleHelpPopup(d *Display) {
	if name := do.frontPageName(d); name == helpPageName {
		do.hidePopup(d, name)
	} else {
		do.showPopup(d, helpPageName, name)
	}
}

func (do *DisplayOperator) ToggleIntroPopup(d *Display) {
	d.dimMainPage()
	d.root.ShowPage(introPageName)
}

//nolint:revive
func (do *DisplayOperator) ToggleStatsPopup(d *Display, b backend.Backend) {
	panic("ToggleStatsPopup is unimplemented")
}

//nolint:revive
func (do *DisplayOperator) ToggleStatusBar(d *Display) {
	panic("ToggleStatusBar is unimplemented")
}

func (do *DisplayOperator) UnfocusFront(d *Display) {
	name := do.frontPageName(d)
	if name == mainPageName || name == "" {
		d.inner.SetFocus(d.root)
	} else {
		do.hidePopup(d, name)
	}
}

func (do *DisplayOperator) frontPageName(d *Display) string {
	name, _ := d.root.GetFrontPage()
	return name
}

func (do *DisplayOperator) showPopup(d *Display, name string, currentFront string) {
	if currentFront == mainPageName {
		do.stashFocus(d)
	} else {
		d.root.HidePage(currentFront)
	}
	d.dimMainPage()
	d.root.ShowPage(name)
}

func (do *DisplayOperator) hidePopup(d *Display, name string) {
	d.root.HidePage(name)
	d.normalizeMainPage()
	if top := do.focusStack; top != nil {
		d.inner.SetFocus(top)
	}
	do.focusStack = nil
}

func (do *DisplayOperator) stashFocus(d *Display) {
	do.focusStack = d.inner.GetFocus()
}

// Ensure DisplayOperator implements Operator.
var _ Operator = new(DisplayOperator)