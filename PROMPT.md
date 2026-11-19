# Prompt

Implementiere in Go (ein Modul, `main.go`, nur Standardbibliothek) das Composite Pattern am Beispiel eines Fensters:

- Component: Interface `Item` mit `Price() float64` und `Display(indent int)`.
- Leafs: `Glass` (Breite × Höhe in m, Preis pro m²) und `Frame` (Rahmen, Preis pro Laufmeter des Umfangs).
- Composite: `Window` enthält 1..* `Item`s (also auch weitere `Window`s); `Price()` = Summe der Kinder, `Display()` gibt den Baum eingerückt mit Teilpreisen aus.
- `main()` baut dieses Fenster: ein Hauptfenster mit Rahmen, 2×2 geteilt; drei Felder sind einfache Glasscheiben, das Feld rechts unten ist ein eigenes Unterfenster mit Rahmen und wiederum 2×2 Glasscheiben. Gib den Baum und den Gesamtpreis aus.

Halte den Code kurz, idiomatisch und kommentiert.
