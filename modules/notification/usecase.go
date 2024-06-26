package notification

import "time"

type UseCase interface {
	StoreNotification(*Notification) (*Notification, error)
	ReadNotification(int) (*Notification, error)
	GetAllNotifications(uint) ([]Notification, error)
	DeleteAllNotifications(uint) error
	CreateCustomizeWateringReminder(*CustomizeWateringReminder) (*CustomizeWateringReminder, error)
}

type notificationUseCase struct {
	notificationRepo Repository
}

func NewUseCase(notificationRepo Repository) *notificationUseCase {
	return &notificationUseCase{
		notificationRepo: notificationRepo,
	}
}

func (u *notificationUseCase) StoreNotification(notification *Notification) (*Notification, error) {
	notification.IsRead = false
	location, _ := time.LoadLocation("Asia/Jakarta")
	notification.CreatedAt = time.Now().In(location)
	notification.UpdatedAt = time.Now().In(location)
	return u.notificationRepo.StoreNotification(notification)
}

func (u *notificationUseCase) ReadNotification(id int) (*Notification, error) {
	return u.notificationRepo.ReadNotification(id)
}

func (u *notificationUseCase) GetAllNotifications(userID uint) ([]Notification, error) {
	return u.notificationRepo.GetAllNotifications(userID)
}

func (u *notificationUseCase) DeleteAllNotifications(userID uint) error {
	return u.notificationRepo.DeleteAllNotifications(userID)
}

func (u *notificationUseCase) CreateCustomizeWateringReminder(reminder *CustomizeWateringReminder) (*CustomizeWateringReminder, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	reminder.CreatedAt = time.Now().In(location)
	reminder.UpdatedAt = time.Now().In(location)
	return u.notificationRepo.CreateCustomizeWateringReminder(reminder)
}
