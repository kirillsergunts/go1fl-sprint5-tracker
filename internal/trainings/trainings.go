// Пакет trainings разбирает строки с данными о тренировках
// и формирует по ним отчёт.
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Виды тренировок, которые умеет обрабатывать пакет.
const (
	walking = "Ходьба"
	running = "Бег"
)

// ErrUnknownTrainingType возвращается, если тип тренировки не поддерживается.
var ErrUnknownTrainingType = errors.New("неизвестный тип тренировки")

// Training хранит данные об одной тренировке.
// Структура personaldata.Personal встроена (у поля нет имени),
// поэтому её поля и метод Print доступны напрямую: t.Weight, t.Print().
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "3456,Ходьба,3h00m" и записывает
// данные в структуру. Получатель — указатель, иначе метод
// изменил бы копию структуры, а не оригинал.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат строки %q: ожидалось 3 значения, получено %d", datastring, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля, получено %d", steps)
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность %q: %w", parts[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть больше нуля, получено %v", duration)
	}

	// Записываем в структуру только после того, как все проверки пройдены:
	// так при ошибке в структуре не останутся данные наполовину.
	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

// ActionInfo возвращает отчёт о тренировке.
// При ошибке возвращает пустую строку и ошибку.
func (t Training) ActionInfo() (string, error) {
	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case walking:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case running:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", ErrUnknownTrainingType
	}
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	), nil
}
