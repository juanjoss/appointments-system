package main

import (
	"appointments-system/handlers"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/template/html"
)

//go:embed views/*
var buildFS embed.FS

func main() {
	app := fiber.New(fiber.Config{
		Views:        html.New("./views", ".html"),
		ErrorHandler: handlers.NotFound,
	})

	build, err := fs.Sub(buildFS, "views")
	if err != nil {
		log.Fatal(err)
	}

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:   http.FS(build),
		MaxAge: 31536000, // 1 year
	}))

	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	/*
	* client routes
	 */
	app.Get("/bookings", handlers.BookingsPage).Name("client_booking")
	app.Get("/booked", handlers.AppointmentBookedPage).Name("appointment_booked")
	app.Post("/bookings", handlers.BookAppointment)

	/*
	* API endpoints
	 */
	apiRouter := app.Group("/api")
	apiRouter.Get("/calendar", handlers.GetCalendar)
	apiRouter.Post("/attentionHours", handlers.GetAttentionHours)
	apiRouter.Post("/appointments", handlers.UpdateAppointment)

	/*
	* authentication routes
	 */
	authRouter := app.Group("/auth")
	authRouter.Get("/login", handlers.LoginPage).Name("login")
	authRouter.Post("/login", handlers.Login)
	authRouter.Post("/logout", handlers.Logout)

	/*
	* admin routes
	 */
	app.Use("/admin/*", handlers.Auth)

	adminRouter := app.Group("/admin")
	adminRouter.Get("/profile", handlers.AdminProfilePage).Name("admin_profile")
	adminRouter.Post("/profile", handlers.UpdateAdminProfile)

	adminRouter.Get("/appointments", handlers.AppointmentsPage).Name("admin_appointments")

	adminRouter.Get("/bookings", handlers.AdminBookingsPage).Name("admin_bookings")
	adminRouter.Post("/bookings", handlers.DeleteBooking)

	adminRouter.Get("/appointmentTypes", handlers.GetAllAppointmentTypes).Name("appointment_types")
	adminRouter.Post("/appointmentTypes", handlers.CreateUpdateAppointmentType)
	adminRouter.Post("/appointmentTypes/:id", handlers.DeleteAppointmentType)

	adminRouter.Get("/clients", handlers.AdminClientsPage).Name("admin_clients")
	adminRouter.Post("/clients", handlers.UpdateClient)

	adminRouter.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
