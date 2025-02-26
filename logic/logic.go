package logic

import (
	"fmt"
	"log"
	"test_task/models"
	"test_task/service"
)

type Logic struct {
	service service.IService
}

func NewLogic(s service.IService) *Logic {
	return &Logic{service: s}
}

func (l *Logic) ShowAllOrders() {
	orders, err := l.service.Order().GetAll()
	if err != nil {
		log.Println("Ошибка при получении заказов:", err)
		return
	}

	if len(orders) == 0 {
		fmt.Println("\nЗаказов нет.")
		return
	}

	fmt.Println("\nВсе заказы:")
	for _, order := range orders {
		fmt.Printf("\nЗаказ %s\n", order.Id)
		fmt.Printf("Сумма: %.2f сомони, Экспресс: %v\n", order.TotalPrice, order.Express)

		for _, item := range order.OrderItems {
			fmt.Printf(" - Услуга: %s, Одежда: %s (%.2f сомони)\n",
				item.ServiceKindName, item.ServiceClothesName, item.Price)
		}
	}
}

func (l *Logic) GetAllServiceKinds() ([]models.ServiceKind, error) {
	return l.service.ServiceKind().GetAll()
}

func (l *Logic) ShowAllServiceKinds() {
	serviceKinds, err := l.service.ServiceKind().GetAll()
	if err != nil {
		log.Println("Ошибка при получении услуг:", err)
		return
	}

	fmt.Println("\nСписок услуг:")
	for _, sk := range serviceKinds {
		fmt.Printf("- %s (%s)\n", sk.Name, sk.Unit)
	}
}

func (l *Logic) GetAllServiceClothes(filter models.GetAllServiceClothesReq) ([]models.ServiceClothes, error) {
	return l.service.ServiceClothes().GetAll(filter)
}

func (l *Logic) ShowAllServiceClothes() {
	clothes, err := l.service.ServiceClothes().GetAll(models.GetAllServiceClothesReq{})
	if err != nil {
		log.Println("Ошибка при получении одежды:", err)
		return
	}

	fmt.Println("\nСписок одежды:")
	for _, item := range clothes {
		fmt.Printf("- %s (Цена: %.2f)\n", item.Name, item.Price)
	}
}

func (l *Logic) CreateOrder(order *models.Order) error {
	_, err := l.service.Order().Create(order)
	return err
}

func (l *Logic) AddServiceKind(serviceKind *models.ServiceKind) error {
	_, err := l.service.ServiceKind().Create(serviceKind)
	return err
}

func (l *Logic) AddServiceClothes(serviceClothes *models.ServiceClothes) error {
	_, err := l.service.ServiceClothes().Create(serviceClothes)
	return err
}
