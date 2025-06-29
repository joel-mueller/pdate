package dates

import (
	"pdate/internal/job"
	"testing"
	"time"
)

func TestFormatDates(t *testing.T) {
	dates := []time.Time{
		time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, time.December, 31, 0, 0, 0, 0, time.UTC),
	}
	format := "{YY}-{mn}-{D}"

	expected := []string{"25-Jan-1", "25-Dec-31"}
	result := FormatDates(dates, format, job.English)

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
}

func TestReplaceDatePlaceholdersWithDate(t *testing.T) {
	date := time.Date(2025, time.March, 5, 0, 0, 0, 0, time.UTC)

	tests := map[string]string{
		"{YYYY}": "2025",
		"{YY}":   "25",
		"{MM}":   "03",
		"{M}":    "3",
		"{DD}":   "05",
		"{D}":    "5",
		"{WD}":   "Wednesday",
		"{wd}":   "Wed",
		"{MN}":   "March",
		"{mn}":   "Mar",
	}

	for placeholder, expected := range tests {
		result := ReplaceDatePlaceholdersWithDate(placeholder, date, job.English)
		if result != expected {
			t.Errorf("Placeholder %s: expected %s, got %s", placeholder, expected, result)
		}
	}

	t.Run("Full custom string", func(t *testing.T) {
		format := "Today is {WD}, {MN} {D}, {YYYY}"
		expected := "Today is Wednesday, March 5, 2025"
		result := ReplaceDatePlaceholdersWithDate(format, date, job.English)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}

func TestGetShortFormName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		lang     job.Language
		expected string
	}{
		{
			name:     "English short form",
			input:    "Monday",
			lang:     job.English,
			expected: "Mon",
		},
		{
			name:     "German short form",
			input:    "Dienstag",
			lang:     job.German,
			expected: "Die",
		},
		{
			name:     "Spanish short form",
			input:    "Martes",
			lang:     job.Spanish,
			expected: "Mar",
		},
		{
			name:     "Dutch short form",
			input:    "Woensdag",
			lang:     job.Dutch,
			expected: "Woe",
		},
		{
			name:     "Hindi no short form",
			input:    "सोमवार",
			lang:     job.Hindi,
			expected: "सोमवार",
		},
		{
			name:     "Chinese no short form",
			input:    "星期一",
			lang:     job.Chinese,
			expected: "星期一",
		},
		{
			name:     "Short input string",
			input:    "Lundi",
			lang:     job.French,
			expected: "Lun",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetShortFormName(tt.input, tt.lang)
			if result != tt.expected {
				t.Errorf("GetShortFormName(%q, %v) = %q; want %q", tt.input, tt.lang, result, tt.expected)
			}
		})
	}
}
