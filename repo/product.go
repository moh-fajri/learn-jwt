package repo

import (
	"context"

	"github.com/moh-fajri/learn-jwt/repo/mysql"
	"github.com/moh-fajri/learn-jwt/util"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	SKU         string  `json:"sku" gorm:"unique"`
	ProductName string  `json:"product_name"`
	Qty         float64 `json:"qty"`
	Price       float64 `json:"price"`
	Unit        string  `json:"unit"`
	Status      int32   `json:"status"`
}

func (p *Product) Create(ctx context.Context) error {
	result := mysql.DB.WithContext(ctx).Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) Update(ctx context.Context) error {
	result := mysql.DB.WithContext(ctx).Model(&p).Updates(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) Delete(ctx context.Context) error {
	result := mysql.DB.WithContext(ctx).Delete(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Product) ListWithPagination(ctx context.Context, pg *util.Pagination, sort string) ([]Product, *util.Pagination, error) {
	var products []Product
	queryBuilder := mysql.DB.WithContext(ctx).Where("sku like ?", "%"+p.SKU+"%").Limit(int(pg.Limit())).Offset(int(pg.Offset())).Order(sort)
	result := queryBuilder.Model(&p).Find(&products)
	if result.Error != nil {
		return []Product{}, nil, result.Error
	}
	var count int64
	resCount := mysql.DB.WithContext(ctx).Model(&p).Where("sku like ?", "%"+p.SKU+"%").Count(&count)
	if resCount.Error != nil {
		return []Product{}, nil, resCount.Error
	}
	paginate := pg.SetTotalPage(int32(count))
	return products, paginate, nil
}
