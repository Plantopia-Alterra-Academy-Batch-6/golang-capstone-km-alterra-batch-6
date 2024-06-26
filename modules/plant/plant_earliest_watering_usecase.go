// usecase.go
package plant

import (
    "time"
    "fmt"
)

type PlantEarliestWateringService interface {
    FindEarliestWateringTime() ([]PlantReminderResponse, error)
}

type plantEarliestWateringService struct {
    repository PlantEarliestWateringRepository
}

func NewPlantEarliestWateringService(repository PlantEarliestWateringRepository) PlantEarliestWateringService {
    return &plantEarliestWateringService{repository}
}

func ConvertToTime(timeStr string) (time.Time, error) {
    layout := "15:04"
    return time.Parse(layout, timeStr)
}

func (s *plantEarliestWateringService) FindEarliestWateringTime() ([]PlantReminderResponse, error) {
   schedules, err := s.repository.GetEarliestWatering()
	if err != nil {
		return nil, err
	}

    if len(schedules) == 0 {
        return []PlantReminderResponse{}, fmt.Errorf("no schedules provided")
    }

    earliestSchedule := schedules[0]
    earliestTime, err := ConvertToTime(earliestSchedule.WateringTime)
    if err != nil {
        return []PlantReminderResponse{}, err
    }

    for _, schedule := range schedules[1:] {
        currentTime, err := ConvertToTime(schedule.WateringTime)
        if err != nil {
            return []PlantReminderResponse{}, err
        }

        if currentTime.Before(earliestTime) {
            earliestTime = currentTime
            earliestSchedule = schedule
        }
    }
    response := NewPlantReminderResponse(earliestSchedule)
    return []PlantReminderResponse{response}, nil
}
