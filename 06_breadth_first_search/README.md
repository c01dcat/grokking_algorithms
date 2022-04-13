# 广度优先搜索

广度优先搜索（breadth-first search，BFS）能够找出两样东西之间的最短距离，不过最短距离的含义有很多。

## 图

最短路径问题（shortest-path problem），解决最短路径问题的算法被称为广度优先搜索。 

* 图由节点（node）和边（edge）组成
* 一个节点可能与众多节点直接相连，这些节点被称为邻居
* 有向图（directed graph），其中的关系是单向的。
* 无向图（undirected graph）没有箭头，直接相连的节点互为邻居。
* 树是一种特殊的图，其中没有往后指的边。

队列是一种先进先出（First In First Out，FIFO）的数据结构，而栈是一种后进先出（Last In First Out，LIFO）的数据结构。

广度优先搜索的运行时间为 O(V + E),其中 V 为顶点（vertices）数，E 为边数。