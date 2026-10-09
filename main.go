package main

import (
	"fmt"

	httpserver "hithab.com/otvs/bd/http_server"
)

func main() {
	fmt.Println("запуск http сервера")

	err := httpserver.StartHTTPserver()

	if err != nil {
		fmt.Println("ошибка во время работы сервера", err)
	} else {
		fmt.Println("сервер завершился успешно")
	}

	/*_, err := os.Create("out/newfile.txt")
	if err != nil {
		panic(err)
	}


	val := os.Getenv("phone_number")
	if val != "" {
		fmt.Println("val", val)
	} else {
		fmt.Println("переменная val не задана")
	}
	ctx := context.Background()

		conn, err := simpleconection.CreateConnection(ctx)
		if err != nil {
			panic(err)
		}
		if err := simplesql.CreateTable(ctx, conn); err != nil {
			panic(err)
		}

		tasks, err := simplesql.SelecktRows(ctx, conn)
		if err != nil {
			panic(err)
		}
		for _, task := range tasks {
			if task.ID == 2 {
				task.Title = "покормить кошку"

				if err := simplesql.UpdateTasks(ctx, conn, task); err != nil {
					panic(err)

				}
				break
			}
		}

		fmt.Println("succeed")
	}*/
}
