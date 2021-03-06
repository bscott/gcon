package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Event
func Test_Event(t *testing.T) {
	e := &models.Event{}
	e.Name = "Bar Crawl"
	e.ShortDescription = "insert short joke here"
	if m := e.String(); !strings.Contains(m, "Bar Crawl") {
		t.Errorf("expected contains %s, got %s", "Great Divide", m)
	}
}

// Test_Events
func Test_Events(t *testing.T) {
	e := &models.Events{
		{
			Name:             "Bar Crawl",
			ShortDescription: "insert short joke here",
		},
	}
	if m := e.String(); !strings.Contains(m, "Bar Crawl") {
		t.Errorf("expected contains %s, got %s", "Bar Crawl", m)
	}
}
