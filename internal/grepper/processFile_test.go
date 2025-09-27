package grepper

import (
	"reflect"
	"testing"
	"wbgrep/internal/parser"
)

func TestGrepper_processFile(t *testing.T) {

	tests := []struct {
		name     string
		flags    *parser.Flags
		pattern  string
		files    []string
		filename string
		want     []ResultLine
		wantErr  bool
	}{
		// Успешные случаи - чтение файлов
		{
			name:     "empty_file",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "test",
			files:    []string{},
			filename: "testdata/empty.txt",
			want:     nil,
			wantErr:  false,
		},
		{
			name:     "single_line_file_with_match",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "hello",
			files:    []string{},
			filename: "testdata/single.txt",
			want:     []ResultLine{{Line: "hello world", Num: 1}},
			wantErr:  false,
		},
		{
			name:     "multiple_lines_with_matches",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "match",
			files:    []string{},
			filename: "testdata/multi.txt",
			want: []ResultLine{
				{Line: "first match", Num: 2},
				{Line: "second match", Num: 4},
			},
			wantErr: false,
		},
		{
			name:     "file_with_context_flags",
			flags:    &parser.Flags{FlagAN: 1, FlagBN: 1, FlagCN: 0},
			pattern:  "target",
			files:    []string{},
			filename: "testdata/context.txt",
			want: []ResultLine{
				{Line: "before", Num: 1},
				{Line: "target line", Num: 2},
				{Line: "after", Num: 3},
			},
			wantErr: false,
		},
		{
			name:     "file_with_ignore_case",
			flags:    &parser.Flags{FlagI: true, FlagF: false, FlagV: false},
			pattern:  "HELLO",
			files:    []string{},
			filename: "testdata/mixed_case.txt",
			want:     []ResultLine{{Line: "hello world", Num: 1}, {Line: "HELLO WORLD", Num: 2}, {Line: "Hello World", Num: 3}},
			wantErr:  false,
		},
		{
			name:     "file_with_fixed_string",
			flags:    &parser.Flags{FlagI: false, FlagF: true, FlagV: false},
			pattern:  "hello.world",
			files:    []string{},
			filename: "testdata/fixed.txt",
			want:     []ResultLine{{Line: "hello.world", Num: 1}},
			wantErr:  false,
		},
		{
			name:     "file_with_invert_match",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: true},
			pattern:  "skip",
			files:    []string{},
			filename: "testdata/invert.txt",
			want: []ResultLine{
				{Line: "keep this line", Num: 1},
				{Line: "and this one", Num: 3},
			},
			wantErr: false,
		},

		// Ошибочные случаи - проблемы с файлами
		{
			name:     "non_existent_file",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "test",
			files:    []string{},
			filename: "testdata/nonexistent.txt",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "directory_instead_of_file",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "test",
			files:    []string{},
			filename: "testdata",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "empty_filename",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "test",
			files:    []string{},
			filename: "",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "permission_denied_file",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "test",
			files:    []string{},
			filename: "/root/protected.txt",
			want:     nil,
			wantErr:  true,
		},

		// Граничные случаи
		{
			name:     "file_with_only_newlines",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "test",
			files:    []string{},
			filename: "testdata/newlines.txt",
			want:     nil,
			wantErr:  false,
		},
		{
			name:     "file_with_empty_lines",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern:  "content",
			files:    []string{},
			filename: "testdata/empty_lines.txt",
			want:     []ResultLine{{Line: "line with content", Num: 2}},
			wantErr:  false,
		},

		// Комбинации флагов
		{
			name:     "all_flags_combined",
			flags:    &parser.Flags{FlagI: true, FlagF: true, FlagV: false, FlagAN: 1, FlagBN: 1, FlagCN: 5},
			pattern:  "TARGET",
			files:    []string{},
			filename: "testdata/combined.txt",
			want: []ResultLine{
				{Line: "before target", Num: 1},
				{Line: "target", Num: 2},
				{Line: "after target", Num: 3},
				{Line: "another line", Num: 4},
			},
			wantErr: false,
		},
		{
			name:     "invert_with_context",
			flags:    &parser.Flags{FlagI: false, FlagF: false, FlagV: true},
			pattern:  "exclude",
			files:    []string{},
			filename: "testdata/invert_context.txt",
			want: []ResultLine{
				{Line: "keep before", Num: 1},
				{Line: "keep after", Num: 3},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got, gotErr := g.processFile(tt.filename)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("processFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("processFile() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
