package budget

import (
	"fmt"
	"time"

	"github.com/shaikhjunaidx/pennywise-backend/models"
	"gorm.io/gorm"
)

type BudgetRepositoryImpl struct {
	DB *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepositoryImpl {
	return &BudgetRepositoryImpl{DB: db}
}

func (r *BudgetRepositoryImpl) Create(budget *models.Budget) error {
	return r.DB.Create(budget).Error
}

func (r *BudgetRepositoryImpl) Update(budget *models.Budget) error {
	return r.DB.Save(budget).Error
}

func (r *BudgetRepositoryImpl) DeleteByID(id uint) error {
	return r.DB.Delete(&models.Budget{}, id).Error
}

func (r *BudgetRepositoryImpl) FindByID(id uint) (*models.Budget, error) {
	var budget models.Budget
	if err := r.DB.First(&budget, id).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepositoryImpl) FindAllByUserID(userID uint) ([]*models.Budget, error) {
	var budgets []*models.Budget
	if err := r.DB.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *BudgetRepositoryImpl) FindByUserIDAndCategoryID(userID uint, categoryID *uint, month string, year int) (*models.Budget, error) {
	var budget models.Budget

	monthFormatted := fmt.Sprintf("%02d", time.Now().Month())

	query := r.DB.Where("user_id = ? AND budget_month = ? AND budget_year = ?", userID, monthFormatted, year)

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	} else {
		query = query.Where("category_id IS NULL")
	}
	if err := query.First(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepositoryImpl) FindAllByUserIDAndMonthYear(userID uint, month string, year int) ([]*models.Budget, error) {
	var budgets []*models.Budget

	monthFormatted := fmt.Sprintf("%02d", time.Now().Month())

	err := r.DB.Where("user_id = ? AND budget_month = ? AND budget_year = ?", userID, monthFormatted, year).Find(&budgets).Error
	if err != nil {
		return nil, err
	}
	return budgets, nil
}
