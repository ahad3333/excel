package excel

func GetExcelByLetters(columnTitle string) int {
    result := 0
    for i := 0; i < len(columnTitle); i++ {
        value := int(columnTitle[i] - 'A' + 1)
        result = result*26 + value
    }
    return result
}
