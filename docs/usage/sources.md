# Emoji Sources

## Overview

- Emoji lists can be loaded from multiple sources.
- Default source: remote JSON file.
- All sources are registered in `sources.json` in `~/.gitmojigo`.

## Adding a Source

Edit `sources.json` to add new sources:

```json
[
  {
    "name": "default",
    "url": "https://gitmoji.dev/api/gitmojis"
  }
]
```

## Caching

- Emoji data is cached in `~/.gitmojigo/sources/default/`.
- Cache is created automatically on first use.

## Parsing

- Each source can have its own parsing logic.
- Default expects a JSON array with fields: emoji, description, category, aliases, tags, version.

---
