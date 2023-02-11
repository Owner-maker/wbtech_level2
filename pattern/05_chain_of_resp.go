package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Поведенческий паттерн

Применимость:
1) Паттерн позволяет создать цепь из обработчиков, и последовательно опрашивая их запускать обработку или нет
2) При необходимости запускать обработчики события в строгом порядке друг за другом
3) Если есть необходимость динамически менять цепь обработчиков

Плюсы:
1) Уменьшение тесной связанности между клиентом и обработчиками -> принцип открытости закрытости

Минусы:
1) Есть вероятность, что запрос останется необработанным (зависит от построения архитектуры)

Примеры:
1) Как уже написано выше данный паттерн идеально подходит там, где есть определенная последовательность действий,
которую необходимо выполнить над запросом, к примеру, когда приходит HTTP - запрос от клиента, и, прежде чем выдать ответ,
необходимо проверить запрос с токеном авторизации на валидность, на правильность запроса, проверить нужный ответ и выдать его

Мой пример:
У нас есть структура Новость, которая должна пройти ряд проверок (сначала ее проверяет сам автор, затем редактор, затем главный редактор
и потом только уже новость попадает на печать).
*/

type News struct {
	Text                string
	isAuthorChecked     bool
	isEditorChecked     bool
	isMainEditorChecked bool
	isPrinted           bool
}

type Check interface {
	Execute(*News)
	SetNext(Check)
}

// первый обработчик - сам автор, поле типа Check, указывающего на следующего обработчика

type Author struct {
	next Check
}

func (a *Author) SetNext(next Check) {
	a.next = next
}

func (a *Author) Execute(n *News) {
	if n.isAuthorChecked { // если этап уже пройден -> переходим к следующему
		a.next.Execute(n)
		return
	}
	n.isAuthorChecked = true // в противном случае этап успешно завершен
	a.next.Execute(n)        // также переходим к следующему
}

// обработчик - редактор

type Editor struct {
	next Check
}

func (e *Editor) SetNext(next Check) {
	e.next = next
}

func (e *Editor) Execute(n *News) {
	if n.isEditorChecked {
		e.next.Execute(n)
		return
	}
	n.isEditorChecked = true
	e.next.Execute(n)
}

// обработчик - главный редактор

type MainEditor struct {
	next Check
}

func (m *MainEditor) SetNext(next Check) {
	m.next = next
}

func (m *MainEditor) Execute(n *News) {
	if n.isMainEditorChecked {
		m.next.Execute(n)
		return
	}
	n.isMainEditorChecked = true
	m.next.Execute(n)
}

// обработчик - отвечающий за печать

type Printer struct {
	next Check
}

func (p *Printer) SetNext(next Check) {
	p.next = next
}

func (p *Printer) Execute(n *News) {
	if n.isPrinted { // если этат пройден -> выходим из функции: следующего этапа нет
		return
	}
	n.isPrinted = true
}

// пример использования

func chainOfRespTest() { // создаем сущности с конца, т.к. необходимо в поля обработчиков при их инициализации указывать сущность для следующей обработки
	printer := Printer{}

	mainEditor := MainEditor{}
	mainEditor.SetNext(&printer)

	editor := Editor{}
	editor.SetNext(&mainEditor)

	author := Author{}
	author.SetNext(&editor)

	news := News{Text: "TEXTTTT"}
	author.Execute(&news)

	fmt.Print(news)
}
