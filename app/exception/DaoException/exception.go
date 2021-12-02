package DaoException

import "higo-framework/app/exception/BusinessException"

func Throw(message string, code int) {
	BusinessException.Throw(message, code)
}
