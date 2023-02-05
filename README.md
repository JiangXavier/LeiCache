# LeiCache
LeiCache模仿了groupcache（Go语言版的memcached）的实现，支持特性有：
- 实现LRU缓存淘汰算法，解决了资源限制的问题
- 实现了单机缓存和基于HTTP的分布式缓存，给用户提供了自定义数据源的回调函数
- 使用Go锁机制实现了singleflight防止缓存击穿
- 使用一致性哈希选择节点，实现负载均衡，解决远程节点的挑选问题
- 使用protobuf库优化节点间二进制通信
