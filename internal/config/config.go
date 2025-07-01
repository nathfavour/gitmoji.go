package config

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ConfigDirName    = ".gitmojigo"
	SourcesFileName  = "sources.json"
	DefaultSourceKey = "default"
	DefaultSourceURL = "https://raw.githubusercontent.com/github/gemoji/0eca75db9301421efc8710baf7a7576793ae452a/db/emoji.json"
	DefaultSourceDir = "sources/default"
	DefaultEmojiFile = "emoji.json"
)

type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Emoji struct {
	Emoji       string   `json:"emoji"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Aliases     []string `json:"aliases"`
	Tags        []string `json:"tags"`
	UnicodeVer  string   `json:"unicode_version"`
	IOSVer      string   `json:"ios_version"`
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func configDir() string {
	return filepath.Join(homeDir(), ConfigDirName)
}

func sourcesFile() string {
	return filepath.Join(configDir(), SourcesFileName)
}

func defaultSourceDir() string {
	return filepath.Join(configDir(), DefaultSourceDir)
}

func defaultEmojiFile() string {
	return filepath.Join(defaultSourceDir(), DefaultEmojiFile)
}

// EnsureConfig ensures config directory and sources.json exist.
func EnsureConfig() error {
	if err := os.MkdirAll(configDir(), 0755); err != nil {
		return err
	}
	if _, err := os.Stat(sourcesFile()); errors.Is(err, os.ErrNotExist) {
		sources := map[string]Source{
			DefaultSourceKey: {Name: DefaultSourceKey, URL: DefaultSourceURL},
		}
		b, _ := json.MarshalIndent(sources, "", "  ")
		return os.WriteFile(sourcesFile(), b, 0644)
	}
	return nil
}

// EnsureDefaultSource ensures the default source directory exists.
func EnsureDefaultSource() error {
	return os.MkdirAll(defaultSourceDir(), 0755)
}

// EnsureDefaultEmojiJSON downloads the emoji JSON if not present.
func EnsureDefaultEmojiJSON() error {
	emojiPath := defaultEmojiFile()
	if _, err := os.Stat(emojiPath); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(DefaultSourceURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return os.WriteFile(emojiPath, b, 0644)
	}
	return nil
}

// LoadEmojis loads emojis from the default emoji JSON.
func LoadEmojis() ([]Emoji, error) {
	b, err := os.ReadFile(defaultEmojiFile())
	if err != nil {
		return nil, err
	}
	var emojis []Emoji
	if err := json.Unmarshal(b, &emojis); err != nil {
		return nil, err
	}
	return emojis, nil
}
