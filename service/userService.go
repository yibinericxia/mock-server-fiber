package service

import (
	datafile "mock-server-fiber/datafile"
	"mock-server-fiber/model"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var allUsers []model.User
var maxID int = 0

func CreateUser(ctx *fiber.Ctx) error {
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Wrong input",
			"message": err.Error(),
		})

	}
	allUsers = getAllUsers()
	maxID += 1
	user.ID = maxID
	allUsers = append(allUsers, user)
	time.Sleep(time.Duration(seconds) * time.Second)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User has created",
		"user":    user,
	})
}

func FindUserByID(ctx *fiber.Ctx) error {
	user, error := getUserByID(ctx, ctx.Params("id"))
	if user == nil || error != nil {
		return error
	}
	time.Sleep(time.Duration(seconds) * time.Second)
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func FindUsers(ctx *fiber.Ctx) error {
	allUsers = getAllUsers()
	time.Sleep(time.Duration(seconds) * time.Second)
	return ctx.Status(fiber.StatusOK).JSON(allUsers)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var newUser model.User
	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Cannot parse",
			"message": err.Error(),
		})
	}
	idx := newUser.ID
	allUsers = getAllUsers()
	if idx <= 0 || idx > len(allUsers) {
		return ctx.Status(fiber.StatusUnprocessableEntity).SendString("Wrong ID in body")
	}
	user := &allUsers[idx-1]
	user.ID = newUser.ID
	user.Name = newUser.Name
	user.Email = newUser.Email
	time.Sleep(time.Duration(seconds) * time.Second)
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func DeleteUserByID(ctx *fiber.Ctx) error {
	user, error := getUserByID(ctx, ctx.Params("id"))
	if user == nil || error != nil {
		return error
	}
	// delete
	allUsers = getAllUsers()
	var newUsers []model.User
	for i := range allUsers {
		if allUsers[i] != *user {
			newUsers = append(newUsers, allUsers[i])
		}
	}
	allUsers = newUsers
	time.Sleep(time.Duration(seconds) * time.Second)
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func getAllUsers() []model.User {
	if len(allUsers) == 0 {
		allUsers = datafile.GetUserData()
		maxID = len(allUsers)
	}
	return allUsers
}

func getUserByID(ctx *fiber.Ctx, id string) (*model.User, error) {
	idx, err := strconv.Atoi(id)
	if err != nil || idx <= 0 {
		return nil, ctx.Status(fiber.StatusBadRequest).SendString("Wrong user id")
	}
	users := getAllUsers()
	for i := range users {
		if users[i].ID == idx {
			return &users[i], err
		}
	}
	return nil, ctx.Status(fiber.StatusNotFound).SendString("User not found")
}
