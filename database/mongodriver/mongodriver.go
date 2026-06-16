package mongodriver

import(
	"fmt"
	"net/url"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Build a data source name string, user can specify if they are on an atlat or self hosted instance
func BuildDSN(host, port, user, password, dbname string, atlastHost bool) string {
	escapedUser := url.QueryEscape(user)
	escapedPass := url.QueryEscape(password)
	if atlastHost {
		return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", escapedUser, escapedPass, host, dbname)
	} else {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", escapedUser, escapedPass, host, port, dbname)
	}
}

func Connect(uri string) (*mongo.Client, error) {
	session, err := options.Client().ApplyURI(uri).SetMaxPoolSize(25)
	if err != nil {
		return nil, err
	}

	return session, nil
}
