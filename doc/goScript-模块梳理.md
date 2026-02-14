# goScript 实现模块梳理

本文基于项目中的 GoScript.g4 语法定义与现有代码结构，梳理解释器/运行时需要实现的模块与能力清单，并标注当前实现状态与建议优先级，便于后续迭代。

## 1. 解析与执行总览
- 语法定义：见文件 GoScript.g4（表达式、语句、类型、构造器等）。
- 解析入口：GoScript.Init 负责创建词法/语法分析器并构建 ParseTree，随后将树交给执行器访问执行。[go_script.go](file:///Users/bytedance/Go/src/wsydehub/goScript/go_script.go#L1-L24)
- 执行器：采用 Visitor 模式，对各语法节点进行解释执行。[executor.go](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go)
- 运行时基础：作用域/符号表、变量与类型、函数实体与栈、与 Go 的互操作数据结构。

## 2. 运行时与基础结构
- 作用域与符号表
  - Scope/ScopeType：管理变量可见性与生命周期。[scope.go](file:///Users/bytedance/Go/src/wsydehub/goScript/scope.go)
  - ScopeStack：维护当前作用域栈。[stack.go](file:///Users/bytedance/Go/src/wsydehub/goScript/stack.go#L3-L33)
- 变量与类型
  - Variable/VariableType：抽象脚本变量与基础类型。[variale.go](file:///Users/bytedance/Go/src/wsydehub/goScript/variale.go)
  - 基础类型映射与默认值、数值/布尔/字符串转换（已实现初版）。[executor.go](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L423-L482)
- 函数系统
  - Function：封装函数名、参数/返回列表、函数体与作用域。[function.go](file:///Users/bytedance/Go/src/wsydehub/goScript/function.go)
  - FuncStack：执行期函数栈。[stack.go](file:///Users/bytedance/Go/src/wsydehub/goScript/stack.go#L35-L65)
  - funcMap/GoVarMap：用户态函数注册/Go 互操作的驻留容器（需补全使用策略）。[executor.go](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L17-L22)
- 执行器（Visitor）
  - 基础遍历、终结符处理、错误节点兜底已实现。[executor.go](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L54-L75)
  - 块级作用域推入/弹出与语句执行分发已就绪。[executor.go](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L93-L101)

## 3. 语句语义实现清单
- 已实现
  - if 语句（条件求值与 then/else 分支）。[executor.go:VisitIfStatement](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L174-L183)
  - 表达式语句（直接求值）。[executor.go:VisitExpressionStatement](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L216-L218)
  - 变量声明：类型解析、初始化表达式求值、变量入作用域。[executor.go:VisitVariableDeclaration](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L103-L124)
  - return 语句：返回表达式值（尚未配合函数返回机制统一处理）。[executor.go:VisitReturnStatement](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L201-L206)
- 待实现/完善
  - for 语句：forControl（Init/Condition/Update）执行流程与作用域控制；break/continue 控制流标志与跳转。[executor.go](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L185-L214)
  - 函数声明注册/调用流程与返回值传递；CreateAndAssign（:=）真正创建变量的行为与多返回值匹配。

## 4. 表达式语义实现清单
- 已实现
  - 字面量：整数、浮点、字符、字符串、布尔、null。[executor.go:VisitLiteral/VisitIntegerLiteral](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L354-L379)
  - 标识符读取：从当前作用域向上查找变量。[executor.go:VisitPrimary](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L337-L352)
  - 一元运算：+x、-x、!x。[executor.go:VisitUnaryExpr](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L278-L287)
  - 二元算术：+、-、*、/、%（含基础的整/浮数混合策略）。[executor.go:VisitAddExpr/VisitMulExpr](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L248-L255)
  - 比较：<=、>=、>、<、==、!=（数值与字符串）。[executor.go:VisitConditionalExpr/compare](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L257-L276)
  - 逻辑：&&、||（短路）。[executor.go:VisitAndExpr/VisitOrExpr](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L232-L238)
  - 三元：cond ? a : b。[executor.go:VisitTernaryExpr](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L329-L335)
  - PrimaryExpr 分派与 ExpressionList 聚合。[executor.go:VisitPrimaryExpr/VisitExpressionList](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L321-L387)
- 待实现/完善
  - 赋值：AssignExpr 将 RHS 写回 LValue（变量/选择器/索引），含多重赋值。
  - CreateAndAssignExpr（:=）：在当前作用域创建新变量并初始化。
  - 选择器/索引：SelectorExpr、IndexExpr 对象字段与数组/Map 访问（含边界/缺失处理）。
  - 调用：CallExpr 支持脚本内函数与 Go 绑定函数；参数求值/形参匹配，返回值绑定。
  - 自增自减：SelfAddExpr 对 LValue 的 ++/-- 写回。
  - new/creator：CreateExpr 与 Map/Array/Primitive/Dynamic/Connector 的构造语义与初始化器（ArrayInitializer/MapInitializer）。

## 5. 类型系统与构造器
- 语法支持
  - primitiveType：int/uint/float/bool/char/string
  - dynamic/error：动态类型与错误占位
  - mapType、connectorType、数组后缀 []（如 int[]）
  - 初始化器：arrayInitializer、mapInitializer
- 运行时策略（当前初版）
  - 简单类型映射与默认值、基本数值/布尔/字符串的转换与二元运算数值提升。[executor.go:helpers](file:///Users/bytedance/Go/src/wsydehub/goScript/executor.go#L483-L682)
  - creator 占位实现：Primitive/DynamicCreator 返回内含表达式的值；其余待完善。
- 待完善
  - 数组/Map 的具体承载（Go slice/map）与下标/键访问及初始化器展开。
  - connector 的构造与方法分派策略（可基于反射或注册表）。
  - 更细的类型检查与错误提示（算术/比较的操作数类型约束）。

## 6. 与 Go 的互操作（建议）
- 变量绑定：通过 GoVarMap 将外部 Go 值注入脚本环境（读取/写回策略）。
- 函数绑定：注册 Go 函数到 funcMap，CallExpr 通过反射分派并完成参数/返回值转换。
- 安全与错误处理：反射调用异常包装为脚本运行时错误返回。

## 7. 错误处理与调试建议
- 语法错误：由 ANTLR 的错误恢复机制提供；可配置监听器收集错误。
- 运行时错误：如未绑定变量、类型不匹配、下标越界、除零等，统一返回错误对象或抛出异常机制（根据 dynamic/error 的设计选择）。
- 调试：为各 Visitor 分支添加统一的 trace/日志钩子（可后续注入，不污染核心逻辑）。

## 8. 模块实施清单（按优先级）
1) 表达式写回链
   - VisitAssignExpr（含多赋值）— 高优
   - VisitIndexExpr/VisitSelectorExpr 的取值与写回— 高优
   - VisitSelfAddExpr 对 LValue 的 ++/--— 中高
2) 集合与构造器
   - Map/Array 初始化器与 new 语义— 高优
   - Index 边界/键缺失处理— 高优
3) 函数与控制流
   - 函数声明注册与调用、返回值匹配— 高优
   - for 循环完整执行流与 break/continue 标志— 高优
4) Go 互操作
   - Go 函数/变量注册与反射分派、类型转换策略— 中高
5) 完善类型系统与错误
   - 更严格的类型检查与友好错误— 中优

## 9. 当前进度速记
- 已能执行最基本的表达式计算（算术/比较/逻辑/三元/字面量/标识符读取），以及变量声明与 if 分支求值。
- 尚未打通：赋值写回、索引/选择器、函数/调用、数组与 Map 初始化、for 语句与循环控制。

如需，我可以基于上述清单逐项补齐实现，并附加示例与简单单测框架，帮助快速回归。 

