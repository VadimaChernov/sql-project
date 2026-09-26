package main

import (
	"fmt"
	"os"
)

func main() {
	val := os.Getenv("phone_number")
	if val != "" {
		fmt.Println("val", val)
	} else {
		fmt.Println("переменная val не задана")
	}
	/*ctx := context.Background()

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
