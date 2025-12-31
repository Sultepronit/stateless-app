package gemini

import (
	"regexp"
	"strings"
)

func filterKanji(input string, exclude string) string {
	rx := regexp.MustCompile(`\p{Han}`)
	matches := rx.FindAllString(input, -1)

	enlisted := make(map[string]bool)
	var sb strings.Builder
	for _, m := range matches {
		if m == exclude || enlisted[m] {
			continue
		}

		sb.WriteString(m)
		enlisted[m] = true
	}

	return sb.String()
}

func GuesssKanji(req string) (string, error) {
	instruction := "You are helping someone with limited knowledge of Japanese. They see a kanji they don't know, but they do visually associate it with something they provide as the input. Examples of input -> expected: 心門文 -> 憫; 亦心 -> 恋; 道ホ -> 述; 私 -> 仏; 周 <-> 彫; 殴 <-> 投; 疲 <-> 痩. Please, provide rather long list of suggestions. Guessing is ok. No explanations, just the list of kanji to choose from."
	resp, err := useGemini(instruction, req, true)
	if err != nil {
		return "", err
	}

	return filterKanji(resp, req), nil
}

func TranslateEnUk(req string) (string, error) {
	instruction := "Переклади українською. Якщо мова слова/виразу/тексту — англійська — не згадуй це додатково; якщо не англійська — вкажи її; якщо українcька — переклади англійською. Якщо є якась помилка — вкажи на неї, запропонуй правильний варіант. Результат подати як HTML article."
	instruction = "Переклади українською. Якщо мова завдання — англійська — не згадуй це додатково; якщо не англійська — вкажи її; якщо українcька — переклади англійською. Якщо є якась помилка — вкажи на неї, запропонуй правильний варіант. Результат подати як HTML article."
	instruction = "Ти — словник останньої надії. Після того, як людина не може зрозуміти/знайти переконливий переклад слова/фрази/тексту, вона звертається сюди. Типово, завдання — англійською, якщо так — переклади українською. Якщо навпаки українська — отже переклади англійською. Якщо мова — якась третя — вкажи її, переклади українською. Якщо є якась помилка — вкажи на неї, запропонуй правильний варіант. Результат подати як HTML article."
	instruction = "Ти — словник останньої надії. Після того, як людина не може зрозуміти/знайти переконливий переклад слова/фрази/тексту, вона звертається сюди. Типово, завдання — англійською, якщо так — переклади українською. Якщо навпаки українська — отже переклади англійською. Якщо мова — якась третя — вкажи її, переклади українською. Якщо є якась помилка — вкажи на неї, запропонуй правильний варіант. Результат подати як HTML article (максимум інформації — мінімум тексту для читання)."
	instruction = "Ти — ШІ-словник. Після того, як людина не може зрозуміти/знайти переконливий переклад слова/фрази/тексту, із більш простими засобами — вона звертається сюди. Типово, завдання — англійською, якщо так — переклади українською. Якщо навпаки українська — отже переклади англійською. Якщо мова — якась третя — вкажи її, переклади українською. Якщо є якась помилка — вкажи на неї, запропонуй правильний варіант. Результат подати як HTML article (лише корисна інформація — жодного зайвого тексту)."

	resp, err := useGemini(instruction, req, false)
	if err != nil {
		return "", err
	}

	rgx := regexp.MustCompile(`(?s)<article.*?>(.*?)</article>`)
	match := rgx.FindStringSubmatch(resp)
	if len(match) > 1 {
		return match[1], nil
	}

	return "No response!", nil
}
