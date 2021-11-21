package main

import(

	"log"
	"context"
	"strconv"
	"app/dbapp"
	"google.golang.org/grpc"

	"net/http"
	"github.com/gin-gonic/gin"
)

const S_A = "localhost:4000"

var stab dbapp.UsersInfoClient

func PostUser (ctx *gin.Context) {

	var newUser dbapp.UserProfile

	ctx.ShouldBindJSON(&newUser)

	requset := &dbapp.PostUserReq {

		NewUser: &dbapp.UserProfile {

			FullName: newUser.FullName,
			UserName: newUser.UserName,
			PhoneNumber: newUser.PhoneNumber,
		},
	}

	result, err := stab.PostUser(context.Background(), requset )

	log.Println(newUser)

	if err != nil {
		log.Fatalf("getting result error: %v", err)
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func GetUser (ctx *gin.Context) {

	id := ctx.Param("id")

	i, _ := strconv.Atoi(id)

	log.Println(i)

	result, err := stab.GetUser(context.Background(), &dbapp.GetUserReq { UserId: int64(i)})

	if err != nil {
		log.Fatalf("getting result error: %v", err)
	}

	ctx.IndentedJSON(http.StatusCreated, result)
}

/*func UpdateUser (ctx *gin.Context) {

	var uUser dbapp.UserProfile

	ctx.ShouldBindJSON(&uUser)

	requset := &dbapp.UpdateUserReq {

		UpdatedUser: &dbapp.UserProfile {

			FullName: uUser.FullName,
			UserName: uUser.UserName,
			PhoneNumber: uUser.PhoneNumber,
		},
	}

	result, err := stab.UpdateUser(context.Background(), requset )

	log.Println(result, "=============")

	if err != nil {
		log.Fatalf("getting result error: %v", err)
	}

	ctx.IndentedJSON(http.StatusCreated, result)
}*/

func DeleteUser (ctx *gin.Context) {

	var id dbapp.DeleteUserReq

	ctx.ShouldBindJSON(&id)

	result, err := stab.DeleteUser(context.Background(), &id)

	if err != nil {
		log.Fatalf("getting result error: %v", err)
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func main () {

	connection, err := grpc.Dial(S_A, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("connection error: %v", err)
	}

	defer connection.Close()

	stab = dbapp.NewUsersInfoClient(connection)

	router := gin.Default()

	router.POST("/postuser", PostUser)
	router.GET("/users/:id", GetUser)
	//router.PUT("/user/editinfo", UpdateUser)
	router.DELETE("/user/clear", DeleteUser)

	router.Run("localhost:4040")
}