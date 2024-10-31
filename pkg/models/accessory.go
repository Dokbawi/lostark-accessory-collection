package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type MarketTrend string

const (
	TrendUp     MarketTrend = "up"
	TrendDown   MarketTrend = "down"
	TrendStable MarketTrend = "stable"
)

type Accessory struct {
	ID           int64
	Name         string
	CurrentPrice float64
	MarketTrend  MarketTrend
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AccessoryRepository struct {
	db *sql.DB
}

func NewAccessoryRepository(db *sql.DB) *AccessoryRepository {
	return &AccessoryRepository{db: db}
}

// UpdatePrices updates prices based on market trends
func (r *AccessoryRepository) UpdatePrices(ctx context.Context) (int64, error) {
	result, err := r.db.ExecContext(ctx, `
        UPDATE accessories 
        SET 
            current_price = CASE
                WHEN market_trend = 'up' THEN current_price * 1.05
                WHEN market_trend = 'down' THEN current_price * 0.95
                ELSE current_price
            END,
            updated_at = NOW()
        WHERE active = true
    `)
	if err != nil {
		return 0, fmt.Errorf("failed to update prices: %v", err)
	}

	return result.RowsAffected()
}

// GetAccessory retrieves a single accessory by ID
func (r *AccessoryRepository) GetAccessory(ctx context.Context, id int64) (*Accessory, error) {
	var acc Accessory
	err := r.db.QueryRowContext(ctx, `
        SELECT id, name, current_price, market_trend, active, created_at, updated_at
        FROM accessories
        WHERE id = ?
    `, id).Scan(
		&acc.ID,
		&acc.Name,
		&acc.CurrentPrice,
		&acc.MarketTrend,
		&acc.Active,
		&acc.CreatedAt,
		&acc.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get accessory: %v", err)
	}
	return &acc, nil
}

// ListAccessories retrieves all active accessories
func (r *AccessoryRepository) ListAccessories(ctx context.Context) ([]*Accessory, error) {
	rows, err := r.db.QueryContext(ctx, `
        SELECT id, name, current_price, market_trend, active, created_at, updated_at
        FROM accessories
        WHERE active = true
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to list accessories: %v", err)
	}
	defer rows.Close()

	var accessories []*Accessory
	for rows.Next() {
		var acc Accessory
		if err := rows.Scan(
			&acc.ID,
			&acc.Name,
			&acc.CurrentPrice,
			&acc.MarketTrend,
			&acc.Active,
			&acc.CreatedAt,
			&acc.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan accessory: %v", err)
		}
		accessories = append(accessories, &acc)
	}
	return accessories, nil
}
