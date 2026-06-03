package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// ── Save file location ────────────────────────────────────────────────────────

func saveFilePath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = "."
	}
	return filepath.Join(dir, "gogotchi", "save.json")
}

// ── Pet ───────────────────────────────────────────────────────────────────────

type Pet struct {
	mu sync.Mutex `json:"-"`

	Name      string    `json:"name"`
	Hunger    int       `json:"hunger"`
	Happiness int       `json:"happiness"`
	Sleep     int       `json:"sleep"`
	Health    int       `json:"health"`
	Age       int       `json:"age"`
	LastSeen  time.Time `json:"last_seen"`
	Dead      bool      `json:"dead"`
}

func NewPet(name string) *Pet {
	return &Pet{
		Name:      name,
		Hunger:    20,
		Happiness: 80,
		Sleep:     80,
		Health:    100,
		LastSeen:  time.Now(),
	}
}

func (p *Pet) applyOfflineDecay() string {
	if p.LastSeen.IsZero() {
		return ""
	}
	elapsed := time.Since(p.LastSeen)
	ticks := int(elapsed.Seconds() / 10)
	if ticks == 0 {
		return ""
	}
	if ticks > 500 {
		ticks = 500
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	for i := 0; i < ticks; i++ {
		p.tick()
	}
	mins := int(elapsed.Minutes())
	return fmt.Sprintf("Away ~%d min — %d ticks of decay applied.", mins, ticks)
}

// tick must be called with mu held.
func (p *Pet) tick() {
	if p.Dead {
		return
	}
	p.Hunger = clamp(p.Hunger+2, 0, 100)
	p.Happiness = clamp(p.Happiness-1, 0, 100)
	p.Sleep = clamp(p.Sleep-1, 0, 100)
	p.Age++
	if p.Hunger >= 80 || p.Sleep <= 10 || p.Happiness <= 10 {
		p.Health = clamp(p.Health-2, 0, 100)
	} else if p.Hunger < 40 && p.Sleep > 60 && p.Happiness > 60 {
		p.Health = clamp(p.Health+1, 0, 100)
	}
	if p.Health <= 0 {
		p.Dead = true
	}
}

func (p *Pet) Feed() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Hunger = clamp(p.Hunger-20, 0, 100)
	p.Happiness = clamp(p.Happiness+5, 0, 100)
	return randomFrom("feed")
}

func (p *Pet) Play() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Happiness = clamp(p.Happiness+15, 0, 100)
	p.Sleep = clamp(p.Sleep-10, 0, 100)
	p.Hunger = clamp(p.Hunger+5, 0, 100)
	return randomFrom("play")
}

func (p *Pet) Clean() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Health = clamp(p.Health+10, 0, 100)
	return randomFrom("clean")
}

func (p *Pet) Rest() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Sleep = clamp(p.Sleep+30, 0, 100)
	p.Hunger = clamp(p.Hunger+5, 0, 100)
	return randomFrom("rest")
}

func (p *Pet) Snapshot() Pet {
	p.mu.Lock()
	defer p.mu.Unlock()
	return *p
}

// ── Sprites & mood ────────────────────────────────────────────────────────────

func getSprite(p Pet) string {
	if p.Dead {
		return "💀"
	}
	switch {
	case p.Age < 10:
		return "🥚"
	case p.Age < 30:
		return "🐣"
	case p.Age < 80:
		return "🐥"
	case p.Age < 200:
		return "🐤"
	case p.Age < 400:
		return "🐔"
	default:
		return "🦅"
	}
}

func getMood(p Pet) string {
	if p.Dead {
		return "💔 has passed away..."
	}
	switch {
	case p.Hunger >= 80:
		return "😫 is starving!"
	case p.Sleep <= 10:
		return "😪 is exhausted!"
	case p.Health <= 30:
		return "🤒 is very sick!"
	case p.Happiness >= 80:
		return "😄 is delighted!"
	case p.Happiness <= 20:
		return "😢 is sad..."
	case p.Hunger <= 20:
		return "😊 is well-fed and happy"
	default:
		return "😐 is doing okay"
	}
}

var flavourEmojis = map[string][]string{
	"feed":  {"🍔", "🍕", "🍎", "🥗", "🍜", "🍣", "🍧", "🍦", "🍉", "🥛", "🍟", "🥪"},
	"play":  {"🎮", "🕹️", "🎲", "🃏", "🎯", "⚽", "🏀", "🏈", "⚾", "🎾", "🏓", "🏸"},
	"clean": {"🧹", "🧼", "🧽", "🧺", "🧻", "🚿", "🛁", "🪣", "🧴", "🧤"},
	"rest":  {"🛌", "😴", "💤", "🛏️", "🦥", "🧸", "🌙", "🌜", "🌛", "🌚", "🌕", "🌟"},
}

func randomFrom(action string) string {
	opts := flavourEmojis[action]
	if len(opts) == 0 {
		return ""
	}
	return opts[rand.Intn(len(opts))]
}

// ── Persistence ───────────────────────────────────────────────────────────────

func loadPet() (*Pet, bool) {
	data, err := os.ReadFile(saveFilePath())
	if err != nil {
		return nil, false
	}
	var p Pet
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, false
	}
	return &p, true
}

func savePet(p *Pet) {
	if p == nil {
		return
	}
	p.mu.Lock()
	p.LastSeen = time.Now()
	p.mu.Unlock()

	path := saveFilePath()
	os.MkdirAll(filepath.Dir(path), 0755)
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(path, data, 0644)
}

func deleteSave() {
	os.Remove(saveFilePath())
}

// ── UI helpers ────────────────────────────────────────────────────────────────

func bar(value, maxVal, width int) string {
	filled := value * width / maxVal
	result := ""
	for i := 0; i < width; i++ {
		if i < filled {
			result += "█"
		} else {
			result += "░"
		}
	}
	return result
}

func statColor(value int, inverted bool) string {
	if inverted {
		if value >= 70 {
			return "red"
		} else if value >= 40 {
			return "yellow"
		}
		return "green"
	}
	if value <= 30 {
		return "red"
	} else if value <= 60 {
		return "yellow"
	}
	return "green"
}

func renderStatus(p Pet, lastAction string) string {
	s := fmt.Sprintf("\n  %s  [white]%s[white] %s\n\n", getSprite(p), p.Name, getMood(p))
	barW := 20
	s += fmt.Sprintf("  [white]Hunger    [%s]%s[white] %3d\n", statColor(p.Hunger, true), bar(p.Hunger, 100, barW), p.Hunger)
	s += fmt.Sprintf("  [white]Happiness [%s]%s[white] %3d\n", statColor(p.Happiness, false), bar(p.Happiness, 100, barW), p.Happiness)
	s += fmt.Sprintf("  [white]Sleep     [%s]%s[white] %3d\n", statColor(p.Sleep, false), bar(p.Sleep, 100, barW), p.Sleep)
	s += fmt.Sprintf("  [white]Health    [%s]%s[white] %3d\n", statColor(p.Health, false), bar(p.Health, 100, barW), p.Health)
	s += fmt.Sprintf("\n  [grey]Age: %d ticks[white]\n", p.Age)
	if lastAction != "" {
		s += fmt.Sprintf("\n  %s\n", lastAction)
	}
	return s
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	app := tview.NewApplication()
	pages := tview.NewPages()

	var pet *Pet
	var lastAction string

	// ── Status panel (left) ───────────────────────────────────────────────────
	statusView := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true)
	statusView.SetBorder(true).SetTitle(" gogotchi ").SetTitleAlign(tview.AlignCenter)

	// ── Log panel (right, top) ────────────────────────────────────────────────
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetWordWrap(true).
		SetScrollable(true)
	logView.SetBorder(true).SetTitle(" log ").SetTitleAlign(tview.AlignCenter)

	// addLog and redrawStatus must only be called from the main tview goroutine
	// (i.e. inside QueueUpdateDraw callbacks or directly in button handlers).
	addLog := func(msg string) {
		fmt.Fprintf(logView, "[grey]%s[white] %s\n", time.Now().Format("15:04:05"), msg)
		logView.ScrollToEnd()
	}

	redrawStatus := func() {
		if pet == nil {
			return
		}
		statusView.Clear()
		fmt.Fprint(statusView, renderStatus(pet.Snapshot(), lastAction))
	}

	// ── Buttons (right, bottom) ───────────────────────────────────────────────
	mkBtn := func(label string, fn func()) *tview.Button {
		return tview.NewButton(label).SetSelectedFunc(fn)
	}

	buttonBar := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(mkBtn("🍔 Feed", func() {
			if pet == nil || pet.Dead {
				return
			}
			lastAction = pet.Feed() + " Fed!"
			addLog(lastAction)
			redrawStatus()
		}), 0, 1, true).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(mkBtn("🎮 Play", func() {
			if pet == nil || pet.Dead {
				return
			}
			lastAction = pet.Play() + " Played!"
			addLog(lastAction)
			redrawStatus()
		}), 0, 1, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(mkBtn("🧼 Clean", func() {
			if pet == nil || pet.Dead {
				return
			}
			lastAction = pet.Clean() + " Cleaned!"
			addLog(lastAction)
			redrawStatus()
		}), 0, 1, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(mkBtn("💤 Rest", func() {
			if pet == nil || pet.Dead {
				return
			}
			lastAction = pet.Rest() + " Rested!"
			addLog(lastAction)
			redrawStatus()
		}), 0, 1, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(mkBtn("💾 Save", func() {
			savePet(pet)
			addLog("[green]Saved!")
		}), 0, 1, false).
		AddItem(tview.NewBox(), 1, 0, false).
		AddItem(mkBtn("☠️  Reset", func() {
			deleteSave()
			pet = nil
			addLog("[red]Save deleted — quit and restart for a new pet.")
		}), 0, 1, false)

	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(logView, 0, 1, false).
		AddItem(buttonBar, 3, 0, true)

	mainFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(statusView, 0, 1, false).
		AddItem(rightPane, 0, 1, true)

	// ── Tick goroutine — started once when game begins ─────────────────────
	startTicker := func() {
		go func() {
			ticker := time.NewTicker(10 * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				pet.mu.Lock()
				pet.tick()
				isDead := pet.Dead
				pet.mu.Unlock()

				app.QueueUpdateDraw(func() {
					if isDead {
						lastAction = "💀 Your pet has died... Reset to start over."
						addLog("[red]Your pet has died.")
					}
					redrawStatus()
				})
			}
		}()

		go func() {
			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				savePet(pet)
			}
		}()
	}

	// ── Name screen ───────────────────────────────────────────────────────────
	nameVal := "Tama"
	nameForm := tview.NewForm().
		AddInputField("Name your pet:", "Tama", 20, nil, func(text string) {
			nameVal = text
		}).
		AddButton("Start!", func() {
			if nameVal == "" {
				nameVal = "Tama"
			}
			pet = NewPet(nameVal)
			pages.SwitchToPage("game")
			app.SetFocus(buttonBar)
			lastAction = fmt.Sprintf("Welcome, %s! Take good care of them.", pet.Name)
			addLog(lastAction)
			redrawStatus()
			startTicker()
		})
	nameForm.SetBorder(true).SetTitle(" 🥚 gogotchi ").SetTitleAlign(tview.AlignCenter)

	pages.AddPage("name", nameForm, true, true)
	pages.AddPage("game", mainFlex, true, false)

	// ── Load existing save ────────────────────────────────────────────────────
	if existing, found := loadPet(); found && !existing.Dead {
		pet = existing
		offlineMsg := pet.applyOfflineDecay()
		if offlineMsg == "" {
			offlineMsg = fmt.Sprintf("Welcome back! %s missed you.", pet.Name)
		}
		pages.SwitchToPage("game")
		app.SetFocus(buttonBar)
		lastAction = offlineMsg
		addLog(offlineMsg)
		redrawStatus()
		startTicker()
	}

	// Handle Ctrl+C / SIGTERM cleanly — defers don't run on signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		app.Stop()
	}()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' || event.Key() == tcell.KeyEscape {
			app.Stop()
			return nil
		}
		return event
	})

	defer savePet(pet)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
