## 单例模式 ##
单例模式是一种常用的软件设计模式。在它的核心结构中只包含一个被称为单例类的特殊类。
通过单例模式可以保证系统中一个类只有一个实例而且该实例易于外界访问，从而方便对实例个数的控制并节约系统资源。

### 示例 ## 
Go的sync包提供了并发控制，如互斥锁，原子操作。
示例提供了2种常见的单例模式的实现，2个都是线程安全的，方法1比较简单。
