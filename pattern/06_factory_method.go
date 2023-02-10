package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Порождающий паттерн

Применимость:
1) Необходимость разделения создания структур от их использования, а именно когда мы заранее не знаем какую именно структуру будем использовать
2) Возможность расширяться в будущем, создавая новые типы, не ломая старой логики создания объектов
3) Повторное использованиеуже существующих объектов

Плюсы:
1) Разделение создания объектов от их использования
2) Упрощенное добавление новых объектов -> принцип закрытости открытости

Минусы:
1) Создание дополнительных иерархий, усложнение читаемости кода

Примеры:
1) Мы используем определенный UI - фреймворк и в определенный момент у нас появилась потребность в создании текстового поля другого оформления

Мой пример:
Небходимо иметь возможность создавать определенную банковскую карту - дебетовую или кредитную, в зависимости от входных параметров.
Так как в Го нет классического наследования, можно реализовать данный паттерн с использованием встраивания
*/

// интерфейс банковского аккаунта

type IBankAccount interface {
	setBalance(float64)
	getBalance() float64
	setMoneyLimit(float64)
	getMoneyLimit() float64
	setAnnualInterest(float64)
	getAnnualInterest() float64
}

// структура банковского аккаунта с нужными полями

type BankAccount struct {
	balance        float64
	moneyLimit     float64
	annualInterest float64
}

// имплементация методов

func (b *BankAccount) setMoneyLimit(f float64) {
	b.moneyLimit = f
}

func (b *BankAccount) getMoneyLimit() float64 {
	return b.moneyLimit
}

func (b *BankAccount) setAnnualInterest(f float64) {
	b.annualInterest = f
}

func (b *BankAccount) getAnnualInterest() float64 {
	return b.annualInterest
}

func (b *BankAccount) setBalance(v float64) {
	b.balance = v
}

func (b *BankAccount) getBalance() float64 {
	return b.balance
}

// структура дебетового аккаунта со встраиванием анонимного поля типа BankAccount

type DebitAccount struct {
	BankAccount
}

// метод, возвращающий конкретный объект дебетового аккаунта

func newDebitAccount(balance, limit, percent float64) IBankAccount {
	return &DebitAccount{
		BankAccount: BankAccount{
			balance:        balance,
			moneyLimit:     limit,
			annualInterest: percent,
		},
	}
}

type CreditAccount struct {
	BankAccount
}

func newCreditAccount(balance, limit, percent float64) IBankAccount {
	return &CreditAccount{
		BankAccount: BankAccount{
			balance:        balance,
			moneyLimit:     limit,
			annualInterest: percent,
		},
	}
}

func getBankAccount(accountType string, balance, limit, percent float64) (IBankAccount, error) {
	if accountType == "debitAccount" {
		return newCreditAccount(balance, limit, percent), nil
	}
	if accountType == "creditAccount" {
		return newDebitAccount(balance, limit, percent), nil
	}
	return nil, errors.New("wrong account type")
}

func factoryMethodTest() {
	creditAcc, _ := getBankAccount("creditAccount", 0, 100_000, 20)
	fmt.Print(creditAcc)
}
