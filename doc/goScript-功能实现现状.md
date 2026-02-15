# GoScript 功能实现现状

## 概览
- 本文基于语法文件与执行器实现进行对照梳理，包含已实现能力与缺失/不完整能力列表。
- 主要依据文件：语法 [GoScript.g4](../GoScript.g4)、执行器 [executor.go](../executor.go)、栈与作用域 [stack.go](../stack.go)、[scope.go](../scope.go)、测试 [executor_test.go](../executor_test.go)。

## 已实现能力
- 解析入口与执行框架：CompilationUnit 由执行器 Visitor 执行 [go_script.go](../go_script.go#L1-L24)、[executor.go](../executor.go#L68-L152)
- 作用域与变量查找：ScopeStack + 逐层查找变量 [scope.go](../scope.go#L1-L34)、[executor.go](../executor.go#L1030-L1050)
- 变量声明（带类型）：默认值、基础类型转换、数组/Map 初始化时的动态类型修正，支持变量声明内 `{}` 初始化器，支持 `map<..>[]` 与 `connector<..>` 类型解析；数组维度运行时元信息可区分多维数组 [executor.go](../executor.go#L168-L237)、[executor.go](../executor.go#L1157-L1235)
- 基础字面量与标识符取值：int/float/char/string/bool/null [executor.go](../executor.go#L754-L812)
- 表达式运算：算术、比较、逻辑、单目、三元 [executor.go](../executor.go#L284-L381)、[executor.go](../executor.go#L689-L744)
- 变量创建并赋值（:=）：支持多 lhs 与多 rhs，支持单一多返回值拆包 [executor.go](../executor.go#L304-L333)、[executor.go](../executor.go#L481-L536)
- 赋值写回（=）：标识符、数组索引、Map 索引、map/struct 选择器写入，支持链式 LValue 与多返回值拆包 [executor.go](../executor.go#L405-L691)
- 索引读：[]interface{}、map[interface{}]interface{}、map[string]interface{} [executor.go](../executor.go#L383-L403)
- 选择器读：map 与 struct 反射读取 [executor.go](../executor.go#L570-L589)
- 自增自减：++/-- 支持标识符、索引、选择器 LValue [executor.go](../executor.go#L599-L733)
- 集合构造：new + array/map creator [executor.go](../executor.go#L912-L999)
- primitive/dynamic creator：支持显式类型转换 [executor.go](../executor.go#L1143-L1155)
- 控制流：if、for、break、continue、return [executor.go](../executor.go#L269-L312)
- 脚本函数：声明注册、调用、参数绑定与 return 结果处理 [executor.go](../executor.go#L107-L200)、[executor.go](../executor.go#L551-L676)
- Go 函数调用：RegisterFunc + 反射调用 [executor.go](../executor.go#L39-L49)、[executor.go](../executor.go#L551-L677)
- connector：RegisterConnector + new connector<> 构造，方法选择器调用 [executor.go](../executor.go#L43-L49)、[executor.go](../executor.go#L551-L591)、[executor.go](../executor.go#L983-L999)
- connector 运行时适配器：GoCaller 方法分派与参数转换 [go_caller.go](../go_caller.go#L1-L62)
- 栈行为：Scope/Function 栈顶在末尾的 LIFO 语义 [stack.go](../stack.go#L1-L73)

## 缺失或不完整能力
- killFlag 保留字段未定义触发路径，仅用于占位 [executor.go](../executor.go#L11-L24)

## 测试覆盖情况
- 已覆盖：算术、:= 创建、数组/Map 索引写入、for+break/continue、脚本函数、--、声明 `{}` 初始化器、connector 运行时适配器与基础调用、复杂 LValue 链式写入、多返回值拆包、类型解析与 creator 强转、路径规划脚本用例、多维数组运行时维度、多维数组路径规划 [executor_test.go](../executor_test.go)
