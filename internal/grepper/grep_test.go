package grepper_test

import (
	"testing"
	"wbgrep/internal/grepper"
	"wbgrep/internal/parser"
)

func TestGrepper_Grep(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		wantErr bool
	}{
		{
			name:    "stdin success",
			flags:   &parser.Flags{},
			pattern: "test",
			files:   []string{}, // Пустой массив = stdin
			wantErr: false,
		},
		{
			name:    "single file success",
			flags:   &parser.Flags{},
			pattern: "hello",
			files:   []string{"./testdata/combined.txt"},
			wantErr: false,
		},
		{
			name:    "single file not found",
			flags:   &parser.Flags{},
			pattern: "hello",
			files:   []string{"nonexistent.txt"},
			wantErr: true,
		},
		{
			name:    "multiple files success",
			flags:   &parser.Flags{},
			pattern: "pattern",
			files:   []string{"./testdata/context.txt", "./testdata/combined.txt"},
			wantErr: false,
		},
		{
			name:    "multiple files first fails",
			flags:   &parser.Flags{},
			pattern: "pattern",
			files:   []string{"nonexistent.txt", "./testdata/combined.txt"}, // Первый файл вызывает ошибку
			wantErr: true,
		},
		{
			name:    "multiple files middle fails",
			flags:   &parser.Flags{},
			pattern: "pattern",
			files:   []string{"./testdata/combined.txt", "nonexistent.txt", "./testdata/single.txt"}, // Средний файл вызывает ошибку
			wantErr: true,
		},
		{
			name:    "multiple files last fails",
			flags:   &parser.Flags{},
			pattern: "pattern",
			files:   []string{"./testdata/combined.txt", "./testdata/single.txt", "nonexistent.txt"}, // Последний файл вызывает ошибку
			wantErr: true,
		},
		{
			name:    "nil files slice",
			flags:   &parser.Flags{},
			pattern: "test",
			files:   nil, // nil вместо пустого slice
			wantErr: false,
		},
		{
			name:    "empty filename in multiple files",
			flags:   &parser.Flags{},
			pattern: "test",
			files:   []string{"./testdata/combined.txt", "", "./testdata/single.txt"}, // Пустое имя файла
			wantErr: true,
		},
		{
			name:    "empty pattern",
			flags:   &parser.Flags{},
			pattern: "", // Пустой паттерн
			files:   []string{"./testdata/combined.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := grepper.New(tt.flags, tt.pattern, tt.files)
			gotErr := g.Grep()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Grep() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Grep() succeeded unexpectedly")
			}
		})
	}
}

func TestGrepper_GrepSpecialCases(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		wantErr bool
	}{
		{
			name:    "stdin with count flag",
			flags:   &parser.Flags{FlagC: true},
			pattern: "test",
			files:   []string{},
			wantErr: false,
		},
		{
			name:    "multiple files with line numbers",
			flags:   &parser.Flags{FlagN: true},
			pattern: "pattern",
			files:   []string{"./testdata/combined.txt", "./testdata/empty.txt"},
			wantErr: false,
		},
		{
			name:    "single file with ignore case",
			flags:   &parser.Flags{FlagI: true},
			pattern: "TEST",
			files:   []string{"./testdata/combined.txt"},
			wantErr: false,
		},
		{
			name:    "stdin with inverted match",
			flags:   &parser.Flags{FlagV: true},
			pattern: "exclude",
			files:   []string{},
			wantErr: false,
		},
		{
			name:    "with context lines flag",
			flags:   &parser.Flags{FlagAN: 2, FlagBN: 1},
			pattern: "pattern",
			files:   []string{"./testdata/combined.txt"},
			wantErr: false,
		},
		{
			name:    "stdin with count flag",
			flags:   &parser.Flags{FlagC: true},
			pattern: "test",
			files:   []string{},
			wantErr: false,
		},
		{
			name:    "multiple files with line numbers",
			flags:   &parser.Flags{FlagN: true},
			pattern: "pattern",
			files:   []string{"./testdata/combined.txt", "./testdata/multi.txt"},
			wantErr: false,
		},
		{
			name:    "single file with ignore case",
			flags:   &parser.Flags{FlagI: true},
			pattern: "TEST",
			files:   []string{"./testdata/combined.txt"},
			wantErr: false,
		},
		{
			name:    "stdin with inverted match",
			flags:   &parser.Flags{FlagV: true},
			pattern: "exclude",
			files:   []string{},
			wantErr: false,
		},
		{
			name:    "with context lines flag",
			flags:   &parser.Flags{FlagAN: 2, FlagBN: 1},
			pattern: "pattern",
			files:   []string{"./testdata/combined.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := grepper.New(tt.flags, tt.pattern, tt.files)
			gotErr := g.Grep()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Grep() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Grep() succeeded unexpectedly")
			}
		})
	}
}
