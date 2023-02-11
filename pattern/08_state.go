package pattern

import (
	"errors"
	"fmt"
	"log"
	"time"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Поведенческий паттерн

Применимость:
1) Необходимость делегировать часть состояния объекта в другие внутренние структуры (в силу того, что 	поведение объекта часто меняется
в зависимости от состояния объекта)
2) Наличие switch или большой условной конструкции, а также дублирующих кусков кода с похожими состояниями объекта

Плюсы:
1) Избавление от больших условных конструкций (или свитчей)
2) Весь код, от сносящийся к одному состоянию объекта находится в одном месте -> упрощение читабельности кода

Минусы:
1) Оверкилл для слишком маленького кол-ва состояний

Примеры:
1) Состояние Документа при его модификации:в состоянии черновика, под редакцией или уже опубликован. В каждом из этих состояний должны
быть осуществлены свои действия
2) Состояние банкомата в зависимости от наличия в нем средств или же от отдельных факторов: к примеру, если средства, есть состояние выдачи, если нет ->
средства выдать невозможно, но, к примеру, можно произвести перевод средств и т.д.

Мой пример:
Есть сервис, который отвечает за хранение файлов, их передачу.
У него есть несколько состояний:
	файл (абстрактный файл) готов к передаче,
	не готов,
	передача выполняется,
	выполняется запись файла,
    сервис не доступен

*/

// константные значения логина и пароля для того, чтобы можно остановить сервис извне

const (
	mockAdminLogin    = "admin"
	mockAdminPassword = "pass"
)

// главный интерфейс Состояния с методами записи файла, получения, проверки доступности сервиса и проверки доступности чтения / записи файла

type State interface {
	writeFile(data string, rewrite bool) error
	getFile() (string, error)
	isAvailable() bool
	isFileReady() bool
}

// структура самого сервиса со всеми возможными состояниями

type FileService struct {
	fileReady           State
	fileNotReady        State
	broadcastInProgress State
	writingFile         State
	notAvailable        State

	curState State // текущее состояние

	data     string // поле информации файла типа string
	lastData string // поле "последней" информации файла типа string для чтения при чьей-то записи
}

// конструктор сервиса, где также настраиваются зависимости

func NewFileService() *FileService {
	s := FileService{
		data: "",
	}
	fileReadyState := &FileReadyState{
		fileService: &s,
	}
	fileNotReadyState := &FileNotReadyState{
		fileService: &s,
	}
	broadcastInProgressState := &BroadcastInProgressState{
		fileService: &s,
	}
	writingFileState := &WritingFileState{
		fileService: &s,
	}
	notAvailableState := &NotAvailableState{
		fileService: &s,
	}

	s.curState = fileNotReadyState
	s.fileReady = fileReadyState
	s.fileNotReady = fileNotReadyState
	s.broadcastInProgress = broadcastInProgressState
	s.writingFile = writingFileState
	s.notAvailable = notAvailableState

	return &s
}

// метод установки состояния

func (f *FileService) setState(state State) {
	f.curState = state
}

// методы, которые вызывают метод определенного состояния (текущего)

func (f *FileService) isFileReady() bool {
	return f.curState.isFileReady()
}

func (f *FileService) isServiceAvailable() bool {
	return f.curState.isAvailable()
}

func (f *FileService) writeToFile(data string, rewrite bool) error {
	return f.curState.writeFile(data, rewrite)
}

func (f *FileService) getFile() (string, error) {
	return f.curState.getFile()
}

// функция, которая "висит" наружу и условно говоря по логину и паролю позволяет "выключить" себя, меняя свое состояние на NotAvailable

func (f *FileService) TurnTheServiceOff(login, password string) error {
	if login != mockAdminLogin {
		return errors.New("login is not correct")
	}
	if password != mockAdminPassword {
		return errors.New("password is not correct")
	}
	f.setState(f.notAvailable)
	return nil
}

// ниже представлены реализации конкретных состояний с имплементацией методов интерфейса State, у каждой реализации состояния в качестве поля есть
// указатель на объект конкретного сервиса для доступа к его полям

//---------------

type FileReadyState struct {
	fileService *FileService
}

func (f *FileReadyState) writeFile(data string, rewrite bool) error {
	if data == "" {
		return errors.New("data can not be empty")
	}
	f.fileService.lastData = data // сохраняем последнюю версию файла для того, чтобы выдать ее когда кто-то захочет прочитать при заблокированном сервисе (когда кто-то уже пишет)

	if rewrite {
		f.fileService.data = data
	} else {
		f.fileService.data += data
	}
	f.fileService.setState(f.fileService.writingFile)
	time.Sleep(time.Second * 10) // представим, что процесс записи занимает какое-то время

	f.fileService.setState(f.fileService.fileReady)
	return nil
}

func (f *FileReadyState) getFile() (string, error) {
	f.fileService.setState(f.fileService.broadcastInProgress)
	time.Sleep(time.Second * 10) // представим, что процесс считывания файла также занимает время

	f.fileService.setState(f.fileService.fileReady)
	return f.fileService.data, nil
}

func (f *FileReadyState) isAvailable() bool {
	return true
}

func (f *FileReadyState) isFileReady() bool {
	return true
}

//---------------

type FileNotReadyState struct {
	fileService *FileService
}

func (f *FileNotReadyState) writeFile(data string, rewrite bool) error {
	return errors.New("file is not ready to write to it")
}

func (f *FileNotReadyState) getFile() (string, error) {
	return "", errors.New("file is not ready to get it")
}

func (f *FileNotReadyState) isAvailable() bool {
	return true
}

func (f *FileNotReadyState) isFileReady() bool {
	return false
}

//---------------

type BroadcastInProgressState struct {
	fileService *FileService
}

func (b *BroadcastInProgressState) writeFile(data string, rewrite bool) error {
	return errors.New("can not write: someone is reading file")
}

func (b *BroadcastInProgressState) getFile() (string, error) {
	time.Sleep(time.Second * 10) // предположим, что здесь клиент встает в очередь (увеличивая счетчики семафора) на чтение
	b.fileService.setState(b.fileService.fileReady)
	return b.fileService.data, nil

}

func (b *BroadcastInProgressState) isAvailable() bool {
	return true
}

func (b *BroadcastInProgressState) isFileReady() bool {
	return false
}

//---------------

type WritingFileState struct {
	fileService *FileService
}

func (w *WritingFileState) writeFile(data string, rewrite bool) error {
	return errors.New("can not write: someone is writing")
}

func (w *WritingFileState) getFile() (string, error) {
	// предположим, что здесь клиент встает в очередь (увеличивая счетчик семафора) на запись, при этом он не сможет его увеличить
	// если есть уже в очереди на чтение (по принципу нового устройства Планировщика горутин)
	time.Sleep(time.Second * 10)
	w.fileService.setState(w.fileService.fileReady)

	return w.fileService.lastData, nil
}

func (w *WritingFileState) isAvailable() bool {
	return true
}

func (w *WritingFileState) isFileReady() bool {
	return false
}

//---------------

type NotAvailableState struct {
	fileService *FileService
}

func (n *NotAvailableState) writeFile(data string, rewrite bool) error {
	return errors.New("service is not available")
}

func (n *NotAvailableState) getFile() (string, error) {
	return "", errors.New("service is not available")
}

func (n *NotAvailableState) isAvailable() bool {
	return false
}

func (n *NotAvailableState) isFileReady() bool {
	return false
}

// пример использования

func stateTest() {
	service := NewFileService()

	err := service.writeToFile("Example text....", false)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	file, err := service.getFile() //если чтение было в течение 10 секунд после начала записи, то получим данные из поля lastData самого сервиса
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(file)

	err = service.TurnTheServiceOff(mockAdminLogin, mockAdminPassword)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = service.writeToFile("another text....", true) // получим ошибку, так как сервис уже недоступен
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
