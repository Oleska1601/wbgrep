package grepper

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
	"wbgrep/internal/parser"
)

/*
func СompareResultLines(t *testing.T, got, want []ResultLine) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("length mismatch: got %d, want %d", len(got), len(want))
		return
	}

	for i := range got {
		if got[i].Line != want[i].Line {
			t.Errorf("line %d content mismatch:\n got: %q\nwant: %q",
				i, got[i].Line, want[i].Line)
		}
		if got[i].Num != want[i].Num {
			t.Errorf("line %d number mismatch: got %d, want %d",
				i, got[i].Num, want[i].Num)
		}
	}
}
*/

func TestGrepper_processLinesNoFlags(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		scanner *bufio.Scanner
		want    []ResultLine
		wantErr bool
	}{
		// Базовые тесты без флагов контекста
		{
			name:    "no_context_flags_single_match",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 0},
			pattern: "hello",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("hello world\n")),
			want:    []ResultLine{{Line: "hello world", Num: 1}},
			wantErr: false,
		},
		{
			name:    "no_context_flags_no_match",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 0},
			pattern: "hello",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("goodbye world\n")),
			want:    nil,
			wantErr: false,
		},
		{
			name:    "no_context_flags_multiple_matches",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 0},
			pattern: "test",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("test one\ntest two\nnot this\ntest three\n")),
			want: []ResultLine{
				{Line: "test one", Num: 1},
				{Line: "test two", Num: 2},
				{Line: "test three", Num: 4},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got, gotErr := g.processLines(tt.scanner)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("processLines() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("processLines() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepper_processLinesFlags(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		scanner *bufio.Scanner
		want    []ResultLine
		wantErr bool
	}{

		// Тесты с FlagA (строки после совпадения)
		{
			name:    "flag_A_2_lines_after",
			flags:   &parser.Flags{FlagAN: 2, FlagBN: 0, FlagCN: 0},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nmatch\nline3\nline4\nline5\n")),
			want: []ResultLine{
				{Line: "match", Num: 2},
				{Line: "line3", Num: 3},
				{Line: "line4", Num: 4},
			},
			wantErr: false,
		},
		{
			name:    "flag_A_match_at_end_of_file",
			flags:   &parser.Flags{FlagAN: 2, FlagBN: 0, FlagCN: 0},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nline2\nmatch\n")),
			want: []ResultLine{
				{Line: "match", Num: 3},
			},
			wantErr: false,
		},

		// Тесты с FlagB (строки до совпадения)
		{
			name:    "flag_B_2_lines_before",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 2, FlagCN: 0},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nline2\nmatch\nline4\n")),
			want: []ResultLine{
				{Line: "line1", Num: 1},
				{Line: "line2", Num: 2},
				{Line: "match", Num: 3},
			},
			wantErr: false,
		},
		{
			name:    "flag_B_match_at_beginning",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 2, FlagCN: 0},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("match\nline2\nline3\n")),
			want: []ResultLine{
				{Line: "match", Num: 1},
			},
			wantErr: false,
		},

		// Тесты с FlagC (контекст вокруг совпадения)
		{
			name:    "flag_C_1_line_context",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 1},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nmatch\nline3\n")),
			want: []ResultLine{
				{Line: "line1", Num: 1},
				{Line: "match", Num: 2},
				{Line: "line3", Num: 3},
			},
			wantErr: false,
		},
		{
			name:    "flag_C_2_lines_context_multiple_matches",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 2},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nline2\nmatch\nline4\nline5\nmatch\nline7\n")),
			want: []ResultLine{
				{Line: "line1", Num: 1},
				{Line: "line2", Num: 2},
				{Line: "match", Num: 3},
				{Line: "line4", Num: 4},
				{Line: "line5", Num: 5},
				{Line: "match", Num: 6},
				{Line: "line7", Num: 7},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got, gotErr := g.processLines(tt.scanner)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("processLines() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("processLines() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepper_processLinesSpecialCases(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		scanner *bufio.Scanner
		want    []ResultLine
		wantErr bool
	}{

		// Комбинации флагов контекста
		{
			name:    "flag_A_and_B_together",
			flags:   &parser.Flags{FlagAN: 1, FlagBN: 1, FlagCN: 0},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nmatch\nline3\n")),
			want: []ResultLine{
				{Line: "line1", Num: 1},
				{Line: "match", Num: 2},
				{Line: "line3", Num: 3},
			},
			wantErr: false,
		},
		{
			name:    "flag_C_overrides_A_and_B",
			flags:   &parser.Flags{FlagAN: 5, FlagBN: 5, FlagCN: 1},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("line1\nmatch\nline3\n")),
			want: []ResultLine{
				{Line: "line1", Num: 1},
				{Line: "match", Num: 2},
				{Line: "line3", Num: 3},
			},
			wantErr: false,
		},

		// Граничные случаи
		{
			name:    "empty_file",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 0},
			pattern: "test",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("")),
			want:    nil,
			wantErr: false,
		},
		{
			name:    "only_newlines",
			flags:   &parser.Flags{FlagAN: 0, FlagBN: 0, FlagCN: 0},
			pattern: "test",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("\n\n\n")),
			want:    nil,
			wantErr: false,
		},
		{
			name:    "consecutive_matches",
			flags:   &parser.Flags{FlagAN: 1, FlagBN: 0, FlagCN: 0},
			pattern: "match",
			files:   []string{},
			scanner: bufio.NewScanner(strings.NewReader("match\nmatch\nline3\n")),
			want: []ResultLine{
				{Line: "match", Num: 1},
				{Line: "match", Num: 2},
				{Line: "line3", Num: 3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got, gotErr := g.processLines(tt.scanner)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("processLines() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("processLines() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
