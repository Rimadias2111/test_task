package cashier

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"test_task/logic"
	"test_task/models"
)

type Cashier struct {
	logic *logic.Logic
}

func NewCashier(l *logic.Logic) *Cashier {
	return &Cashier{logic: l}
}

func (c *Cashier) Menu() {
	for {
		fmt.Println("\nМеню кассира:")
		fmt.Println("1 - Оформить заказ")
		fmt.Println("2 - Показать все заказы")
		fmt.Println("3 - Назад")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			c.CreateOrder()
		case 2:
			c.logic.ShowAllOrders()
		case 3:
			return
		default:
			fmt.Println("Некорректный ввод.")
		}
	}
}

func (c *Cashier) CreateOrder() {
	order := models.Order{
		Id:         uuid.New().String(),
		TotalPrice: 0,
		OrderItems: []models.OrderItem{},
	}

	fmt.Println("\nОформление нового заказа...")

	totalWeight := 0.0
	totalBasePrice := 0.0

	for {
		serviceKinds, err := c.logic.GetAllServiceKinds()
		if err != nil {
			log.Println("Ошибка загрузки услуг:", err)
			return
		}

		fmt.Println("\nВыберите услугу:")
		for i, s := range serviceKinds {
			fmt.Printf("%d - %s (%s)\n", i+1, s.Name, s.Unit)
		}

		fmt.Print("Введите номер услуги (0 - завершить выбор): ")
		var serviceChoice int
		fmt.Scan(&serviceChoice)

		if serviceChoice == 0 {
			break
		}

		if serviceChoice < 1 || serviceChoice > len(serviceKinds) {
			fmt.Println("Некорректный ввод.")
			continue
		}

		selectedService := serviceKinds[serviceChoice-1]

		clothesFilter := models.GetAllServiceClothesReq{ServiceKindId: selectedService.Id}
		serviceClothes, err := c.logic.GetAllServiceClothes(clothesFilter)
		if err != nil {
			log.Println("Ошибка загрузки одежды:", err)
			continue
		}

		fmt.Println("\nВыберите одежду для услуги:")
		for i, item := range serviceClothes {
			fmt.Printf("%d - %s (Цена: %.2f)\n", i+1, item.Name, item.Price)
		}

		fmt.Print("Введите номер одежды: ")
		var clothesChoice int
		fmt.Scan(&clothesChoice)

		if clothesChoice < 1 || clothesChoice > len(serviceClothes) {
			fmt.Println("Некорректный ввод.")
			continue
		}

		selectedClothes := serviceClothes[clothesChoice-1]

		var quantity int
		var weight float64
		var price float64

		if selectedService.Unit == "item" {
			fmt.Print("Введите количество: ")
			fmt.Scan(&quantity)
			if quantity < 1 {
				fmt.Println("Количество должно быть больше 0.")
				continue
			}
			price = selectedClothes.Price * float64(quantity)
		} else {
			fmt.Print("Введите вес (кг): ")
			fmt.Scan(&weight)
			if weight <= 0 {
				fmt.Println("Вес должен быть больше 0.")
				continue
			}
			price = selectedClothes.Price * weight
			totalWeight += weight
		}

		fmt.Print("Эта вещь детская? (1 - Да, 0 - Нет): ")
		var isChildItem int
		fmt.Scan(&isChildItem)

		if isChildItem == 1 {
			price *= 0.5
			fmt.Println("🟢 Применена скидка 50% на детскую вещь.")
		}

		totalBasePrice += price
		order.TotalPrice += price

		orderItem := models.OrderItem{
			Id:               uuid.New().String(),
			OrderId:          order.Id,
			ServiceKindId:    selectedService.Id,
			ServiceClothesId: selectedClothes.Id,
			Quantity:         &quantity,
			Weight:           &weight,
			Price:            price,
		}
		order.OrderItems = append(order.OrderItems, orderItem)
	}

	fmt.Print("\nСрочный заказ? (1 - Да, 0 - Нет): ")
	var isExpress int
	discount := 0.0
	fmt.Scan(&isExpress)

	if isExpress == 1 {
		order.Express = true
		order.TotalPrice *= 1.5
	} else {
		fmt.Print("Клиент может оставить вещи на 5+ дней? (1 - Да, 0 - Нет): ")
		var isDelayed int
		fmt.Scan(&isDelayed)

		if isDelayed == 1 {
			discount += 0.3
			fmt.Println("🟢 Применена скидка 30% за долгосрочное хранение.")
		}
	}

	if totalWeight > 10 {
		discount += 0.2
		fmt.Println("🟢 Применена скидка 20% за большой вес (>10 кг).")
	}

	if discount > 0 {
		order.TotalPrice *= (1 - discount)
		fmt.Printf("💰 Итоговая скидка на весь заказ: %.0f%%\n", discount*100)
	}

	if err := c.logic.CreateOrder(&order); err != nil {
		log.Println("Ошибка при сохранении заказа:", err)
		return
	}

	fmt.Printf("\n✅ Заказ оформлен. Итоговая сумма после всех скидок: %.2f сомони\n", order.TotalPrice)
}
