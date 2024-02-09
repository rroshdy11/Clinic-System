package main

import (
	SlotEntity "Clinic_System/entity/Slot"
	UserEntity "Clinic_System/entity/User"
	repository "Clinic_System/repository"
	service "Clinic_System/services"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.New()
	slotRepository repository.SlotRepository = repository.NewSlotRepository()
	UserService    service.UserService       = service.New(userRepository)
	slotService    service.SlotService       = service.NewSlotService(slotRepository)
)

func main() {
	// Default With the Logger and Recovery middleware already attached

	server := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	server.Use(cors.New(config))

	server.POST("/SignUp", func(ctx *gin.Context) {
		var user UserEntity.User
		ctx.BindJSON(&user)
		if _, err := UserService.SignUp(ctx, user); err == nil {
			ctx.JSON(200, gin.H{"message": "user created Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}

	})
	server.GET("/GetAllUsers", func(ctx *gin.Context) {
		if users, err := UserService.GetAll(ctx); err == nil {
			ctx.JSON(200, users)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})
	server.GET("/GetAllDoctors", func(ctx *gin.Context) {
		if users, err := UserService.GetAllDoctors(ctx); err == nil {
			ctx.JSON(200, users)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.POST("/SignIn", func(ctx *gin.Context) {
		var user UserEntity.User
		ctx.BindJSON(&user)
		if user1, err := UserService.LogIn(ctx, user.Email, user.Password); err == nil {
			ctx.JSON(200, user1)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.POST("/CreateNewSlot", func(ctx *gin.Context) {
		var slot SlotEntity.Slot
		ctx.BindJSON(&slot)
		if _, err := slotService.CreateNewSlot(ctx, slot); err == nil {
			ctx.JSON(200, gin.H{"message": "slot created Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.GET("/GetAllSlots", func(ctx *gin.Context) {
		if slots, err := slotService.GetAll(ctx); err == nil {
			ctx.JSON(200, slots)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})
	server.GET("/GetAllDoctorSlots/:doctorId", func(ctx *gin.Context) {
		print(ctx.Param("doctorId"))
		if slots, err := slotService.GetAllDoctorSlots(ctx, ctx.Param("doctorId")); err == nil {
			ctx.JSON(200, slots)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})
	server.POST("/DoctorUpdateSlot/:slotId", func(ctx *gin.Context) {
		var slot SlotEntity.Slot
		ctx.BindJSON(&slot)
		//convert the slotId from string to int
		intSlotId, _ := strconv.Atoi(ctx.Param("slotId"))
		if _, err := slotService.DoctorUpdateSlot(ctx, intSlotId, slot); err == nil {
			ctx.JSON(200, gin.H{"message": "slot updated Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.DELETE("/DoctorDeleteSlot/:slotId", func(ctx *gin.Context) {
		//convert the slotId from string to int
		intSlotId, _ := strconv.Atoi(ctx.Param("slotId"))
		if _, err := slotService.DoctorDeleteSlot(ctx, intSlotId); err == nil {
			ctx.JSON(200, gin.H{"message": "slot deleted Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.GET("/GetAllAvailableSlots/:doctorId", func(ctx *gin.Context) {
		if slots, err := slotService.GetAllAvailableSlots(ctx, ctx.Param("doctorId")); err == nil {
			ctx.JSON(200, slots)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.POST("/PatientReserveSlot", func(ctx *gin.Context) {
		var slot SlotEntity.Slot
		ctx.BindJSON(&slot)
		if _, err := slotService.PatientReserveSlot(ctx, slot); err == nil {
			ctx.JSON(200, gin.H{"message": "slot reserved Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.DELETE("/PatientCancelSlot/:slotId", func(ctx *gin.Context) {
		//convert the slotId from string to int
		intSlotId, _ := strconv.Atoi(ctx.Param("slotId"))
		if _, err := slotService.PatientCancelSlot(ctx, intSlotId); err == nil {
			ctx.JSON(200, gin.H{"message": "slot canceled Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.GET("/GetAllPatientSlots/:patientId", func(ctx *gin.Context) {
		if slots, err := slotService.PatientGetAllReservedSlots(ctx, ctx.Param("patientId")); err == nil {
			ctx.JSON(200, slots)
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})

	server.PUT("/UpdateSlotStatus/:slotIdToCancel/:slotIdToReserve", func(ctx *gin.Context) {
		//convert the slotId from string to int
		intSlotIdToCancel, _ := strconv.Atoi(ctx.Param("slotIdToCancel"))
		intSlotIdToReserve, _ := strconv.Atoi(ctx.Param("slotIdToReserve"))
		if _, err := slotService.UpdateSlotStatus(ctx, intSlotIdToCancel, intSlotIdToReserve); err == nil {
			ctx.JSON(200, gin.H{"message": "slot updated Secussfully"})
		} else {
			ctx.JSON(403, gin.H{"error": err.Error()})
		}
	})
	// Configure the CORS middleware with the "content-type" header
	SERVER_PORT := os.Getenv("SERVER_PORT")
	server.Run(":" + SERVER_PORT)
}
