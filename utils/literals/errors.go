package mainLiterals

const (
	LogEnvFileNotFound                          = "Не найден .env файл"
	LogRequestError                             = "Запрос выполнился неуспешно"
	LogOpenLogFileError                         = "Не удалось открыть файл или директорию для логирования, будет использоваться дефолтное os.Stderr"
	LogEnvVarIsNil                              = "Переменная \"%s\" окружения не обнаружена"
	LogConnDBFailed                             = "Не удалось подключиться к БД: %w"
	LogConnDBTimeout                            = "Время на попытку подключиться к БД вышло"
	LogStructNotSatisfyInterface                = "Реализация структуры не удовлетворяет интерфейсу: %s"
	LogErrorOccurredBeforeResponseWriterMethods = "Возникла ошибка до исполнения методов http.ResponseWriter: :%w"
	LogPanicOccured                             = "В процессе обработки запроса возникла паника: %v"
)
