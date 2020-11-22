# cloudgo

## 这是什么？

这是中山大学服务计算课(2020)的[ 开发 web 服务程序 ](https://pmlpml.gitee.io/service-computing/post/ex-cloudgo-start/)作业的我的仓库。

cloudgo是一个 go web 服务程序，提供：
- 静态文件服务
- 简单 js 访问
- 表单提取服务

**注意：** 如果你是TA，请前往 [报告文档](report.md) 查看curl、apache的测试。  
**注意：** 这是一个初学者项目，不值得开发者参考。

## 部署和使用

**部署**

```
git clone http://https://github.com/chenguofan1999/cloudgo.git
```

进入 `test` 文件夹，

```
go run main.go
```

默认的端口是5990，您可以在 `main.go` 中更改。在浏览器中输入地址：`localhost:5990`, 浏览器显示以下页面：

![](pics/mainPage.png)

终端显示：

```
$ go run main.go
[negroni] listening on :5990
[negroni] 2020-11-22T11:48:19+08:00 | 200 |      2.9369ms | localhost:5990 | GET /
```

## 实现细节

### 静态文件服务

静态文件服务的 url 是 `localhost:5990/static/` , 主页面提供了跳转按钮。

这里的静态文件服务通过 `html.FileServer`建立在 `assets/testStaticFiles` 上。

![](pics/static1.png)

![](pics/static2.png)

![](pics/static3.png)


### 简单的 js 访问

此部分的 url 是 `localhost:5990/assets/testJS/` , 主页面提供了跳转按钮。

事实上在 `"/"`上也通过 `html.FileServer` 建立了文件服务，对应的 Dir 就是`main.go` 所在的位置，因此 `localhost:5990/assets/testJS/` 将进入 `/assets/testJS/` 文件夹。这里放置了演示的 `index.html`，因此在通过 go 提供的静态文件服务访问该文件夹时会直接进入以下 web 页面：

![](pics/js.png)

原始的 `index.html` 文件中并无 ID 和 Content 冒号后的内容，

`assets/testJS/index.html`
```html
<!DOCTYPE html>
<html>

<head>
    ...
    <script src="js/hello.js"></script>
</head>

<body>
...
    <p class="greeting-id">ID :</p>
    <p class="greeting-content">Content :</p>
...
</body>

</html>
```

这部分内容实际上是由 **web server 控制访问 js 来添加到 html 文件**的。

- web server 中对应代码：

    `server.go`
    ```go
    router.HandleFunc("/js", jsonHandler)
    ```


    `Handlers.go`
    ```go
    func jsonHandler(w http.ResponseWriter, req *http.Request) {
        formatter := render.New(render.Options{
            IndentJSON: true,
        })
        formatter.JSON(w, http.StatusOK, struct {
            ID      string `json:"id"`
            Content string `json:"content"`
        }{ID: "18342008", Content: "Contents from the back-end"})
    }
    ```

    `jsonHandler` 将一个 json 结构传给了 .js 文件, 可通过 `localhost:5990/js` 或此页面中的按钮跳转查看该 json 结构：

    ![](pics/json.png)

- 对应的 js 代码:  
(由上面的`index.html` 中的 `<script src="js/hello.js"></script>` 引入)

    `assets/testJS/js/hello.js`

    ```js
    $(document).ready(function() {
        $.ajax({
            url: "/js"
        }).then(function(data) {
        $('.greeting-id').append(data.id);
        $('.greeting-content').append(data.content);
        });
    });
    ```

### 表单服务

此部分的 url 是 `localhost:5990/login` , 主页面提供了跳转按钮。

![](pics/login1.png)

随便填入一组数据：(chenguofan1, 18342008)

![](pics/login2.png)