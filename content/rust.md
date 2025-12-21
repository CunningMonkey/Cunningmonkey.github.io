---
title: "Rust学习"
date: 2025-12-21
summary: "这里有一些Rust."
---

## Trick

### 整形上溢/下溢
实际使用中，`uszie`容易出现下溢（0 - 1），rust提供了`saturating_sub`，能够处理各种整数类型返回边界值
`usize::saturating_sub`（以及其它整数类型的 `saturating_sub`）做的是“饱和减法”：正常计算 `self - rhs`，但如果会发生溢出/下溢则返回类型的边界值而不是 panic 或环绕。

要点：

签名：`fn saturating_sub(self, rhs: Self) -> Self`（在所有整型上可用）。
对无符号类型（如 usize）：下溢时结果饱和为 0。例如：
``` rust
5 as usize.saturating_sub(3) == 2
0 as usize.saturating_sub(1) == 0 // 不会下溢 panic
```
对有符号类型（如 i8）：若计算超出范围则饱和到 MIN 或 MAX：
``` rust
i8::MIN.saturating_sub(1) == i8::MIN
i8::MAX.saturating_sub(-1) == i8::MAX
``` 
和其它方法对比：

`checked_sub`：返回 Option，溢出时返回 None。
`wrapping_sub`：按模 2^N 环绕（不会 panic，但会得到环绕结果）。
`saturating_sub`：溢出时返回边界值（方便避免下溢 panic，同时保留一个有意义的最小/最大值）。

