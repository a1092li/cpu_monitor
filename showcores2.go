package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CPUStats struct {
	User    uint64
	Nice    uint64
	System  uint64
	Idle    uint64
	IOWait  uint64
	IRQ     uint64
	SoftIRQ uint64
	Steal   uint64
	Guest   uint64
}

func readCPUStats() (map[string]CPUStats, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := make(map[string]CPUStats)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "cpu") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		cpuName := fields[0]
		var cpuStats CPUStats

		// Парсим числовые значения
		values := make([]uint64, len(fields)-1)
		for i := 1; i < len(fields); i++ {
			val, err := strconv.ParseUint(fields[i], 10, 64)
			if err != nil {
				continue
			}
			values[i-1] = val
		}

		// Заполняем структуру в зависимости от количества доступных полей
		if len(values) > 0 {
			cpuStats.User = values[0]
		}
		if len(values) > 1 {
			cpuStats.Nice = values[1]
		}
		if len(values) > 2 {
			cpuStats.System = values[2]
		}
		if len(values) > 3 {
			cpuStats.Idle = values[3]
		}
		if len(values) > 4 {
			cpuStats.IOWait = values[4]
		}
		if len(values) > 5 {
			cpuStats.IRQ = values[5]
		}
		if len(values) > 6 {
			cpuStats.SoftIRQ = values[6]
		}
		if len(values) > 7 {
			cpuStats.Steal = values[7]
		}
		if len(values) > 8 {
			cpuStats.Guest = values[8]
		}

		stats[cpuName] = cpuStats
	}

	return stats, scanner.Err()
}

func calculateCPUUsage(stats1, stats2 CPUStats) float64 {
	// Вычисляем общее время CPU
	total1 := stats1.User + stats1.Nice + stats1.System + stats1.Idle +
		stats1.IOWait + stats1.IRQ + stats1.SoftIRQ + stats1.Steal + stats1.Guest

	total2 := stats2.User + stats2.Nice + stats2.System + stats2.Idle +
		stats2.IOWait + stats2.IRQ + stats2.SoftIRQ + stats2.Steal + stats2.Guest

	// Вычисляем время простоя
	idle1 := stats1.Idle + stats1.IOWait
	idle2 := stats2.Idle + stats2.IOWait

	// Вычисляем разницу
	totalDiff := total2 - total1
	idleDiff := idle2 - idle1

	if totalDiff == 0 {
		return 0.0
	}

	// Процент использования = 100% - процент простоя
	usage := 100.0 * (float64(totalDiff-idleDiff) / float64(totalDiff))
	return usage
}

func parseCPUCores(args []string) ([]int, error) {
	var cores []int

	if len(args) == 0 {
		// По умолчанию мониторим ядра 0,1,2,3
		return []int{0, 1, 2, 3}, nil
	}

	for _, arg := range args {
		core, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("Incorrect Core number: %s", arg)
		}
		if core < 0 {
			return nil, fmt.Errorf("Core Number should be above 0: %d", core)
		}
		if core > 999 {
			return nil, fmt.Errorf("Core Number too much (max. 999): %d", core)
		}
		cores = append(cores, core)
	}

	return cores, nil
}

func getCPUName(core int) string {
	return fmt.Sprintf("cpu%d", core)
}

func getDisplayName(core int) string {
	// Для корректного выравнивания используем фиксированную ширину
	return fmt.Sprintf("CPU %3d", core)
}

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  cpu_monitor [cores...]")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  cpu_monitor              # Cores 0,1,2,3 by default")
	fmt.Println("  cpu_monitor 0 2          # Monitor Cores 0 and 2")
	fmt.Println("  cpu_monitor 1 2 5 7      # Monitor Cores 1,2,5,7")
	fmt.Println("  cpu_monitor 0 12 128     # Monitor Cores 0, 12, 128")
	fmt.Println("  cpu_monitor 0            # Monitor Core 0 only")
	fmt.Println()
	fmt.Println("Ctrl+C to exit")
}

func main() {
	var coresToMonitor []int
	var err error

	// Парсим аргументы командной строки
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			showUsage()
			return
		}
		coresToMonitor, err = parseCPUCores(os.Args[1:])
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			showUsage()
			return
		}
	} else {
		// Используем ядра по умолчанию
		coresToMonitor = []int{0, 1, 2, 3}
	}

	fmt.Printf("CPU Cores Monitor: %v\n", coresToMonitor)
	fmt.Println("Ctrl+C to exit")

	// Динамически вычисляем ширину разделителя
	separatorWidth := 20 + len(coresToMonitor)*15
	if separatorWidth < 60 {
		separatorWidth = 60
	}
	fmt.Println(strings.Repeat("-", separatorWidth))

	for {
		// Первое измерение
		stats1, err := readCPUStats()
		if err != nil {
			fmt.Printf("Error read CPU stat: %v\n", err)
			return
		}

		// Ждем 1 секунду
		time.Sleep(1 * time.Second)

		// Второе измерение
		stats2, err := readCPUStats()
		if err != nil {
			fmt.Printf("Error read CPU stat: %v\n", err)
			return
		}

		// Выводим загрузку для указанных ядер
		fmt.Printf("\rTime: %s ", time.Now().Format("15:04:05"))

		for _, core := range coresToMonitor {
			cpuName := getCPUName(core)
			displayName := getDisplayName(core)

			if stats1[cpuName].User == 0 && stats2[cpuName].User == 0 {
				fmt.Printf("| %s: ---  ", displayName)
				continue
			}

			usage := calculateCPUUsage(stats1[cpuName], stats2[cpuName])
			fmt.Printf("| %s: %5.1f%% ", displayName, usage)
		}

		// Также показываем общую загрузку CPU
		//if stats1["cpu"] != (CPUStats{}) && stats2["cpu"] != (CPUStats{}) {
		//	totalUsage := calculateCPUUsage(stats1["cpu"], stats2["cpu"])
		//	fmt.Printf("| Всего: %5.1f%%", totalUsage)
		//}

		// Используем \r для обновления строки вместо добавления новой
		fmt.Print("   ")
	}
}
