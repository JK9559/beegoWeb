### 项目执行步骤
1. 进入main.go编译程序
2. 进入router.go路由器，查找需要执行的路由以及对应的控制器
3. 找到对应的controller控制器和执行的方法函数
4. 进入controller使用的模型
5. 根据控制器需要执行的方法函数，执行渲染模板页面，
6. 模板用到的静态文件，执行完成的业务逻辑，执行成功后生成一个对应项目名的exe文件


关于构建一个beego 以及部分beego的简单原理
1. 在$GOPATH/src 执行 bee new 项目名会新建一个beego的初始项目
2. main函数是入口函数，go的执行过程为递归执行每个包的init函数
3. 在routers包内的init函数，包含了一个路由注册函数 beego.Router，其功能是映射URL到controller
4. beego.Run函数的内部所做的事情：
    1. 自动解析在conf下的配置文件 app.conf
    2. 执行用户的hookfunc，默认存在了注册mime，用户可通过函数AddAPPStartHook注册资金的启动函数
    3. 是否开启session 如果开启的话 就初始化全局的session
    4. 是否编译模板，beego会在启动时候根据配置把views目录下的所有模板预编译，存在map里
    5. 是否开启内置的文档路由功能
    6. 是否启动管理模块，应用内监控模块，会在8088端口做一个内部监听，通过此接口可以查询到QPS、CPU、内存、GC、goroutine、
    thread等信息
    7. 监听服务端口，调用ListenAndServe
5. controller的运行机制
   1. 首先声明了一个MainController控制器，其中内嵌了beego.Controller，也就是MainController中自动拥有了全部的beego.Controller的方法
   2. beego.Controller包含了很多方法，Init、Prepare、Post、Get、Delete、Head。可以通过重写的方式实现这些方法
   3. 其中的代码需要执行的逻辑，this.Data是一个用来存储数据的map，可以赋值任意类型的值
   4. 最后一个就是需要去渲染的模板，this.TplName 就是需要渲染的模板，这里指定了 index.tpl，如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找，例如上面的方法会去 maincontroller/get.tpl (文件、文件夹必须小写)。
   5. 用户设置了模板之后 系统自动调用Render函数，无需用户自己去渲染
   6. 也可以直接用 this.Ctx.WriteString 输出字符串
6. model逻辑
   1. Web应用中我们使用最多的就是数据库操作，model层一般来做这些操作
   2. 逻辑简单controller就可以替代model，如果需要重用，则需要使用model来抽象
7. view渲染
   1. https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.4.md
8. 