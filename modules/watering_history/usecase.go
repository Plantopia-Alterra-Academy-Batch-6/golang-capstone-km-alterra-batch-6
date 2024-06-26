package wateringhistory

import "time"

type WateringHistoryUseCase interface {
	StoreWateringHistory(*WateringHistory) (*WateringHistory, error)
	GetAllWateringHistories(uint) ([]WateringHistory, error)
	GetLateWateringHistories(uint) (Notification, error)
}

type wateringHistoryUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *wateringHistoryUseCase {
	return &wateringHistoryUseCase{
		repo: repo,
	}
}

func (uc *wateringHistoryUseCase) StoreWateringHistory(wh *WateringHistory) (*WateringHistory, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	wh.CreatedAt = time.Now().In(location)
	wh.UpdatedAt = time.Now().In(location)
	wh, err := uc.repo.StoreWateringHistory(wh)
	if err != nil {
		return nil, err
	}
	return wh, nil
}

func (uc *wateringHistoryUseCase) GetAllWateringHistories(userID uint) ([]WateringHistory, error) {
	wh, err := uc.repo.GetAllWateringHistories(userID)
	if err != nil {
		return nil, err
	}
	return wh, nil
}

func (uc *wateringHistoryUseCase) GetLateWateringHistories(userID uint) (Notification, error) {
	notification, err := uc.repo.GetLateWateringHistories(userID)
	if err != nil {
		return notification, err
	}
	return notification, nil
}
