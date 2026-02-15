# GoScript 功能实现现状

## 概览
- 本文基于语法文件与执行器实现进行对照梳理，包含已实现能力与缺失/不完整能力列表。
- 主要依据文件：语法 [GoScript.g4](../GoScript.g4)、执行器 [executor.go](../executor.go)、栈与作用域 [stack.go](../stack.go)、[scope.go](../scope.go)、测试 [executor_test.go](../executor_test.go)。

## 已实现能力
- 解析入口与执行框架：CompilationUnit 由执行器 Visitor 执行 [go_script.go](../go_script.go#L1-L24)、[executor.go](../executor.go#L58-L114)
- 作用域与变量查找：ScopeStack + 逐层查找变量 [scope.go](../scope.go#L1-L30)、[executor.go](../executor.go#L804-L824)
- 变量声明（带类型）：默认值、基础类型转换、数组/Map 初始化时的动态类型修正 [executor.go](../executor.go#L117-L153)、[executor.go](../executor.go#L734-L804)
- 基础字面量与标识符取值：int/float/char/string/bool/null [executor.go](../executor.go#L642-L684)
- 表达式运算：算术、比较、逻辑、单目、三元 [executor.go](../executor.go#L284-L381)、[executor.go](../executor.go#L593-L599)
- 变量创建并赋值（:=）：支持多 lhs 与多 rhs [executor.go](../executor.go#L304-L333)
- 赋值写回（=）：支持标识符、数组索引、Map 索引、map 选择器写入 [executor.go](../executor.go#L405-L489)
- 索引读：[]interface{}、map[interface{}]interface{}、map[string]interface{} [executor.go](../executor.go#L383-L403)
- 选择器读：map 与 struct 反射读取 [executor.go](../executor.go#L491-L511)
- 集合构造：new + array/map creator [executor.go](../executor.go#L513-L675)
- 控制流（基础）：if、for、break [executor.go](../executor.go#L212-L249)
- Go 函数调用：RegisterFunc + 反射调用（只支持 Go 绑定函数） [executor.go](../executor.go#L36-L38)、[executor.go](../executor.go#L551-L591)
- 栈行为：Scope/Function 栈顶在末尾的 LIFO 语义 [stack.go](../stack.go#L11-L61)

## 缺失或不完整能力
- 函数体系：语法支持函数声明与返回值，但执行器未实现函数声明/参数绑定/返回流程 [GoScript.g4](../GoScript.g4#L6-L33)、[executor.go](../executor.go#L91-L101)
- return/continue 语义未贯穿语句序列：return 不会中断后续语句，continue 标志未在 for 中生效 [executor.go](../executor.go#L236-L278)、[executor.go](../executor.go#L62-L71)
- `--` 未实现：SelfAddExpr 仅实现自增，且仅支持标识符 [GoScript.g4](../GoScript.g4#L148-L149)、[executor.go](../executor.go#L520-L545)
- 变量声明中的 `{}` 初始化器：语法允许 array/map initializer，但 VisitVariableInitializer 未处理 array/map 分支 [GoScript.g4](../GoScript.g4#L64-L88)、[executor.go](../executor.go#L157-L162)
- 选择器写入仅支持 map：struct 字段写入未实现 [executor.go](../executor.go#L443-L463)、[executor.go](../executor.go#L491-L511)
- 复杂 LValue 链式写入：仅回溯到 base Identifier 再写，无法覆盖多级链式路径 [executor.go](../executor.go#L412-L474)
- connector/error/dynamic 深层语义：connectorCreator 未实现，error 仅映射为动态类型 [executor.go](../executor.go#L680-L682)、[executor.go](../executor.go#L717-L721)
- 脚本函数调用：CallExpr 仅按名称从 GoVarMap 取函数，不支持脚本函数或方法调用 [executor.go](../executor.go#L551-L556)

## 测试覆盖情况
- 已覆盖：算术、:= 创建、数组/Map 索引写入、for+break、Go 函数调用 [executor_test.go](../executor_test.go)
- 未覆盖：函数声明与 return、continue、--、声明中的 `{}` initializer、selector 写入 struct、复杂 LValue、多返回值绑定等
