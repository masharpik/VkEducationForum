package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	forum "github.com/masharpik/ForumVKEducation/forum/pkg"
	mainLiterals "github.com/masharpik/ForumVKEducation/utils/literals"
	"github.com/masharpik/ForumVKEducation/utils/logger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogOperationFatal(errors.New(mainLiterals.LogEnvFileNotFound))
	}

	debugOnStr := os.Getenv("DEBUG_ON")
	if debugOnStr == "" {
		logger.LogOperationFatal(errors.New(fmt.Sprintf(mainLiterals.LogEnvVarIsNil, "DEBUG_ON")))
	}

	logger.DebugOn, err = strconv.ParseBool(os.Getenv("DEBUG_ON"))
	if err != nil {
		logger.LogOperationFatal(err)
	}

	if logger.DebugOn {
		logger.InitLogger()
	}
}

func main() {
	r, err := forum.RegisterUrls()
	if err != nil {
		logger.LogOperationFatal(err)
	}

	err = forum.StartServer(r)
	if err != nil {
		logger.LogOperationFatal(err)
	}
}
