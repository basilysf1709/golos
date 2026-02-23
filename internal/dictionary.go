package internal

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Dictionary struct {
	entries map[string]string // lowercase spoken phrase → replacement
}

type dictionaryFile struct {
	Words map[string]string `toml:"words"`
}

func LoadDictionary() *Dictionary {
	d := &Dictionary{entries: make(map[string]string)}

	home, err := os.UserHomeDir()
	if err != nil {
		return d
	}

	path := filepath.Join(home, ".config", "golos", "dictionary.toml")
	if _, err := os.Stat(path); err != nil {
		return d
	}

	var f dictionaryFile
	if _, err := toml.DecodeFile(path, &f); err != nil {
		return d
	}

	for phrase, replacement := range f.Words {
		d.entries[strings.ToLower(phrase)] = replacement
	}

	return d
}

func dictionaryPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "golos", "dictionary.toml")
}

func (d *Dictionary) Add(phrase, replacement string) error {
	d.entries[strings.ToLower(phrase)] = replacement
	return d.save()
}

func (d *Dictionary) Delete(phrase string) bool {
	key := strings.ToLower(phrase)
	if _, ok := d.entries[key]; !ok {
		return false
	}
	delete(d.entries, key)
	d.save()
	return true
}

func (d *Dictionary) List() map[string]string {
	out := make(map[string]string, len(d.entries))
	for k, v := range d.entries {
		out[k] = v
	}
	return out
}

func (d *Dictionary) save() error {
	path := dictionaryPath()
	os.MkdirAll(filepath.Dir(path), 0755)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(dictionaryFile{Words: d.entries})
}

// Import merges entries from a TOML file into the dictionary and saves.
func (d *Dictionary) Import(path string) (int, error) {
	var f dictionaryFile
	if _, err := toml.DecodeFile(path, &f); err != nil {
		return 0, err
	}

	count := 0
	for phrase, replacement := range f.Words {
		d.entries[strings.ToLower(phrase)] = replacement
		count++
	}

	return count, d.save()
}

func (d *Dictionary) Replace(text string) string {
	if len(d.entries) == 0 {
		return text
	}

	lower := strings.ToLower(text)
	for phrase, replacement := range d.entries {
		for {
			idx := strings.Index(strings.ToLower(lower), phrase)
			if idx == -1 {
				break
			}
			text = text[:idx] + replacement + text[idx+len(phrase):]
			lower = strings.ToLower(text)
		}
	}

	return text
}
