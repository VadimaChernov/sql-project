package main

import (
	"context"
	"fmt"

	simpleconection "hithab.com/otvs/bd/simple_conection"
	simplesql "hithab.com/otvs/bd/simple_sql"
)

func main() {
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
}
