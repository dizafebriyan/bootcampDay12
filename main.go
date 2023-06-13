package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

// nama dari structnya adalah Blog
type Blog struct {
	Title     string
	Content   string
	StartDate string
	EndDate   string
	box1      string
	box2      string
	box3      string
	box4      string
}

var dataProject = []Blog{}

func main() {
	// e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })

	// e.Logger.Fatal(e.Start("localhost:5000"))
	e := echo.New()

	e.Static("/public", "public")

	//Routing

	//get
	e.GET("/", home)
	e.GET("/myproject", myproject)
	e.GET("/contact", contact)
	e.GET("/testimonial", testimonial)
	e.GET("/myproject/:id", blogDetail)

	//post
	e.POST("/addmyproject", addmyproject)
	e.POST("/project-delete/:id", deleteproject)

	e.Logger.Fatal(e.Start("localhost:5500"))

}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	// data := map[string]interface{}{
	// 	"login": true,
	// }

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"project": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
}

func myproject(c echo.Context) error {

	// data := map[string]interface{}{
	// 	"login": true,
	// }

	var tmpl, err = template.ParseFiles("views/form-project.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	// data := map[string]interface{}{
	// 	"login": true,
	// }
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {

	// data := map[string]interface{}{
	// 	"login": true,
	// }

	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil { //nil == null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{}{
	// 	"id":      id,
	// 	"title":   "Duino-Coin, Sebagai pembelajaran Mining",
	// 	"content": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Neque, eveniet quia in asperiores ducimus, quam maiores excepturi at voluptate quos ipsum? Labore quis dolore id optio nemo placeat voluptas sunt! Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus ratione, pariatur a quo quis quas in atque magnam, eveniet dolores, animi excepturi quidem dignissimos voluptates expedita repellendus aliquid quasi sed.Lorem ipsum dolor sit, amet consectetur adipisicing elit. Est harum illo vero! Quibusdam esse quidem mollitia amet necessitatibus voluptates inventore sapiente eius expedita. Qui magni placeat error? Libero, debitis minus. Lorem ipsum dolor sit amet consectetur adipisicing elit. Eveniet nam corporis vitae omnis soluta, deserunt reprehenderit temporibus debitis error, voluptatem aliquam corrupti necessitatibus aperiam voluptas aut! Deserunt alias nobis dolor? Lorem ipsum dolor sit amet consectetur adipisicing elit. Repudiandae maiores consectetur non nesciunt rerum unde blanditiis ea, eum ducimus libero nulla corporis in a, consequuntur dolorem voluptatum alias. Explicabo, veritatis.",
	// }

	var ProjectDetail = Blog{}

	//intinya disini membangun sebuah variabel "Blog" Baru bernama ProjectDetail untuk menampung 1 data Detailnya.
	//perulangan dibawah digunakan untuk mencari index yang sesuai yang terdapat pada inputan dan nantinya jika index
	//sudah sesuai maka akan langsung di salin datanya ke variabel ProjectDetail.
	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Blog{
				Title:     data.Title,
				Content:   data.Content,
				StartDate: data.StartDate,
				EndDate:   data.EndDate,
				box1:      data.box1,
				box2:      data.box2,
				box3:      data.box3,
				box4:      data.box4,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": ProjectDetail,
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

	var newProject = Blog{
		Title:     title,
		Content:   content,
		StartDate: startDate,
		EndDate:   endDate,
		box1:      cbox1,
		box2:      cbox2,
		box3:      cbox3,
		box4:      cbox4,
	}

	dataProject = append(dataProject, newProject)

	println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteproject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index: ", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
