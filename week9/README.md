# 作业要求
1.总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用  
2.实现一个从 socket connection 中解码出 goim 协议的解码器。
---
# 1.粘包的解包方式总结
## 1.1 fix length
规定发送方和接收方使用固定大小的缓冲区，当字符长度不够时使用空字符填充。

优点：可以很简单解决粘包和半包的问题。
缺点：当发送的数据比较小时，使用固定缓冲区大小的方式会增加不必要的数据传输，因此会增加网络传输的负担。
应用：netty FixedLengthFrameDecoder

## 1.2 delimiter based
使用特定字符结尾，例如可以通过"\n"结尾代表流的边界，读取时候可以按行读取。

应用：netty DelimiterBasedFrameDecoder

## 1.3 length field based frame decoder
又称为自定义长度解码器，通过自定义长度（长度域）来解决粘包问题。
数据包长度 = 长度域的值 + lengthFieldOffset + lengthFieldLength + lengthAdjustment

应用：netty LengthFieldBasedFrameDecoder

