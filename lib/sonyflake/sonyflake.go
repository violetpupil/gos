// sonyflake 分布式唯一id生成器
//
// A Sonyflake ID is composed of
// 39 bits for time in units of 10 msec
// 8 bits for a sequence number
// 16 bits for a machine id
//
// As a result, Sonyflake has the following advantages and disadvantages:
// The lifetime (174 years) is longer than that of Snowflake (69 years)
// It can work in more distributed machines (2^16) than Snowflake (2^10)
// It can generate 2^8 IDs per 10 msec at most in a single machine/thread (slower than Snowflake)
package sonyflake
