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
	"bufio"
	"fmt"
	"os"
	"hw13/internal/passwords"
)

func main() {
	pm := passwords.NewPassword("passwords.json")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nPassword Manager")
		fmt.Println("1. List saved passwords")
		fmt.Println("2. Save a new password")
		fmt.Println("3. Retrieve a password")
		fmt.Println("4. Exit")

		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			names := pm.ListPasswords()
			if len(names) == 0 {
				fmt.Println("No passwords saved.")
			} else {
				fmt.Println("Saved passwords:")
				for _, name := range names {
					fmt.Println("-", name)
				}
			}
		case "2":
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter password: ")
			scanner.Scan()
			password := scanner.Text()

			pm.SavePassword(name, password)
		case "3":
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()

			password, exists := pm.GetPassword(name)
			if !exists {
				fmt.Println("Password not found.")
			} else {
				fmt.Printf("Password for %s: %s\n", name, password)
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
