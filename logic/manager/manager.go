package manager

import (
	"fmt"
	"log"
	"strings"
	"test_task/logic"
	"test_task/models"
)

type Manager struct {
	logic *logic.Logic
}

func NewManager(l *logic.Logic) *Manager {
	return &Manager{logic: l}
}

func (m *Manager) Menu() {
	for {
		fmt.Println("\nМеню менеджера:")
		fmt.Println("1 - Добавить услугу")
		fmt.Println("2 - Добавить одежду")
		fmt.Println("3 - Посмотреть услуги")
		fmt.Println("4 - Посмотреть одежду")
		fmt.Println("5 - Назад")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			m.AddServiceKind()
		case 2:
			m.AddServiceClothes()
		case 3:
			m.logic.ShowAllServiceKinds()
		case 4:
			m.logic.ShowAllServiceClothes()
		case 5:
			return
		default:
			fmt.Println("Некорректный ввод.")
		}
	}
}

func (m *Manager) AddServiceKind() {
	fmt.Print("\nВведите название услуги(%20 для пробела): ")
	var name string
	fmt.Scan(&name)
	nameResult := strings.ReplaceAll(name, "%20", " ")

	fmt.Print("Введите единицу измерения (item/kg): ")
	var unit string
	fmt.Scan(&unit)

	serviceKind := models.ServiceKind{
		Name: nameResult,
		Unit: unit,
	}

	if err := m.logic.AddServiceKind(&serviceKind); err != nil {
		log.Println("Ошибка при добавлении услуги:", err)
		return
	}

	fmt.Println("✅ Услуга успешно добавлена!")
}

func (m *Manager) AddServiceClothes() {
	serviceKinds, err := m.logic.GetAllServiceKinds()
	if err != nil {
		log.Println("Ошибка при загрузке услуг:", err)
		return
	}

	if len(serviceKinds) == 0 {
		fmt.Println("Список услуг пуст. Сначала добавьте услугу.")
		return
	}

	fmt.Println("\nВыберите услугу:")
	for i, sk := range serviceKinds {
		fmt.Printf("%d - %s (%s)\n", i+1, sk.Name, sk.Unit)
	}

	fmt.Print("Введите номер услуги: ")
	var serviceChoice int
	fmt.Scan(&serviceChoice)

	if serviceChoice < 1 || serviceChoice > len(serviceKinds) {
		fmt.Println("Некорректный выбор.")
		return
	}

	selectedService := serviceKinds[serviceChoice-1]

	fmt.Print("Введите название одежды(%20 для пробела): ")
	var name string
	fmt.Scan(&name)
	nameResult := strings.ReplaceAll(name, "%20", " ")

	fmt.Print("Введите цену: ")
	var price float64
	fmt.Scan(&price)
	if price < 0 {
		fmt.Println("Цена не может быть отрицательной.")
		return
	}

	serviceClothes := models.ServiceClothes{
		ServiceKindId: selectedService.Id,
		Name:          nameResult,
		Price:         price,
	}

	if err := m.logic.AddServiceClothes(&serviceClothes); err != nil {
		log.Println("Ошибка при добавлении одежды:", err)
		return
	}

	fmt.Println("✅ Одежда успешно добавлена!")
}
