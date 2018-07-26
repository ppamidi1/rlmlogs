package logfiles

import (
	"testing"
)

func TestShowValidItems(t *testing.T) {
	ShowValidItems()
}

func TestValidItem(t *testing.T) {
	t.Logf("Valid Item CS = %v\n", ValidItem("CS"))
	t.Logf("Valid Item CPUTemp = %v\n", ValidItem("CPUTemp"))
}

func TestIndex(t *testing.T) {
	t.Logf("Index of CS = %d\n", Index("CS"))
	t.Logf("Index of CPUTemp = %d\n", Index("CPUTemp"))
}
