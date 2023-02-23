# supernet  

从零开始写的框架. 用时1个月. 实现了通讯这块,调通了游戏基本功能, 其中红黑,龙虎斗已经ok. 待完善.  因为疫情项目被放弃。

```shell
git submodule update --init  --recursive
git submodule init  # 初始化本地.gitmodules文件
git submodule update  # 同步远端submodule源码
git submodule add https://github.com/superufo/doc.git   doc

git submodule add https://github.com/superufo/hall.bojiu.com.git   src/hall.bojiu.com

git submodule add https://github.com/superufo/hhgame.bojiu.com.git   src/hhgame.bojiu.com

git submodule add https://github.com/superufo/center.bojiu.com.git   src/center.bojiu.com

git submodule  add https://github.com/superufo/common.bojiu.com.git   src/common.bojiu.com

git submodule  add https://github.com/superufo/gateway.bojiu.com.git   src/gateway.bojiu.com

git submodule add https://github.com/superufo/longhu.bojiu.com.git   src/longhu.bojiu.com

git submodule add https://github.com/superufo/yuxiaxie.bojiu.com.git   src/yuxiaxie.bojiu.com

git add .
git commit -m  "submodule"
git push
git submodule update
```

下载clone

```shell
git   clone  --recurse-submodules   https://github.com/superufo/supernet.git
```


#系统流程图
![supernet excalidraw](https://user-images.githubusercontent.com/20591332/220812915-464b280b-3cec-4c73-b032-d9c7b0420a77.png)



MIT License


