package controller

import (
	cli_ui "app/cli-ui"
	"app/pkg"
	"github.com/sirupsen/logrus"
)

func ProcessChoice(tourController *TourController,
	saleController *SaleController,
	accController *AccountController,
	reqController *RequestController,
	clntController *ClientController,
	mngrController *ManagerController,
	lg *logrus.Logger) (int, error) {
	var err error
	var choice string
	fl := true

	for fl {
		choice = cli_ui.GetSectionChoice()
		switch choice {
		case "1":
			for choice != "0" {
				choice = cli_ui.GetTourChoice()
				err = tourController.ProcessTours(choice, lg)
				if err != nil {
					pkg.ProcessErrors(err)
				}
			}
		case "2":
			for choice != "0" {
				choice = cli_ui.GetSaleChoice()
				err = saleController.ProcessSales(choice, lg)
				if err != nil {
					pkg.ProcessErrors(err)
				}
			}
		case "6":
			choice = cli_ui.GetAccountChoice()
			err = accController.ProcessAccounts(choice, lg)
			if err != nil {
				pkg.ProcessErrors(err)
			}
		case "3":
			for choice != "0" {
				choice = cli_ui.GetRequestChoice()
				err = reqController.ProcessRequests(choice, lg)
				if err != nil {
					pkg.ProcessErrors(err)
				}
			}
		case "4":
			for choice != "0" {
				choice = cli_ui.GetManagerChoice()
				err = mngrController.ProcessManagers(choice, lg)
				if err != nil {
					pkg.ProcessErrors(err)
				}
			}
		case "5":
			for choice != "0" {
				choice = cli_ui.GetClientChoice()
				err = clntController.ProcessClients(choice, lg)
				if err != nil {
					pkg.ProcessErrors(err)
				}
			}
		case "0":
			fl = false
		}
	}

	return 0, err
}
