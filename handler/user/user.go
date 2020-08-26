package user

/**
* @Author: super
* @Date: 2020-08-26 18:37
* @Description: 统一将Request和Response放在同一文件中，以便管理与维护
**/

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}