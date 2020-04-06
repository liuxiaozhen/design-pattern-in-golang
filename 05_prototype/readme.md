## 原型模式 ##
原型模式（Prototype Pattern）用于创建重复的对象，同时又能保证性能。它属于创建型设计模式，它提供了一种创建对象的最佳方法。
这种模式是实现了一个原型接口，该接口用于创建当前对象的克隆。当直接创建对象的代价比较大时，则采用这种模式。

## 示例 ##
go中并没有从语言级别支持clone, 可以借助第三方库来实现深度复制
- https://github.com/jinzhu/copier
- https://github.com/huandu/go-clone