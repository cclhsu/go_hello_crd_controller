package resources

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Resources struct {
	// gorm.Model
	Id      int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name    string `gorm:"not null" form:"name" json:"name"`
	Enable  bool   `gorm:"not null" form:"enable" json:"enable"`
	State   string `gorm:"not null" form:"state" json:"state"`
	Message string `json:"message,omitempty"`

	// Health ResourceHealthState `json:"health,omitempty"`
}

// type ResourceHealthState struct {
// 	// gorm.Model
// 	Enable  bool   `json:"enable"`
// 	State   string `json:"state"`
// 	Message string `json:"message,omitempty"`
// }

func InitDb() *gorm.DB {

	dir := "./data"
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	db, err := gorm.Open("sqlite3", dir+"/"+"gorm_sqlite3.db")
	// db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&Resources{}) {
		db.CreateTable(&Resources{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Resources{})
		// db.CreateTable(&Resources{}, &ResourceHealthState{})
		// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Resources{}, &ResourceHealthState{})
	}

	return db
}

// var db = InitDb()

func CreateResourceTable(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	if !db.HasTable(&Resources{}) {
		db.CreateTable(&Resources{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Resources{})
		// db.CreateTable(&Resources{}, &ResourceHealthState{})
		// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Resources{}, &ResourceHealthState{})
	}
}

func DropResourceTable(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	if db.HasTable(&Resources{}) {
		// db.DropTableIfExists(&Resources{})
		db.DropTable(&Resources{})
	}
}

func GetResources(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	// NOTE: Get items from database
	var resources []Resources
	// SELECT * FROM resources
	db.Find(&resources)

	c.JSON(http.StatusOK, resources)
	// curl -i http://localhost:8080/api/v1/resources
}

func GetResource(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")

	// NOTE: Get item from database
	var resource Resources
	// SELECT * FROM resources WHERE id = 1;
	db.First(&resource, id)

	if resource.Id != 0 {
		c.JSON(http.StatusOK, resource)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
	}
	// curl -i http://localhost:8080/api/v1/resources/1
}

func PostResource(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var resource Resources
	c.Bind(&resource)

	if resource.Name != "" {
		// NOTE: Do create action and insert new item to database.
		// INSERT INTO "resources" (name) VALUES (resource.Name);
		db.Create(&resource)
		c.JSON(http.StatusCreated, resource)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Fields are empty"})
	}
	// curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"A\",\"enable\":false,\"state\":\"\"}" http://localhost:8080/api/v1/resources
	// curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"A\",\"health\":{\"enable\":false,\"state\":\"\"}}" http://localhost:8080/api/v1/resources
}

func UpdateResource(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var currentResource Resources
	// SELECT * FROM resources WHERE id = 1;
	db.First(&currentResource, id)

	if currentResource.Name != "" {
		// NOTE: Do update action and update fields for item in database.
		if currentResource.Id != 0 {
			var updateResource Resources
			c.Bind(&updateResource)

			resource := Resources{
				Id:      currentResource.Id,
				Name:    updateResource.Name,
				Enable:  updateResource.Enable,
				State:   updateResource.State,
				Message: updateResource.Message,
				// Health: updateResource.Health,
			}

			// UPDATE resources SET state='updateResource.Status', name='updateResource.Name' WHERE id = currentResource.Id;
			db.Save(&resource)

			c.JSON(http.StatusOK, resource)

		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		}
	} else {
		// Display JSON error
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Fields are empty"})
	}
	// curl -i -X PUT -H "Content-Type: application/json" -d "{\"name\":\"B\",\"enable\":false,\"state":\"\"}" http://localhost:8080/api/v1/resources/1
	// curl -i -X PUT -H "Content-Type: application/json" -d "{\"name\":\"B\",\"health\":{\"enable\":false,\"state\":\"\"}}" http://localhost:8080/api/v1/resources/1
}

func DeleteResource(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var resource Resources
	// SELECT * FROM resources WHERE id = 1;
	db.First(&resource, id)

	if resource.Id != 0 {
		// NOTE: Do delete action and delete item from database
		// DELETE FROM resources WHERE id = resource.Id
		db.Delete(&resource)
		c.JSON(http.StatusOK, gin.H{"id #" + id: " deleted"})

	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
	}
	// curl -i -X DELETE http://localhost:8080/api/v1/resources/1
}

func OptionsResource(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func GetResourceHealth(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")

	// NOTE: Get items health state from environment/database
	var resource Resources
	// SELECT * FROM resources WHERE id = 1;
	db.First(&resource, id)

	if resource.Id != 0 {
		c.JSON(http.StatusOK, resource)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
	}
	// curl -i http://localhost:8080/api/v1/resources/1/health
}

func UpdateResourceHealth(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	// NOTE: update items health state from environment/database
	id := c.Params.ByName("id")
	var currentResource Resources
	// SELECT * FROM resources WHERE id = 1;
	db.First(&currentResource, id)

	if currentResource.Enable == false {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Resource Update is disable"})
	} else if currentResource.Name != "" {

		if currentResource.Id != 0 {
			var updateResource Resources
			c.Bind(&updateResource)

			resource := Resources{
				Id:      currentResource.Id,
				Name:    updateResource.Name,
				Enable:  updateResource.Enable,
				State:   updateResource.State,
				Message: updateResource.Message,
				// Health: updateResource.Health,
			}

			// UPDATE resources SET state='updateResource.Status', name='updateResource.Name' WHERE id = currentResource.Id;
			db.Save(&resource)

			c.JSON(http.StatusOK, resource)

		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		}
	} else {
		// Display JSON error
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Fields are empty"})
	}
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\":"A\",\"enable\":true,\"state\":"1\" }" http://localhost:8080/api/v1/resources/1
	// curl -i -X PUT -H "Content-Type: application/json" -d "{\"name\":\"A\",\"health\":{\"enable\":true,\"state\":\"\"}}" http://localhost:8080/api/v1/resources/1
}
