package main

import (
	"fmt"
	"strconv"
	"strings"
)

// структура
type IPv4 struct {
	parts []int // Части IP-адреса
}

// метод для проверки корректности IP-адреса
func (ip *IPv4) IsValid() bool {
	// проверяем, что IP-адрес состоит из 4 частей
	if len(ip.parts) != 4 {
		return false
	}

	// проверяем, что каждая часть находится в диапазоне 0-255
	for _, part := range ip.parts {
		if part < 0 || part > 255 {
			return false
		}
	}
	return true
}

// парсинг IP-адреса из строки
func (ip *IPv4) Parse(ipStr string) error {
	// разбиваем строку на части по точке
	partsStr := strings.Split(ipStr, ".")
	if len(partsStr) != 4 {
		return fmt.Errorf("IPадрес должен состоять из 4 частей")
	}

	// парсим каждую часть
	ip.parts = make([]int, 4)
	for i, partStr := range partsStr {
		value, base, err := parsePart(partStr)
		if err != nil {
			return fmt.Errorf("ошибка в части %d: %v", i+1, err)
		}

		// проверяем диапазон для восьмеричных чисел
		if base == 8 && value > 377 {
			return fmt.Errorf("восьмеричная часть %d превышает 377", i+1)
		}

		// Проверяем диапазон для всех чисел
		if value > 255 {
			return fmt.Errorf("часть %d превышает 255", i+1)
		}

		ip.parts[i] = value
	}
	return nil
}

// вспомогательная функция для парсинга
func parsePart(part string) (int, int, error) {
	if len(part) > 1 && part[0] == '0' && part[1] == 'x' {
		// шестнадцатеричное число
		value, err := strconv.ParseInt(part[2:], 16, 32)
		if err != nil {
			return 0, 0, err
		}
		return int(value), 16, nil
	} else if len(part) > 1 && part[0] == '0' {
		// восьмеричное число
		value, err := strconv.ParseInt(part[1:], 8, 32)
		if err != nil {
			return 0, 0, err
		}
		return int(value), 8, nil
	} else {
		// десятичное число
		value, err := strconv.Atoi(part)
		if err != nil {
			return 0, 0, err
		}
		return value, 10, nil
	}
}

// метод для преобразования IP-адреса в строку
func (ip *IPv4) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.parts[0], ip.parts[1], ip.parts[2], ip.parts[3])
}

func main() {
	fmt.Println("введите IPv4-адрес для проверки:")
	var ipStr string
	fmt.Scanln(&ipStr)

	// создаем объект IPv4
	ip := IPv4{}

	// парсим IP-адрес
	err := ip.Parse(ipStr)
	if err != nil {
		fmt.Printf("oшибка: %v\n", err)
		return
	}

	// проверяем корректность IP-адреса
	if ip.IsValid() {
		fmt.Printf("IP-адрес действителен.\n", ip.String())
	} else {
		fmt.Println("IP-адрес недействителен.")
	}
}
