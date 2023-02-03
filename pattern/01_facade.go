package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
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

func (f Factory) start() { // главный метод, предоставляющий доступ к запуску / работе всех сущностей
	f.loader.loadObject()
	f.conveyor.startWorking()
	f.handler.startHandling()
}
