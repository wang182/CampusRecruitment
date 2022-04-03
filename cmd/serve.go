package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"CampusRecruitment/pkg/config"
	"CampusRecruitment/pkg/infra/db"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/web"
)

var (
	cfgFile string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run web service",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Load(cfgFile); err != nil {
			return err
		}
		mysql := config.Get().Mysql
		dbDSN := fmt.Sprintf(
			"mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.DB,
		)
		if err := db.Init(dbDSN); err != nil {
			return err
		}
		if err := models.Init(db.Get()); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return web.Start()
	},
}

func init() {
	serveCmd.Flags().StringVarP(&cfgFile, "config", "c", "./config.yaml", "specify config file")
}
