package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Поведенческий паттерн

Применимость:
1) При необходимости над объектами разных классов (структур, интерфейсов) совершить определенную операцию
2) При необходимости совершить разные операции над объектами и при этом не хочется добавлять эти методы в классы (структуры)
3) Нам нужно новое поведение, но не для всех объектов, а для конкретных

Плюсы:
1) Упрощение добавления новых операций (поддержание принципа открытости - закрытости)
2) Объединение схожих операций в одной структуре
3) Возможность накопления состояния при обходе объектов

Минусы:
1) Если иерархия меняется довольно часто, то паттерн становится неудобен из-за того, что придется отслеживать, где и кому какая
операция теперь необходима и необходима ли
2) Нарушение инкапсуляции (спорный момент)

Примеры:
1) Я думаю, что основное применение, когда у нас уже готова реализация какой-то структуры, а нам вдруг нужно добавить новую логику, при это не
ломая старую. Например, есть структура Заказ, все заказы находятся в древовидной структуре или обычном списке, нам необходимо добавить возможность
подсчета суммы одного заказа и его запись в файл, при добавлении данного метода в интерфейс, который имплементируют Заказ, то обратная совместимость может быть
сломана, вот здесь нам и может помочь Посетитель.
Структуру заказа, конечно, придется изменить, добавив метод для реализации паттерна, но только единожды.

Мой пример:
Предположим, у нас есть несколько отправителей - сервисов, которые отправляют данные в некий брокер сообщений. Это сервис погоды, новостей и, к примеру,
рекламы. Каждый из этих сервисов уже умеет отправлять нужные данные, но вдруг нам понадобилась логика, которая позволит нам собирать информацию с этих сервисов,
к примеру, для отладки. Нам нужен метод сбора данных - getMessagesPerDay().
*/

type WeatherService struct { // сервис погоды
	url    string
	briefs int
}

func (s *WeatherService) accept(v Visitor) { // метод, который принимает объект типа Visitor и выполняет его команду
	v.visitForWeatherService(s)
}

type NewsService struct {
	url     string
	authors []string
	news    int
}

func (s *NewsService) accept(v Visitor) {
	v.visitForNewsService(s)
}

type AdsService struct {
	url         string
	topAdsOfDay string
	ads         int
}

func (s *AdsService) accept(v Visitor) {
	v.visitForAdsService(s)
}

type Visitor interface { // интерфейс посетителя, с методами, необходимыми для "посещения" каждого нужного объекта, при этом разных типов
	visitForWeatherService(*WeatherService) int
	visitForNewsService(*NewsService) int
	visitForAdsService(*AdsService) int
}

type MessagesPerDayCalculator struct { // конкретная реализация посетителя - рассчитывающего значение сообщений в день конкретного сервиса
}

func (m *MessagesPerDayCalculator) visitForWeatherService(s *WeatherService) int {
	// какие - то вычисления
	return s.briefs
}

func (m *MessagesPerDayCalculator) visitForNewsService(s *NewsService) int {
	// какие - то вычисления
	return s.news
}

func (m *MessagesPerDayCalculator) visitForAdsService(s *AdsService) int {
	// какие - то вычисления
	return s.ads
}

func testVisitor() {
	s1 := WeatherService{"asd.com", 123}
	s2 := NewsService{"asdsad.com", nil, 10}
	s3 := AdsService{"asdsad.com", "xv", 5}

	calc := &MessagesPerDayCalculator{}

	s1.accept(calc)
	s2.accept(calc)
	s3.accept(calc)
}
