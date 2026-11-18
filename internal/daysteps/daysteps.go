// Пакет daysteps разбирает строки с данными о прогулках
// и формирует по ним отчёт.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные об одной прогулке.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "678,0h50m" и записывает данные в структуру.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат строки %q: ожидалось 2 значения, получено %d", datastring, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля, получено %d", steps)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность %q: %w", parts[1], err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше нуля, получено %v", duration)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo возвращает отчёт о прогулке.
// При ошибке возвращает пустую строку и ошибку.
func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
