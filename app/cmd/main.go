package main

import (
	"app/controller"
	"app/pkg"
	"app/repo"
	"app/route"
	"app/services"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func initLogger(logFile string) *logrus.Logger {
	// Создаем экземпляр логгера
	lg := logrus.New()

	// Открываем файл для записи логов
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		lg.Fatalf("Ошибка при открытии файла для логов: %v", err)
	}

	// Устанавливаем вывод логов в файл
	lg.SetOutput(file)

	// Настройка форматирования логов
	lg.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return lg
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, err := pgx.Connect(ctx, "postgres://creator:password@127.0.0.1:5432/tour-db")
	defer pool.Close(context.Background())
	if err != nil {
		panic(err)
	}
	//
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "127.0.0.1:6379",
	//	Password: "pahst13_82", // no password set
	//	DB:       0,            // use default DB
	//})
	//
	//rr := repo.NewRedisAuthRepo(rdb, time.Second*60)
	//authS := services.CreateNewAuthService(rr, time.Second)
	//
	tr := repo.NewPostgresTourRepo(pool)
	//sr := repo.NewPostgresSaleRepo(pool)
	ts := services.NewTourService(tr, time.Second)
	//ss := services.NewSaleService(sr, time.Second)
	//
	//ar := repo.NewPostgresAccountRepo(pool)
	//as := services.CreateNewAccountService(ar, time.Second)
	//
	//pa := services.NewPayAdapter()
	//lg := initLogger("logs.txt")
	//
	//cr := repo.NewPostgresClientRepo(pool)
	//cs := services.CreateNewClientService(cr, time.Second)
	//cc := controller.ClientController{
	//	ClientService: cs,
	//	AuthService:   authS,
	//}
	//
	//mr := repo.NewPostgresManagerRepo(pool)
	//ms := services.CreateNewManagerService(mr, time.Second)
	//mc := controller.ManagerController{
	//	ManagerService: ms,
	//	AuthService:    authS,
	//}
	//
	//done := make(chan bool, 1)
	//reqR := repo.NewPostgresRequestRepo(pool)
	//rs := services.CreateNewRequestService(reqR, pa, time.Second, done, 10, time.Second*60, lg)
	//tc := controller.TourController{ts, rs, authS}
	//sc := controller.SaleController{ss, authS}
	//
	//reqCntr := controller.RequestController{
	//	RequestService: rs,
	//	AuthService:    authS,
	//}
	//ac := controller.AccountController{as, authS}
	//
	//for {
	//	code, _ := controller.ProcessChoice(&tc, &sc, &ac, &reqCntr, &cc, &mc, lg)
	//	if code == 0 {
	//		break
	//	}
	//}

	lg := initLogger("log.txt")
	pager, _ := pkg.CreateGinPager(ts, lg)

	tc := controller.TourController{ts, pager, lg}

	router := gin.Default()
	router.Static("/css", "./templates/css")
	router.Static("/static", "./templates/static")
	router.Static("/js", "./templates/js")
	router.LoadHTMLGlob("./templates/*.tmpl")
	route.MainRouter(router)
	route.TourRouter(router, &tc)

	router.Run()
}
