package mgutil

type MongoData interface {
	GetInsertData() any
	GetUpdateData() any
	GetUpsertData() any
}
