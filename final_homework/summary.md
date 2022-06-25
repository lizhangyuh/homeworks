## 概述

不知不觉13周的Go训练营就结营了，人生总是在感叹时间走的太快，但好在过去的这几个月时间让我收获了很多，也算没有白白浪费这段光阴。



我是一名非科班出身的程序猿，在参加Go训练营之前，我所有的编程知识都是通过自学的方式来获得的，所以我一直不太接受在线编程课程这类东西，总觉得文档和各种博文网上都有，为什么还要花钱去参加编程课程呢？有一天我无意间看见了极客时间Go训练营的广告，因为自己对Go语言比较感兴趣，也正在自学，所以就留意了一下极客时间Go训练营的介绍，要来了课程大纲之后，我来了兴趣。和我之前想的不同，训练营的毛剑老师是B站的技术大佬，在Go语言方面经验丰富，而且课程大纲中的知识要点全都是我感兴趣和有计划学习的内容，还有很多实战经验的部分，感觉干货满满，于是抱着试一试的态度报了名，开始了第一周的试课。



## 第一周

第一周结束后，我就决定留下来了。第一周微服务的讲解，其实解答了我很多的困惑。因为我目前工作的公司项目规模不大，也是一个有一定年龄的老项目，所以日常工作中接触的都是单体应用架构，对微服务只是大概懂个概念，还并不能理解到其背后的工程复杂性。在这一周的课程里，比较全面的认识了微服务的概念，优缺点，设计思路，也了解API Gateway，BFF这些之前都不知道的概念。在实际应用方面，认识了gRPC，明白了服务发现的实现原理，毛大还讲了通过流量染色实现多租户这类的“骚操作”~我算是真正打开了微服务世界的大门。



## 第二周

异常处理是Go语言的陈年老梗，使用Error而不是Exception来处理错误确实逻辑更简单直白，也强制要求你在编码的时候就考虑错误的处理和流程的控制，但是我确实还是不太习惯这种工程实践，之后在实践中应该会再慢慢总结和抽象自己的Error使用规则。毛大在这一周也详细介绍了Error的实现原理以及业内的一些处理Error的标准范式，我觉得最有用的一句就是：error只应处理一次，要么log返回nil，要么return err抛给上一层。



## 第三周

并发编程是Go的一大亮点，但是对我这个Ruby程序员来说，也是一大难点。Go语言一开始就被设计成可以很直观友好的进行并发编程，不过要想用好这个特性，必须先理解Goroutine，channel，context这些概念，刚开始接触起来确实还是有些抽象。当然，对于并发编程，Go还有很多工具包也是必须吃透的，比如sync，errgroup，Pool...现在回顾起来，这一周的课程还是值得去反复温习吃透的。



## 第四周

这一周主要说的是工程化实践了，我总觉得毛大是在夹带私货，哈哈哈，基本就是在介绍Kratos的实现。参加训练营之前就看过Kratos的文档，说老实话没怎么看懂...（不知道是我技术太菜，还是文档不够友好，可能都有？但主要是前者吧）但是这周学习之后，虽然不能说我完全懂了Kratos，但是基本概念还是都摸清了，也能上手用kratos搭一个项目了。很喜欢Kratos用protobuf来定义数据的理念，不管是api，还是error，甚至是config都统一规范的定义好了。但是，Kratos V2的结构我还是不能完全理解它的设计用意，大概是我不太懂DDD的缘故吧，后续DDD这块还需要再补起来~



## 第五周

这周主要介绍微服务的可用性设计，干货满满，毛大介绍了很多B站业务中遇到的问题和解决方案。微服务架构确实比我之前想象的要复杂的多，项目模块拆分过后，所有流量不再是在一个程序内部流转了，无形中增加了很多的不确定因素，但是也正是因为这种拆分，引入了隔离，限流，熔断，重试，负载均衡这些概念，让原来的单体应用似乎变成了一个服务“种群”，通过制定合理的内部规则，让整个项目能拥有旺盛的生命力，甚至还能拥有自我修复这样的超能力，这就是微服务架构的魅力所在吧。



## 第六周&第七周

这两周就主要是讲的两个实际业务的设计，讲了几周的概念终于落地了一把。



评论架构的设计案例完全就是为了高并发大数据量来做的，针对读和写都做了异步处理，写数据库和从数据库回源的操作全部发到了Kafka，让下游以最大能力消费，也保证消息最终会被存储起来。



历史播放记录架构的设计案例就有意思一些了，类似的功能每天都会用到，但之前其实没有怎么仔细想过实现原理的（看来还是对技术细节不够敏锐啊）。历史记录从理论上来说其实是一个高频数据，播放的每一秒钟都是历史记录，但是last-write-win，只有最后一次数据进行持久化即可。这里有一个很有意思的设计，用户的历史记录都会先进入缓存，读写都直接在缓存，因为高频访问的历史数据都是近期的，所以高频数据的有效性就不用担心。另外，会把uid聚合发送给kafka，kafka通过uid让job来批量从缓存中读取最后一条有效数据写入持久化存储（HBASE）。这样一来，近期热点数据都在缓存里了，而远期的冷数据也可以稳定的被持久化，相当有意思的方案啊。



## 第八周

分布式缓存和分布式事务，面试高频考点。讲老实话，这章我有点划水，课程里面通俗易懂的讲了分布式事务的一些实现方式，也用支付的例子做了解释，当时是看懂了，现在回想起来还是有些模糊。而分布式缓存这一块，我发现我的笔记竟然是缺失的，看来这一周的课程是需要好好再复习一次的了。



## 第九周

这一周毛大又在“夹带私货”，主要讲了GoIM的实现，跟着学完课程也算是跟着实现了一个IM系统后台API的设计。IM系统的消息同步模型也很有意思，读扩散写扩散，单信箱多信箱，推拉模式，以及这些概念与私信群聊甚至公众号的关系。我也终于明白了，为啥每次打开微信都要读取几秒的消息。



## 第十周&第十一周

这两章主要是介绍微服务架构的基础设施的建设。日志，监控指标，链路追踪，DNS，CDN以及多活架构。这些内容可能需要在之后的工作中花一定时间去实践了才能有很深入的理解。但不管怎么说，这两周的学习还是帮我拓宽了很多知识面，尤其是不同平台的多活架构的设计，之前听到多活架构四个字脑海里只有多集群三个字，但其实集群怎么部署，跟业务逻辑有很大关系，也是很灵活多变的，门道还是很深的。



## 第十二周

深入浅出的讲解了热门组件Kafka。这个组件在各种技术博文中的出现频率奇高，这次算是跟着毛大把Kafka的底层原理捋了一遍，了解这些底层实现原理相当重要，可以毫不夸张的说，明白了底层原理，你也就知道了怎么用好Kafka，但是不明白底层原理，你看再多的Kafka入门教程，组件本身用的再溜，真的要解决一些问题或者做优化的时候也只有两眼抓瞎。



Kafka的很多细节这里就不一一回顾了，总结一点自己的感受就是，消息队列中，Kafka是一个较全面的成熟解决方案，不管是性能，持久化，扩展性，高可用性，各个方面Kafka都有很细致的考虑。最后再说一个自己以前的误区：原来磁盘的读取速度并不慢，主要还是看读取数据的方式。



## 第十三周

最后一周是说的Go Runtime的内容，用大明老师的话来说就是面试必考，但是实际工作中基本遇不到的问题。这一周的内容其实还是需要有比较扎实的计算机原理和操作系统知识才能理解的。不管怎么说，这些设计机制和原理还是很值得我们去学习的，很多思想其实在其他架构的设计中也是有借鉴意义的，道理都是一样的。



好吧，流水账式的回顾了整个训练营的学习，其实也没写个啥东西出来，更多的是自己的一些感性认识。很不凑巧，最近这几周的确实工作生活上事情太多，毕业项目本来打算好好弄一弄，现在可能只来得及实现一个简版了。已经决定下个月要开始的项目用Kratos了，温故而知新，学过的知识一定要用起来才能真正转化成自己的东西啊。



加油吧，GO！



2022.6.24 00:35