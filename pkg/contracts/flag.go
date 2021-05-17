package contracts

import "cowin/utils"

type InputFlags struct {
	Districts   string
	Pincodes    string
	Filters     utils.DuplicateStringFlag
	Interval    string
	Help        bool
	MinAlertVal int
}
