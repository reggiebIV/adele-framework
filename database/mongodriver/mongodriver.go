package mongodriver

import(
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/upper/db/v4/adapter/mongo"
)

// Build a data source name string, user can specify if they are on an atlat or self hosted instance
func BuildDSN(host, port, user, password, dbname string, atlastHost bool) string {
	if atlastHost {
		return fmt.Sprintf("mongodb+srv://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	} else {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	}
}

func Session(pool *mongo.DB) (db.Session, error) {
	session, err := mongo.New(pool)
	if err != nil {
		return nil, err
	}

	return session, nil
}
