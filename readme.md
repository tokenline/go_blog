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



