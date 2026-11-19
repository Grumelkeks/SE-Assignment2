// Composite Pattern am Beispiel eines Fensters.
//
//	Item (Component):  Price(), Display()
//	Glass, Frame:      Leafs
//	Window:            Composite, enthält 1..* Items (auch weitere Windows)
package main

import (
	"fmt"
	"strings"
)

// Item ist die gemeinsame Schnittstelle für Leafs und Composites.
type Item interface {
	Price() float64
	Display(indent int)
}

// ---------- Leaf: Glass ----------

// Glass ist eine einzelne Glasscheibe.
type Glass struct {
	Name          string
	Width, Height float64 // in m
	PricePerSqm   float64 // € pro m²
}

func (g Glass) Price() float64 { return g.Width * g.Height * g.PricePerSqm }

func (g Glass) Display(indent int) {
	fmt.Printf("%sGlass %s (%.2f x %.2f m): %.2f €\n",
		pad(indent), g.Name, g.Width, g.Height, g.Price())
}

// ---------- Leaf: Frame ----------

// Frame ist der Rahmen eines Fensters; Preis nach Umfang.
type Frame struct {
	Width, Height float64 // in m
	PricePerMeter float64 // € pro Laufmeter
}

func (f Frame) Price() float64 { return 2 * (f.Width + f.Height) * f.PricePerMeter }

func (f Frame) Display(indent int) {
	fmt.Printf("%sFrame (%.2f x %.2f m): %.2f €\n",
		pad(indent), f.Width, f.Height, f.Price())
}

// ---------- Composite: Window ----------

// Window besteht aus 1..* Items (Rahmen, Scheiben, Unterfenster).
type Window struct {
	Name     string
	children []Item
}

// NewWindow erzwingt die Multiplizität 1..*.
func NewWindow(name string, first Item, more ...Item) *Window {
	return &Window{Name: name, children: append([]Item{first}, more...)}
}

func (w *Window) Add(items ...Item) { w.children = append(w.children, items...) }

func (w *Window) Price() float64 {
	sum := 0.0
	for _, c := range w.children {
		sum += c.Price()
	}
	return sum
}

func (w *Window) Display(indent int) {
	fmt.Printf("%sWindow %s: %.2f €\n", pad(indent), w.Name, w.Price())
	for _, c := range w.children {
		c.Display(indent + 1)
	}
}

func pad(n int) string { return strings.Repeat("  ", n) }

// ---------- Client ----------

func main() {
	const (
		glassPrice = 80.0 // €/m²
		framePrice = 25.0 // €/m
	)

	// Unterfenster rechts unten: Rahmen + 2x2 Scheiben
	sub := NewWindow("rechts-unten",
		Frame{Width: 0.6, Height: 0.6, PricePerMeter: framePrice},
		Glass{"RU-1", 0.3, 0.3, glassPrice},
		Glass{"RU-2", 0.3, 0.3, glassPrice},
		Glass{"RU-3", 0.3, 0.3, glassPrice},
		Glass{"RU-4", 0.3, 0.3, glassPrice},
	)

	// Hauptfenster: Rahmen + 3 Scheiben + Unterfenster (2x2 geteilt)
	window := NewWindow("Hauptfenster",
		Frame{Width: 1.2, Height: 1.2, PricePerMeter: framePrice},
		Glass{"links-oben", 0.6, 0.6, glassPrice},
		Glass{"rechts-oben", 0.6, 0.6, glassPrice},
		Glass{"links-unten", 0.6, 0.6, glassPrice},
	)
	window.Add(sub)

	var item Item = window // Client arbeitet nur mit dem Interface
	item.Display(0)
	fmt.Printf("\nGesamtpreis: %.2f €\n", item.Price())
}
