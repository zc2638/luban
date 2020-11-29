/**
 * Created by zc on 2020/11/28.
 */
package util

import "time"

func MilliTimestampToTime(milli int64) time.Time {
	return time.Unix(0, 0).Add(time.Millisecond * time.Duration(milli))
}

func TimeToMilliTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func SecTimestampToTime(sec int64) time.Time {
	return time.Unix(0, 0).Add(time.Second * time.Duration(sec))
}

func SecTimestampToDateTime(sec int64) string {
	return SecTimestampToTime(sec).Format("2006-01-02 15:04:05")
}