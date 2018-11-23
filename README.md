# filecollector
### 1. 说明
```
1. 按日期整理文件的命令行工具  
2. 开发语言：
        Golang
3. 目前支持如下平台：  
        - Mac  
        - Linux (测试环境：Ubuntu)
        // Windows目前未做适配
```

> 特殊说明
```
1. 本程序调用的 rename() 实现文件移动（不是复制）
2. 本程序将按照文件的"创建时间"进行按年、按月归类
    2.1. 由于 linux 平台不存在创建时间，取文件的 Mtim(编辑时间) 进行文件归类
    2.2. Mac 平台适用文件创建时间 Birthtimespec 进行文件归类
```


### 2. 适用场景
> 整理个人相册（家庭照片）  
    
### 3. 参数说明
- 有参启动
```
./filecollector -h
```
```
    -bs int
        批次大小(一个批次处理的文件数量) (default 10)
    -s string
        源文件夹
    -t string
        目标文件夹
```

- 无参启动
```
./filecollector
```

### 4. 系统命令行界面
```
--------------------------------------------------------------
请选择操作:
        0: 批次大小 10
        1: 设置源目录 /home/kouki/a
        2: 设置目标目录 /home/kouki/b
        3: 执行整理程序
        9: 退出系统
--------------------------------------------------------------
>
```
> 设置源目录、目标目录后，选择 [ 3 ] 执行文件整理

### 5. 执行结果示例
> 
> ```
> /home/kouki/b/
>     2018/
>         2018-10/
>             a.jpg
>             b.jpg
>         2018-12/
>             c.jpg
>     2019/
>         2019-01/
>             e.jpg
>             f.jpg
> ```
