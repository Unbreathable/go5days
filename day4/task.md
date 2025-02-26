# Tag 4: Fiber & ein einfacher Counter

Die letzten Tage haben wir ja nur irgendwelche komischen Beispiele in Go gemacht, woran sich natürlich auch nichts ändern wird, aber heute geht es mal richtig los mit etwas was man wirklich mal einsetzen kann: Dem Web-Framework [Fiber](https://gofiber.io). Falls du irgendwelche Sachen brauchst, einfach dort nachschauen.

## Ganz einfache Sachen mit Fiber

Dazu brauchts nicht viel Erklärung glaube ich, deswegen mal einfach ein kleines Beispiel:

```go

// Nur damit du die richtigen Sachen importierst (Go macht das manchmal falsch)
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// ...

// Eine neue App mit Fiber erstellen
app := fiber.New()

// Den Logger, damit man weiß welche Endpoints gehittet wurden
app.Use(logger.New())

// Einen neuen Endpoint erstellen
app.Get("/hello", func(c *fiber.Ctx) error {
    return c.SendString("Hello, world!")
})

// Den Server starten
app.Listen("127.0.0.1:3000")

```

Wenn du jetzt die das alles in eine `main` Funktion packst, dann solltest du einfach zu `http://localhost:3000/hello` im Browser gehen können. Dort sollte dann einfach `Hello, world!` stehen.

Natürlich gibt es in Fiber noch so viel mehr, aber viel mehr als diese einfache Struktur brauchen wir eigentlich für unsere Projekte in diesem kleinen Kurs nicht. Deswegen scheiß drauf.

## Deine heutige Aufgabe

Wie schon vorhin gesagt, wir bauen heute eigentlich eine ganz einfache API. Hier ist ein einfacher Überblick:

- `POST /[id]`: Erhört den Counter `id` um eins und gibt den Wert des Counters zurück.
- `GET /[id]`: Gibt den Wert des Counters `id` zurück.

Damit kann man also mehrere Counter erstellen. Hier ein Beispiel wie man es aufrufen könnte:

```sh
# Tipp: curl geht nur im alten CMD (nicht in PowerShell)
curl -X POST /hello # 1
curl /hello # 1
curl /other # 0
curl -X POST /other # 1
curl /other # 0
```

Du musst dann auch noch in der Fiber Documentation nachgucken, wie man Path Variablen hinzufügt. Also so damit man auch abfragen kann, was nach dem Slash steht.

### Aufgabe

**1.** Als erstes bearbeitest du `counter.go` und implementierst dort die elementaren Elemente der Counter. Dazu musst du auch mal online recherchieren, wie man eine `sync.Map` verwendet, damit einem nicht die Concurrent-Exceptions um die Ohren gehauen werden, wenn man gleichzeitig etwas speichert und liest.

**2.** Teste deinen Counter indem du `go test` in die Kommandozeile eingibst.

**3.** Erstelle jetzt in `app.go` die Endpoints wie oben angegeben und lasse sie durch meine Tests laufen. Und dann gucken wir mal ob du alles richtig gemacht hast ^^

**4.** Überprüfe deine API indem du in `app.go` gehst und dort `APITests` auf `true` setzt und dann `go test` ausführst.

Und dann ist auch für heute alles geschafft. Good job ^^
