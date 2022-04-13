# 递归

* Leigh Caldwell 在 Stack Overflow 上说的一句话: "如果使用循环,程序的性能可能更高;如果使用递归,程序可能更容易理解。如何选择要看什么对你来说更重要。 ” 
* 每个递归函数都有两部分：基线条件（base case）和递归条件（recursive case）。递归条件指的是函数调用自己，而基线条件则指的是函数不再调用自己，从而避免形成无限循环。
* 调用另一个函数时,当前函数暂停并处于未完成状态。该函数的所有变量的值都还在内存中。

## 尾递归

尾调用是指一个函数里的最后一个动作是返回一个函数的调用结果的情形，即这个调用的返回值直接被当前的函数返回的情形。若这个函数在尾部位置调用本身，则称为尾递归。