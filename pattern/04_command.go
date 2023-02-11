package pattern

import "time"

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Поведенческий паттерн

Применимость:
1) При необходимости добавить дополнительные параметры с конкретным действием: паттерн позволяет превратить операцию в объект, добавить
параметры и передать этот объект дальней для обработки
2) Сохранять команды, путем сериализации и десерелизации превращать в строку или в объект, чтобы выполнить позднее по таймауту или передать по сети
3) Возможность сохранения состояний (в совокупности с паттерном Снимок), для дальнейшего роллбэка при необходимости.

Плюсы:
1) Добавление посредничества между отправителем "действия" и получателем, что позволяет на лету менять реализации команд, т.к.
отправитель не знает кому отправляет, а получатель -> от кого приходит команда
2) Возможность реализации отмены
3) Благодаря тому, что команда - отдельный объект,в который мы добавляем нужную нам логику, не изменяя оригинальной структуры,
мы следуем принципу открытости - закрытости
4) Возможность реализовать действие по таймауту
5) Возможность агрегации команд из более простых

Минусы:
1) Усложнение кода из-за вода дополнительных структур

Примеры:
1) Банковский аккаунт. Простой функционал -> мы можем пополнить баланс, можем снять средства. Но пользователь не единственный, кто
может осуществлять напрямую эти операции (к примеру покупка в Интернет магазине, продление подписки и т.д.)

Мой пример:
У нас есть структура Машины. Ее двигатель можно завести несколькими способами:
1) Используя ключ дистанционного зажигания
2) Нажав кнопку в машине
3) Используя мобильное приложение (к примеру)

Также мы к примеру хотим иметь возможность осуществить зажигание двигателя для прогрева машины по истечению времени, заранее
*/

type Machine interface { // интерфейс получателя (в нашем случае - любая структура, двигатель к-ой можно вкл и выкл)
	turnEngineOn()
	turnEngineOff()
}

type Car struct { // структура Машины
	isEngineRunning bool
}

// методы включения и выключения двигателя Машины

func (c Car) turnEngineOn() {
	if !c.isEngineRunning {
		c.isEngineRunning = true
	}
}

func (c Car) turnEngineOff() {
	if c.isEngineRunning {
		c.isEngineRunning = false
	}
}

type Command interface { // интерфейс команды
	execute()
}

type TimeCommand interface { // интерфейс команды с заданием времени
	execute(int)
}

// конкретная реализация команды, полем которой является Machine

type OnCommand struct {
	machine Machine
}

func (o *OnCommand) execute() {
	o.machine.turnEngineOn()
}

type OffCommand struct {
	machine Machine
}

func (o *OffCommand) execute() {
	o.machine.turnEngineOff()
}

// конкретная реализация команды, у к-ой также полем является Machine, но теперь вкл и выкл двигатель можно с заданием времени

type OnTimeCommand struct {
	machine Machine
}

func (o *OnTimeCommand) execute(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
	o.machine.turnEngineOn()
}

type OffTimeCommand struct {
	machine Machine
}

func (o *OffTimeCommand) execute(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
	o.machine.turnEngineOff()
}

type Keys struct { // Отправитель - ключи машины
	command Command
}

func (k *Keys) press() {
	k.command.execute()
}

type SmartPhone struct {
	command TimeCommand
}

func (s *SmartPhone) press(value int) { // Отправитель - смартфон
	s.command.execute(value)
}

// пример использования

func testCommand() {
	car := Car{}

	onCommand := OnCommand{machine: car}
	offTimeCommand := OffTimeCommand{machine: car}

	onCommand.execute()
	offTimeCommand.execute(1000)
}
