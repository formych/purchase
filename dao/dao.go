package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // 为了调用sql.Register()
)

// DB 全局的DB
var DB *sql.DB

// 如果有框架的话，可以放在框架里面
func Init() {
	DB = NewDB()
}

// NewDB 建立一个全局的db连接
func NewDB() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./db/purchase.db")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// CloseDB 关闭数据库连接
func CloseDB(db *sql.DB) (err error) {
	err = db.Close()
	if err != nil {
		log.Printf("Close DB failed, err:[%v]\n", err)
	} else {
		log.Printf("Close DB success\n")
	}
	return
}

// PurchaseInfo 记录采购信息
type PurchaseInfo struct {
	ID           int
	User         string
	Company      string
	Tel          string
	PurchaseNum  int
	PurchaseTime string
	CreatedTime  string
	UpdatedTime  string
}
type PurchaseInfoDAO struct {
	tableName string
	cloumes   string
}

var PurchaseInfoDao = &PurchaseInfoDAO{
	tableName: "purchase_info",
	cloumes:   "id, user, company, tel, purchase_num, purchase_time, created_time, updated_time",
}

func (p *PurchaseInfoDAO) Add(r *PurchaseInfo) (err error) {
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES(null, '%s', '%s', '%s', %d, '%s', '%s', '%s')", p.tableName, p.cloumes, r.User, r.Company, r.Tel, r.PurchaseNum, r.PurchaseTime, r.CreatedTime, r.UpdatedTime)
	_, err = DB.Exec(sql)
	if err != nil {
		log.Printf("Add db failed, sql:[%s], err:[%v]\n", sql, err)
	}
	return
}

func (p *PurchaseInfoDAO) Get(start, limit uint64) (res []*PurchaseInfo, err error) {
	sql := fmt.Sprintf("SELECT * from %s LIMIT %d, %d", p.tableName, start, limit)
	rows, err := DB.Query(sql)
	if err != nil {
		log.Printf("Get db rows failed, sql:[%s], err:[%v]\n", sql, err)
		return
	}
	for rows.Next() {
		pInfo := &PurchaseInfo{}
		if err = rows.Scan(&pInfo.ID, &pInfo.User, &pInfo.Company, &pInfo.Tel, &pInfo.PurchaseNum, &pInfo.PurchaseTime, &pInfo.CreatedTime, &pInfo.UpdatedTime); err != nil {
			log.Printf("Scan rows failed, err:[%v]\n", err)
			return
		}
		res = append(res, pInfo)
	}
	rows.Close()
	return
}

func (p *PurchaseInfoDAO) GetAll(start, limit int64) (res []*PurchaseInfo, err error) {
	sql := fmt.Sprintf("SELECT * from %s", p.tableName)
	rows, err := DB.Query(sql)
	if err != nil {
		log.Printf("Get db rows failed, sql:[%s], err:[%v]\n", sql, err)
		return
	}
	for rows.Next() {
		pInfo := &PurchaseInfo{}
		if err = rows.Scan(&pInfo.ID, &pInfo.User, &pInfo.Company, &pInfo.Tel, &pInfo.PurchaseNum, &pInfo.PurchaseTime, &pInfo.CreatedTime, &pInfo.UpdatedTime); err != nil {
			log.Printf("Scan rows failed, err:[%v]\n", err)
			return
		}
		res = append(res, pInfo)
	}
	rows.Close()
	return
}

func (p *PurchaseInfoDAO) Count() (total int64, err error) {
	sql := fmt.Sprintf("SELECT COUNT(*) from %s", p.tableName)
	rows, err := DB.Query(sql)
	for rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			log.Printf("Scan rows failed, err:[%v]\n", err)
		}
	}
	if err != nil {
		log.Printf("Count rows failed, err:[%v]", err)
	}
	rows.Close()
	return
}
