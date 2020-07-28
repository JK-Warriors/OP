package common

import (
)


type Dr struct {
	Id          int    `xorm:"int 'id'"`
	Bs_Name     string `xorm:"varchar(200) 'bs_name'"`
	Db_Id_P     int    `xorm:"int 'db_id_p'"`
	Db_Type_P   int    `xorm:"int 'db_type_p'"`
	Host_P      string `xorm:"varchar(20) 'host_p'"`
	Port_P      int    `xorm:"int 'port_p'"`
	Alias_P     string `xorm:"varchar(200) 'alias_p'"`
	Inst_Name_P string `xorm:"varchar(50) 'inst_name_p'	"`
	Db_Name_P   string `xorm:"varchar(50) 'db_name_p'"`
	Db_Id_S     int    `xorm:"int 'db_id_s'"`
	Db_Type_S   int    `xorm:"int 'db_type_s'"`
	Host_S      string `xorm:"varchar(20) 'host_s'"`
	Port_S      int    `xorm:"int 'port_s'"`
	Alias_S     string `xorm:"varchar(200) 'alias_s'"`
	Inst_Name_S string `xorm:"varchar(50) 'inst_name_s'"`
	Db_Name_S   string `xorm:"varchar(50) 'db_name_s'"`
	Is_Shift    int    `xorm:"int 'is_shift'"`
}
