package grepper

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/Oleska1601/wbgrep/internal/parser"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		flags    *parser.Flags
		pattern  string
		expected []string
	}{
		{
			name:    "basic match",
			input:   "apple\nbanana\norange\napple",
			flags:   &parser.Flags{},
			pattern: "apple",
			expected: []string{
				"apple",
				"apple",
			},
		},
		{
			name:     "no match",
			input:    "banana\norange\ngrape",
			flags:    &parser.Flags{},
			pattern:  "apple",
			expected: nil,
		},
		{
			name:    "flag C count only",
			input:   "apple\nbanana\norange\napple",
			flags:   &parser.Flags{FlagC: true},
			pattern: "apple",
			expected: []string{
				"2",
			},
		},
		{
			name:    "flag C count no match",
			input:   "banana\norange\ngrape",
			flags:   &parser.Flags{FlagC: true},
			pattern: "apple",
			expected: []string{
				"0",
			},
		},
		{
			name:    "flag N line numbers",
			input:   "apple\nbanana\norange\napple",
			flags:   &parser.Flags{FlagN: true},
			pattern: "apple",
			expected: []string{
				"1:apple",
				"4:apple",
			},
		},
		{
			name:    "flag I ignore case",
			input:   "Apple\nbanana\norange\nAPPLE",
			flags:   &parser.Flags{FlagI: true},
			pattern: "apple",
			expected: []string{
				"Apple",
				"APPLE",
			},
		},
		{
			name:    "flag V invert match",
			input:   "apple\nbanana\norange\napple",
			flags:   &parser.Flags{FlagV: true},
			pattern: "apple",
			expected: []string{
				"banana",
				"orange",
			},
		},
		{
			name:    "flag F fixed string",
			input:   "apple\napple pie\npineapple",
			flags:   &parser.Flags{FlagF: true},
			pattern: "apple",
			expected: []string{
				"apple",
			},
		},
		{
			name:    "flag B before context",
			input:   "line1\nline2\nTARGET\nline4\nline5",
			flags:   &parser.Flags{FlagBN: 2},
			pattern: "TARGET",
			expected: []string{
				"line1",
				"line2",
				"TARGET",
			},
		},
		{
			name:    "flag A after context",
			input:   "line1\nline2\nTARGET\nline4\nline5",
			flags:   &parser.Flags{FlagAN: 2},
			pattern: "TARGET",
			expected: []string{
				"TARGET",
				"line4",
				"line5",
			},
		},
		{
			name:    "flag C context around",
			input:   "line1\nline2\nTARGET\nline4\nline5",
			flags:   &parser.Flags{FlagCN: 1},
			pattern: "TARGET",
			expected: []string{
				"line2",
				"TARGET",
				"line4",
			},
		},
		{
			name:    "multiple matches with context",
			input:   "line1\nTARGET1\nline3\nTARGET2\nline5",
			flags:   &parser.Flags{FlagCN: 1},
			pattern: "TARGET",
			expected: []string{
				"line1",
				"TARGET1",
				"line3",
				"TARGET2",
				"line5",
			},
		},
		{
			name:    "flag B at start of file",
			input:   "TARGET\nline2\nline3",
			flags:   &parser.Flags{FlagBN: 2},
			pattern: "TARGET",
			expected: []string{
				"TARGET",
			},
		},
		{
			name:    "flag A at end of file",
			input:   "line1\nline2\nTARGET",
			flags:   &parser.Flags{FlagAN: 2},
			pattern: "TARGET",
			expected: []string{
				"TARGET",
			},
		},
		{
			name:    "combination flags N and C",
			input:   "line1\nline2\nTARGET\nline4\nline5",
			flags:   &parser.Flags{FlagN: true, FlagCN: 1},
			pattern: "TARGET",
			expected: []string{
				"2:line2",
				"3:TARGET",
				"4:line4",
			},
		},
		{
			name:     "empty input",
			input:    "",
			flags:    &parser.Flags{},
			pattern:  "apple",
			expected: nil,
		},
		{
			name:    "regex pattern",
			input:   "test123\nabc456\nhello789",
			flags:   &parser.Flags{},
			pattern: "[0-9]+",
			expected: []string{
				"test123",
				"abc456",
				"hello789",
			},
		},
		{
			name:    "overlapping context",
			input:   "line1\nTARGET1\nline3\nTARGET2\nline5",
			flags:   &parser.Flags{FlagAN: 2, FlagBN: 2},
			pattern: "TARGET",
			expected: []string{
				"line1",
				"TARGET1",
				"line3",
				"TARGET2",
				"line5",
			},
		},
		// more hard test cases
		{
			name:    "A and B flags together",
			input:   "1\n2\nTARGET\n4\n5\n6\nTARGET\n8\n9",
			flags:   &parser.Flags{FlagAN: 1, FlagBN: 2},
			pattern: "TARGET",
			expected: []string{
				"1", "2", "TARGET", "4",
				"5", "6", "TARGET", "8",
			},
		},
		{
			name:    "C flag with line numbers",
			input:   "start\nmiddle\nERROR\nnext\nend",
			flags:   &parser.Flags{FlagCN: 1, FlagN: true},
			pattern: "ERROR",
			expected: []string{
				"2:middle",
				"3:ERROR",
				"4:next",
			},
		},
		{
			name:    "multiple flags: i, n, v",
			input:   "ERROR: one\ninfo: two\nError: three\nDEBUG: four",
			flags:   &parser.Flags{FlagI: true, FlagN: true, FlagV: true},
			pattern: "error",
			expected: []string{
				"2:info: two",
				"4:DEBUG: four",
			},
		},
		{
			name:     "F flag with regex chars",
			input:    "a.b\nab\nabc\na+b",
			flags:    &parser.Flags{FlagF: true},
			pattern:  "a.b",
			expected: []string{"a.b"},
		},
		{
			name:     "count with invert",
			input:    "match1\nskip\nmatch2\nskip\nmatch3",
			flags:    &parser.Flags{FlagC: true, FlagV: true},
			pattern:  "match",
			expected: []string{"2"},
		},

		// Простые регулярки с флагами
		{
			name:    "simple regex with A flag",
			input:   "test1\ntest2\n12345\ntest4\ntest5",
			flags:   &parser.Flags{FlagAN: 2},
			pattern: `^[0-9]+$`,
			expected: []string{
				"12345", "test4", "test5",
			},
		},
		{
			name:    "regex with B flag and line numbers",
			input:   "apple\nbanana\n123\norange\ngrape",
			flags:   &parser.Flags{FlagBN: 1, FlagN: true},
			pattern: `^[0-9]+$`,
			expected: []string{
				"2:banana",
				"3:123",
			},
		},
		{
			name:    "regex char class with i flag",
			input:   "ERROR\nError\nerror\nSUCCESS",
			flags:   &parser.Flags{FlagI: true},
			pattern: `^[a-z]+$`,
			expected: []string{
				"ERROR", "Error", "error", "SUCCESS",
			},
		},

		// Комбинации контекстных флагов
		{
			name:    "A=2 B=1 with multiple matches",
			input:   "1\n2\nT1\n4\n5\nT2\n7\n8",
			flags:   &parser.Flags{FlagAN: 2, FlagBN: 1},
			pattern: "T",
			expected: []string{
				"2", "T1", "4", "5", "T2", "7", "8",
			},
		},
		{
			name:    "C=2 with line numbers",
			input:   "a\nb\nc\nd\ne\nf\ng\nh",
			flags:   &parser.Flags{FlagCN: 2, FlagN: true},
			pattern: "e",
			expected: []string{
				"3:c", "4:d", "5:e", "6:f", "7:g",
			},
		},

		// Fixed string комбинации
		{
			name:    "F with i and n",
			input:   "Apple\napple\nAPPLE\nbanana",
			flags:   &parser.Flags{FlagF: true, FlagI: true, FlagN: true},
			pattern: "apple",
			expected: []string{
				"1:Apple", "2:apple", "3:APPLE",
			},
		},
		{
			name:     "F with v and c",
			input:    "test.exe\nconfig.txt\nreadme.md\nscript.sh",
			flags:    &parser.Flags{FlagF: true, FlagV: true, FlagC: true},
			pattern:  ".exe",
			expected: []string{"4"},
		},

		// Граничные кейсы с контекстом
		{
			name:     "large context values",
			input:    "1\n2\nTARGET\n4\n5",
			flags:    &parser.Flags{FlagCN: 10}, // Больше чем строк в файле
			pattern:  "TARGET",
			expected: []string{"1", "2", "TARGET", "4", "5"},
		},
		{
			name:     "zero context",
			input:    "1\n2\nMATCH\n4\n5",
			flags:    &parser.Flags{FlagCN: 0},
			pattern:  "MATCH",
			expected: []string{"MATCH"},
		},

		// Специальные символы
		{
			name:     "dots in regex",
			input:    "a.b\na+b\nab\nacb",
			flags:    &parser.Flags{},
			pattern:  `a.b`,
			expected: []string{"a.b", "a+b", "acb"},
		},
		{
			name:     "plus in regex",
			input:    "ab\nabb\nabbb\nacb",
			flags:    &parser.Flags{},
			pattern:  `ab+`,
			expected: []string{"ab", "abb", "abbb"},
		},
		{
			name:     "question mark regex",
			input:    "color\ncolour\ncolr",
			flags:    &parser.Flags{},
			pattern:  `colou?r`,
			expected: []string{"color", "colour"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grepper := New(tt.flags, tt.pattern)
			reader := strings.NewReader(tt.input)
			ctx := context.Background()

			output := grepper.Grep(ctx, reader)

			var results []string
			for line := range output {
				results = append(results, line)
			}

			if !reflect.DeepEqual(results, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, results)
			}
		})
	}
}
