// Консольний менеджер паролів 
// Написати консольну програму для зберігання паролів (спрощений аналог утиліти pass в UNIX). Шифрування паролів в цій роботі не реалізуємо.
// Функціонал: 
// * вивести назви збережених паролів
// * зберегти пароль за назвою (введення паролю у через fmt.Scan)
// * дістати збережений пароль
// Додаткові умови:
// * використовуємо tracer bullet development, тобто пишемо ітеративно
// * зберігати стан у файлі (щоб паролі можна було дивитися між запусками)
// * використати рекомендовану структуру пакетів (cmd, internal, …)

package main

import (
	"fmt"
	"os"
	"hw13/internal/passwords"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("./pass-manager set <name> <password>")
		fmt.Println("./pass-manager get <name>")
		fmt.Println("./pass-manager get-all")
		fmt.Println("./pass-manager help")
		return
	}

	pm := passwords.NewPassword("passwords.json")
	command := os.Args[1]

	switch command {
	case "set":
		if len(os.Args) != 4 {
			fmt.Println("Usage: ./pass-manager set <name> <password>")
			return
		}
		name := os.Args[2]
		password := os.Args[3]
		pm.SavePassword(name, password)

	case "get":
		if len(os.Args) != 3 {
			fmt.Println("Usage: ./pass-manager get <name>")
			return
		}
		name := os.Args[2]
		password, exists := pm.GetPassword(name)
		if !exists {
			fmt.Println("Password not found.")
		} else {
			fmt.Printf("Password for %s: %s\n", name, password)
		}

	case "get-all":
		names := pm.ListPasswords()
		if len(names) == 0 {
			fmt.Println("No passwords saved.")
		} else {
			fmt.Println("Saved passwords:")
			for _, name := range names {
				fmt.Println("-", name)
			}
		}

	case "help":
		fmt.Println("Usage:")
		fmt.Println("./pass-manager set <name> <password>")
		fmt.Println("./pass-manager get <name>")
		fmt.Println("./pass-manager get-all")
		fmt.Println("./pass-manager help")

	default:
		fmt.Println("Unknown command. Use 'help' to see the list of available commands.")
	}
}
