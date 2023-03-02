package env

import "fmt"

func init() {
	if err := initEnvironmentConfig(); err != nil {
		panic(
			fmt.Sprintf("error: config.initEnvironmentConfig failed: %v", err),
		)
	}
	// mock環境の場合mysqlの環境変数設定を行わない
	if CONFIG.ENVIRONMENT == "mock" {
		return
	}
	if err := initMYSQLConfig(); err != nil {
		panic(
			fmt.Sprintf("error: config.initMYSQLConfig failed: %v", err),
		)
	}
}
