package databaseConn

import(
	"testing"
	"log"
)

func TestConnect(t *testing.T){
	db, err := GetConn()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMaxConnection(t *testing.T){
	count := 0
	for i:=0; i<10000; i++ {
		db, err := GetConn()
		if err != nil {
			log.Fatal(err)
		}
		err = db.Ping()
		if err != nil {
			log.Printf("ping fail for:", err.Error())
		}
		count++
	}
	log.Printf("max conn is %d\n", count)
//	t.Errorf("stop")
//	CloseDB()

}




