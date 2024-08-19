// паект keyboard используется для чтения ввода с клавиатуры
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat читаеь число с плавающей точкой введенное с клавитауры
// возвращает число и и ошибку
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, err
}
