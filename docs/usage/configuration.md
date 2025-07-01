# Configuration

`gitmoji.go` supports multiple configuration methods:

## Configuration Sources

1. `gitmoji` key in your `package.json`
2. `.gitmojirc.json` file in your project
3. Parent directories (recursively)
4. Global configuration via CLI

## Example `.gitmojirc.json`

```json
{
  "autoAdd": false,
  "emojiFormat": "code | emoji",
  "scopePrompt": false,
  "messagePrompt": false,
  "capitalizeTitle": true,
  "gitmojisUrl": "https://gitmoji.dev/api/gitmojis"
}
```

## Configuration Directory

- Located at `~/.gitmojigo`
- Contains `sources.json` and cached emoji data

## Setting Preferences

Run:

```bash
gitmoji -g
```

to interactively set preferences.

---
