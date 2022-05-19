Gin  +  Gorm

备忘录

项目结构

![image-20220302230522758](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220302230522758.png)

+ api

  接口函数分为用户相关（注册和登录）和任务相关（CRUD）

+ conf

  服务器和Mysql的相关配置

+ docs

  使用swagger生成的api接口文档，可以用来调试，需要在main函数和各个api函数按照响应的标准写注释

+ middleware

  中间件，进行CRUD时，都需要验证用户身份，避免重复操作，将其放在路由组中

+ model

  init.go 数据库连接函数

  migrate.go 将定义的字段进行迁移至数据库,以及添加外键，将user和task连接起来

  tasks.go 定义一个任务所需要的字段

  user.go 定义一个用户所需的字段

+ routes

  设置路由，get,post,delete,put请求

+ serializer(序列化)

  common.go 定义服务器响应的内容

  user.go task.go 分别是对用户和任务进行序列化

+ service

  task.go 定义各个功能需要传给服务器的字段，以及返回服务器的响应

  user.go 登录和注册函数，登录时需要验证token

+ utils

  jwt.go 

  ![image-20220303161852139](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303161852139.png)

  定义一个密钥，用于签发token

  同时也包括token验证函数

+ go mod 表明项目所引用的包

1. 先写配置文件（.ini）提高通用性，更改时只需要更改配置文件不需要重新更改代码

2. 添加数据库Mysql相关内容，切记要导入数据库驱动

   ForeignKey 外键

   index 普通索引

   longtext 一种比较长的字符串类型

   on delete和on update , 可设参数cascade(跟随外键改动), restrict(限制外表中的外键改动),set Null(设空值）,set Default（设默认值）,[默认]no action

   配置文件

   ![image-20220302230210788](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220302230210788.png)

   ![image-20220302230218263](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220302230218263.png)

   先导入读取配置文件的包

   ```go
   "gopkg.in/ini.v1"  //需要事先安装依赖 go get gopkg.in/ini.v1
   ```

   ![image-20220302225348633](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220302225348633.png)

   将字符串切片拼接在一起，按照配置文件的内容，拼接成连接数据库时所需要的语句

   ```go
   path := strings.Join([]string{DbUser,":",DbPassWord,"@tcp(",DbHost,":",DbPort,")/",DbName,"?charset=utf8mb4&parseTime=True"},"")
   //"user:password@tcp(127.0.0.1:3306)/sql_testcharset=utf8mb4&parseTime=True"
   ```

   通过配置文件读取出服务器和Mysql的相应配置

   ![image-20220302225931206](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220302225931206.png)

3. 连接数据库

   连接数据库并设置一些限制条件

   需要导入Gin + Gorm 框架

   **切记一定要导入数据库驱动**

   ![image-20220302230154964](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220302230154964.png)

   4. 定义user和task结构体

      ![image-20220303162529410](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303162529410.png)

      ![image-20220303162540290](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303162540290.png)

   5. 编写函数对用户密码进行加密，以及登录时的密码验证

      ![image-20220303162713464](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303162713464.png)

   6. 将定义好的结构体各个字段映射到数据库中进行迁移

   <img src="https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303162359644.png" alt="image-20220303162359644"  />

   7. 在serializer写服务器响应的结构体内容

      ![image-20220303163039393](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303163039393.png)

   8. 在service写用户的登录和注册函数，编写生成token和验证token函数![image-20220303163829872](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303163829872.png)

      ![image-20220303163837421](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303163837421.png)

   9. 在路由中添加注册登录路由

   10. 编写备忘录增删改查相应的函数

   11. 写中间件，用户进行操作时需要验证身份，在路由组中加入中间件函数

       ![image-20220303164135901](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303164135901.png)

   12. 生成接口文档

       ![image-20220303164215414](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303164215414.png)

       一定要导入swagger包

   使用Runapi进行测试

   ![image-20220303164321918](https://raw.githubusercontent.com/Aurora-Galaxy/image/master/img/image-20220303164321918.png)
