# Tag 2: Goroutines, Channels & die Mutex

Gestern haben wir ja basically einfach nur einen Replacer auf steroids gemacht. Heute soll der aber noch mehr Steroide nehmen. Punkt ist, wir geben ihm die Möglichkeit mehrere Wörter gleichzeitig auszurechnen. Also kopieren wir einfach nur den Code von gestern, setzen aber noch einen drauf.

## Goroutines

Wie du vielleicht schon weißt, gibt es in Go keine Threads oder sonstiges in die Richtung. Es gibt nur Goroutines. Goroutines sind quasi etwas ähnliches wie Threads, die aber von Go für dich gemanaget werden. Wenn also zum Beispiel mal ein richtiger Thread auf deinem Computer eine Goroutine laufen lässt, die dann aber blockt, weil sie zum Beispiel eine Anfrage an einen Server schickt, kann jetzt eine andere Goroutine auf diesem Thread ausgeführt werden. Er wartet ja gerade nur. Sowas geht mit Threads nur sehr schwer und deswegen ist es sehr cool, dass Go dass einfach für uns macht. Nie wieder mit Threads rumhantieren.

Aber wir müssen uns natürlich immer noch um die ganzen Probleme im Multi-Threading kümmern, mehr dazu später.

## Channels

Bevor wir zur Aufgabe kommen, hier noch eine kleine Erklärung zu Channels. Channel sind ein Weg um Daten zwischen Goroutines hin- und herzuschicken. Da kann man sich einfach vorstellen, dass man auf der einen Seite was reinwirft, was dann auf der anderen Seite wieder rauskommt. Mit denen kann man natürlich auch richtig komplexe Scheiße bauen, aber das braucht man eigentlich (fürs erste) nicht zu wissen.

So kann man einen Channel erstellen:

```go
resultChan := make(chan string)
```

Wenn man jetzt will auslesen will was rauskommt, kann man einfach eine neue Variable auf den Wert vom Channel setzen, indem man ihn mit `<-` aus dem Channel rausholt.

```go
result := <-resultChan
```

Man kann damit zum Beispiel einen Receiver bauen, der die ganze Zeit das Ergebnis von einem Channel in eine Variable speichert und dann damit etwas machen könnte.

```go
for {
    result := <-resultChan
    // Und irgendein Code hier
}
```

Hierbei blockt auf das Abfragen des Ergebnisses eines Channels immer die jetzige Goroutine. Das heißt, dass der Code darunter nicht ausgeführt wird, bis etwas aus dem Channel kommt.

## Die heutige Aufgabe

Wie schon vorher erwähnt, machen wir heute den Grammatikrechner schneller und lassen ihn mehr Wörter ausrechnen. Ich hab dir wieder die `grammar.txt` vom letzten Mal hier in den Ordner gepackt. Generell soll das Programm dieses Mal wie folgt funktionieren:

1. Die Regeln werden nach wie vor einfach genauso wie gestern geladen.
2. Jetzt erstellst du `routineAmount` Goroutines (das soll eine Variable am Anfang des Programmes sein, damit man sie später ändern kann), die dann jeweils ein Wort ausrechnen (mit dem Code von gestern).
3. Dann brauchen wir noch einen Channel, kannst du auch genauso machen wie oben, der dann alle Ergebnisse sammelt und am Ende der `main` Funktion noch einen Infinite-Loop der quasi die ganzen Ergebnisse "einsammelt" und sie dann in die Konsole ausgibt (hab ich dir ja auch oben basically schon vorgegeben).

Jetzt wollen wir aber auch wissen, aus welcher Goroutine welches Ergebnis kam. Wir könnten natürlich auch einfach in der Goroutine printen, aber das wäre langweilig. Dazu zuerst ein paar Definitionen.

## Maps

Vielleicht kennst du noch die gute alte `HashMap` aus Java. Genau das gibt es in Go auch. Bloß mit ein paar mehr Problemen, aber dazu unten gleich mehr. In Golang funktioniert das Ding dann so:

```go
// Erstellen einer Map
resultMap := map[int]string{}

// Etwas in die Map hinzufügen
resultMap[0] = "test"
resultMap[100] = "glatze"

// Einen Wert aus der Map auslesen
fmt.Println(resultMap[0]) // test

// Um zu überprüfen ob etwas in einer Map ist, gibt es dir auch einen Boolean zurück
// Hab es hier nur hinzugefügt, weil du es in Zukunft vielleicht brauchst
value, valid := resultMap[10]
if valid {
    fmt.Println("it exists!!")
}
```

So dann versuch doch jetzt einfach mal eine Map zu benutzen um jeweils jedes Ergebnis der jeweiligen Goroutine zuzuordnen. Also eine Map wie oben mit einem Integer und einem String, und dann wird eben das Ergebnis von der ersten Goroutine zum Beispiel so in die Map gepackt:

```go
resultMap[1] = result
```

Wenn du jetzt das ganze mal mit so 100-1000 Goroutines laufen lässt, merkst du schnell, dass da etwas nicht funktioniert. Dafür brauchen wir eine Mutex.

## Mutex

Eine Mutex ist im Prinzip einfach ein Schloss, dass man auf und zu schließen kann. Das funktioniert genau so:

```go
// Erstellen einer neuen Mutex
mutex := &sync.Mutex{} // Ein Pointer, sonst funktioniert die Mutex nicht

// Aufschließen der Mutex
mutex.Unlock() // Nur nach mutex.Lock() ausführen

// Schließen der Mutex
mutex.Lock() // Blockt die jetzige Goroutine bis mutex.Unlock() ausgeführt wird
```

Finde jetzt einmal selber heraus, wie du damit die Fehler mit der Map fixen kannst. Ich hoffe, mit der kleinen Erklärung verstehst du jetzt wie das geht. Ansonsten ist das mein Issue.

Aber das wars dann auch für heute. Ich weiß bissl viel Text, aber hoffe du hast bissl was über Goroutines gelernt. Wir sehen uns morgen!
