#   Wolfram Mathematica

##  1.高频常用命令       
    Shift+Enter                     #计算之前的输入
    Ctrl+L                          #重新前一次的输入
    Shift+Ctrl+L                    #重新前一次输出
    ?FuncXXX                        #查询FuncXXX函数说明
    \[Alpha                         #输入希腊字母α 
    \[CapitalDelta                  #输入大写的希腊字母Δ
    ESC a                           #输入小写希腊字母α //这是已别名的形式输入
    ESC D                           #以别名的形式输入大写Δ
    x Ctrl+^ y Ctrl+Space           #以x^y指数右上角指数形式输入
    Ctrl+@ x                        #以sqrt(x)形式输入
    x Ctrl+/ y Ctrl+Space           #以x/y分数形式输入
    x Ctrl+_ y                      #以x(y)右下角标形式输入   
    x=.                             #清除变量x的定义

##  2.高频常用常数                
    Pi                  #希腊字母pi 
    E                   #自然对数
    I                   #复数虚数单位
    Infinity            #无穷
    GoldenRadio         #黄金分割比(1+sqrt(5))/2
    EulerGamma          #欧拉gamma常数
    Catalan             #卡特兰常数 

##  3.高频常用函数    
    Clear["Global`*"]                   #清除之前所有变量
    Clear[x,y,...]                      #清除指定变量
    Factor                              #因式分解  
    Plot3D                              #画3维图像
    Sin                                 #正弦函数
    N                                   #提供指定精度输出
    Re[z]                               #复数z的实部
    Im[z]                               #复数z的虚部
    Conjugate[z]                        #共轭复数
    Arg[z]                              #幅角  
    Expand[fn]                          #将因子乘积fn展开为多项式
    Coefficient[fn, x^i]                #求fn多项式的x^i的系数  
    Product                             #连乘积形式
    D[fn, x]                            #求导 
    Integrate                           #求积分  
    Sum                                 #求和
    Solve[lhs==rhs, x]                  #求解方程
    Series[f, {x,x0,order}]             #求Taylor幂级数展开 
    Limit[f, x->x0]                     #求极限
    Minimize[f, x]                      #求最小值
    TrigExpand[Sin[a+b]]                #三角函数展开 
    TrigReduce[2Sin[x]Cos[x]]           #三角函数约化
    
##  4.核心语法     
    规则相关：     
        expr /.x->3                     #将expr中的x替换为3    
        expr //.rules                   #反复使用规则，直到结果不在变化     
        patt /;test                     #仅当test测试为true才匹配模式patt
            例子: sig[x_ /;x>0] := 1
         
    流程控制：    
        Do[expr, {imax}]                #对expr，计算imax次
        Do[expr, {i,imax}]              #对expr，变量i=1,2,3...imax计算
        While[test, body]
        For[start, test, incr, body]    #
        Table[expr, {}]                 #根据后面条件，通过expr计算出列表
        Nest[f, expr, n]                #嵌套n次   f[f[...f[expr]]]
        If[condi, t, f]                 #conid为ture给出t，否则给出f

    映射：
        Map[f, list]                    #将函数f作于于列表list的每个元素，生成新列表

    定义过程：
        myfunc[x_] := (expr1; expr2; ...; exprn) #最后结果是exprn   

    定义列表：向量   
        Range[n]                        #{1,2,...n}
        Range[imin, imax]               #{imin, imin+1, ...imax}
        Range[imin, imax, di]           #
        Table[expr, {i,n}]
        Table[expr, {i,m},{j,m}]        #生成mxn的矩阵
        Array[f, n]                     #{f[1], f[2], ...f[n]}
        Array[f, {m,n}]                 #{{f[1,1], ... f[1,n]}, ....} 构造mxn矩阵
        Part[v,i] 和 v[[i]]             #取列表v的第i个元素，从1开始数
        v[[i;;j]]                       #取子列表，从i取到j
        v[[-i;;]]                       #取末尾i个元素
        Partition[v, n]                 #分割列表v，为长度为n的多个子列表
        Join[v0,v1]                     #将列表v0,放在v1前，构成新列表
        c v                             #向量v的数量积
        v1.v2 和 Dot[v1, v2]            #点积
        Cross[a, b]                     #叉积
        v1 [ESC]cross[ESC] v2           #叉积
        m[[r1;;r2, c1;;c2]]             #取矩阵m的子矩阵从r1-r2行，c1-c2列
        DiagonalMatrix[v]               #给出对角矩阵，diag(M) = v
        MatrixForm[M]                   #以矩阵形式显示M
        Tr[M]                           #矩阵M的迹 sum(diag(M))
        Det[M]                          #矩阵M的行列式
        MatrixRank[M]                   #矩阵M的秩
        RowReduce[M]                    #给出矩阵行约减
        LinearSolve[M,b]                #求解MX=b
        Inverse[M]                      #矩阵的逆
        m1.m2 和 Dot[m1, m2]            #矩阵乘法 !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
        MatrixPower[M, n]               #矩阵的幂
        expr(#)&@v                      #相当于expr(v), '#'是参数占位符
        Riffle[list, ele]               #在list中交错放置ele
        Total[list]                     #列表求和
        Sum[expr, {i, imax}]            #表达式求和
        Accumulate[list]                #逐次累加生成新列表 {list[1], list[1]+list[2], ..., list[1]+list[2]+...+list[n]}

    定义置换：(群论)   










