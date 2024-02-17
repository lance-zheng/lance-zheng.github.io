<!-- customize-category:Java -->

# JUC

- [JUC](#juc)
  - [AQS](#aqs)
  - [ReentrantLock](#reentrantlock)
    - [可重入](#可重入)
    - [公平锁](#公平锁)
    - [锁超时](#锁超时)
    - [锁中断](#锁中断)

## AQS

`java.util.concurrent.locks.AbstractQueuedSynchronizer`
提供一个用于实现阻塞锁和相关同步器的框架

## ReentrantLock

**特性：**

- 可重入
- 可中断 `ReentrantLock#lockInterruptibly`
- 支持锁超时 `ReentrantLock#tryLock(long, java.util.concurrent.TimeUnit)`
- 支持公平锁 `new ReentrantLock(true)`
- 多条件变量

### 可重入

可以重入是指同一个线程可以多次获取同一个锁，而不会被阻塞住

```java
public class Main {
    static ReentrantLock reentrantLock = new ReentrantLock();

    public static void main(String[] args) {
        foo(10);
    }

    public static void foo(int c) {
        if (c == 0) {
            return;
        }
        try {
            reentrantLock.lock();
            System.out.println("获取到了锁");
            foo(c - 1);
        } finally {
            reentrantLock.unlock();
        }

    }
}
```

**源码：**
通过 state 记录获取锁的状态，如果是当前线程尝试再次获取锁则会让 `state + 1`。

```java
final void lock() {
    if (!initialTryLock())
        acquire(1);
}

final boolean initialTryLock() {
    Thread current = Thread.currentThread();
    if (compareAndSetState(0, 1)) { // first attempt is unguarded
        setExclusiveOwnerThread(current);
        return true;
    } else if (getExclusiveOwnerThread() == current) {
        int c = getState() + 1;
        if (c < 0) // overflow
            throw new Error("Maximum lock count exceeded");
        setState(c);
        return true;
    } else
        return false;
}
```

### 公平锁

公平锁是指会按照获取的顺序分配锁资源
ReentrantLock 中同步器有两个实现，分别是`FairSync` `NonfairSync` 其中 `FairSync` 就是公平锁的实现。  
如果需要使用公平锁需要使用 `new ReentrantLock(true)` 创建对象

**源码：**
实现原理就是在尝试获取锁时会查看有没有前驱节点。若存在前驱节点则说明在这之前有线程尝试去获取锁了。

```java
protected final boolean tryAcquire(int acquires) {
    if (getState() == 0 && !hasQueuedPredecessors() &&
        compareAndSetState(0, acquires)) {
        setExclusiveOwnerThread(Thread.currentThread());
        return true;
    }
    return false;
}
```

### 锁超时

通过 `ReentrantLock#tryLock(long, java.util.concurrent.TimeUnit)` 可以尝试获取锁，如果超过时间则会 return false。

`AQS` 内部使用的是 LockSupport.park 实现阻塞等待的
如果使用的是 tryLock(long, java.util.concurrent.TimeUnit) 获取锁，则最后会使用 LockSupport.parkNanos(Object blocker, long nanos)，这个 API 可以指定阻塞时间。
**源码：**

```java
if (!timed)
    LockSupport.park(this);
else if ((nanos = time - System.nanoTime()) > 0L)
    LockSupport.parkNanos(this, nanos);
```

具体实现细节可以参考：  
AbstractQueuedSynchronizer#acquire(java.util.concurrent.locks.AbstractQueuedSynchronizer.Node, int, boolean, boolean, boolean, long)

### 锁中断

锁中断是指在等待锁的过程中可以被其他线程中断
需要使用 `lockInterruptibly()` 获取锁，其他线程可以用 `thread.interrupt()` 打断阻塞状态

```java
public final void acquireInterruptibly(int arg)
    throws InterruptedException {
    if (Thread.interrupted() ||
        (!tryAcquire(arg) && acquire(null, arg, false, true, false, 0L) < 0))
        throw new InterruptedException();
}
```
