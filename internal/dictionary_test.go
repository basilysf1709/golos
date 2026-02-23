package internal

import "testing"

func TestReplaceEmpty(t *testing.T) {
	d := &Dictionary{entries: map[string]string{}}
	got := d.Replace("hello world")
	if got != "hello world" {
		t.Errorf("got %q, want %q", got, "hello world")
	}
}

func TestReplaceSingleWord(t *testing.T) {
	d := &Dictionary{entries: map[string]string{
		"new line": "\n",
	}}
	got := d.Replace("hello new line world")
	if got != "hello \n world" {
		t.Errorf("got %q, want %q", got, "hello \n world")
	}
}

func TestReplaceCaseInsensitive(t *testing.T) {
	d := &Dictionary{entries: map[string]string{
		"period": ".",
	}}
	got := d.Replace("end of sentence Period")
	if got != "end of sentence ." {
		t.Errorf("got %q, want %q", got, "end of sentence .")
	}
}

func TestReplaceMultipleOccurrences(t *testing.T) {
	d := &Dictionary{entries: map[string]string{
		"comma": ",",
	}}
	got := d.Replace("one comma two comma three")
	if got != "one , two , three" {
		t.Errorf("got %q, want %q", got, "one , two , three")
	}
}

func TestReplaceNoMatch(t *testing.T) {
	d := &Dictionary{entries: map[string]string{
		"exclamation": "!",
	}}
	got := d.Replace("nothing to replace here")
	if got != "nothing to replace here" {
		t.Errorf("got %q, want %q", got, "nothing to replace here")
	}
}

func TestLoadDictionaryReturnsNonNil(t *testing.T) {
	d := LoadDictionary()
	if d == nil {
		t.Fatal("expected non-nil dictionary")
	}
}
