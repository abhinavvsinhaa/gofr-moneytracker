package main

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type User struct {
	Email string `json:"email"`
}

type Record struct {
	Id          int     `json:"id"`
	Email       string  `json:"email"`
	Date        string  `json:"date"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	// POST user
	app.POST("/user", func(ctx *gofr.Context) (interface{}, error) {
		var user User
		err := ctx.Bind(user)

		if err != nil {
			return nil, err
		}

		// checks if user already exists with the same email
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM user WHERE email=?", user.Email)

		if err != nil {
			return nil, err
		}

		rowsSize := 0
		for rows.Next() {
			rowsSize++
			break
		}

		// user exists
		if rowsSize > 0 {
			return nil, &errors.Response{
				StatusCode: 400,
				Reason:     "User already exists",
			}
		} else {
			// add user into the table, since the user doesnt exists
			_, err := ctx.DB().ExecContext(ctx, "INSERT INTO user (email) VALUES (?)", user.Email)

			if err != nil {
				return nil, &errors.Response{
					StatusCode: 400,
					Reason:     "Error creating user",
				}
			}

			return "User added successfully", nil

		}
	})

	// GET record
	app.GET("/record", func(ctx *gofr.Context) (interface{}, error) {
		type RequestBody struct {
			Email string `json:"email"`
		}

		var body RequestBody
		err := ctx.Bind(body)

		if err != nil {
			return nil, err
		}

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM record WHERE email=?", body.Email)

		if err != nil {
			return nil, err
		}

		var records []Record

		for rows.Next() {
			var record Record

			if err := rows.Scan(&record.Amount, &record.Date, &record.Description); err != nil {
				return nil, err
			}

			records = append(records, record)
		}

		return records, nil

	})

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
