package router

import (
	adminRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/admin"
	bioRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/bio"
	contactRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/contacts"
	experienceRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/experience"
	projectRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/projects"
	resumeRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/resume"
	skillRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/skills"
	socialRoute "github.com/2SSK/portfolio_terminal/backend/internals/routes/socials"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	adminRoute.SetupAdminRoutes(api)
	bioRoute.SetupBioRoutes(api)
	skillRoute.SetupSkillRoute(api)
	contactRoute.SetupContactRoute(api)
	socialRoute.SetupSocailRoute(api)
	resumeRoute.SetupResumeRoute(api)
	projectRoute.SetupProjectRoute(api)
	experienceRoute.SetupExperienceRoute(api)
}
