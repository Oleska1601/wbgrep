package grepper

import (
	"testing"
	"wbgrep/internal/parser"
)

func TestGrepper_isMatchNoFlags(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		line string
		want bool
	}{
		// Базовые regexp тесты (без флагов)
		{
			name:    "regexp_exact_match",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern: "hello",
			files:   []string{},
			line:    "hello",
			want:    true,
		},
		{
			name:    "regexp_partial_match",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern: "hello",
			files:   []string{},
			line:    "hello world",
			want:    true,
		},
		{
			name:    "regexp_no_match",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern: "hello",
			files:   []string{},
			line:    "goodbye",
			want:    false,
		},
		{
			name:    "regexp_special_chars",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern: "h.llo",
			files:   []string{},
			line:    "hello",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got := g.isMatch(tt.line)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepper_isMatchFlagF(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		line string
		want bool
	}{
		// FlagF тесты (точное совпадение)
		{
			name:    "fixed_exact_match",
			flags:   &parser.Flags{FlagI: false, FlagF: true, FlagV: false},
			pattern: "hello",
			files:   []string{},
			line:    "hello",
			want:    true,
		},
		{
			name:    "fixed_no_match_partial",
			flags:   &parser.Flags{FlagI: false, FlagF: true, FlagV: false},
			pattern: "hello",
			files:   []string{},
			line:    "hello world",
			want:    false,
		},
		{
			name:    "fixed_empty_pattern_empty_line",
			flags:   &parser.Flags{FlagI: false, FlagF: true, FlagV: false},
			pattern: "",
			files:   []string{},
			line:    "",
			want:    true,
		},
		{
			name:    "fixed_empty_pattern_non_empty_line",
			flags:   &parser.Flags{FlagI: false, FlagF: true, FlagV: false},
			pattern: "",
			files:   []string{},
			line:    "text",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got := g.isMatch(tt.line)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepper_isMatchFlagV(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		line string
		want bool
	}{
		// FlagV тесты (инвертирование)
		{
			name:    "invert_regexp_match",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: true},
			pattern: "hello",
			files:   []string{},
			line:    "hello",
			want:    false,
		},
		{
			name:    "invert_regexp_no_match",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: true},
			pattern: "hello",
			files:   []string{},
			line:    "goodbye",
			want:    true,
		},
		{
			name:    "invert_fixed_match",
			flags:   &parser.Flags{FlagI: false, FlagF: true, FlagV: true},
			pattern: "hello",
			files:   []string{},
			line:    "hello",
			want:    false,
		},
		{
			name:    "invert_fixed_no_match",
			flags:   &parser.Flags{FlagI: false, FlagF: true, FlagV: true},
			pattern: "hello",
			files:   []string{},
			line:    "hello world",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got := g.isMatch(tt.line)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrepper_isMatchSpecialCases(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		flags   *parser.Flags
		pattern string
		files   []string
		// Named input parameters for target function.
		line string
		want bool
	}{
		// Комбинации флагов
		{
			name:    "ignore_case_invert_regexp",
			flags:   &parser.Flags{FlagI: true, FlagF: false, FlagV: true},
			pattern: "HELLO",
			files:   []string{},
			line:    "hello",
			want:    false,
		},
		{
			name:    "ignore_case_invert_fixed",
			flags:   &parser.Flags{FlagI: true, FlagF: true, FlagV: true},
			pattern: "HELLO",
			files:   []string{},
			line:    "hello",
			want:    false,
		},
		{
			name:    "ignore_case_invert_no_match",
			flags:   &parser.Flags{FlagI: true, FlagF: true, FlagV: true},
			pattern: "HELLO",
			files:   []string{},
			line:    "hell",
			want:    true,
		},

		// Комбинация всех флагов
		{
			name:    "all_flags_enabled_match",
			flags:   &parser.Flags{FlagI: true, FlagF: true, FlagV: true},
			pattern: "TEST",
			files:   []string{},
			line:    "test",
			want:    false,
		},
		{
			name:    "all_flags_enabled_no_match",
			flags:   &parser.Flags{FlagI: true, FlagF: true, FlagV: true},
			pattern: "TEST",
			files:   []string{},
			line:    "different",
			want:    true,
		},

		// Граничные случаи
		{
			name:    "empty_pattern_regexp",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern: "",
			files:   []string{},
			line:    "any text",
			want:    true,
		},
		{
			name:    "empty_pattern_regexp_empty_line",
			flags:   &parser.Flags{FlagI: false, FlagF: false, FlagV: false},
			pattern: "",
			files:   []string{},
			line:    "",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.flags, tt.pattern, tt.files)
			got := g.isMatch(tt.line)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
