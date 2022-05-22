### 1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用

fix length: 每次发送固定长度的数据。

delimiter based: 使用分隔符标记数据包的边界。举例：HTTP header和body中间的间隔就是一个空行。

length field based frame: 在头部中指定数据包的长度，解包的时候根据长度信息来确定数据包。举例：HTTP头部的content-length信息。

### 2.实现一个从 socket connection 中解码出 goim 协议的解码器
