## Tag 3: Die ersten Schritte für AirPlay

So, heute bauen wir endlich mal was sinnvolles. Und zwar für unseren AirPlay Clone. Das Ziel ist heute einfach schonmal die scheiß basics auf die niemand Bock hat zu schreiben. Das muss man nämlich auch machen. Aber erstmal die Basics für heute.

## Structs

Zuerst mal ganz einfach Objekte, kann man sich vorstellen wie Klassen in Java oder so, aber halt mit viel weniger Features und weniger Zeug. Hier zum Beispiel mal wie man einen Kopf repräsentieren kann:

```go
type Head struct {
    HairAmount int
    SkinColor  Color
}
```

Um einen Struct dann auch zu benutzen hier als Beispiel mal dein Kopf:

```go
head := Head{
    HairAmount: 0,
    SkinColor: colors.White, // Existiert nicht, aber egal
}
```

Man kann auch ein Feld einfach leer lassen, dann wird alles auf seine default value gesetzt:

```go
head2 := Head{
    SkinColor: colors.White,
}

// Die default value von einem Int ist zum Beispiel 0, deswegen wird hier 0 ausgegeben
fmt.Println(head2.HairAmount) // 0
```

Deswegen füllen wir am Besten die Structs immer ganz auf.

## AirPlay Codes

Was du auch noch für heute brauchst: Wie funktioniert eigentlich AirPlay (von einer Benutzer Perspektive)? Wir wollen nämlich jetzt die Grundsteine für nächste Woche legen und werden nach und nach darauf aufbauen.

Unser AirPlay wird wie folgt funktionieren:

1. Ein `Receiver` registriert sich beim Server und bekommt einen Token zurück.
2. Jetzt kann ein anderer Client (der `Sender`) eine Request an den Server schicken und damit auch wieder einen Token für sich generieren.
3. Der `Receiver` zeigt jetzt den generierten Code an, den dann der `Sender` abschreiben muss und an den Server senden muss.

## Heutige Aufgabe

Deine Aufgabe ist eine ganze einfache und zwar genau dieses Zeug in einfachen Go-Code umzusetzen. Natürlich auch mit der Hilfe von Structs. Ich hab dir diesmal in `util.go` schon ein paar Methoden vorbereitet und auch ein paar Tests, damit deine Methode auch genau das machen, was sie sollen (wir bauen ja darauf und ich will, dass du am Ende das selbe hast wie ich, weil sonst mehr Arbeit).

Hier noch ein paar Requirements:

- Der Code des `Sender` soll nur aus Zahlen bestehen und genau 6 Zeichen lang sein
- Die beiden Tokens sollen jeweils 12 Zeichen lang sein

Die Tests dafür sind in `receiver_test.go` (natürlich AI-Generated, hab aber noch bissl geedited und so) und ein paar Methoden (ohne Lösung natürlich) sind in `receiver.go`. Du musst heute einfach nur diese Methoden mit der obigen Beschreibung und den Tests zusammen lösen. Viel Glück.
