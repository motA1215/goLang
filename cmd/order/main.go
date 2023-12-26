package order

import (
	"database/sql";
	"fmt";
	"strconv";
	"github.com/motA1215/goLang/internal/infra/database";
	"github.com/motA1215/goLang/internal/usecase";
	_ "github.com/mattn/go-sqlite3";
);

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3");
	if err != nil{
		panic(err);	
	}

	uc := usecase.NewCalculateFinalPrice(database.NewOrderRepository(db));

	input := usecase.OrderInput{
		ID: "1235",
		Price: 10.0,
		Taxa: 1.0,
	}

	output, err := uc.Execute(input);
	if err != nil{
		panic(err);	
	}

	rowTotals, err := uc.GetTotalRowsTable();
	if err != nil{
		panic(err);
	}

	fmt.Println(output);
	fmt.Println("Quantidade de preços finais inseridos na tabela de orders: " + strconv.Itoa(rowTotals));
}