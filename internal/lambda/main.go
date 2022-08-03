package lambda

import (
	"context"
	"fmt"
	"github.com/kvendingoldo/aws-cognito-backup/internal/config"
	"os"
)

func Execute(config config.Config) {
	client, err := cloud.New(context.TODO(), config.Region)
	if err != nil {
		log.Error(fmt.Sprintf("Could not create AWS client"), "error", err)
		os.Exit(1)
	}

	ctx := context.TODO()

	//

}
