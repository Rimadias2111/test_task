package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"test_task/models"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) Order {
	return &OrderRepo{
		db: db,
	}
}

func (o OrderRepo) Create(order *models.Order) (string, error) {
	id := uuid.New()
	order.Id = id.String()

	err := o.db.Create(order).Error
	if err != nil {
		return "", err
	}

	return order.Id, nil
}

func (o OrderRepo) Update(order *models.Order) error {
	err := o.db.Save(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (o OrderRepo) Delete(req models.RequestId) error {
	err := o.db.Delete(&models.Order{}, req.Id).Error
	if err != nil {
		return err
	}

	return nil
}

func (o OrderRepo) Get(req models.RequestId) (*models.Order, error) {
	var order models.Order
	err := o.db.Preload("OrderItems").First(&order, req.Id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o OrderRepo) GetAll() ([]models.Order, error) {
	var orders []models.Order

	query := `
		SELECT 
			o.id AS order_id, o.total_price, o.express,
			oi.id AS order_item_id, oi.service_kind_id, oi.service_clothes_id, oi.quantity, oi.weight, oi.price,
			sk.name AS service_kind_name, sc.name AS service_clothes_name
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		LEFT JOIN service_kinds sk ON oi.service_kind_id = sk.id
		LEFT JOIN service_clothes sc ON oi.service_clothes_id = sc.id
	`

	type OrderWithItems struct {
		OrderId            string
		TotalPrice         float64
		Express            bool
		OrderItemId        string
		ServiceKindId      string
		ServiceClothesId   string
		Quantity           *int
		Weight             *float64
		Price              float64
		ServiceKindName    string
		ServiceClothesName string
	}

	var results []OrderWithItems
	err := o.db.Raw(query).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	orderMap := make(map[string]*models.Order)
	for _, row := range results {
		if _, exists := orderMap[row.OrderId]; !exists {
			orderMap[row.OrderId] = &models.Order{
				Id:         row.OrderId,
				TotalPrice: row.TotalPrice,
				Express:    row.Express,
				OrderItems: []models.OrderItem{},
			}
		}

		orderItem := models.OrderItem{
			Id:                 row.OrderItemId,
			OrderId:            row.OrderId,
			ServiceKindId:      row.ServiceKindId,
			ServiceClothesId:   row.ServiceClothesId,
			Quantity:           row.Quantity,
			Weight:             row.Weight,
			Price:              row.Price,
			ServiceKindName:    row.ServiceKindName,
			ServiceClothesName: row.ServiceClothesName,
		}

		orderMap[row.OrderId].OrderItems = append(orderMap[row.OrderId].OrderItems, orderItem)
	}

	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders, nil
}
