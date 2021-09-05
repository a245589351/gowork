# 本周作业
1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
---

# 1. 测试redis get set 性能
## 1.1 测试10字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 10 -t get,set
```
执行结果
> SET: 33311.12 requests per second  
GET: 33783.79 requests per second

## 1.2 测试20字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 20 -t get,set
```
执行结果
> SET: 33288.95 requests per second  
GET: 32970.66 requests per second

## 1.3 测试50字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 50 -t get,set
```
执行结果
> SET: 33726.81 requests per second  
GET: 34164.67 requests per second

## 1.4 测试100字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 100 -t get,set
```
执行结果
> SET: 33658.70 requests per second  
GET: 34002.04 requests per second

## 1.5 测试200字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 200 -t get,set
```
执行结果
> SET: 33079.72 requests per second  
GET: 33590.86 requests per second

## 1.6 测试1k字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 1024 -t get,set
```
执行结果
> SET: 33300.03 requests per second  
GET: 33658.70 requests per second

## 1.7 测试5k字节
执行命令如下
```
redis-benchmark -n 100000 -q -d 5120 -t get,set
```
执行结果
> SET: 32446.46 requests per second  
GET: 31466.33 requests per second

# 2. 分析key占用的内存空间
## 2.1 初始状态
used_memory_dataset:582480

## 2.2 插入30w条10字节数据
used_memory_dataset:7785280
每个key占用大小：(7785280-582480)/300000=24字节

## 2.3 插入30w条20字节数据
used_memory_dataset:17386712
每个key占用大小：(17386712-7785280)/300000=32字节

## 2.4 插入30w条50字节数据
used_memory_dataset:36588624
每个key占用大小：(36588624-17386712)/300000=64字节

## 2.5 插入30w条100字节数据
used_memory_dataset:72591376
每个key占用大小：(72591376-36588624)/300000=120字节

## 2.6 插入30w条200字节数据
used_memory_dataset:142197248
每个key占用大小：(142180032-72591376)/300000=232字节

## 2.7 插入1w条1024字节数据
used_memory_dataset:155060032
每个key占用大小：(155060032-142180032)/10000=1288字节

## 2.2 插入1w条5120字节数据
used_memory_dataset:17386712
每个key占用大小：(216586048-155060032)/10000=6152字节
