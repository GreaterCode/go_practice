package main

import (
	"fmt"
	"log"
	"time"
)
import bolt  "go.etcd.io/bbolt"


func main() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}

 		if err = b.Put([]byte("answer"), []byte("42"));err != nil {
			 return err
		}

		if err = b.Put([]byte("zero"), []byte("0")); err != nil {
			return  nil
		}
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		//v := b.Get([]byte("noexist"))
		//fmt.Println(reflect.DeepEqual(v, nil))
		//fmt.Println(v == nil)
		//
		//v = b.Get([]byte("zero"))
		//fmt.Println(reflect.DeepEqual(v,nil))
		//fmt.Println(v == nil)

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key: %v, value: %v\n", string(k),string(v))
			return nil
		})
		return nil
	})
}
