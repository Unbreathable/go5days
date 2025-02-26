# Tag 1: Basic Control Flow

Heute lernst du, wie man generell mit Go umgeht, indem wir ein kleines Programm bauen, was eigentlich total scheiße und überhaupt nicht nützlich ist, aber warum auch nicht? Wir müssen ja irgendwas machen. Dabei lernst du die absoluten Basics und generell wie man gescheiten Go Code schreibt, und wenn der am Ende nicht gut ist, dann komme ich persönlich vorbei.

## Grammatiken

Ich hatte dieses Programm noch bei mir rumfliegen von der Uni und hab mir gerade eben so gedacht, dass es eigentlich keine schlechte Aufgabe für den Anfang ist. Deswegen warum nicht? Das braucht aber zuerst mal ein bisschen Kontext.

### Was sind eigentlich Grammatiken?

Eigentlich total einfach, du hast Variablen (meist groß geschrieben: A, B, ..) und ein Alphabet (a, b, c, ..). Und jetzt gibt es eben Regeln wie man die Scheißer zusammenschreiben darf. Diese Regeln sind die Grammatik einer Sprache. Das hatten wir in der Uni und das Thema ist übel Arsch, aber das Output von diesem Code ist manchmal echt statisfying. Hier ist mal ein Beispiel wie man die Regeln einer Grammatik formulieren kann (hier jetzt alles mal schön als Strings formuliert, nur damit "" Sinn macht):

```sh
S -> "A" | "B" | ""
A -> "aS"
B -> "bS"
```

Was das heißt in Sprache, die wir verstehen:

```sh
S ist der Start und wird zu "A", "B" oder "".
A wird zu "aS".
B wird zu "bS".
```

Die Idee ist jetzt, dass man mit S anfängt und dann immer wieder durch die Regeln gehen darf. Lange Rede, kurzer Schwanz, hier ist wie das eigentlich aussehen kann:

```sh
S -> A -> aS -> aA -> aA -> aaS -> aaB -> aabS -> aab
```

Ich hoffe es ist einigermaßen klar, wie man das konvertieren kann, falls du Fragen hast, schreib mir einfach auf Discord oder so. Dann können wir das gerne noch mal abklären.

### Dein Programm

Was du jetzt in Go machen sollst, ist eigentlich ganz einfach: Eine Grammatik in der Form wie sie oben steht als Input nehmen und dann zufällig ein Output der Grammatik generieren. Dabei kannst du davon ausgehen, dass die erste Variable in der Datei immer der Start ist und dass jede Variable nur genau einmal in der Datei vorkommt. Dann soll das Programm zufällig zwischen den Outputs die herauskommen entscheiden und die nächste Variable im Output wieder umwandeln bis keine Variablen mehr übrig sind und dann stoppen.

Hab dir dafür auch schon eine `grammar.txt` vorbereitet mit einem bissl komplexeren Beispiel. Da ist auch alles ohne Leerzeichen als man muss nicht so viel beachten.

Ob deine Lösung richtig ist, kannst du glaube ich selber sehen, ansonsten besprechen wir es auch nochmal heute Abend. Hoffe es ist nicht zu annoying.
