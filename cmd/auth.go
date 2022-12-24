package cmd

import (
	"log"
	"time"

	"github.com/dev-hyunsang/clone-twitter-backend/auth"
	"github.com/dev-hyunsang/clone-twitter-backend/database"
	"github.com/dev-hyunsang/clone-twitter-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func JoinUserHandler(c *fiber.Ctx) error {
	req := new(models.RequestJoinUser)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusInternalServerError,
				Status:     "error",
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	userUUID := uuid.New()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusInternalServerError,
				Status:     "error",
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	userData := models.User{
		UserUUID:        userUUID,
		UserEmail:       req.UserEmail,
		UserPhoneNumber: req.UserPhoneNumber,
		UserPassword:    string(hashPassword),
		UserBirthday:    req.UserBirthday,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = database.CreateUser(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusInternalServerError,
				Status:     "error",
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		MetaData: models.MetaData{
			StatusCode: fiber.StatusOK,
			Status:     "success",
			Success:    true,
			Message:    "새로운 유저를 성공적으로 생성했습니다.",
		},
		ResponsedAt: time.Now(),
	})
}

func LoginUserHandler(c *fiber.Ctx) error {
	req := new(models.RequestLoginUser)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusInternalServerError,
				Status:     "error",
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	data, err := database.QueryUser(req.UserEmail)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusBadRequest,
				Status:     "bad request",
				Success:    false,
				Message:    "입력하신 메일 혹은 비밀번호를 찾아볼 수 없습니다. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.UserPassword), []byte(req.UserPassword)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusBadRequest,
				Status:     "bad request",
				Success:    false,
				Message:    "입력하신 메일 혹은 비밀번호를 찾아볼 수 없습니다. 확인 후 다시 시도해 주세요.",
			},
			ResponsedAt: time.Now(),
		})
	}

	token, err := auth.NewAuthJWT(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			MetaData: models.MetaData{
				StatusCode: fiber.StatusInternalServerError,
				Status:     "error",
				Success:    false,
				Message:    err.Error(),
			},
			ResponsedAt: time.Now(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})

	return c.Status(fiber.StatusOK).JSON(models.SuccessLoginResponse{
		MetaData: models.MetaData{
			StatusCode: fiber.StatusOK,
			Status:     "success",
			Success:    true,
			Message:    "성공적으로 로그인 했습니다.",
		},
		UserData: models.UserData{
			Token: token,
		},
		ResponsedAt: time.Now(),
	})

}

func LogoutUserHandler(c *fiber.Ctx) error {
	c.ClearCookie("jwt")

	return c.Status(fiber.StatusOK).JSON(models.SuccessResponse{
		MetaData: models.MetaData{
			StatusCode: fiber.StatusOK,
			Status:     "success",
			Success:    true,
			Message:    "성공적으로 로그아웃 되었습니다.",
		},
		ResponsedAt: time.Now(),
	})
}
