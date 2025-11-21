package validators

import "testing"

func TestIsValidTitle(t *testing.T) {
	correctTitle := "Normal Title 123"
	result := IsValidTitle(correctTitle)

	if !result {
		t.Error("Корректная строка не прошла валидацию", result)
	}

	incorrectTitle := "Title@With#Symbols"
	result = IsValidTitle(incorrectTitle)

	if result {
		t.Error("Некорректная строка прошла валидацию", result)
	}

	shortTitle := "A"
	result = IsValidTitle(shortTitle)

	if result {
		t.Error("Слишком короткая строка прошла валидацию", result)
	}

	longTitle := "Это очень длинный заголовок который точно превышает лимит в пятьдесят символов"
	result = IsValidTitle(longTitle)
	if result {
		t.Error("Слишком длинная строка прошла валидацию", result)
	}
}
