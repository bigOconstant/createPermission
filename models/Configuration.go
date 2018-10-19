package models

type Configuration struct {
	Name                    string
	Description             string
	Param                   string
	Value                   string
	ConfigurationDataTypeId int
	SecurityActivityId      int
	ValidValues             string
}

type ConfigurationDataType_Enum struct {
	ConfigurationDataTypeId int
	Name                    string
}
