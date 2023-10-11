## 第二次作业
* 实现edit接口的错误返回
![img.png](img.png)
* 问题：profile业务上是展示接口，我不认为profile接口应该返回错误信息，因为不能用错的信息修改数据库
  * 因此只加了profile的jwt有效时间延长
![img_1.png](img_1.png)