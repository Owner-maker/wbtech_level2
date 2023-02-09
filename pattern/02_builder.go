package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Применимость:
1) Необходимость избавиться от большого кол-ва маленьких конструкторов сущности, которые занимаются вызовом друг друга
2) Когда нужно чтобы код создавал разные представления какого-то объекта, т.е. когда создание объектов состоит из
нескольких одинаковых этапов, отличающихся в деталях
3) Когда нужно собрать сложные объекты по частям, иногда рекурсивно (деревья)
4) Единый порядок создания объекта

Плюсы:
1) Пошаговое создание
2) Переиспользование кода
3) Разделение сложного кода сборки объекта от основной бизнес - логики

Минусы:
1) Повышается сложность из-за дополнительных классов
2) Привязка клиентов к конретным билдерам (тесное связывание)

Примеры:
1) Логика "сбора" заказа, будь-то сфера питания, услуг, и т.д., где необходима постепенная настройка (создание) объекта

Мой пример:
Есть структура Игрока с рядом его характеристик (имя, кол-во жизней, вида оружия и т.д.), необходимо создать билдер, который позволит
создавать объект игрока, добавляя ему характеристики постепенно.
*/

type Player struct { // структура игрока
	name         string
	health       int
	gunType      string
	ammo         int
	ammoToReload int
}

type PlayerBuilder struct { // билдера для игрока
	player *Player
}

type PlayerInfoBuilder struct { // структура билдера, отвечающего за задание персоналной информации игрока
	PlayerBuilder
}

type PlayerGunBuilder struct { // структура билдера, отвечающего за задание инормации об оружии игрока
	PlayerBuilder
}

func NewPlayerBuilder() *PlayerBuilder { // конструктор создания билдера игрока (внтури создается объект Игрока с дефолтными значениями)
	return &PlayerBuilder{player: &Player{}}
}

func (p *PlayerBuilder) PlayerInfo() *PlayerInfoBuilder { // метод, возвращающий указатель на Билдер игрока, отвечающего за персональную информацию
	return &PlayerInfoBuilder{*p}
}

func (p *PlayerBuilder) GunInfo() *PlayerGunBuilder { // метод, возвращающий указатель на Билдер оружия игрока
	return &PlayerGunBuilder{*p}
}

// ряд методов билдера PlayerInfoBuilder, которые задают персональную информацию об игроке

func (p *PlayerInfoBuilder) Name(name string) *PlayerInfoBuilder {
	p.player.name = name
	return p
}

func (p *PlayerInfoBuilder) Health(health int) *PlayerInfoBuilder {
	p.player.health = health
	return p
}

// ряд методов билдера PlayerGunBuilder, которые задают информацию об оружии игрока

func (p *PlayerGunBuilder) GunType(gunType string) *PlayerGunBuilder {
	p.player.gunType = gunType
	return p
}

func (p *PlayerGunBuilder) Ammo(ammo int) *PlayerGunBuilder {
	p.player.ammo = ammo
	return p
}

func (p *PlayerGunBuilder) AmmoToReload(ammoToReload int) *PlayerGunBuilder {
	p.player.ammoToReload = ammoToReload
	return p
}

// конечный метод Build билдера игрока, возвращающего настроенный объект игрока

func (p *PlayerBuilder) Build() *Player {
	return p.player
}
