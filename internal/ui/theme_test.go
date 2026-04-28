package ui

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestDefaultTheme_ReturnsNonZeroStyles(t *testing.T) {
	theme := DefaultTheme()

	if theme.Title.GetBold() != true {
		t.Error("expected Title style to be bold")
	}
	if theme.Serving.GetBold() != true {
		t.Error("expected Serving style to be bold")
	}
	if theme.NotServing.GetBold() != true {
		t.Error("expected NotServing style to be bold")
	}
}

func TestDefaultTheme_ServingColor(t *testing.T) {
	theme := DefaultTheme()
	got := theme.Serving.GetForeground()
	want := lipgloss.Color("#00D26A")
	if got != want {
		t.Errorf("Serving foreground = %v, want %v", got, want)
	}
}

func TestDefaultTheme_NotServingColor(t *testing.T) {
	theme := DefaultTheme()
	got := theme.NotServing.GetForeground()
	want := lipgloss.Color("#FF4D4D")
	if got != want {
		t.Errorf("NotServing foreground = %v, want %v", got, want)
	}
}

func TestDefaultTheme_UnknownIsNotBold(t *testing.T) {
	theme := DefaultTheme()
	if theme.Unknown.GetBold() {
		t.Error("expected Unknown style to not be bold")
	}
}

func TestDefaultTheme_TitleHasPadding(t *testing.T) {
	theme := DefaultTheme()
	_, right, _, left := theme.Title.GetPadding()
	if left == 0 && right == 0 {
		t.Error("expected Title style to have horizontal padding")
	}
}

func TestDefaultTheme_AlertColors(t *testing.T) {
	theme := DefaultTheme()

	if theme.AlertInfo.GetForeground() != lipgloss.Color("#00BFFF") {
		t.Error("unexpected AlertInfo foreground colour")
	}
	if theme.AlertWarn.GetForeground() != lipgloss.Color("#FFA500") {
		t.Error("unexpected AlertWarn foreground colour")
	}
}

func TestDefaultTheme_BorderStyleHasBorder(t *testing.T) {
	theme := DefaultTheme()
	top, _, _, _ := theme.Border.GetBorder()
	if top == "" {
		t.Error("expected Border style to have a border defined")
	}
}
