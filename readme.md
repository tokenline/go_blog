# Go_blog开发日志

### 3月13日 Delete接口的开发

A.实现Delete功能应该仅需要一个ID就行，是否要为了这一个ID写一个Dto?

还是说这个ID通过其他途径获取？

软删除的本质是一个update?

本次先采用软删除的方法。

按照上节课的逻辑，实现删除这个功能应该按照如下顺序：

1.首先获取user的列表(list)

2.update对应ID的user的deleted_at字段的内容

3.再次获取user列表，查询对应ID的user是否在其中

4.如果是则软删除成功，如果否责软删除失败

B.Update功能存在一个缺陷，比如用户仅需更新自己的邮箱或者电话，不更新其他选项，
如果更新项不填写，则原有信息会被覆写成nil.

C.关于list接口的问题

list这个接口的输入和输出应该是什么？这个接口是应该获取全部满足指定条件user
还事仅仅是返回user表全体内容？返回全体内容不实际。

那么根据A中描述的问题，应该是软删除指定ID的user而不用list，所以A中提到的功能
逻辑可能存在问题，那么下列逻辑是否可行：

1.update对应ID的user的deleted_at字段的内容

2.应该不用再次查询user表中对应IDuser deleted_at字段的内容。

D.其余问题

1.detail接口需要返回怎样的detail?

2.query接口需要返回怎样的query？根据ID返回对应的user在user表中的全部信息？


### 3月15日 delete接口实现 list接口实现

如果gorm模型包含一个gorm.deleted_at字段，调用Delete方法会自动使用软删除功能
，不用造轮子。

delete接口实现，测试通过。

list接口大体上实现了不过存在两个问题：

1.list接口应该传入什么参数？

2.list的返回应该是什么格式？

### 3月16日 update接口改进

原本update接口更新时会将没有填写的字段更新为空，今天改进这一缺点。

为了得到原本的数据，需要实现一个单一查询的功能。

修改原本query案例，用来实现这个功能。

重写了update接口的一部分，但是遇到一个问题，对于类似int这种整型变量，如何去表示一个nil呢？在输入的时候总会有没有输入值，而导致一个空的。

gorm.Find()这个方法和Java中的方法不一样，它不用穿出它的处理结果，这大概是用指针的好处。


### 3月18日 课程记录

1、使用公共IdInfoDto

2、已有常量的字段，要使用常量来查询

3、Update

(1)Update返回error

(2)update优化

4、包名constant改为constants

5、List

(1)要有放回数据

(2)定义返回结构体vo

6、Query

Gorm数据库字段相关使用小写

7、Detail也是返回Vo

8、方法顺序

TODO：待实现的功能

### 3月18日 分页查询的实现和SWAG

A.分页查询

通过gorm中的Limit()和Offset()方法来实现分页查询的功能

eg: db.Limit(10).offset(5) 

执行上述语句的结果为获取跳过前五条数据，获得从第六条开始的十条数据。

实现分页查询有几个需要传入的参数：

1.PageNum 页码

2.Offset 需要跳过的数据条数

3.PageSize 页面大小 页面大小是固定的

Offset = (PageNum - 1) * PageSize 

偏移量 = (当前页码 - 1) * 页面大小

db.Limit(10).Offset(Offset)

Limit() 和 Offset()方法的不同调用顺序似乎不会影响结果。

B.SWAG

SWAG安装 初始化失败

### 3月21日

写完文章和文章目录的路由

创建对应API和Service的空函数

