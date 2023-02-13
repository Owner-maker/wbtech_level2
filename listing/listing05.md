Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error

У нас имеется кастомная ошибка, которая имплементирует интерфейс Error.
После объвления переменной типа error она бы прошла проверку на соотвествие nil, т.к.
поле tab и data этого интерфейса были бы nil.

Но после присвоения значения, к-е вернула функция test() -> значения типа *customError,
поле tab теперь не будет nil: в нем будут метаданные о типе customError, соответственно
теперь переменная типа error не пройдет проверку на nil -> err != nil выдаст true.
```
