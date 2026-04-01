## 目标范围
基于当前语法与执行器缺口，补齐以下能力：
1) 脚本函数系统（声明/调用/返回）
2) continue/return 控制流语义完善
3) 变量声明中的 array/map initializer
4) ++/-- 完整语义（含 -- 与复杂 LValue）
5) connector creator 与方法分派最小实现

## 实现步骤
### 1. 脚本函数体系
- 完成 VisitFunctionDeclaration：注册 Function 到 Executor.funcMap，并建立函数签名与参数列表。
- 完成 VisitFormalParameters/VisitFormalParameterDecl/VisitReturnType：解析参数类型与返回类型。
- 在 VisitCallExpr 中：
  - 优先解析脚本函数（funcMap）并执行；
  - 保留 GoVarMap 反射调用作为后备。
- 执行函数时：
  - 新建函数作用域 + 参数绑定；
  - 运行函数体 block；
  - 处理 returnFlag/returnValue；
  - 清理作用域。

### 2. return/continue 控制流
- return：VisitReturnStatement 置 returnFlag 并记录 returnValue。
- continue：在 VisitForStatement 中检测 continueFlag，立即进入 update/下一轮，并清理 flag。
- 统一 block/statement 遍历：在 VisitBlock/VisitChildren 中遇到 returnFlag/breakFlag/continueFlag 时中止遍历。

### 3. 变量声明 initializer
- VisitVariableInitializer 增加 arrayInitializer/mapInitializer 分支，复用 VisitArrayInitializer/VisitMapInitializer 返回值。

### 4. ++/-- 完整语义
- 识别 ++/-- 操作符，区分增减。
- 支持标识符、索引、选择器 LValue 写回。

### 5. connector 最小实现
- 实现 VisitConnectorCreator：基于注册表或反射构造（先按现有结构最小化可用）。
- 扩展 VisitSelectorExpr/VisitCallExpr：支持 struct 方法反射调用（MethodByName）。

## 验证
- 扩展 executor_test.go：
  - 函数声明/调用/return
  - continue 生效
  - 变量声明 initializer
  - -- 和复杂 LValue 的 ++/--
  - connector（若提供可构造结构）
- 运行 go test ./...

确认后我会开始修改代码并补齐测试。