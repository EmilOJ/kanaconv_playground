package kana

import (
	"github.com/miiton/kanaconv"
	"testing"
)

func TestHiraganaToKatakana(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "あいうえお",
			want:  "アイウエオ",
		},
		{
			input: "アイウエオ",
			want:  "アイウエオ",
		},
		{
			input: "あいウえお",
			want:  "アイウエオ",
		},
		{
			input: "あいuえお",
			want:  "アイuエオ",
		},
		{
			input: "あいUえお",
			want:  "アイUエオ",
		},
		{
			input: "あい雨えお",
			want:  "アイ雨エオ",
		},
	}

	for i, test := range tests {
		got := kanaconv.HiraganaToKatakana(test.input)
		if got != test.want {
			t.Errorf("[%d] got %s, but want %s", i, got, test.want)
		}
	}
}

func TestKatakanaToHiragana(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "アイウエオ",
			want:  "あいうえお",
		},
		{
			input: "あいうえお",
			want:  "あいうえお",
		},
		{
			input: "アイうエオ",
			want:  "あいうえお",
		},
		{
			input: "アイuエオ",
			want:  "あいuえお",
		},
		{
			input: "アイUエオ",
			want:  "あいUえお",
		},
		{
			input: "アイ雨エオ",
			want:  "あい雨えお",
		},
	}

	for i, test := range tests {
		got := kanaconv.KatakanaToHiragana(test.input)
		if got != test.want {
			t.Errorf("[%d] got %s, but want %s", i, got, test.want)
		}
	}
}

func TestZenToHan(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "アイウエオ",
			want:  "ｱｲｳｴｵ",
		},
		{
			input: "あいうえお",
			want:  "あいうえお",
		},
		{
			input: "アイうエオ",
			want:  "ｱｲうｴｵ",
		},
		{
			input: "アイｕエオ",
			want:  "ｱｲｕｴｵ",
		},
		{
			input: "アイuエオ",
			want:  "ｱｲuｴｵ",
		},
		{
			input: "アイ雨エオ",
			want:  "ｱｲ雨ｴｵ",
		},
	}

	for i, test := range tests {
		got := kanaconv.ZenkakuToHankaku(test.input)
		if got != test.want {
			t.Errorf("[%d] got %s, but want %s", i, got, test.want)
		}
	}
}

func TestHanToZen(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},
		{
			input: "ｱｲｳｴｵ",
			want:  "アイウエオ",
		},
		{
			input: "あいうえお",
			want:  "あいうえお",
		},
		{
			input: "ｱｲうｴｵ",
			want:  "アイうエオ",
		},
		{
			input: "ｱｲｕｴｵ",
			want:  "アイｕエオ",
		},
		{
			input: "ｱｲuｴｵ",
			want:  "アイuエオ",
		},
		{
			input: "ｱｲ雨ｴｵ",
			want:  "アイ雨エオ",
		},
	}

	for i, test := range tests {
		got := kanaconv.HankakuToZenkaku(test.input)
		if got != test.want {
			t.Errorf("[%d] got %s, but want %s", i, got, test.want)
		}
	}
}
