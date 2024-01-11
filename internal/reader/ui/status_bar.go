// Copyright (c) 2024 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package ui

import (
	"fmt"
	"time"

	"github.com/bow/neon/internal/entity"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const iconAllRead = "✔"

type statusBar struct {
	tview.Flex

	theme *Theme

	eventsWidget     *tview.TextView
	readStatusWidget *tview.TextView
	lastPullWidget   *tview.TextView
}

func newStatusBar(theme *Theme) *statusBar {

	var (
		eventsWidget     = tview.NewTextView().SetTextAlign(tview.AlignLeft)
		readStatusWidget = tview.NewTextView().SetTextAlign(tview.AlignCenter)
		lastPullWidget   = tview.NewTextView().SetTextAlign(tview.AlignRight)
	)

	quickStatusFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(readStatusWidget, 1, 0, false).
		AddItem(lastPullWidget, len(shortDateFormat)+1, 0, true)

	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn)

	bar := statusBar{
		Flex:             *flex,
		theme:            theme,
		eventsWidget:     eventsWidget,
		readStatusWidget: readStatusWidget,
		lastPullWidget:   lastPullWidget,
	}
	bar.AddItem(eventsWidget, 0, 1, false).
		AddItem(quickStatusFlex, len(shortDateFormat)+2, 1, false)
	bar.refreshColors()

	return &bar
}
func (b *statusBar) setStats(stats *entity.Stats) {
	if stats.NumFeeds < 1 {
		return
	}
	b.setLastPullTime(stats.LastPullTime)
	if stats.NumEntriesUnread == 0 {
		b.setAllRead()
	}
}

func (b *statusBar) refreshColors() {
	b.readStatusWidget.SetTextColor(b.theme.statusBarFG)
	b.lastPullWidget.SetTextColor(b.theme.statusBarFG)
}

func (b *statusBar) setAllRead() {
	b.readStatusWidget.SetText(iconAllRead)
}

func (b *statusBar) setLastPullTime(value *time.Time) {
	if value != nil {
		b.lastPullWidget.SetText(value.Local().Format(shortDateFormat))
	}
}

func (b *statusBar) showEvent(ev *event) {
	var color tcell.Color
	switch ev.level {
	case eventLevelInfo:
		color = b.theme.eventInfoFG
	case eventLevelWarn:
		color = b.theme.eventWarnFG
	case eventLevelErr:
		color = b.theme.eventErrFG
	default:
		panic(fmt.Sprintf("unsupported event level: %v", ev.level))
	}
	b.eventsWidget.SetTextColor(color).Clear()
	fmt.Fprintf(b.eventsWidget, "%s\n", ev.text)
}
