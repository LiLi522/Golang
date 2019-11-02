package utils
import (
	"fmt"
	"time"
)

type FamilyAccount struct {
	//声明一个变量，用来接收用户的选项
	option string
	//声明一个变量，用来判断是否退出软件
	flag bool
	//声明账户余额
	balance float64
	//声明收支的金额
	money float64
	//声明收支的说明
	note string
	//收支的详情 使用字符串来拼接
	details string
	//记录是否有收支行为
	hasMoney bool
	//时间
	createTime string
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
	option : "",
	flag : true,
	balance : 10000.0,
	money : 0.0,
	note : "",
	details : "收支\t账户余额\t收支金额\t说    明\t记账时间",
	hasMoney : false,
	createTime : "",
	}
}

//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("----------------当前收支明细----------------")
	if this.hasMoney {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前还没有收支明细。。。 来一笔吧")
	}
	
}

//将登记收入写成一个方法
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	now := time.Now()
	this.createTime = fmt.Sprintf(now.Format("2006/01/02 15:04:05"))
	//拼接details
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v\t\t%v", this.balance, this.money , this.note, this.createTime)
	this.hasMoney = true
}

//将登记支出写成一个方法
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	now1 := time.Now()
	this.createTime = fmt.Sprintf(now1.Format("2006/01/02 15:04:05"))
	//拼接details
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v\t\t%v", this.balance, this.money , this.note, this.createTime)
	this.hasMoney = true
}

func (this *FamilyAccount) exit () {
	fmt.Println("你确定要退出吗？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.flag = false
	}
}

//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------家庭收支记账软件--------------\n\n")
		fmt.Println("               1  收支明细\n")
		fmt.Println("               2  登记收入\n")
		fmt.Println("               3  登记支出\n")
		fmt.Println("               4  退出软件\n")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&this.option)
		fmt.Println()
		switch this.option {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.flag {
			fmt.Println("已退出家庭记账软件")
			break
		}

	}
}