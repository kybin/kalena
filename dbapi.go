package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AddSchedule 함수는 DB에 Schedule 을 추가한다.
func AddSchedule(session *mgo.Session, s Schedule) error {
	// DB에 추가되는 모든 시간은 UTC 시간을 가진다.
	err := s.ToUTC()
	if err != nil {
		return err
	}
	// 데이터를 넣기전 Excel시간도 셋팅한다.
	err = s.SetTimeNum()
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(s.Collection)
	err = c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}

// RmSchedule 함수는 DB에서 id가 일치하는 Schedule을 삭제한다.
func RmSchedule(session *mgo.Session, Collection string, id bson.ObjectId) error {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	err := c.RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}

// allSchedules는 DB에서 전체 스케쥴 정보를 가져오는 함수입니다.
func allSchedules(session *mgo.Session, Collection string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	var result []Schedule
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// SearchMonth 함수는 Collection, Year, Month을 입력받아 start의 값이 일치하면 반환한다.
func SearchMonth(session *mgo.Session, Collection, year, month string) ([]Schedule, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	var all []Schedule
	var results []Schedule
	y, err := strconv.Atoi(year)
	m, err := strconv.Atoi(month)
	start, err := BeginningOfMonth(y, m)
	end, err := EndOfMonth(y, m)
	s, err := TimeToNum(start.String())
	e, err := TimeToNum(end.String())
	query := []bson.M{}
	query.append(query, bson.M{"Startnum":{$gt: e}})
	query.append(query, bson.M{"Endnum":{$lt: s}})
	q := bson.M{"$nor": query}
	err = c.Find(q).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetLayers 함수는 DB Collection 에서 사용되는 모든 layer값을 반환한다.
func GetLayers(session *mgo.Session, Collection string) ([]string, error) {
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(Collection)
	var layers []string
	err := c.Find(bson.M{}).Distinct("layer", &layers)
	if err != nil {
		return nil, err
	}
	sort.Strings(layers)
	return layers, nil
}
