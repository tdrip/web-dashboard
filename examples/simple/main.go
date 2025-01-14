package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tdrip/web-dashboard/pkg/v1/bootstrap"
	navigation "github.com/tdrip/web-dashboard/pkg/v1/navigation"
	pages "github.com/tdrip/web-dashboard/pkg/v1/pages"
	render "github.com/tdrip/web-dashboard/pkg/v1/render"
)

func main() {
	router := chi.NewRouter()

	fileServer := http.StripPrefix("/public", http.FileServer(http.Dir("./public")))
	router.Handle("/public/*", fileServer)

	pg := pages.Page{}
	pg.Title = "Test"
	pg.HasThemeSwicther = true
	pg.HeaderRawItem = bootstrap.DashBoardStyle
	pg.GetBodyMain = pg.GetEmptyMain
	//pg.BodyRawItem = `
	//<canvas class="my-4 w-100" id="myChart" width="1153" height="487" style="display: block; box-sizing: border-box; height: 487px; width: 1153px;"></canvas>
	//`
	pg.UseEmbeddedBootstrapCSS = false
	pg.UseEmbeddedBootstrapJS = false
	pg.UseEmbeddedDashBoardCSS = false
	pg.UseEmbeddedDashBoardJS = false
	pg.HeaderScripts = []string{
		"/public/color-modes.js",
	}

	pg.HeaderLinks = []pages.Link{
		{Item: "/public/bootstrap.min.css", Type: "stylesheet"},
		{Item: "/public/bootstrap-icons.min.css", Type: "stylesheet"},
		{Item: "/public/css@3.css", Type: "stylesheet"},
		{Item: "/public/dashboard.css", Type: "stylesheet"},
	}

	pg.BodyScripts = []string{
		"/public/bootstrap.bundle.min.js",
		//"/public/chart.umd.js",
		"/public/dashboard.js",
	}

	pg.GetNavMenu = navigation.GetNavigation
	pg.GetBodyHeader = pg.SimpleNav
	pg.NavMenuItems = []navigation.NavMenuItem{
		{
			Title: "A Link",
			Type:  2,
			HREF:  "/",
		},
		{
			Title: "A Link",
			Type:  2,
			HREF:  "/",
		},
		{
			Title: "Test",
			Type:  1,
		},
	}
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		render.RenderPage(request, writer, pg.GetPage)
	})

	http.ListenAndServe(":3000", router)
}
