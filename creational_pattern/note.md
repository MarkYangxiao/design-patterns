# 创建者设计模式

## 工厂方法模式
### 意图
父类提供一个创建对象的方法，允许子类绝决定实例化对象的类型
### 模式结构
1. 产品：对接口进行声明，对所有由创建者及其子类的构建对对象，这些接口都是通用对
2. 具体产品：不同产品不同实现具体产品接口
3. 创建者：类声明返回产品对象的 *工厂方法*，该接口返回的对象必须与产品接口相匹配，虽然是工厂类，但是不具体生产产品
4. 具体创建者：重写创建者的工厂方法，使其返回不同的类
### 适用场景
1. 工厂方法将产品创建和产品实际使用代码分离，从而能在不影响其他代码的情况下扩展产品创建部分代码
2. 希望用户能方便的扩展软件库或框架内部组件
3. 复用现有对象

## 抽象工作模式
一般是类产品，有多个不同的具体产品，抽象工厂是一系列这类产品工厂的抽象
### 模式结构
1. 抽象产品：为构成一组产品相关但不相同都接口声明
2. 具体产品：为抽象产品都多种不同具体实现
3. 抽象工厂：为创建不同抽象产品的方法声明
4. 具体工厂：实现抽象工厂的具体方法
5. 尽管具体工厂会对具体产品进行初始化，但是其构造方法必须返回各种抽象产品，这样客户端就不用与工厂创建特定对产品进行耦合
## 生成器模式/建造者模式
构造复杂对象，使用相同的创建代码生成不同类型的对象。一般用于要构造的对象比较复杂，属性比较多，同时属性不是单纯的赋值，还包含其他的逻辑
### 模式结构
1. 生成器：声明所有产品/属性的构造方法
2. 具体生成器：提供构造过程不同的实现
3. 产品：最终生成的产品对象
4. 主管：定义调用构造产品的步骤书讯






