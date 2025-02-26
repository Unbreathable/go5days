# Tag 5: Die API für AirPlay

Heute bauen wir endlich die API, die wir als Basis von unserem kleinen AirPlay Projekt brauchen. Du kennst ja jetzt Fiber und Go ziemlich gut hoffentlich. Deswegen sind die Beschreibungen heute nicht sehr detailreich, wenn irgendwas unklar ist, einfach fragen. Ich hab in der Uni bestimmt noch ein bisschen Zeit für deine Fragen, da ist eh im Moment die ganze Zeit alles andere als Stoff angesagt..

## Noch eine kurze Sache zu Fiber

Du hast zwar bereits eine Menge über Fiber gelernt, aber ein paar Sachen fehlen noch. Und eine der wichtigsten Sachen ist wie man JSON entgegen nimmt und auch wieder zurück gibt, deswegen hier ein ganz kleiner lustiger Endpoint.

```go
// Route: /farmer/get
func getFarmer(c *fiber.Ctx) error {

    // Parse the request
    var req struct {
        Token string `json:"token"` // Wenn eine JSON rein kommt, wird das Feld "token" hier rein gepackt
    }
    if err := c.BodyParser(&req); err != nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }

    // Get the farmer
    farmer := getFarmer(req.Token)
    if farmer == nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }

    // Return JSON
    return c.JSON(fiber.Map{
        "token": farmer.Token,
    })
}
```

## Layout der API

Ich hab dir natürlich wie immer schon sehr viel von der einfachen Datei Struktur vorbereitet. Du musst nur noch deine Lösung von Tag 3 nochmal in die jeweiligen Methoden einfüllen und eventuell anpassen. Wunder dich nicht wegen den Mutexes, die braucht man halt im asynchronen Kontext, hab ich dir jetzt aber zumindest bei den Methoden mal gespart. Aber vergiss nicht sie auch überall bei den Endpoints, wo du auf den Receiver/Sender zugreifst einzubauen, sonst bringt es ja nix.

Ich hab dir natürlich auch wieder Tests gebaut, also einfach immer `go test` ausführen, falls du dir unsicher bist, ob es geht. Da sollten eigentlich alle Sachen überprüft werden.

Hier sind natürlich alles einfach `POST` Requests, ganz im Liphium Style. Sorry ist mir actually erst jetzt aufgefallen (nach dem Machen) und jetzt verändern wir es nicht mehr. In Firmen würde man aus ein paar Endpoints hier auf jeden Fall ein `GET` machen oder so, aber für jetzt juckt es erstmal nicht. Aber egal, hier ist das API Layout:

### `POST /receiver/create`

Erstellt einen neuen Receiver und gibt folgende JSON zurück:

```json
{
  "token": "the_token"
}
```

Falls schon ein Receiver existiert, kannst du einfach `c.SendStatus(fiber.StatusConflict)` zurücksenden.

### `POST /receiver/check_state`

Jetzt wird es schon etwas komplizierter. Dieser Endpoint wird vom Client benutzt, um die ganze Zeit zu checken, ob ein Client verbunden ist und falls einer verbunden ist wird dann auch der Name und Code zurückgegeben. Er gibt dann folgende JSON zurück (du kannst auch alles außer `"exists": false` weglassen, falls kein Sender existiert):

```json
{
  "exists": true, // Ob ein Sender existiert oder nicht
  "connected": true, // Ob der Sender die Verbindung abgeschlossen hat (einfach false returnen)
  "accepted": false, // Ob der Sender den Code richtig eingegeben hat
  "name": "Name", // Name des Senders
  "code": "123456" // Code den der Sender eingeben musst
}
```

Du musst natürlich auch wieder Fehler überprüfen, hier eine kleine Checkliste:

- Der Receiver darf nicht `nil` sein.
- Der Token vom Receiver muss stimmen.

### `POST /sender/create`

So, dieser Endpoint erstellt jetzt einen neuen Sender und funktioniert jetzt quasi identisch wie der für den Receiver. Zurückgegebene JSON sieht dann wieder so aus:

```json
{
  "token": "the_token"
}
```

### `POST /sender/attempt`

Jetzt brauchen wir nur noch einen Endpoint, um auch den eingegeben Code vom Sender zu überprüfen. Das ist dann dieser hier. Falls der Versuch erfolgreich ist, soll einfach eine 200 zurückgegeben werden. Also als Status Code mit `c.SendStatus(fiber.StatusOK)`.

Ich hoffe du hast alles hier verstanden, falls Fragen da sein sollten, mir gerne schreiben.
