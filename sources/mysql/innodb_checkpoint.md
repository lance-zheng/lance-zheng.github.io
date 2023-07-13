<!-- customize-category:MySQL-->

# InnoDB Checkpoint

- [InnoDB Checkpoint](#innodb-checkpoint)
  - [Sharp Checkpoint](#sharp-checkpoint)
  - [Fuzzy Checkpoint](#fuzzy-checkpoint)
    - [Master Thread Checkpoint](#master-thread-checkpoint)
    - [FLUSH_LRU_LIST Checkpoint](#flush_lru_list-checkpoint)
    - [Async/Sync Flush Checkpoint](#asyncsync-flush-checkpoint)
    - [Drity Page too much Checkpoint](#drity-page-too-much-checkpoint)

Checkpoint 所做的事情就是讲缓冲池中的脏页刷新回磁盘。
当数据库发生宕机时，数据库不需要重做所有的日志，因为 Checkpoint 之前的页都已经刷新到磁盘了。

Checkpoint 有以下两种：

- **Sharp Checkpoint**：刷新全部脏页数据
- **Fuzzy Checkpoint**：刷新部分脏页

## Sharp Checkpoint

Sharp Checkpoint 在 MySQL 实例关闭时执行，以保证所有脏页数据刷新到磁盘。

## Fuzzy Checkpoint

### Master Thread Checkpoint

Master Thread 会定期的进行 Checkpoint

### FLUSH_LRU_LIST Checkpoint

InnoDB 需要让 LRU List 中有一定的空闲页可以使用，通过 `innodb_lru_scan_depth` 可指定。若空闲页不足就会淘汰 LRU List 尾部的页，如果淘汰的是脏页则会进行 Checkpoint。

### Async/Sync Flush Checkpoint

当 **redo log** 出现不可用时会触发 **Async/Sync Flush Checkpoint**，Async/Sync Flush Checkpoint 是为了保证 redo log 的循环使用的可用性。

### Drity Page too much Checkpoint

当缓冲池中出现太多脏页时会触发 Checkpoint。 通过 `innodb_max_dirty_pages_pct` 可配置。
