# go-sensitive
## 基于ac自动机实现的字符串匹配算法
初步实现简单的ac自动机，并基于此实现一个关键词匹配的微服务

##启动
cmd 目录下执行 go run main.go
##请求
 curl 127.0.0.1:8081/filter -d "content=地中海&group_id=1,2"


## 待完善
- [x] 单元测试 
- [x] 错误捕捉 
- [x] 返回内容标准化
- [x] 配置更新 AC自动机自动更新,通过etcd，监听更新
- [x] 敏感词从数据库读取
- [ ] 日志标准格式化
- [ ] 字典树优化
- [ ] Code Review，代码结构调整
- [x] 信号处理机制



## 数据库
```sql
CREATE TABLE `words` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `word` varchar(128) NOT NULL DEFAULT '' COMMENT '单词',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组id',
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4
```