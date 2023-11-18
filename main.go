package mergeMKVnMKA

package main

import (
"bufio"
"fmt"
"os"
"os/exec"
"path/filepath"
"sort"
"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите путь к папке с видеофайлами: ")
	videoFolder, _ := reader.ReadString('\n')
	videoFolder = strings.TrimSpace(videoFolder)

	fmt.Print("Введите путь к папке с аудиофайлами: ")
	audioFolder, _ := reader.ReadString('\n')
	audioFolder = strings.TrimSpace(audioFolder)

	fmt.Print("Введите путь к папке для сохранения результатов: ")
	outputFolder, _ := reader.ReadString('\n')
	outputFolder = strings.TrimSpace(outputFolder)

	videoFiles, err := ReadFiles(videoFolder, ".mkv")
	if err != nil {
		panic(err)
	}

	audioFiles, err := ReadFiles(audioFolder, ".mka")
	if err != nil {
		panic(err)
	}

	if len(videoFiles) != len(audioFiles) {
		fmt.Println("Количество видео и аудиофайлов не совпадает!")
		return
	}

	for i, videoFile := range videoFiles {
		videoPath := filepath.Join(videoFolder, videoFile)
		audioPath := filepath.Join(audioFolder, audioFiles[i])
		outputPath := filepath.Join(outputFolder, fmt.Sprintf("output_%s", videoFile))

		cmd := exec.Command("ffmpeg", "-i", videoPath, "-i", audioPath, "-c", "copy", "-map", "0:v", "-map", "1:a", outputPath)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Ошибка при объединении файла %s: %v\n", videoFile, err)
		} else {
			fmt.Printf("Файл %s успешно обработан\n", videoFile)
		}
	}

	fmt.Println("Объединение файлов завершено.")
}

func ReadFiles(folder string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
			files = append(files, info.Name())
		}
		return nil
	})
	sort.Strings(files)
	return files, err
}
