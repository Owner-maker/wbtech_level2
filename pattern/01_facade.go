package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

Применимость:
1) Простой доступ к сложной системе
2) Возможность разложить систему на слои

Плюсы:
1) Изолирует клиентов от поведения сложной системы

Минусы:
1) Сам интерфейс фасада (или его реализация) может стать супер - объектом, но в то же время есть возможность создать "под-фасад"

Примеры:
1) Есть структуры Банк, Карта пользователя, Магазин, к примеру метод sell() у самого магазина может являться фасадом, предоставляющим
обращение к подсистемам -> к карте пользователя (узнать ее данные), к банку (сверить карту пользователя, баланс) и т.д.
2) Есть фреймфорк, позволяющий заниматься конвертацией видео, но его использование крайне неудобно и запутанно, можно создать фасад - метод convertVideo(),
который будет обращаться к структурам, которые к примеру занимаются аудио, получением файла видео, декодироанием видео и т.д. И данный метод предоставить в
качестве использования системы.

Мой пример:
Есть структуры Loader, Conveyor и Handler (элементы условного завода), в качестве фасада будет выстпуать метод start() у структуры
Factory, к-я содержит в качестве полей все остальные структуры
*/

type Loader struct { // структура - Погрузчик
}

func (l *Loader) loadObject() {
	fmt.Println("Loaders has load an object to the conveyor")
}

type Conveyor struct { // структура - Конвейер
}

func (c Conveyor) startWorking() {
	fmt.Println("The conveyor has started the line")
}

type Handler struct { // структура Обработчик
}

func (h Handler) startHandling() {
	fmt.Println("The handler has started its work")
}

type Factory struct { // структура - Завод, содержащая в себе в качестве полей остальные структуры
	loader   *Loader
	conveyor *Conveyor
	handler  *Handler
}

func NewFactory() *Factory { // конструктор создания
	return &Factory{
		loader:   &Loader{},
		conveyor: &Conveyor{},
		handler:  &Handler{},
	}
}

func (f Factory) Start() { // главный метод, предоставляющий доступ к запуску / работе всех сущностей
	f.loader.loadObject()
	f.conveyor.startWorking()
	f.handler.startHandling()
}

// пример использования

func testFacade() {
	f := NewFactory()
	f.Start()
}
