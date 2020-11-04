package yunke

import (
	`strconv`

	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// GetLastVersion 取得某个客户端类型的最新版本
func (a *Admin) GetLastVersion(clientType yunke.ClientType) (client *yunke.BaseClient, err error) {
	client = new(yunke.BaseClient)
	err = a.request(yunke.AdminApiGetFinalClientByTypeUrl, gox.HttpMethodGet, nil, map[string]string{
		"clientType": strconv.Itoa(int(clientType)),
	}, client)

	return
}