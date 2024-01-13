package main

//NEED OPERATOR PRECEDENCE - this doesn't work: select(company.Employees, "emp", emp.Department.Name=="Sales" && emp.JobTitle == "Salesperson")
//...because it evaluates JobTitle=="Salesperson" as false & then tries to "and" it with "Sales".

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rbrumby/xex"
)

type Organization struct {
	Name      string
	Employees []*Employee
}

type Department struct {
	Organization *Organization
	Name         string
}

func (d *Department) Employees() (employees []*Employee) {
	for _, emp := range d.Organization.Employees {
		if emp.Department == d {
			employees = append(employees, emp)
		}
	}
	return
}

type Employee struct {
	FirstName  string
	LastName   string
	JobTitle   string
	IsActive   bool
	Department *Department
	Manager    *Employee
}

func main() {
	// l := capnslog.NewPackageLogger("github.com/rbrumby/xex", "xex")
	// logLvl := xex.DEBUG.String()
	// repolog := capnslog.MustRepoLogger("github.com/rbrumby/xex")
	// cfg, err := repolog.ParseLogLevelConfig(fmt.Sprintf("xex=%s", logLvl))
	// if err != nil {
	// 	panic(err)
	// }
	// repolog.SetLogLevel(cfg)
	// xex.SetLogger(l)

	company := &Organization{Name: "ACME, inc"}

	// none := &Department{company, "None"}
	purchasing := &Department{company, "Purchasing"}
	sales := &Department{company, "Sales"}
	accounts := &Department{company, "Accounts"}

	smith := &Employee{"David", "Smith", "CEO", true, nil, nil}
	jones := &Employee{"David", "Jones", "Buyer", true, purchasing, smith}
	booth := &Employee{"David", "Booth", "Buyer", false, purchasing, smith}
	robson := &Employee{"David", "Robson", "Sales Team Lead", true, sales, smith}
	johnson := &Employee{"David", "Johnson", "Salesperson", true, sales, robson}
	marshall := &Employee{"David", "Marshall", "CFO", true, accounts, smith}
	davies := &Employee{"David", "Davies", "Accounting Clerk", true, accounts, marshall}

	company.Employees = []*Employee{smith, jones, booth, robson, johnson, marshall, davies}

	fmt.Println("The company is accessible by top level variable \"company\"")
	for {
		scanr := bufio.NewScanner(os.Stdin)
		//scan.Split(bufio.ScanLines) //default
		fmt.Print("Enter expression: ")
		if scanr.Scan() {
			ex, err := xex.NewStr(scanr.Text())
			if err != nil {
				fmt.Printf("Expression syntax error: %s\n", err)
				continue
			}
			val, err := ex.Evaluate(xex.Values{"company": company})
			if err != nil {
				fmt.Printf("Expression evalutaion error: %s\n", err)
				continue
			}
			fmt.Println(val)
		}
	}
}
