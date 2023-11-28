package repository

import (
	"database/sql"
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type ProductData struct {
	ID          int             `json:"id"`
	Name        sql.NullString  `json:"name"`
	Description sql.NullString  `json:"description"`
	Tag         sql.NullString  `json:"tag"`
	Price       sql.NullFloat64 `json:"price"`
	Quantity    sql.NullInt32   `json:"quantity"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (p *ProductData) MapToRepoNew(product entity.ProductItems) {
	p.ID = product.ID()
	p.Name.Valid = true
	p.Name.String = product.Name().Val
	p.Description.Valid = true
	p.Description.String = product.Description().Val
	p.Tag.Valid = true
	p.Tag.String = product.Tag().Val
	p.Price.Valid = true
	p.Price.Float64 = product.Price().Val
	p.Quantity.Valid = true
	p.Quantity.Int32 = int32(product.Quantity().Val)
	p.CreatedAt = time.Now().UTC()
}

func (p *ProductData) MapToRepoNewList(products []entity.ProductItems) []ProductData {
	var result []ProductData
	for _, product := range products {
		productData := ProductData{}
		productData.MapToRepoNew(product)
		result = append(result, productData)
	}
	return result
}

func (p *ProductData) MapToRepo(product entity.ProductItems) {
	p.ID = product.ID()
	p.Name.Valid = true
	p.Name.String = product.Name().Val
	p.Description.Valid = true
	p.Description.String = product.Description().Val
	p.Tag.Valid = true
	p.Tag.String = product.Tag().Val
	p.Price.Valid = true
	p.Price.Float64 = product.Price().Val
	p.Quantity.Valid = true
	p.Quantity.Int32 = int32(product.Quantity().Val)
}

func (p *ProductData) MapToRepoList(products []entity.ProductItems) []ProductData {
	var result []ProductData
	for _, product := range products {
		productData := ProductData{}
		productData.MapToRepo(product)
		result = append(result, productData)
	}
	return result
}

func (p *ProductData) MapToEntity() entity.ProductItems {
	var product entity.ProductItems
	product.SetID(p.ID)
	if p.Name.Valid {
		product.SetName(commonentity.CustomString{
			Val:   p.Name.String,
			Valid: p.Name.Valid,
		})
	}
	if p.Description.Valid {
		product.SetDescription(commonentity.CustomString{
			Val:   p.Description.String,
			Valid: p.Description.Valid,
		})
	}
	if p.Tag.Valid {
		product.SetTag(commonentity.CustomString{
			Val:   p.Tag.String,
			Valid: p.Tag.Valid,
		})
	}
	if p.Price.Valid {
		product.SetPrice(commonentity.CustomFloat{
			Val:   p.Price.Float64,
			Valid: p.Price.Valid,
		})
	}
	if p.Quantity.Valid {
		product.SetQuantity(commonentity.CustomInt{
			Val:   int(p.Quantity.Int32),
			Valid: p.Quantity.Valid,
		})
	}
	product.SetCreatedAt(p.CreatedAt)
	return product
}

func (p *ProductData) MapToEntityList(products []ProductData) []entity.ProductItems {
	var result []entity.ProductItems
	for _, product := range products {
		result = append(result, product.MapToEntity())
	}
	return result
}
