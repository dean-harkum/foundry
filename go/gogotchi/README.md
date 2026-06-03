# 🥚 gogotchi

A terminal Tamagotchi written in Go

This is for fun only, not a serious project. It is in active development, with more improvements to come

## Requirements

- [Go 1.21+](https://go.dev/dl/)

## Running

```bash
git clone <your-repo-url>
cd gogotchi

go run .
```

That's it. Go will automatically download dependencies on first run

If you'd rather download dependencies separately before running:

```bash
go mod download
go run .
```

To build a standalone binary instead:

```bash
go build -o gogotchi
./gogotchi
```

## Controls

| Key          | Action                |
| ------------ | --------------------- |
| `Tab`        | Cycle between buttons |
| `Enter`      | Press focused button  |
| `q` or `Esc` | Save and quit         |
| `Ctrl+C`     | Save and quit         |

Mouse clicks also work on buttons

## Actions

| Button   | Effect                                    |
| -------- | ----------------------------------------- |
| 🍔 Feed  | Reduces hunger, small happiness boost     |
| 🎮 Play  | Boosts happiness, costs sleep & hunger    |
| 🧼 Clean | Restores health                           |
| 💤 Rest  | Restores sleep, increases hunger slightly |
| 💾 Save  | Manually save your pet                    |
| ☠️ Reset | Delete save file (restart for a new pet)  |

## Stats

| Stat      | Good           | Bad                     |
| --------- | -------------- | ----------------------- |
| Hunger    | Low (0 = full) | High (100 = starving)   |
| Happiness | High           | Low                     |
| Sleep     | High           | Low                     |
| Health    | High           | Low — reaches 0 = death |

Health slowly **recovers** when your pet is thriving (low hunger, high sleep & happiness)
Health **degrades** when your pet is starving, exhausted, or miserable

## Persistence

State is saved to `~/.config/gogotchi/save.json`. When you reopen the game,
offline decay is calculated based on how long you were away (1 tick = 10 real
seconds, capped at 500 ticks). Your pet can die while the app is closed if their
stats were already critical
