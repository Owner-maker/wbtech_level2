package pattern

import (
	"errors"
	"fmt"
)

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
*/

type PlayerBuilder interface { // интерфейс билдера для игрока
	SetHealth(int) error
	SetGun(*Gun) error
	GetPlayer() *Player
}

type PlayerBuilderImpl struct { // конкретная имлементация билдера игрока с полем типа *Player
	player *Player
}

// имплементация методов билдера игрока

func (p *PlayerBuilderImpl) SetHealth(h int) error {
	if h <= 0 {
		return errors.New("health can not be negative or equals zero")
	}
	p.player.health = h
	return nil
}

func (p *PlayerBuilderImpl) SetGun(g *Gun) error {
	if g == nil {
		return errors.New("gun can not be nil")
	}
	p.player.gun = g
	return nil
}

func (p *PlayerBuilderImpl) GetPlayer() *Player {
	return &Player{
		health: p.player.health,
		gun:    p.player.gun,
	}
}

type GunBuilder interface { // интерфейс билдера для оружия игрока
	SetGunType(string) error
	SetAmmo(int) error
	SetAmmoToReload(int) error
	GetGun() *Gun
}

type GunBuilderImpl struct { // конретная имплементация билдера для оружия игрока
	gun *Gun
}

// имплементация методов билдера оружия

func (g *GunBuilderImpl) GetGun() *Gun {
	return &Gun{
		gunType:      g.gun.gunType,
		ammo:         g.gun.ammo,
		ammoToReload: g.gun.ammoToReload,
	}
}

func (g *GunBuilderImpl) SetGunType(t string) error {
	if t == "" {
		return errors.New("gun type can not be empty")
	}
	g.gun.gunType = t
	return nil
}

func (g *GunBuilderImpl) SetAmmo(a int) error {
	if a <= 0 {
		return errors.New("ammo value can not be negative or equals zero")
	}
	g.gun.ammo = a
	return nil
}

func (g *GunBuilderImpl) SetAmmoToReload(a int) error {
	if a <= 0 {
		return errors.New("ammo to reload value can not be negative or equals zero")
	}
	g.gun.ammoToReload = a
	return nil
}

// структура самого игрока, где есть поле типа *Gun

type Player struct {
	health int
	gun    *Gun
}

func (p Player) ShowInfo() {
	fmt.Printf("Player: [health] %d\n", p.health)
	fmt.Printf("Gun [gun type] %s ;[ammo] %d; [ammo to reload] %d\n", p.gun.gunType, p.gun.ammo, p.gun.ammoToReload)
}

type Gun struct {
	gunType      string
	ammo         int
	ammoToReload int
}

// структура Game с двумя билдерами (для игрока и его оружия) интерфейсных типов

type Game struct {
	playerBuilder PlayerBuilder
	gunBuilder    GunBuilder
}

// основной метод создания объекта игрока и его оружия

func (g Game) build(health int, gunType string, ammo, ammoReload int) (*Player, error) {
	err := g.gunBuilder.SetGunType(gunType)
	if err != nil {
		return nil, err
	}

	err = g.gunBuilder.SetAmmo(ammo)
	if err != nil {
		return nil, err
	}

	err = g.gunBuilder.SetAmmoToReload(ammoReload)
	if err != nil {
		return nil, err
	}

	err = g.playerBuilder.SetHealth(health)
	if err != nil {
		return nil, err
	}

	err = g.playerBuilder.SetGun(g.gunBuilder.GetGun())
	if err != nil {
		return nil, err
	}
	return g.playerBuilder.GetPlayer(), nil
}
