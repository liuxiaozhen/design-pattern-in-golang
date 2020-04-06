## 观察者模式 ##
当对象间存在一对多关系时，则使用观察者模式（Observer Pattern）。比如，当一个对象被修改时，则会自动通知它的依赖对象。观察者模式属于行为型模式。
参与者有：
被观察者(Subject)：知道它的通知对象，事件发生后会通知所有它知道的对象，提供添加删除观察者的接口。
观察者(Observer)：提供通知后的更新事件
具体被观察者(ConcreteSubject)：被观察者具体的实例，存储观察者感兴趣的状态。
具体观察者(ConcreteObserver)：观察者的具体实现。