# 快速排序

## 分而治之

分而治之（divide and conquer，D&C）——一种著名的递归式问题解决方法。  

使用 D&C 解决问题的过程包括两个步骤：
* 找出基线条件，这种条件必须尽可能简单
* 不断将问题分解（或者说缩小规模），直到符合基线条件

编写涉及数组的递归函数时, 基线条件通常是数组为空或只包含一个元素。

## 快速排序

归纳证明是一种证明算法行之有效的方式, 它分两步: 基线条件和归纳条件。  

快速排序的独特之处在于,其速度取决于选择的基准值。  

* 合并排序 (merge sort) 的排序算法，其运行时间为 O(n * log<sub>2</sub> n)
* 快速排序的情况比较棘手，在最糟情况下，其运行时间为 O(n<sup>2</sup>)。在平均情况下，快速排序的运行时间为 O(n * log<sub>2</sub> n)