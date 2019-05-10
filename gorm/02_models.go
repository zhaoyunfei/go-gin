package main

import (
	"database/sql"
	"gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        	string
	Age         	sql.NullInt64
	Birthday		*time.Time
	Email    		string		`gorm:"type:varchar(100);unique_index"`
	Role     		string		`gorm:"size:255"`
	MemberNumber	*string		`gorm:"unique;not null"`
	Num          	int			`gorm:"AUTO_INCREMENT"`
	Address      	string 		`gorm:"index:addr"`
	IgnoreMe     	int  		`gorm:"-"`
}

func main() {
	db, _ := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	//create
	user := User{Name: "test", Age: 18, Birthday: time.Now()}
	db.NewRecord(user)
	db.Create(&user)
	//db.Set("gorm:insert_option", "ON CONFLICT").Create(&user)
	//INSERT INTO user() values() ON CONFLICT;
	//	select query
	db.First(&user) //get first record;
	db.Take(&user) // get one record, no specfied order
	db.Last(&user) // order by id desc
	db.Find(&users) // get all record
	db.First(&user, 10) //select * from users where id = 10;

	//select where
	db.Where("name = ?", "test").First(&user)  //select * from users where name = 'test' limit 1;
	db.Where("name = ?", "test").Find(&users)  //select * from users where name = 'test';
	db.Where("name <> ?", "test").Find(&users)  //select * from users where name <> 'test';
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)  //in
	db.Where("name LIKE ?", "%jin%").Find(&users) //like
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users) //and
	db.Where("updated_at > ?", lastWeek).Find(&users) //time
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users) // between

	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user) //use struct
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users) // use Map
	db.Where([]int64{20, 21, 22}).Find(&users) // 主键 slice of primary keys
	db.Not("name", "test").First(&user) //not in
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&user)

	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10) // select for update
	db.FirstOrInit(&user, User{Name: "zhangsan"}) //查不到则根据条件初始化一个新的记录

	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user)
	//SELECT * FROM USERS WHERE name = 'non_existing';
	//unfound: user -> User{Name: "non_existing", Age: 20}

	db.Where(User{Name: "lisi"}).Assign(User{Age: 12}).FirstOrInit(&user)//无论能否查到记录，都将参数赋予记录
	//user -> User{Name: "non_existing", Age: 20}

	db.FirstOrCreate(&user, User{Name: "lisi"})// 查不到则创建根据条件创建一条新记录
	//subquery 子查询
	db.Where("amount > ?", gorm.DB.Table("orders").Select("ANG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)
	//SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	//order by
	db.Order("age desc, name").Find(&users)
	db.Limit(3).Find(&users)
}

