package controllers

import (

)

const COMMON_ERROR int = -1
const AUTH_ERROR int = -2
const LOGIN_SERVER_TEST = "http://10.12.129.43:9000"


const PARAM_ERROR_TEXT = "parm error!"

func Get_JsonOutput()(map[string]interface{}){
	var json_output = map[string]interface{}{
		"errorcode": 0,
		"msg" : "ok",
	}
	return json_output
}



