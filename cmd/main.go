package cmd

import (
	c "context" // alias (именнованный тип пакета)
	. "fmt"     // к данным пакетам обращаться не надо, сразу функцию

	_ "github.com/Veylin2175/alotofmodules/pkg" // вызывается только init из данного пакета
	"golang.org/x/sync/errgroup"
)

func Priiint() {
	Println("Функция была вызвана!") // не пишем fmt., т.к. fmt импортировано с точкой
}

func End() {
	Println("Меня должны вызвать в самом конце!")
}

func main() {
	g, ctx := errgroup.WithContext(c.Background())
	Println(g, ctx)
	defer End()
	Priiint()
	Haha()
	Println("Еще немного выводов...")
}

// std - внутренние пакеты Go (fmt, context, os etc.)
// external (github) => go get ...
// local - пакеты проекта (alotofmodules/pkg)

// builtin - пакет, из которого можно брать методы без обращения к пакету (print, все типы и т.д.)

// Маленькая / заглавная буквы - модификаторы доступа

// go mod init module_name - инициализация go.mod
// go mod tidy - актуализация go.mod
// go mod vendor - скачивает все зависимости в папку проекта

// go mod init github.com/$username/$project_name
