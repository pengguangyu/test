package main

import (
	"fmt"
	"strconv"
	"time"

	// "github.com/jinzhu/copier"
	"github.com/ybzhanghx/copier"
	// "hub.fastgit.org/jinzhu/copier"
)

type User struct {
	NameT string `json:"name"`
	Role  string
	Age   int32
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name      string `json:"name"`
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRule string
}

func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func main() {
	CopyByTag2()

	var (
		user  = User{NameT: "Jinzhu", Age: 18, Role: "Admin"}
		users = []User{{NameT: "Jinzhu", Age: 18, Role: "Admin"},
			{NameT: "jinzhu 2", Age: 30, Role: "Dev"}}
		employee  = Employee{}
		employees = []Employee{}
	)

	// copier.Copy(&employee, &user)
	copier.CopyByTag(&employee, &user, "json")

	fmt.Printf("%#v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRule: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	// copier.Copy(&employees, &user)
	copier.CopyByTag(&employees, &user, "json")

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"}
	// }

	// Copy slice to slice
	employees = []Employee{}
	// copier.Copy(&employees, &users)
	copier.CopyByTag(&employees, &users, "json")

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, DoubleAge: 60, EmployeId: 0, SuperRule: "Super Dev"},
	// }
}

///////////////////////

type structA struct {
	Items   string `mson:"Item_string"`
	UserId  string `mson:"UserId_string"`
	PubTime int64  `mson:"PubTime_int64"`
}

type structB struct {
	Item    string    `mson:"Item_string"`
	UserId  int64     `mson:"UserId_int64"`
	PubTime time.Time `mson:"PubTime_time_Time"`
}

func (p *structB) UserId_string(t string) {
	p.UserId, _ = strconv.ParseInt(t, 10, 64)
}

func (p *structB) PubTime_int64(t int64) {
	p.PubTime = time.Unix(t, 0)
}

func CopyByTag2() {
	obj1 := structA{Items: "233", UserId: "5433", PubTime: time.Now().Unix()}
	obj2 := &structB{}
	err := copier.CopyByTag(obj2, &obj1, "mson")
	if err != nil {
		fmt.Println("Should not raise error")
	}

	if obj2.Item != obj1.Items {
		fmt.Println("Field A should be copied")
	}
	if strconv.FormatInt(obj2.UserId, 10) != obj1.UserId {
		fmt.Println("Field B should be copied")
	}
	if obj2.PubTime.Unix() != obj1.PubTime {
		fmt.Println("Field C should be copied")
	}
}
