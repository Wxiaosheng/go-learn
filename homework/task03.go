package homework

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type student struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Age   int    `json:"age" gorm:"not null"`
	Grade string `json:"grade" gorm:"not null"`
}

var dsn = "root:root@tcp(localhost:3306)/golearn?charset=utf8mb4&parseTime=True&loc=Local"

func Task03() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&student{})

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// INSERT INTO students (name, age, grade) VALUES ("张三", 20, "三年级");
	// gorm 增 Create
	u := student{
		Name:  "张三1",
		Age:   20,
		Grade: "三年级1",
	}
	db.Create(&u)

	// SELECT * FROM students WHERE age > 18;
	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// gorm 查 Find、First、Last、Take
	var u2 student
	db.Where("age > ?", 18).Find(&u2)
	fmt.Printf("%+v\n", u2)

	// UPDATE students SET grade = "四年级" WHERE name = "张三"
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// gorm 改 Update
	db.Model(&student{}).Where("name = ?", "张三").Update("grade", "四年级1")

	// DELETE FROM students WHERE age < 15
	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	// gorm 删 Delete
	// 软删除（表结构中要有 deleted_at 字段）
	db.Where("age < ?", 15).Delete(&student{})

	// 使用 Unscoped() 进行硬删除
	// db.Unscoped().Delete(&student{})
}

/*
假设有两个表：

  accounts 表（包含字段 id 主键， balance 账户余额）
  和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，

  如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
  如果余额不足，则回滚事务。
*/

type Account struct {
	Id      int     `json:"id" gorm:"primaryKey"`
	Balance float64 `json:"balance" gorm:"not null"`
}

type Transaction struct {
	Id            int     `json:"id" gorm:"primaryKey"`
	FromAccountId int     `json:"from_account_id" gorm:"not null"`
	ToAccountId   int     `json:"to_account_id" gorm:"not null"`
	Amount        float64 `json:"amount" gorm:"not null"`
}

func Test03_2() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Account{}, &Transaction{})

	// db.Create(&Account{
	//  Id:      1,
	//  Balance: 500.0,
	// })
	// db.Create(&Account{
	//  Id:      2,
	//  Balance: 0.0,
	// })

	accountA := 1
	accountB := 2
	amount := 100.0

	err = TransferMoney(db, accountA, accountB, amount)
	if err != nil {
		panic(err)
	}

}

func TransferMoney(db *gorm.DB, fromId int, toId int, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var from Account
		// 1、查询账户余额
		if err := tx.Where("id = ?", fromId).First(&from).Error; err != nil {
			return err
		}
		// 2、检查账户余额是否足够
		if from.Balance < amount {
			return errors.New("from account balance is not enough")
		}
		// 3. 扣除账户 A 的余额
		if err := tx.Model(&Account{}).
			Where("id = ?", fromId).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return fmt.Errorf("扣除账户 A 余额失败: %w", err)
		}
		// 4. 增加账户 B 的余额
		if err := tx.Model(&Account{}).
			Where("id = ?", toId).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return fmt.Errorf("增加账户 B 余额失败: %w", err)
		}
		// 5. 记录转账信息
		trans := Transaction{
			FromAccountId: fromId,
			ToAccountId:   toId,
			Amount:        amount,
		}
		if err := tx.Create(&trans).Error; err != nil {
			return fmt.Errorf("记录转账信息失败: %w", err)
		}

		// 返回 nil 表示成功，会自动提交事务
		return nil
	})
}
