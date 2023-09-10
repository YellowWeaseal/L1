package main

// 1 способ с помощью канала для остановок
//Можно использовать канал для отправки сигнала о завершении выполнения горутины.

/*func Worker(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println("Горутина завершает работу.")
			return
		default:
			// Ваша логика выполнения работы
		}
	}
}

func main() {
	stopCh := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go Worker(stopCh, &wg)

	// В какой-то момент вы можете отправить сигнал для завершения горутины
	// close(stopCh)

	wg.Wait() // Ожидание завершения работы горутины
}*/

//2 способ с помощью context
//Можно передатьсигнал о завершении горутины через context
//Вызов cancel() завершает выполнение горутины через context

/*func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина завершает работу.")
			return
		default:
			// Ваша логика выполнения работы
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	// В какой-то момент вы можете отправить сигнал для завершения горутины
	// cancel()

	wg.Wait() // Ожидание завершения работы горутины
}*/

//3 способ с помощью sync.Once
//sync.Once гарантирует, что функция будет вызвана только один раз,
//даже если она будет вызвана из разных горутин.

/*func worker(done *sync.Once, wg *sync.WaitGroup) {
	defer wg.Done()
	done.Do(func() {
		fmt.Println("Горутина завершает работу.")
	})
	// Ваша логика выполнения работы
}

func main() {
	var (
		done sync.Once
		wg   sync.WaitGroup
	)

	wg.Add(1)
	go worker(&done, &wg)

	// В какой-то момент вы можете вызвать done.Do() для завершения горутины
	// done.Do(func() {
	//     fmt.Println("Основная горутина завершает работу.")
	// })

	wg.Wait() // Ожидание завершения работы горутины
}*/

//4 способ c помощью runtime.Goexit()
//Функция runtime.Goexit() может использоваться для немедленного завершения выполнения текущей горутины.
//Она завершает только текущую горутину, не влияя на другие горутины

/*func worker() {
	// Выполнение работы
	fmt.Println("Горутина: Выполнение работы...")

	// Завершение текущей горутины
	runtime.Goexit()
}

func main() {
	go worker()

	// Ожидаем завершения всех горутин
	runtime.Goexit()
	fmt.Println("Основная горутина завершает работу.")
}*/

//5 способ с помощью sync.Mutex
//Мьютекс можно использовать для контроля выполнения горутин
//Здесь мьютекс используется для предотвращения выполнения кода в горутине более одного раза.

/*var mu sync.Mutex
var done bool

func worker() {
	mu.Lock()
	if done {
		mu.Unlock()
		return
	}
	// Ваша логика выполнения работы
	fmt.Println("Горутина: Выполнение работы...")
	done = true
	mu.Unlock()
}

func main() {
	go worker()
	go worker()

	// Ожидание завершения всех горутин
	time.Sleep(2 * time.Second)
	fmt.Println("Основная горутина завершает работу.")
}*/

//6 способ с помощью os.Exit
//os.Exit немедленно завершает работу прогрммы включая горутины

/*

func worker() {
    // Ваша логика выполнения работы
    fmt.Println("Горутина: Выполнение работы...")
}

func main() {
    go worker()
    go worker()

    // Немедленное завершение всей программы
    os.Exit(0)
}*/

// 7 способ с помощью context.WithTimeout
//Горутина завершит выполнение, если она не завершится в течение заданного

/*func worker(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Горутина: Выполнение работы...")
	case <-ctx.Done():
		fmt.Println("Горутина завершает работу.")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go worker(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Основная горутина завершает работу.")
	}
}*/

//Пожалуй это все основные методы которые я знаю
