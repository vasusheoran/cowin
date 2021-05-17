package contracts

import "cowin/utils"

type InputFlags struct {
	Districts string
	Filters   utils.DuplicateStringFlag
	Interval  string
	Help      bool
}
