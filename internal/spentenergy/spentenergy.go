package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть выше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	callories := (weight * averageSpeed * durationInMinutes) / minInH
	editedCallories := callories * walkingCaloriesCoefficient
	return editedCallories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть выше 0")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть больше 0")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	callories := (weight * averageSpeed * durationInMinutes) / minInH
	return callories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	return distance / hours
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	return distanceKm
}
