package storage

import (
	"gorm.io/gorm"
	"log"
	"test_task/models"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.ServiceKind{},
		&models.ServiceClothes{},
		&models.Order{},
		&models.OrderItem{},
	)

	var count int64
	db.Model(&models.ServiceKind{}).Count(&count)
	if count == 0 {
		serviceKinds := []models.ServiceKind{
			{Id: "11111111-1111-1111-1111-111111111111", Name: "Химчистка", Unit: "item"},
			{Id: "22222222-2222-2222-2222-222222222222", Name: "Ручная стирка", Unit: "item"},
			{Id: "33333333-3333-3333-3333-333333333333", Name: "Общие услуги по стирке", Unit: "kg"},
			{Id: "44444444-4444-4444-4444-444444444444", Name: "Гладильные услуги", Unit: "item"},
			{Id: "55555555-5555-5555-5555-555555555555", Name: "Ремонт одежды", Unit: "item"},
			{Id: "66666666-6666-6666-6666-666666666666", Name: "Удаление пятен", Unit: "item"},
		}
		db.Create(&serviceKinds)
		log.Println("✅ Добавлены услуги в ServiceKind")
	}

	db.Model(&models.ServiceClothes{}).Count(&count)
	if count == 0 {
		serviceClothes := []models.ServiceClothes{
			{Id: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", ServiceKindId: "11111111-1111-1111-1111-111111111111", Name: "Пальто", Price: 20},
			{Id: "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb", ServiceKindId: "11111111-1111-1111-1111-111111111111", Name: "Брюки", Price: 10},
			{Id: "cccccccc-cccc-cccc-cccc-cccccccccccc", ServiceKindId: "11111111-1111-1111-1111-111111111111", Name: "Костюм", Price: 15},
			{Id: "dddddddd-dddd-dddd-dddd-dddddddddddd", ServiceKindId: "11111111-1111-1111-1111-111111111111", Name: "Сюртук", Price: 15},
			{Id: "eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee", ServiceKindId: "11111111-1111-1111-1111-111111111111", Name: "Вечернее Платье", Price: 15},
			{Id: "ffffffff-ffff-ffff-ffff-ffffffffffff", ServiceKindId: "11111111-1111-1111-1111-111111111111", Name: "Свадебное Платье", Price: 20},

			{Id: "11112222-3333-4444-5555-666677778888", ServiceKindId: "22222222-2222-2222-2222-222222222222", Name: "Обычная вещь", Price: 5},

			{Id: "99998888-7777-6666-5555-444433332222", ServiceKindId: "33333333-3333-3333-3333-333333333333", Name: "Белые", Price: 10},
			{Id: "88887777-6666-5555-4444-333322221111", ServiceKindId: "33333333-3333-3333-3333-333333333333", Name: "Цветные", Price: 8},
			{Id: "77776666-5555-4444-3333-222211110000", ServiceKindId: "33333333-3333-3333-3333-333333333333", Name: "Шерсть", Price: 12},
			{Id: "66665555-4444-3333-2222-11110000ffff", ServiceKindId: "33333333-3333-3333-3333-333333333333", Name: "Шелк", Price: 15},
			{Id: "55554444-3333-2222-1111-0000eeeeeeee", ServiceKindId: "33333333-3333-3333-3333-333333333333", Name: "Мягкие Игрушки", Price: 9},
			{Id: "44443333-2222-1111-0000-ffffffffffff", ServiceKindId: "33333333-3333-3333-3333-333333333333", Name: "Постельное Бельё", Price: 10},

			{Id: "aaaa0000-bbbb-cccc-dddd-eeeeffff0000", ServiceKindId: "44444444-4444-4444-4444-444444444444", Name: "Рубашки", Price: 5},
			{Id: "bbbb1111-cccc-dddd-eeee-ffff00001111", ServiceKindId: "44444444-4444-4444-4444-444444444444", Name: "Брюки", Price: 3},
			{Id: "cccc2222-dddd-eeee-ffff-000011112222", ServiceKindId: "44444444-4444-4444-4444-444444444444", Name: "Юбки", Price: 2},
			{Id: "dddd3333-eeee-ffff-0000-111122223333", ServiceKindId: "44444444-4444-4444-4444-444444444444", Name: "Платья", Price: 6},
			{Id: "eeee4444-ffff-0000-1111-222233334444", ServiceKindId: "44444444-4444-4444-4444-444444444444", Name: "Костюмы", Price: 8},

			{Id: "ffff5555-0000-1111-2222-333344445555", ServiceKindId: "55555555-5555-5555-5555-555555555555", Name: "Исправление шва / штопка", Price: 3.50},

			{Id: "00006666-1111-2222-3333-444455556666", ServiceKindId: "66666666-6666-6666-6666-666666666666", Name: "Пятна от масла", Price: 5},
			{Id: "11117777-2222-3333-4444-555566667777", ServiceKindId: "66666666-6666-6666-6666-666666666666", Name: "Пятна от крови", Price: 3},
			{Id: "22228888-3333-4444-5555-666677778888", ServiceKindId: "66666666-6666-6666-6666-666666666666", Name: "Общая грязь", Price: 2},
		}
		db.Create(&serviceClothes)
		log.Println("✅ Добавлена одежда в ServiceClothes")
	}

	return err
}
