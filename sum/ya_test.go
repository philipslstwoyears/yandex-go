package main

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	// набор тестов
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		values []byte
		// желаемый результат
		want int
		// ожидаемая ошибка
		isError error
	}{
		// тестовые данные № 1
		{
			name:    "positive values",
			values:  []byte("AaPa"),
			want:    4,
			isError: nil,
		},
		{
			name:    "pa",
			values:  []byte(""),
			want:    0,
			isError: nil,
		},
		{
			name:    "positive values",
			values:  []byte{0xff, 0xfe, 0xfd},
			want:    0,
			isError: ErrInvalidUTF8,
		},
	}
	// перебор всех тестов
	for _, tc := range cases {
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got, err := GetUTFLength(tc.values)
			// проверим полученное значение
			if got != tc.want {
				t.Errorf("Sum(%v) = %v; want %v", tc.values, got, tc.want)
			}
			if !errors.Is(err, tc.isError) {
				t.Errorf("Sum(%v) = %v; want %v", tc.values, err, tc.isError)
			}
		})
	}
}
