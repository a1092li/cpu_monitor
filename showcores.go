package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/proc/stat")
	if err != nil {
		fmt.Printf("Ошибка открытия /proc/stat: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Текущая загрузка ядер CPU:")
	fmt.Println(strings.Repeat("-", 40))

	scanner := bufio.NewScanner(file)
	coresPrinted := 0

	for scanner.Scan() && coresPrinted < 4 {
		line := scanner.Text()

		// Ищем строки с ядрами cpu0, cpu1, cpu2, cpu3
		if strings.HasPrefix(line, "cpu") && !strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)
			if len(fields) < 5 {
				continue
			}

			cpuName := fields[0]

			// Проверяем, что это одно из нужных ядер
			if cpuName == "cpu0" || cpuName == "cpu1" || cpuName == "cpu2" || cpuName == "cpu3" {
				// Парсим значения
				user, _ := strconv.ParseUint(fields[1], 10, 64)
				nice, _ := strconv.ParseUint(fields[2], 10, 64)
				system, _ := strconv.ParseUint(fields[3], 10, 64)
				idle, _ := strconv.ParseUint(fields[4], 10, 64)

				total := user + nice + system + idle
				if len(fields) > 5 {
					iowait, _ := strconv.ParseUint(fields[5], 10, 64)
					total += iowait
				}

				// Вычисляем процент использования
				var usage float64
				if total > 0 {
					used := user + nice + system
					usage = float64(used) / float64(total) * 100
				}

				fmt.Printf("%s: %6.2f%% загрузки\n", cpuName, usage)
				coresPrinted++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
	}
}
