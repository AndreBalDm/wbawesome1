package wbawesome1

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте

//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.
import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// функция, которая будет запускаться в качетве рабочей (читающей) горутины

func worker(id int, vals <-chan int, signal chan os.Signal) {
	defer close(signal) // заранее объявляем о закрытии канала с сигналом
	for {
		// конструкция select, позволяющая выбирать какой кейс будет запщуен (обладает приоритезацией: сначала
		// неблокирующая операция, затем блокирующая, затем дефолтный кейс
		select {
		case val := <-vals: // если горутина считала значение с основного канала
			fmt.Printf("goroutine id=%d value=%d\n", id, val)

		case <-signal: // если горутина считала значение с канала с сигналом - значит программа была вынужденно остановлена
			fmt.Printf("goroutnie id=%d stopped\n", id)
			return
		}
	}
}

func Wb4() {
	// вызов метода Seed() из пакета rand -> позволяет каждый раз при вызове метода Intn() из пакета rand каждый раз новое псевдослучайное число
	rand.Seed(time.Now().UnixNano())
	var amount int // кол-во горутин

	_, err := fmt.Scan(&amount) // считывание с консоли кол- во горутин воркеров
	if err != nil {
		return
	}
	// создание буферизированного канала для основных значений, где размер буфера равен кол-ву горутин
	// (можно и не буферизированный, но тогда горутины будут часто блокироваться, что только замедлит программу
	values := make(chan int, amount)
	defer close(values)                    // заранее объявляем о закрытии основного канала
	mainSigChan := make(chan os.Signal, 1) // буф. канал для сигнала
	defer close(mainSigChan)
	// объявление о том, какие сигналы необходимо "слушать" методом Notify, чтобы затем передать в "сигнальный" канал
	signal.Notify(mainSigChan, syscall.SIGINT, syscall.SIGTERM)
	// в цикле происходит также создание сигнального канала для каждой горутины и ее запуск
	for i := 0; i < amount; i++ {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		go worker(i, values, sigChan)
	}
	// основной цикл потока main
	for {
		select { // либо по умолчанию записывается новое значение в основной поток, либо при получении сигнала осн. поток останавливается
		case <-mainSigChan:
			fmt.Printf("main thread is stopped\n%d values are left in main channel", len(values))
			return
		default:
			values <- rand.Intn(12)
			time.Sleep(time.Second / 4)
		}
	}

}
