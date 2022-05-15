### 1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能


`redis-benchmark -t get,set -d 10`

测试结果： benchmark.csv


### 2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同value 大小下，平均每个 key 的占用内存空间


| value size | count | before memory size | after memory size | size per kv |
| ---------- | ----- | ------------------ | ----------------- | ----------- |
| 10B | 10k | 63264 | 446112 | 38.28 Byte |
| 20B | 10k | 46112 | 606112 | 56 Byte |
| 50B | 10k | 75280 | 937824 |  86.25 Byte |
| 100B | 10k | 46112 | 1406112 |  136 Byte |
| 200B | 10k | 46304 | 2366896 | 232.06 Byte |
| 1KB | 10k | 47104 | 10368864 | 1032.18 Byte |
| 5KB | 10k | 57456 | 51493349 | 5143.59 Byte |


| value size | count | before memory size | after memory size | size per kv |
| ---------- | ----- | ------------------ | ----------------- | ----------- |
| 10B | 50k | 75248 | 2042064 | 39.34 Byte |
| 20B | 50k | 42256 | 2843248 | 56.02 Byte |
| 50B | 50k | 43440 | 4444176 |  88.01 Byte |
| 100B | 50k | 44384 | 6845200 |  136.02 Byte |
| 200B | 50k | 45408 | 11646608 | 232.02 Byte |
| 1KB | 50k | 46816 | 51649808 | 1032.06 Byte |
| 5KB | 50k | 50016 | 257251808 | 5144.04 Byte |


