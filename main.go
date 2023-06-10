package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	// e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })

	// e.Logger.Fatal(e.Start("localhost:5000"))
	e := echo.New()

	e.Static("/public", "public")

	//Routing

	e.GET("/", home)
	e.GET("/myproject", myproject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/myproject/:id", blogDetail)
	e.POST("/addmyproject", addmyproject)
	e.Logger.Fatal(e.Start("localhost:5000"))

}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	data := map[string]interface{}{
		"login": true,
	}

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func myproject(c echo.Context) error {

	data := map[string]interface{}{
		"login": true,
	}

	var tmpl, err = template.ParseFiles("views/form-project.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func contact(c echo.Context) error {
	data := map[string]interface{}{
		"login": true,
	}
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonial(c echo.Context) error {

	data := map[string]interface{}{
		"login": true,
	}

	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func blogDetail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      id,
		"title":   "Duino-Coin, Sebagai pembelajaran Mining",
		"content": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Neque, eveniet quia in asperiores ducimus, quam maiores excepturi at voluptate quos ipsum? Labore quis dolore id optio nemo placeat voluptas sunt! Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus ratione, pariatur a quo quis quas in atque magnam, eveniet dolores, animi excepturi quidem dignissimos voluptates expedita repellendus aliquid quasi sed.Lorem ipsum dolor sit, amet consectetur adipisicing elit. Est harum illo vero! Quibusdam esse quidem mollitia amet necessitatibus voluptates inventore sapiente eius expedita. Qui magni placeat error? Libero, debitis minus. Lorem ipsum dolor sit amet consectetur adipisicing elit. Eveniet nam corporis vitae omnis soluta, deserunt reprehenderit temporibus debitis error, voluptatem aliquam corrupti necessitatibus aperiam voluptas aut! Deserunt alias nobis dolor? Lorem ipsum dolor sit amet consectetur adipisicing elit. Repudiandae maiores consectetur non nesciunt rerum unde blanditiis ea, eum ducimus libero nulla corporis in a, consequuntur dolorem voluptatum alias. Explicabo, veritatis.",
		"login":   true,
	}

	var tmpl, err = template.ParseFiles("views/myProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}

func addmyproject(c echo.Context) error {
	title := c.FormValue("inputTitle")
	content := c.FormValue("inputContent")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	cbox1 := c.FormValue("iot")
	cbox2 := c.FormValue("ui")
	cbox3 := c.FormValue("full")
	cbox4 := c.FormValue("ml")

	println("Title : " + title)
	println("Content : " + content)
	println("Start Date : " + startDate)
	println("End Date : " + endDate)
	println("Box IoT : " + cbox1)
	println("Box UI UX : " + cbox2)
	println("Box FullStack : " + cbox3)
	println("Box Machine Learning : " + cbox4)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
