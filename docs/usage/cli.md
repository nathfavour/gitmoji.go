# CLI Usage

All commands are run via the `gitmoji` CLI.

## Basic Usage

```bash
gitmoji [option] [command]
```

## Options

| Option           | Alias | Description                                      |
|------------------|-------|--------------------------------------------------|
| `--commit`       | `-c`  | Interactively commit using prompts               |
| `--config`       | `-g`  | Setup preferences                                |
| `--init`         | `-i`  | Initialize as a commit hook                      |
| `--list`         | `-l`  | List all available emojis                        |
| `--remove`       | `-r`  | Remove a previously initialized commit hook       |
| `--search`       | `-s`  | Search emojis                                    |
| `--update`       | `-u`  | Sync emoji list                                  |
| `--version`      | `-v`  | Print installed version                          |

## Commands

- `commit` – Start interactive commit prompts.
- `config` – Configure preferences.
- `init` – Set up as a commit hook.
- `list` – Show all emojis.
- `remove` – Remove commit hook.
- `search` – Search for emojis.
- `update` – Update emoji list.

## Examples

```bash
gitmoji -c --title="Fix bug" --message="Fixed a bug" --scope="core"
gitmoji -i
gitmoji -l
gitmoji -s "fix"
gitmoji -u
```

---
