package enerbit

type DatabaseServiceImpl struct {
}

func NewDatabaseServiceImpl() *DatabaseServiceImpl {
	return &DatabaseServiceImpl{}
}

func (e *DatabaseServiceImpl) GetAllAdapter() string {
	return "alexander"
}
