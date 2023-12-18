package main

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type User struct {
	Email string `json:"email"`
}

type Record struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
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
		err := ctx.Bind(&user)

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
		err := ctx.Bind(&body)

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

			if err := rows.Scan(&record.Id, &record.Email, &record.Date, &record.Amount, &record.Description); err != nil {
				return nil, err
			}

			records = append(records, record)
		}

		return records, nil

	})

	app.POST("/record", func(ctx *gofr.Context) (interface{}, error) {
		type RequestBody struct {
			Email       string `json:"email"`
			Date        string `json:"date"`
			Amount      string `json:"amount"`
			Description string `json:"description"`
		}

		var body RequestBody
		err := ctx.Bind(&body)

		if err != nil {
			return nil, err
		}

		// check if user exists, with given email address
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM user WHERE email=?", body.Email)

		if err != nil {
			return nil, err
		}

		count := 0
		for rows.Next() {
			count++
			break
		}

		// user exists
		if count != 0 {
			// insert into records table
			_, err := ctx.DB().ExecContext(ctx, "INSERT INTO record (email, date, amount, description) VALUES (?, ?, ?, ?)", body.Email, body.Date, body.Amount, body.Description)

			if err != nil {
				return nil, err
			}

			return "Record successfully added", nil
		} else {
			return nil, &errors.Response{
				StatusCode: 400,
				Reason:     "User does not exists",
			}
		}
	})

	app.PATCH("/record/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		type RequestBody struct {
			Email       string `json:"email"`
			Date        string `json:"date"`
			Amount      string `json:"amount"`
			Description string `json:"description"`
		}

		var body RequestBody
		err := ctx.Bind(&body)

		if err != nil {
			return nil, err
		}

		_, err = ctx.DB().ExecContext(ctx, "UPDATE record SET email=?, date=?, amount=?, description=? WHERE id=?", body.Email, body.Date, body.Amount, body.Description, id)

		if err != nil {
			return nil, err
		}

		return "Data updated success", nil
	})

	app.DELETE("/record/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")

		_, err := ctx.DB().ExecContext(ctx, "DELETE FROM record WHERE id=?", id)

		if err != nil {
			return nil, err
		}

		return "Deleted successfully", nil
	})

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
