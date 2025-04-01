package main

import (
	"fmt"
	"math"
	"time"
)

type Coffee struct {
	Type        string
	Size        string
	Temperature float64
	Milk        bool
	Sugar       int
	Extras      []string
	Price       float64
	OrderTime   time.Time
}

func main() {
	// Создаем латте
	latte := Coffee{
		Type:        "Латте",
		Size:        "Большой",
		Temperature: 85.5,
		Milk:        true,
		Sugar:       2,
		Extras:      []string{"Корица", "Ванильный сироп"},
		Price:       4.99,
		OrderTime:   time.Now(),
	}

	// Простой вывод с различными спецификаторами
	fmt.Printf("Заказ: %s\n", latte.Type)
	fmt.Printf("Размер: %s\n", latte.Size)
	fmt.Printf("Температура: %.1f°C\n", latte.Temperature)
	fmt.Printf("С молоком: %t\n", latte.Milk)
	fmt.Printf("Сахар: %d ложки\n", latte.Sugar)
	fmt.Printf("Добавки: %v\n", latte.Extras)
	fmt.Printf("Цена: $%.2f\n", latte.Price)

	// Использование ширины поля и выравнивания
	fmt.Println("\n--- Чек ---")
	fmt.Printf("%-15s %s\n", "Продукт:", latte.Type)
	fmt.Printf("%-15s %s\n", "Размер:", latte.Size)
	fmt.Printf("%-15s %.1f°C\n", "Температура:", latte.Temperature)
	fmt.Printf("%-15s %t\n", "С молоком:", latte.Milk)
	fmt.Printf("%-15s %d\n", "Сахар:", latte.Sugar)

	// Использование спецификатора для среза
	fmt.Printf("%-15s ", "Добавки:")
	if len(latte.Extras) == 0 {
		fmt.Println("нет")
	} else {
		for i, extra := range latte.Extras {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(extra)
		}
		fmt.Println()
	}

	// Форматирование цены
	fmt.Printf("%-15s $%6.2f\n", "Цена:", latte.Price)
	fmt.Printf("%-15s $%6.2f\n", "Налог (7%):", latte.Price*0.07)
	fmt.Printf("%-15s $%6.2f\n", "Итого:", latte.Price*1.07)

	// Использование %v и %+v для всей структуры
	fmt.Println("\nОтладочная информация:")
	fmt.Printf("Заказ: %v\n", latte)
	fmt.Printf("Подробно: %+v\n", latte)

	// Более сложное форматирование
	fmt.Println("\nАнализ заказа:")
	caloriesPerSugar := 16
	calories := 180 + (latte.Sugar * caloriesPerSugar)
	if latte.Milk {
		calories += 120
	}
	for _, extra := range latte.Extras {
		if extra == "Ванильный сироп" {
			calories += 35
		} else if extra == "Корица" {
			calories += 5
		}
	}

	// Шестнадцатеричное, восьмеричное и двоичное представление
	fmt.Printf(
		"Калорийность: %d ккал (шестнадцатеричное: 0x%x, восьмеричное: 0%o, двоичное: %b)\n",
		calories,
		calories,
		calories,
		calories,
	)

	// Использование дополнительных симоволов для заполнения
	fmt.Printf("Уровень калорийности: [%8s]\n", getCalorieLevel(calories))

	// Форматирование времени
	fmt.Printf("Заказ оформлен: %s\n", latte.OrderTime.Format("15:04:05"))
	readyTime := latte.OrderTime.Add(3 * time.Minute)
	fmt.Printf("Будет готов в: %s (осталось %.0f секунд)\n",
		readyTime.Format("15:04:05"),
		math.Ceil(time.Until(readyTime).Seconds()))

	// Использование %q для строк с кавычками
	fmt.Printf("Тип напитка (в кавычках): %q\n", latte.Type)

	// Печать указателя
	coffeePtr := &latte
	fmt.Printf("Адрес заказа в памяти: %p\n", coffeePtr)
}

func getCalorieLevel(calories int) string {
	if calories < 100 {
		return "Низкая"
	} else if calories < 200 {
		return "Средняя"
	} else if calories < 300 {
		return "Высокая"
	} else {
		return "Очень высокая"
	}
}
