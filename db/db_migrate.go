package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/we-task/Todo-as-a-Service/x/util"
)

const (
	verPerm = 0666
)

var (
	reVer = regexp.MustCompile(`^\d+`)
)

// Using Database Migrations with Golang
// https://blog.stackademic.com/using-database-migrations-with-golang-7f6736f580c8
// How to Perform Database Migrations using Go Migrate
// https://www.freecodecamp.org/news/database-migration-golang-migrate/

func dbMigrate(db *sql.DB, dir string) error {
	if dir == "" {
		return errors.New("migrationDir目录必须指定")
	}
	lockFile := filepath.Join(dir, ".lock")
	verFile := filepath.Join(dir, ".version")

	// 锁定文件
	flock, err := util.FlockCreate(lockFile)
	if err != nil {
		return err
	}
	defer flock.Release()
	if err := flock.Lock(); err != nil {
		return err
	}
	defer flock.Unlock()

	// 执行迁移
	currVer, err := getCurrVer(verFile)
	if err != nil {
		return err
	}
	for _, name := range getSqlFiles(dir) {
		migVer := getMigVer(name)
		if migVer <= currVer {
			// log.Println("[db] 跳过", name)
			continue
		}
		log.Println("[db] 迁移中...", name)
		bytes, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return err
		}
		_, err = db.Exec(string(bytes))
		if err != nil {
			return err
		}
		currVer = migVer
		err = os.WriteFile(verFile, []byte(strconv.Itoa(currVer)), verPerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func getSqlFiles(dir string) (res []string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".sql") && reVer.MatchString(name) {
			res = append(res, name)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		v1, v2 := getMigVer(res[i]), getMigVer(res[j])
		return v1 < v2
	})
	return
}

func getCurrVer(file string) (int, error) {
	bytes, _ := os.ReadFile(file)
	str := string(bytes)
	str = strings.Trim(str, "\n")
	if str == "" { // 文件为空或不存在
		return -1, nil
	}
	return strconv.Atoi(str)
}

func getMigVer(name string) int {
	s := reVer.FindString(name)
	v, _ := strconv.Atoi(s)
	return v
}
